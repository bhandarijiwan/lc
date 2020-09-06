#include "lib.h"
#include "fmt/core.h"
#include "gtest/gtest.h"

TEST(LibTest, TestCase1) {
  const std::vector<std::vector<int>> input{{},   {100}, {80}, {60},
                                            {70}, {60},  {75}, {85}};
  const std::vector<int> expected{0, 1, 1, 1, 2, 1, 4, 6};
  StockSpanner ss{};
  for (int i = 1; i < input.size(); i++) {
    ASSERT_EQ(ss.next(input[i][0]), expected[i]);
  }
}

TEST(LibTest, TestCase2) {
  const std::vector<std::vector<int>> input{{},  {11}, {3}, {9},
                                            {5}, {6},  {4}, {7}};
  const std::vector<int> expected{0, 1, 1, 2, 1, 2, 1, 4};
  StockSpanner ss{};
  for (int i = 1; i < input.size(); i++) {
    ASSERT_EQ(ss.next(input[i][0]), expected[i]);
  }
}
