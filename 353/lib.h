// https://leetcode.com/problems/design-snake-game/submissions/
#ifndef LIB_HPP
#define LIB_HPP

#include "fmt/core.h"
#include "fmt/ranges.h"

#include <list>
#include <stack>
#include <string>
#include <vector>

class SnakeGame {
public:
  SnakeGame(int width, int height, std::vector<std::vector<int>> &p_food)
      : board(height, std::vector<int>(width)), food(p_food), score(0),
        length(1), over(false) {

    if (food.size() > 0) {
      std::reverse(food.begin(), food.end());
      const auto v = food.back();
      food.pop_back();
      if (v[0] < static_cast<int>(board.size()) &&
          v[1] < static_cast<int>(board[v[0]].size())) {
        board[v[0]][v[1]] = 1;
      }
    }
    snake.push_back(std::make_pair(0, 0));
    board[0][0] = -1;
    // fmt::print("{}\n", board);
  }
  int move(std::string direction) {
    if (over) {
      return -1;
    }
    auto head = snake.back();

    // fmt::print("head ={}\n", head);
    if (direction == "U") {
      head.first -= 1;
    } else if (direction == "L") {
      head.second -= 1;
    } else if (direction == "R") {
      head.second += 1;
    } else if (direction == "D") {
      head.first += 1;
    }
    // fmt::print("next head ={}\n", head);
    const auto [i, j] = head;
    if (i < 0 || i >= static_cast<int>(board.size()) || j < 0 ||
        j >= static_cast<int>(board[i].size())) {
      // ran against the the wall
      //// fmt::print("ran into the wall\n");
      over = true;
      return -1;
    }
    if (board[i][j] == 1) {
      score += 1;
      length += 1;
      while (!food.empty()) {
        const auto v = food.back();
        food.pop_back();
        if (v[0] < static_cast<int>(board.size()) &&
            v[1] < static_cast<int>(board[v[0]].size())) {
          board[v[0]][v[1]] = 1;
          break;
        }
      }
    } else {
      const auto [tail_i, tail_j] = snake.front();
      board[tail_i][tail_j] = 0;
      snake.pop_front();
      if (board[i][j] == -1) {
        over = true;
        return -1;
      }
    }
    snake.push_back(head);
    board[i][j] = -1;
    // fmt::print("snake={}\n", snake);
    // fmt::print("board= {} \n", board);
    return score;
  }

private:
  std::vector<std::vector<int>> board;
  std::vector<std::vector<int>> food;
  std::list<std::pair<int, int>> snake;
  int score;
  int length;
  bool over;
};

#endif /* LIB_HPP */