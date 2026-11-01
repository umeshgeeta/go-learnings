package lessons_test

import (
	"go-learnings/pkg/lessons"
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *lessons.TreeNode
		expected []int
	}{
		{
			name: "Example 1",
			root: &lessons.TreeNode{
				Val: 1,
				Left: &lessons.TreeNode{
					Val: 2,
					Left: &lessons.TreeNode{
						Val: 4,
					},
					Right: &lessons.TreeNode{
						Val: 5,
					},
				},
				Right: &lessons.TreeNode{
					Val: 3,
				},
			},
			expected: []int{1, 2, 4, 5, 3},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name: "Single node",
			root: &lessons.TreeNode{
				Val: 1,
			},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lessons.PreorderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
