#ifndef LIB_HPP
#define LIB_HPP
#include <fmt/core.h>
#include <fmt/ranges.h>

#include <array>
#include <vector>

class Solution {

private:
  int square_num(int i, int j, int m, int n) {
    if (m % 2 == 0) {
      int k = (i | 1) != i ? j : (n - (j + 1));
      return (m - i) * m - k;
    } else {
      int k = (i | 1) == i ? j : (n - (j + 1));
      return (m - i) * m - k;
    }
  }
  void square_index(int num, int m, int n, std::array<int, 2> &indices) {
    indices[0] = (m - 1) - (num - 1) / m;
    indices[1] = ((m - indices[0]) * m) - num;
    indices[1] =
        (indices[0] | 1) != indices[0] ? indices[1] : (n - (indices[1] + 1));
    indices[1] = (m | 1) != m ? indices[1] : n - indices[1] - 1;
  }
  int min_unit(int num, int max_sqaure_num,
               std::vector<std::vector<int>> &visited, int m, int n) {
    std::array<int, 2> indices{0, 0};
    int min_steps = max_sqaure_num;
    int square_num = num;
    int counter = 1;
    fmt::print("checking for num {}\n", num);
    while (square_num < max_sqaure_num && counter <= 6) {
      square_num += 1;
      square_index(square_num, m, n, indices);
      auto [i, j] = indices;
      min_steps = std::min(min_steps, visited[i][j]);
      fmt::print(
          "square_num ={} i ={} , j= {}, visited={} min steps={} counter "
          "= {}\n ",
          square_num, i, j, visited[i][j], min_steps, counter);
      counter++;
    }
    return counter > 1 ? min_steps + 1 : 0;
  }

public:
  int snakesAndLadders(std::vector<std::vector<int>> &board) {
    const int m = board.size();
    const int n = board[0].size();
    std::vector<std::vector<int>> visited(m, std::vector<int>(n));
    const auto e = m % 2 == 0;
    const auto max_sqaure_num = m * n;
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        int k = j;
        if (e) {
          if (i % 2 != 0) {
            k = n - j - 1;
          }
        } else {
          if (i % 2 == 0) {
            k = n - j - 1;
          }
        }
        int num = square_num(i, k, m, n);
        if (board[i][k] == -1 || board[i][k] == num) {
          visited[i][k] = min_unit(num, max_sqaure_num, visited, m, n);
        } else if (board[i][k] < num) {
          visited[i][k] = max_sqaure_num + 1;
        } else {
          std::array<int, 2> indices{};
          square_index(board[i][k], m, n, indices);
          const auto target = board[indices[0]][indices[1]];
          // Start of another ladder
          if (target != -1) {
            fmt::print("checking for num {}\n", num);
            visited[i][k] =
                min_unit(board[i][k], max_sqaure_num, visited, m, n);
          } else {
            visited[i][k] = visited[indices[0]][indices[1]];
          }
        }
      }
    }
    fmt::print("{}\n", visited);
    return visited[m - 1][0] > 0 && visited[m - 1][0] < max_sqaure_num
               ? visited[m - 1][0]
               : -1;
  }
};
#endif /* LIB_HPP */