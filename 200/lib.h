#ifndef LIB_HPP
#define LIB_HPP

#include "fmt/core.h"
#include "fmt/ranges.h"

#include <vector>

class Solution {
private:
  void check_islands(std::vector<std::vector<char>> &grid, int i, int j, int m,
                     int n) {
    if (i >= m || i < 0 || j >= n || j < 0) {
      return;
    }
    if (grid[i][j] == '1') {
      grid[i][j] = '2';
      check_islands(grid, i + 1, j, m, n);
      check_islands(grid, i - 1, j, m, n);
      check_islands(grid, i, j + 1, m, n);
      check_islands(grid, i, j - 1, m, n);
    }
  }

public:
  int numIslands(std::vector<std::vector<char>> &grid) {
    size_t count = 0;
    const auto m = grid.size();
    if (m < 1) {
      return 0;
    }
    const auto n = grid[0].size();
    if (n < 1) {
      return 0;
    }

    for (size_t i = 0; i < m; i++) {
      for (size_t j = 0; j < n; j++) {
        if (grid[i][j] == '1') {
          count++;
          check_islands(grid, i, j, m, n);
        }
      }
    }
    return count;
  }
};

#endif /* LIB_HPP */
