package merkle

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	inputs [][]byte
}

var (
	testCases = []testCase{
		{
			inputs: [][]byte{
				[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"),
			},
		},
		{
			inputs: [][]byte{
				[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"),
				[]byte("6"), []byte("7"), []byte("8"), []byte("9"),
			},
		},
		{
			inputs: [][]byte{
				[]byte("1"), []byte("2"), []byte("3"), []byte("4"),
				[]byte("5"), []byte("6"), []byte("7"), []byte("8"),
			},
		},
	}
)

func Test_NewTree(t *testing.T) {
	sh := sha256.New()

	tc := testCases[0]
	tree := NewTree(sh, tc.inputs...)
	assert.Equal(t, 4, tree.Height())
	assert.Equal(t, 5, len(tree.Leafs()))
	assert.NotNil(t, tree.Root())
	printTree(tree)
}

func Test_NewTree_Proof_leaf_not_found(t *testing.T) {
	sh := sha256.New()
	tree := NewTree(sh, testCases[0].inputs...)

	assert.Nil(t, tree.Proof([]byte("foo")))
}

func Test_NewTree_Proof(t *testing.T) {
	sh := sha256.New()

	for _, tc := range testCases {
		tree := NewTree(sh, tc.inputs...)
		for _, node := range tree[0] {
			proof := tree.Proof(node)
			assert.True(t, IsValidProof(sh, node, proof))
		}
	}
}

func printTree(tree Tree) {
	for _, row := range tree {
		for _, item := range row {
			fmt.Printf("%x  ", item[26:])
		}
		fmt.Println()
	}
}
