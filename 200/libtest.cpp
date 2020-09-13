#include "lib.h"
#include "gtest/gtest.h"
#include <vector>

TEST(LibTest, TestCase1) {
  std::vector<std::vector<char>> input = {{'1', '1', '1', '1', '0'},
                                          {'1', '1', '0', '1', '0'},
                                          {'1', '1', '0', '0', '0'},
                                          {'0', '0', '0', '0', '0'}};
  const auto actual = (Solution{}).numIslands(input);
  const auto expected = 1;
  ASSERT_EQ(expected, actual);
}
TEST(LibTest, TestCase2) {

  std::vector<std::vector<char>> input = {{'1', '1', '0', '0', '0'},
                                          {'1', '1', '0', '0', '0'},
                                          {'0', '0', '1', '0', '0'},
                                          {'0', '0', '0', '1', '1'}};
  const auto actual = (Solution{}).numIslands(input);
  const auto expected = 3;
  ASSERT_EQ(expected, actual);
}
