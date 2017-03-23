package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TrieInsert(t *testing.T) {
	assert := assert.New(t)

	trie := NewTrie()
	err := trie.Insert("")
	assert.NotNil(err, "Error should not be nil")
	assert.Equal(err.Error(), "Empty string inserted.", "Emtpy string added")

	trie.Insert("ankur")
	trie.Insert("ana")

	_, err = trie.Search("")
	assert.NotNil(err, "Error should not be nil")
	assert.Equal(err.Error(), "Empty string inserted.", "Emtpy string added")

	found, _ := trie.Search("ankur")
	assert.True(found, "Ankur was in trie")
	found, _ = trie.Search("viv")
	assert.False(found, "viv should not be in trie")

	found, _ = trie.Search("anku")
	assert.False(found, "anku should not be in trie")
}
