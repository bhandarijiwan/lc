// https:// leetcode.com/problems/online-stock-span/
#ifndef LIB_HPP
#define LIB_HPP
#include "fmt/core.h"
#include <vector>

class StockSpanner {
public:
  StockSpanner() {}
  int next(int price) {
    int c = 1;
    while (stocks.size() > 0 && price >= stocks.back().first) {
      c += stocks.back().second;
      stocks.pop_back();
    }
    stocks.push_back(std::make_pair(price, c));
    return c;
  }

private:
  std::vector<std::pair<int, int>> stocks;
};

#endif /* LIB_HPP */