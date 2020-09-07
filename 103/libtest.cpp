#include "lib.h"
#include "fmt/core.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"

using namespace testing;
TEST(LibTest, TestCase1) {
  auto a1 = TreeNode{3};
  auto a2 = TreeNode{9};
  auto a3 = TreeNode{20};
  auto a4 = TreeNode{12};
  auto a5 = TreeNode{16};
  auto a6 = TreeNode{15};
  auto a7 = TreeNode{7};
  auto a8 = TreeNode{10};
  auto a9 = TreeNode{77};
  auto a10 = TreeNode{89};
  a1.left = &a2;
  a1.right = &a3;
  a2.left = &a4;
  a2.right = &a5;
  a3.left = &a6;
  a3.right = &a7;
  a4.left = &a8;
  a4.right = &a9;
  a5.left = &a10;

  Solution sol{};
  const auto actual = sol.zigzagLevelOrder(&a1);
  const std::vector<std::vector<int>> expected = {
      {3}, {20, 9}, {12, 16, 15, 7}, {89, 77, 10}};
  ASSERT_THAT(actual, ContainerEq(expected));
}

TEST(LibTest, TestCase2) {
  auto a1 = TreeNode{3};
  auto a2 = TreeNode{20};
  auto a3 = TreeNode{15};
  auto a4 = TreeNode{7};

  a1.right = &a2;
  a2.left = &a3;
  a2.right = &a4;
  Solution sol{};
  const auto actual = sol.zigzagLevelOrder(&a1);
  const std::vector<std::vector<int>> expected = {{3}, {20}, {15, 7}};
  ASSERT_THAT(actual, ContainerEq(expected));
}

TEST(LibTest, TestCase3) {
  auto a1 = TreeNode{3};
  auto a2 = TreeNode{20};
  auto a3 = TreeNode{15};
  auto a4 = TreeNode{7};

  a1.left = &a2;
  a2.left = &a3;
  a3.left = &a4;
  Solution sol{};
  const auto actual = sol.zigzagLevelOrder(&a1);
  const std::vector<std::vector<int>> expected = {{3}, {20}, {15}, {7}};
  ASSERT_THAT(actual, ContainerEq(expected));
}
TEST(LibTest, TestCase4) {
  Solution sol{};
  const auto actual = sol.zigzagLevelOrder(nullptr);
  const std::vector<std::vector<int>> expected = {};
  ASSERT_THAT(expected, ContainerEq(actual));
}
TEST(LibTest, TestCase5) {
  auto a1 = TreeNode{0};
  auto a2 = TreeNode{2};
  auto a3 = TreeNode{4};
  auto a4 = TreeNode{1};
  auto a5 = TreeNode{3};
  auto a6 = TreeNode{-1};
  auto a7 = TreeNode{5};
  auto a8 = TreeNode{1};
  auto a9 = TreeNode{6};
  auto a10 = TreeNode{8};
  a1.left = &a2;
  a1.right = &a3;
  a2.left = &a4;
  a3.left = &a5;
  a3.right = &a6;
  a4.left = &a7;
  a4.right = &a8;
  a5.right = &a9;
  a6.right = &a10;

  Solution sol{};
  const auto actual = sol.zigzagLevelOrder(&a1);
  const std::vector<std::vector<int>> expected = {
      {0}, {4, 2}, {1, 3, -1}, {8, 6, 1, 5}};
  ASSERT_THAT(expected, ContainerEq(actual));
}
