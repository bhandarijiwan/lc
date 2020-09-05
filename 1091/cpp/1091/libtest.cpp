#include "fmt/core.h"
#include "gtest/gtest.h"
#include "lib.h"

using namespace std;

TEST(LibTest, TestCase1) {
  Solution sol{};
  vector<vector<int>> input{{0, 1}, {1, 0}};
  const auto actual = sol.shortestPathBinaryMatrix(input);
  const auto expected = 2;
  ASSERT_EQ(expected, actual);
}

TEST(LibTest, TestCase2) {
  Solution sol{};
  vector<vector<int>> input{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}};
  const auto actual = sol.shortestPathBinaryMatrix(input);
  const auto expected = 4;
  ASSERT_EQ(expected, actual);
}
TEST(LibTest, TestCase3) {
  Solution sol{};
  vector<vector<int>> input{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}};
  const auto actual = sol.shortestPathBinaryMatrix(input);
  const auto expected = -1;
  ASSERT_EQ(expected, actual);
}
TEST(LibTest, TestCase4) {
  Solution sol{};
  vector<vector<int>> input{{0, 0, 1, 0, 0, 0, 0}, {0, 1, 0, 0, 0, 0, 1},
                            {0, 0, 1, 0, 1, 0, 0}, {0, 0, 0, 1, 1, 1, 0},
                            {1, 0, 0, 1, 1, 0, 0}, {1, 1, 1, 1, 1, 0, 1},
                            {0, 0, 1, 0, 0, 0, 0}};
  const auto actual = sol.shortestPathBinaryMatrix(input);
  const auto expected = 10;
  ASSERT_EQ(expected, actual);
}
TEST(LibTest, TestCase5) {
  Solution sol{};
  vector<vector<int>> input{{0, 0, 0}, {1, 0, 1}, {0, 0, 0}};
  const auto actual = sol.shortestPathBinaryMatrix(input);
  const auto expected = 3;
  ASSERT_EQ(expected, actual);
}
TEST(LibTest, TestCase6) {
  Solution sol{};
  vector<vector<int>> input{
      {0, 0, 0, 1},
      {0, 0, 1, 0},
      {0, 1, 0, 0},
      {1, 0, 0, 0},
  };
  const auto actual = sol.shortestPathBinaryMatrix(input);
  const auto expected = 4;
  ASSERT_EQ(expected, actual);
}

TEST(LibTest, TestCase7) {
  Solution sol{};
  vector<vector<int>> input{
      {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 1, 0, 0, 0, 0, 1},
      {1, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 1, 1, 0},
      {0, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 0, 0, 0, 0, 0, 0},
      {0, 0, 0, 1, 1, 1, 0, 0}, {1, 0, 1, 1, 1, 0, 0, 0},
  };
  const auto actual = sol.shortestPathBinaryMatrix(input);
  const auto expected = 9;
  ASSERT_EQ(expected, actual);
}
TEST(LibTest, TestCase8) {
  Solution sol{};
  vector<vector<int>> input{{0, 0, 0}, {1, 1, 1}, {1, 0, 0}};
  const auto actual = sol.shortestPathBinaryMatrix(input);
  const auto expected = -1;
  ASSERT_EQ(expected, actual);
}
TEST(LibTest, TestCase9) {
  Solution sol{};
  vector<vector<int>> input{
      {0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0},
      {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0},
  };
  const auto actual = sol.shortestPathBinaryMatrix(input);
  const auto expected = 13;
  ASSERT_EQ(expected, actual);
}
