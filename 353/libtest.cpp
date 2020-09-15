#include "lib.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"

TEST(LibTest, TestCase1) {
  const auto width = 3;
  const auto height = 2;
  std::vector<std::vector<int>> food{{1, 2}, {0, 1}};
  SnakeGame sg{width, height, food};
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(1, sg.move("R"));
  ASSERT_EQ(1, sg.move("U"));
  ASSERT_EQ(2, sg.move("L"));
  ASSERT_EQ(-1, sg.move("U"));
  ASSERT_EQ(-1, sg.move("D"));
}
TEST(LibTest, TestCase2) {
  const auto width = 2;
  const auto height = 2;
  std::vector<std::vector<int>> food{{1, 1}, {0, 0}};
  SnakeGame sg{width, height, food};
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(1, sg.move("D"));
  ASSERT_EQ(1, sg.move("L"));
  ASSERT_EQ(2, sg.move("U"));
  ASSERT_EQ(2, sg.move("R"));
}

TEST(LibTest, TestCase3) {
  const auto width = 3;
  const auto height = 3;
  std::vector<std::vector<int>> food{{2, 0}, {0, 0}, {0, 2}, {2, 2}};
  SnakeGame sg{width, height, food};
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(1, sg.move("D"));
  ASSERT_EQ(1, sg.move("R"));
  ASSERT_EQ(1, sg.move("U"));
  ASSERT_EQ(1, sg.move("U"));
  ASSERT_EQ(2, sg.move("L"));
  ASSERT_EQ(2, sg.move("D"));
  ASSERT_EQ(2, sg.move("R"));
  ASSERT_EQ(2, sg.move("R"));
  ASSERT_EQ(3, sg.move("U"));
  ASSERT_EQ(3, sg.move("L"));
  ASSERT_EQ(3, sg.move("D"));
}
TEST(LibTest, TestCase4) {
  const auto width = 3;
  const auto height = 3;
  std::vector<std::vector<int>> food{{2, 0}, {0, 0}, {0, 2},
                                     {0, 1}, {2, 2}, {0, 1}};
  SnakeGame sg{width, height, food};
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(1, sg.move("D"));
  ASSERT_EQ(1, sg.move("R"));
  ASSERT_EQ(1, sg.move("U"));
  ASSERT_EQ(1, sg.move("U"));
  ASSERT_EQ(2, sg.move("L"));
  ASSERT_EQ(2, sg.move("D"));
  ASSERT_EQ(2, sg.move("R"));
  ASSERT_EQ(2, sg.move("R"));
  ASSERT_EQ(3, sg.move("U"));
  ASSERT_EQ(4, sg.move("L"));
  ASSERT_EQ(4, sg.move("L"));
  ASSERT_EQ(4, sg.move("D"));
  ASSERT_EQ(4, sg.move("R"));
  ASSERT_EQ(-1, sg.move("U"));
}
TEST(LibTest, TestCase5) {
  const auto width = 1;
  const auto height = 1;
  std::vector<std::vector<int>> food{};
  SnakeGame sg{width, height, food};
  ASSERT_EQ(-1, sg.move("U"));
}
TEST(LibTest, TestCase6) {
  const auto width = 100;
  const auto height = 30;
  std::vector<std::vector<int>> food{{11, 0}, {58, 7}};
  SnakeGame sg{width, height, food};
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("R"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("U"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("L"));
  ASSERT_EQ(0, sg.move("D"));
  ASSERT_EQ(1, sg.move("L"));
  ASSERT_EQ(1, sg.move("D"));
  ASSERT_EQ(1, sg.move("R"));
  ASSERT_EQ(1, sg.move("U"));
  ASSERT_EQ(1, sg.move("L"));
  ASSERT_EQ(1, sg.move("U"));
  ASSERT_EQ(1, sg.move("R"));
  ASSERT_EQ(1, sg.move("R"));
  ASSERT_EQ(1, sg.move("U"));
  ASSERT_EQ(1, sg.move("L"));
  ASSERT_EQ(1, sg.move("D"));
  ASSERT_EQ(1, sg.move("L"));
  ASSERT_EQ(1, sg.move("D"));
  ASSERT_EQ(1, sg.move("D"));
}
