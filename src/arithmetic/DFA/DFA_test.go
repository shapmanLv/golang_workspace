package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CanMatch_SingleWord(t *testing.T) {
	assert.Equal(t, 1, Detection("我爱我家", BuildSearchTree([]string{"爱", "家"})))
	assert.Equal(t, 1, Detection("我爱我家", BuildSearchTree([]string{"爱"})))
	assert.Equal(t, 0, Detection("我爱我家", BuildSearchTree([]string{"我"})))
}
