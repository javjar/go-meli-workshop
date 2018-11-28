package node_test

import (
	"testing"

	"github.com/javjar/go-meli-workshop/class2/third/node"
)

func TestAdd(t *testing.T) {
	tt := []struct {
		name     string
		node     *node.Node
		args     []int
		expected string
	}{
		{
			name:     "When arguments are [1], the expected output is [1]",
			node:     &node.Node{Value: 1},
			args:     []int{1},
			expected: "1",
		},
		{
			name:     "When arguments are [1 2 3], the expected output is [1 2 3]",
			node:     &node.Node{Value: 1},
			args:     []int{1, 2, 3},
			expected: "1 2 3",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.args[1:] {
				node.Add(tc.node, v)
			}

			result := node.TraverseInOrder(tc.node)
			if tc.expected != result {
				t.Errorf("Expected: %s, but got: %s.", tc.expected, result)
			}
		})
	}
}
