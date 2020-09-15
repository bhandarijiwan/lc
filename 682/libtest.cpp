#include "lib.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"

#include "fmt/core.h"
#include "fmt/ranges.h"
#include <string>
#include <vector>

TEST(LibTest, TestCase1) {
  std::vector<std::string> input{"5", "2", "C", "D", "+"};
  const auto expected = 30;
  Solution sol{};
  ASSERT_EQ(expected, sol.calPoints(input));
}
TEST(LibTest, TestCase2) {
  std::vector<std::string> input{"5", "-2", "4", "C", "D", "9", "+", "+"};
  const auto expected = 27;
  Solution sol{};
  ASSERT_EQ(expected, sol.calPoints(input));
}