package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	target := &TreeNode{Val: 5, Left: &TreeNode{Val: 6, Right: &TreeNode{Val: 18}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}
	tree := &TreeNode{Val: 3, Left: target, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}}
	root := &TreeNode{Val: 9, Left: &TreeNode{Val: 11, Right: &TreeNode{Val: 19, Left: tree}}}
	actual := distanceK(root, target, 2)
	expected := []int{18, 7, 4, 1, 19}
	areEqual := reflect.DeepEqual(actual, expected)
	if !areEqual {
		t.Fatalf("Expected %v, go %v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	// [0,null,1,null,2,null,3]
	// 1
	// 2
	target := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	tree := &TreeNode{Val: 0, Right: target}
	expected := []int{3}
	actual := distanceK(tree, target, 2)
	areEqual := reflect.DeepEqual(actual, expected)
	if !areEqual {
		t.Fatalf("Expected  %v, got %v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	// [0,2,1,null,null,3]
	// 3
	// 3
	target := &TreeNode{Val: 3}
	root := &TreeNode{Val: 0, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1, Left: target}}
	expected := []int{2}
	actual := distanceK(root, target, 3)
	areEqual := reflect.DeepEqual(expected, actual)
	if !areEqual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	// [0,1,null,null,2,null,3,null,4]
	// 3
	// 0
	target := &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}
	root := &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: target}}}
	actual := distanceK(root, target, 0)
	expected := []int{3}
	areEqual := reflect.DeepEqual(actual, expected)
	if !areEqual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}
