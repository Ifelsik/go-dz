package main

import (
	"fmt"
	"sort"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	in := make(chan interface{})
	wg := &sync.WaitGroup{}
	for _, function := range cmds {
		out := make(chan interface{})
		wg.Add(1)
		go func(wg *sync.WaitGroup, cmd cmd, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			// cmd can be blocked by channel reciving or sending operation inside itself
			// in this case whole goroutine blocks
			cmd(in, out)
		}(wg, function, in, out)
		in = out
	}
	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	// 	in - string
	// 	out - User
	var ids = &sync.Map{}
	var wg = &sync.WaitGroup{}
	for email := range in {
		wg.Add(1)
		go func(wg *sync.WaitGroup, out chan interface{}, email interface{}) {
			defer wg.Done()
			user := GetUser(email.(string))

			if _, ok := ids.Load(user.ID); !ok {
				ids.Store(user.ID, true)
				out <- user
			}
		}(wg, out, email)
	}
	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	// 	in - User
	// 	out - MsgID
	batch := make([]User, 0, GetMessagesMaxUsersBatch)
	wg := &sync.WaitGroup{}
	asyncBatchedGetMessages := func(wg *sync.WaitGroup, out chan interface{}, batch []User) {
		defer wg.Done()
		
		res, err := GetMessages(batch...)
		if err != nil {
			fmt.Printf("in SelectMessages %v", err)
			return
		}

		for _, message := range res {
			out <- message
		}
	}

	for user := range in {
		if user, ok := user.(User); ok {
			batch = append(batch, user)
		}
		if len(batch) == GetMessagesMaxUsersBatch {
			wg.Add(1)
			go asyncBatchedGetMessages(wg, out, batch)
			batch = nil
		}
	}

	if len(batch) > 0 {
		wg.Add(1)
		go asyncBatchedGetMessages(wg, out, batch)
	}
	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData
	wg := &sync.WaitGroup{}
	for i := 0; i < HasSpamMaxAsyncRequests; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, in, out chan interface{}) {
			defer wg.Done()

			for msgId := range in {
				id, ok := msgId.(MsgID); 

				if !ok {
					fmt.Printf("in CheckSpam can't convert msgIds; got %T", msgId)
					continue
				}

				hasSpam, err := HasSpam(id)

				if err != nil {
					fmt.Printf("in CheckSpam got %v", err)
				}

				out <- MsgData{
					ID:      id,
					HasSpam: hasSpam,
				}
			}
		}(wg, in, out)
	}
	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string
	msgDatas := make(msgDataSlice, 0)
	for msgDataInterface := range in {
		msgDatas = append(msgDatas, msgDataInterface.(MsgData))
	}
	sort.Sort(msgDatas)
	for _, msgData := range msgDatas {
		out <- fmt.Sprintf("%t %d", msgData.HasSpam, msgData.ID)
	}
}

type msgDataSlice []MsgData

// implementation of Sort interface for []MsgData
func (m msgDataSlice) Len() int      { return len(m) }
func (m msgDataSlice) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m msgDataSlice) Less(i, j int) bool {
	if m[i].HasSpam == m[j].HasSpam {
		return m[i].ID < m[j].ID
	}
	return m[i].HasSpam && !m[j].HasSpam
}
