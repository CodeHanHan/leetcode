package topic031

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextPermutation(t *testing.T) {
	var nums []int

	nums = []int{1, 2, 3}
	nextPermutation(nums)
	require.Equal(t, []int{1, 3, 2}, nums)

	nums = []int{3, 2, 1}
	nextPermutation(nums)
	require.Equal(t, []int{1, 2, 3}, nums)

	nums = []int{1}
	nextPermutation(nums)
	require.Equal(t, []int{1}, nums)

	nums = []int{1, 1, 5}
	nextPermutation(nums)
	require.Equal(t, []int{1, 5, 1}, nums)

	nums = []int{1, 3, 2}
	nextPermutation(nums)
	require.Equal(t, []int{2, 1, 3}, nums)

	nums = []int{1, 3, 2, 4}
	nextPermutation(nums)
	require.Equal(t, []int{1, 3, 4, 2}, nums)

	nums = []int{2, 3, 1}
	nextPermutation(nums)
	require.Equal(t, []int{3, 1, 2}, nums)
}
