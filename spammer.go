package main

import (
	"fmt"
	"sort"
)

func RunPipeline(cmds ...cmd) {
	chans := make([]chan interface{}, len(cmds)+1)
	for i := range chans {
		chans[i] = make(chan interface{})
	}

	for i, foo := range cmds {
		go func(i int, cmd cmd) {
			defer close(chans[i+1])
			// cmd can be blocked by channel reciving or sending operation inside itself
			// in this case whole goroutine blocks
			cmd(chans[i], chans[i+1])
		}(i, foo)
	}

	for line := range chans[len(chans)-1] {
		fmt.Println(line)
	}
}

func SelectUsers(in, out chan interface{}) {
	// 	in - string
	// 	out - User
	var ids = map[uint64]bool{}
	for email := range in {
		user := GetUser(email.(string))

		if _, ok := ids[user.ID]; !ok {
			out <- user
		}

		ids[user.ID] = true
	}
}

func SelectMessages(in, out chan interface{}) {
	// 	in - User
	// 	out - MsgID
	for user := range in {
		message, err := GetMessages(user.(User))

		if err != nil {
			fmt.Printf("error %v", err)
		}

		out <- message
	}
}

func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData
	for idsInterface := range in {
		for _, id := range idsInterface.([]MsgID) {
			hasSpam, err := HasSpam(id)

			if err != nil {
				fmt.Printf("error %v", err)
			}

			out <- MsgData{
				ID:      id,
				HasSpam: hasSpam,
			}
		}
	}
}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string
	msgDatas := make(msgDatas, 0)
	for msgDataInterface := range in {
		msgDatas = append(msgDatas, msgDataInterface.(MsgData))
	}
	sort.Sort(msgDatas)
	for _, msgData := range msgDatas {
		out <- fmt.Sprintf("%t %d", msgData.HasSpam, msgData.ID)
	}
}

type msgDatas []MsgData

func (m msgDatas) Len() int      { return len(m) }
func (m msgDatas) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m msgDatas) Less(i, j int) bool {
	if m[i].HasSpam == m[j].HasSpam {
		return m[i].ID < m[j].ID
	}
	return m[i].HasSpam && !m[j].HasSpam
}
