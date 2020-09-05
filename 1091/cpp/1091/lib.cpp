
#include "lib.h"

#include <deque>

#include "fmt/core.h"
#include "fmt/ranges.h"

std::vector<std::pair<size_t, size_t>> get_next_edges(size_t i, size_t j,
                                                      size_t m, size_t n) {
  std::vector<std::pair<size_t, size_t>> edges;
  edges.reserve(8);
  // up
  if (i > 0) {
    edges.push_back(std::make_pair(i - 1, j));
  }
  // down
  if (i + 1 < m) {
    edges.push_back(std::make_pair(i + 1, j));
  }
  // left
  if (j > 0) {
    edges.push_back(std::make_pair(i, j - 1));
  }
  // right
  if (j + 1 < n) {
    edges.push_back(std::make_pair(i, j + 1));
  }
  // up left
  if (i > 0 && j > 0) {
    edges.push_back(std::make_pair(i - 1, j - 1));
  }
  // up right
  if (i > 0 && j + 1 < n) {
    edges.push_back(std::make_pair(i - 1, j + 1));
  }
  // down left
  if (i + 1 < m && j > 0) {
    edges.push_back(std::make_pair(i + 1, j - 1));
  }
  // down right
  if (i + 1 < m && j + 1 < n) {
    edges.push_back(std::make_pair(i + 1, j + 1));
  }
  return edges;
}

int Solution::shortestPathBinaryMatrix(std::vector<std::vector<int>> &grid) {
  const size_t m = grid.size();
  const size_t n = grid[0].size();
  if (m < 1 || n < 1 || grid[0][0] != 0) {
    return -1;
  }
  std::vector<std::vector<int>> visited(m, std::vector<int>(n, 0));
  visited[0][0] = 1;
  std::deque<std::pair<int, int>> queue{std::make_pair(0, 0)};
  while (!queue.empty()) {
    auto [x, y] = queue.front();
    queue.pop_front();
    auto edges = get_next_edges(x, y, m, n);
    for (auto e : edges) {
      auto [u, v] = e;
      if (grid[u][v] != 0) {
        visited[u][v] = -1;
      }
      if (visited[u][v] == 0) {
        visited[u][v] = visited[x][y] + 1;
        queue.emplace_back(e);
      }
    }
  }
  // fmt::print("{}\n", visited);
  if (visited[m - 1][n - 1] > 0) {
    return visited[m - 1][n - 1];
  }
  return -1;
}