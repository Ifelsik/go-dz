package calc

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	t.Run("Testing default", func(t *testing.T) {
		s := NewStack()
		
		require.True(t, s.IsEmpty())  // stack must be empty
		
		s.Push("1")
		require.True(t, !s.IsEmpty())

		got, err := s.Top()
		require.Equal(t, nil, err)
		require.Equal(t, "1", got)

		got, err = s.Pop()
		require.Equal(t, nil, err)
		require.Equal(t, "1", got)

		require.True(t, s.IsEmpty())

		got, err = s.Pop()
		require.Equal(t, "", got)
		require.Error(t, err, "stack is empty")
	})
	t.Run("Multiple push and pops", func(t *testing.T) {
		s := NewStack()

		input1    := []string{"1", "2", "3", "4", "5"}
		expected1 := []string{"5", "4"}

		for _, value := range input1 {
			err := s.Push(value)
			require.Equal(t, nil, err)

			require.True(t, !s.IsEmpty())

			got, err := s.Top()
			require.Equal(t, nil, err)
			require.Equal(t, value, got)
		}

		for _, value := range expected1 {
			got, err := s.Pop()
			require.Equal(t, nil, err)
			require.Equal(t, value, got)
		}  // removed 5 and 4. Have 3, 2 and 1 in stack

		require.True(t, !s.IsEmpty())

		input2    := []string{"6", "7"}
		expected2 := []string{"7", "6", "3", "2", "1"}

		for _, value := range input2 {
			err := s.Push(value)
			require.Equal(t, nil, err)

			require.True(t, !s.IsEmpty())

			got, err := s.Top()
			require.Equal(t, nil, err)
			require.Equal(t, value, got)
		}

		for _, value := range expected2 {
			got, err := s.Pop()
			require.Equal(t, nil, err)
			require.Equal(t, value, got)
		}

		require.True(t, s.IsEmpty())
	})
}
