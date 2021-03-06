package dicethrow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiceThrow_DP1(t *testing.T) {
	counter := DiceThrow_DP(4, 2, 1)
	assert.Equal(t, counter, 0)

	counter = DiceThrow_DP(2, 2, 3)
	assert.Equal(t, counter, 2)

	counter = DiceThrow_DP(6, 3, 8)
	assert.Equal(t, counter, 21)

	counter = DiceThrow_DP(4, 2, 5)
	assert.Equal(t, counter, 4)

	counter = DiceThrow_DP(4, 3, 5)
	assert.Equal(t, counter, 6)
}
