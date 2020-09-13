#ifndef LIB_HPP
#define LIB_HPP

#include "fmt/core.h"
#include "fmt/ranges.h"

#include <sstream>
#include <string>
#include <unordered_set>
#include <vector>

class Solution {
private:
  void check_islands(std::vector<std::vector<int>> &grid, int i, int j, int m,
                     int n, std::ostringstream &oss, char c) {
    if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1) {
      return;
    }
    grid[i][j] = 2;
    oss << c;
    check_islands(grid, i - 1, j, m, n, oss, 'u');
    check_islands(grid, i + 1, j, m, n, oss, 'd');
    check_islands(grid, i, j - 1, m, n, oss, 'l');
    check_islands(grid, i, j + 1, m, n, oss, 'r');
    oss << 'b';
  }

public:
  int numDistinctIslands(std::vector<std::vector<int>> &grid) {
    std::unordered_set<std::string> path_set{};
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
        if (grid[i][j] == 1) {
          std::ostringstream oss{""};
          check_islands(grid, i, j, m, n, oss, 'o');
          path_set.insert(oss.str());
        }
      }
    }
    return path_set.size();
  }
};

#endif /* LIB_HPP */