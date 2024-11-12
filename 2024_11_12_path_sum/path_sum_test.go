package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	assert.Equal(t, true, HasPathSum(tree, 22))

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	assert.Equal(t, false, HasPathSum(tree, 5))

	tree = nil

	assert.Equal(t, false, HasPathSum(tree, 0))
}
