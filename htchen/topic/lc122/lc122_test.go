package lc122

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	require.Equal(t, 4, maxProfit([]int{1, 2, 3, 4, 5}))

	require.Equal(t, 7, maxProfit([]int{7, 1, 5, 3, 6, 4}))

	require.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}

func Test_maxProfit1(t *testing.T) {
	require.Equal(t, 4, maxProfit1([]int{1, 2, 3, 4, 5}))

	require.Equal(t, 7, maxProfit1([]int{7, 1, 5, 3, 6, 4}))

	require.Equal(t, 0, maxProfit1([]int{7, 6, 4, 3, 1}))
}
