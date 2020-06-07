package _211_添加与搜索单词_数据结构设计

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	w := Constructor()
	w.AddWord("a")
	w.AddWord("a")
	assert.Equal(t, true, w.Search("."))
	assert.Equal(t, true, w.Search("a"))
	assert.Equal(t, false, w.Search("aa"))
	assert.Equal(t, true, w.Search("a"))
	assert.Equal(t, false, w.Search(".a"))
	assert.Equal(t, false, w.Search("a."))
}