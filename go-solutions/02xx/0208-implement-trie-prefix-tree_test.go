package _2xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//["Trie","insert","search","search","startsWith","insert","search"]
//[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]
//[null,null,true,false,true,null,true]
func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	require.True(t, trie.Search("apple"))
	require.False(t, trie.Search("app"))
	require.True(t, trie.StartsWith("app"))
	trie.Insert("app")
	require.True(t, trie.Search("app"))
}

func TestTrie_StartsWith(t *testing.T) {
	tests := []struct {
		name   string
		insert []string
		want   map[string][2]bool
	}{
		{
			name:   "",
			insert: []string{"appl", "apple", "apbarat"},
			want: map[string][2]bool{
				"apple": {true, true},
				"appl":  {true, true},
				"app":   {false, true},
			},
		},
		{
			name:   "",
			insert: []string{"appl", "apple", "app", "apbarat"},
			want: map[string][2]bool{
				"abb":   {false, false},
				"ap":    {false, true},
				"apple": {true, true},
				"appl":  {true, true},
				"app":   {true, true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			for _, i := range tt.insert {
				this.Insert(i)
			}
			for k, v := range tt.want {
				got0 := this.Search(k)
				got1 := this.StartsWith(k)
				require.Equal(t, v[0], got0, k)
				require.Equal(t, v[1], got1, k)
			}
		})
	}
}
