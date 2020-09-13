#include "lib.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"

#include <vector>

TEST(LibTest, TestCase1) {

  std::vector<std::vector<int>> input{
      {1, 1, 0, 0, 0},
      {1, 1, 0, 0, 0},
      {0, 0, 0, 1, 1},
      {0, 0, 0, 1, 1},
  };
  const auto expected = 1;
  const auto actual = Solution{}.numDistinctIslands(input);
  ASSERT_EQ(expected, actual);
}