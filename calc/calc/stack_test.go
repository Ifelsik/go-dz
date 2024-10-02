package calc

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	t.Parallel()
	t.Run("Testing default", func(t *testing.T) {
		t.Parallel()

		s := NewStack()
		
		require.True(t, s.IsEmpty())  // stack must be empty
		
		s.Push("1")
		require.True(t, !s.IsEmpty())

		got, ok := s.Top()
		require.True(t, ok)
		require.Equal(t, "1", got)

		got, ok = s.Pop()
		require.True(t, ok)
		require.Equal(t, "1", got)

		require.True(t, s.IsEmpty())

		got, ok = s.Pop()
		require.Equal(t, "", got)
		require.True(t, !ok)
	})
	t.Run("Multiple push and pops", func(t *testing.T) {
		t.Parallel()
		
		s := NewStack()

		input1    := []string{"1", "2", "3", "4", "5"}
		expected1 := []string{"5", "4"}

		for _, value := range input1 {
			ok := s.Push(value)
			require.True(t, ok)

			require.True(t, !s.IsEmpty())

			got, ok := s.Top()
			require.True(t, ok)
			require.Equal(t, value, got)
		}

		for _, value := range expected1 {
			got, ok := s.Pop()
			require.True(t, ok)
			require.Equal(t, value, got)
		}  // removed 5 and 4. Have 3, 2 and 1 in stack

		require.True(t, !s.IsEmpty())

		input2    := []string{"6", "7"}
		expected2 := []string{"7", "6", "3", "2", "1"}

		for _, value := range input2 {
			ok := s.Push(value)
			require.True(t, ok)

			require.True(t, !s.IsEmpty())

			got, ok := s.Top()
			require.True(t, ok)
			require.Equal(t, value, got)
		}

		for _, value := range expected2 {
			got, ok := s.Pop()
			require.True(t, ok)
			require.Equal(t, value, got)
		}

		require.True(t, s.IsEmpty())
	})
}
