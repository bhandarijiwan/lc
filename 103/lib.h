// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
#ifndef LIB_HPP
#define LIB_HPP
#include "fmt/core.h"
#include "fmt/ranges.h"

// Definition for a binary tree node.
struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right)
      : val(x), left(left), right(right) {}
};

#include <vector>

class Solution {
private:
  void level_order(TreeNode *root, TreeNode *left, TreeNode *right,
                   size_t level, std::vector<std::vector<int>> &output) {
    if (root == nullptr && left == nullptr && right == nullptr) {
      return;
    }
    size_t next_level = level + 1;
    TreeNode *next_right_left = nullptr;
    TreeNode *next_right_right = nullptr;
    TreeNode *next_left_left = nullptr;
    TreeNode *next_left_right = nullptr;
    if (root != nullptr) {
      if (output.size() <= level) {
        output.push_back({root->val});
      } else {
        output[level].push_back({root->val});
      }
      if (root->left != nullptr) {
        next_right_left = root->left->left;
        next_right_right = root->left->right;
      }
      if (root->right != nullptr) {
        next_left_left = root->right->left;
        next_left_right = root->right->right;
      }
    } else {
      if (output.size() <= level) {
        output.push_back({});
      }
    }
    if (level % 2 == 0) {
      this->level_order(right, next_right_left, next_right_right, next_level,
                        output);
      this->level_order(left, next_left_left, next_left_right, next_level,
                        output);
    } else {
      this->level_order(left, next_left_left, next_left_right, next_level,
                        output);
      this->level_order(right, next_right_left, next_right_right, next_level,
                        output);
    }
  }

public:
  std::vector<std::vector<int>> zigzagLevelOrder(TreeNode *root) {
    std::vector<std::vector<int>> output{};
    if (root != nullptr) {
      this->level_order(root, root->left, root->right, 0, output);
    }
    return output;
  }
};

#endif /* LIB_HPP */