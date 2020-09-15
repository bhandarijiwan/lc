#include "lib.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"

#include "fmt/core.h"
#include "fmt/ranges.h"

#include <string>
#include <vector>

TEST(LibTest, TestCase1) {
  std::vector<std::string> input{"practice", "makes", "perfect", "coding",
                                 "makes"};
  const auto expected = 1;
  Solution sol{};
  ASSERT_EQ(expected, sol.shortestWordDistance(input, "makes", "coding"));
}
TEST(LibTest, TestCase2) {
  std::vector<std::string> input{"practice", "makes", "perfect", "coding",
                                 "makes"};
  const auto expected = 3;
  Solution sol{};
  ASSERT_EQ(expected, sol.shortestWordDistance(input, "makes", "makes"));
}
