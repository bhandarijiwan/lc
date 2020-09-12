#include "lib.h"
#include <gmock/gmock.h>
#include <gtest/gtest.h>
#include <vector>

// TEST(LibTest, TestCase1) {
//   std::vector<std::vector<int>> input = {
//       {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1},
//       {-1, -1, -1, -1, -1, -1}, {-1, 35, -1, -1, 13, -1},
//       {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}};
//   const auto expected = 4;
//   Solution sol{};
//   const auto actual = sol.snakesAndLadders(input);
//   ASSERT_EQ(expected, actual);
// }
// TEST(LibTest, TestCase2) {
//   std::vector<std::vector<int>> input = {{-1, -1}, {-1, 3}};
//   const auto expected = 1;
//   Solution sol{};
//   const auto actual = sol.snakesAndLadders(input);
//   ASSERT_EQ(expected, actual);
// }
// TEST(LibTest, TestCase3) {
//   std::vector<std::vector<int>> input = {{-1, 4, -1}, {6, 2, 6}, {-1, 3,
//   -1}}; const auto expected = 2; Solution sol{}; const auto actual =
//   sol.snakesAndLadders(input); ASSERT_EQ(expected, actual);
// }

// TEST(LibTest, TestCase4) {
//   std::vector<std::vector<int>> input = {{1, 1, -1}, {1, 1, 1}, {-1, 1, 1}};
//   const auto expected = -1;
//   Solution sol{};
//   const auto actual = sol.snakesAndLadders(input);
//   ASSERT_EQ(expected, actual);
// }
// TEST(LibTest, TestCase5) {
//   std::vector<std::vector<int>> input = {
//       {-1, -1, 27, 13, -1, 25, -1}, {-1, -1, -1, -1, -1, -1, -1},
//       {44, -1, 8, -1, -1, 2, -1},   {-1, 30, -1, -1, -1, -1, -1},
//       {3, -1, 20, -1, 46, 6, -1},   {-1, -1, -1, -1, -1, -1, 29},
//       {-1, 29, 21, 33, -1, -1, -1}};
//   const auto expected = 4;
//   Solution sol{};
//   const auto actual = sol.snakesAndLadders(input);
//   ASSERT_EQ(expected, actual);
// }
// TEST(LibTest, TestCase6) {

//   std::vector<std::vector<int>> input = {{-1, -1, 2, 21, -1},
//                                          {16, -1, 24, -1, 4},
//                                          {2, 3, -1, -1, -1},
//                                          {-1, 11, 23, 18, -1},
//                                          {-1, -1, -1, 23, -1}};
//   const auto expected = 2;
//   Solution sol{};
//   const auto actual = sol.snakesAndLadders(input);
//   ASSERT_EQ(expected, actual);
// }
TEST(LibTest, TestCase7) {
  std::vector<std::vector<int>> input = {
      {-1, -1, -1, -1, 48, 5, -1}, {12, 29, 13, 9, -1, 2, 32},
      {-1, -1, 21, 7, -1, 12, 49}, {42, 37, 21, 40, -1, 22, 12},
      {42, -1, 2, -1, -1, -1, 6},  {39, -1, 35, -1, -1, 39, -1},
      {-1, 36, -1, -1, -1, -1, 5}};
  const auto expected = 3;
  Solution sol{};
  const auto actual = sol.snakesAndLadders(input);
  ASSERT_EQ(expected, actual);
}
