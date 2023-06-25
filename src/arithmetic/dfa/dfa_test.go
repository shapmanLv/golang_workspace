package dfa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 单个简单词 可以匹配的情况
func Test_CanMatch_SingleWord(t *testing.T) {
	assert.Equal(t, 1, Detection("我爱我家", BuildSearchTree([]string{"爱"})))
	assert.Equal(t, 0, Detection("我爱我家", BuildSearchTree([]string{"我"})))
	assert.Equal(t, 1, Detection("我爱我家", BuildSearchTree([]string{"爱", "家"})))
}

// 容易混淆的词 可以匹配的情况
func Test_CanMatch_ConfuseWord(t *testing.T) {
	assert.Equal(t, 2, Detection("王八王八蛋", BuildSearchTree([]string{"王八蛋"})))
	assert.Equal(t, 3, Detection("王八王八蛋", BuildSearchTree([]string{"八蛋"})))
	assert.Equal(t, 0, Detection("王八王八蛋", BuildSearchTree([]string{"王八蛋", "王八"})))
}

// 单个简单词 不可以匹配的情况
func Test_FailToMatchMatch_SingleWord(t *testing.T) {
	assert.Equal(t, -1, Detection("我爱我家", BuildSearchTree([]string{"你"})))
	assert.Equal(t, -1, Detection("我爱我家", BuildSearchTree([]string{"好"})))
	assert.Equal(t, -1, Detection("我爱我家", BuildSearchTree([]string{"他", "它"})))
}

// 容易混淆的词 不可以匹配的情况
func Test_FailToMatch_ConfuseWord(t *testing.T) {
	assert.Equal(t, -1, Detection("王八王八蛋", BuildSearchTree([]string{"王八蛋王"})))
	assert.Equal(t, -1, Detection("王八王八蛋", BuildSearchTree([]string{"八蛋子"})))
	assert.Equal(t, -1, Detection("王八王八蛋", BuildSearchTree([]string{"王八蛋子", "王八八"})))
}
