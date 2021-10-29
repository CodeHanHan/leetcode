package topic136

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singleNumber(t *testing.T) {
	require.Equal(t, 1, singleNumber([]int{2, 2, 1}))

	require.Equal(t, 1, singleNumber([]int{1}))

	require.Equal(t, 4, singleNumber([]int{4, 1, 2, 1, 2}))

	require.Equal(t, 5, singleNumber([]int{4, 1, 2, 1, 2, 4, 5}))
}
