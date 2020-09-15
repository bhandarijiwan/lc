#ifndef LIB_HPP
#define LIB_HPP

#include "fmt/core.h"
#include "fmt/ranges.h"

#include <algorithm>
#include <stack>
#include <string>
#include <vector>

class Solution {
public:
  int calPoints(std::vector<std::string> &ops) {
    std::stack<int> s{};
    int sum = 0;
    for (const auto &v : ops) {
      if (v == "D") {
        if (s.size() > 0) {
          const auto t = s.top() * 2;
          sum += t;
          s.push(t);
        }
      } else if (v == "C") {
        if (s.size() > 0) {
          sum -= s.top();
          s.pop();
        }
      } else if (v == "+") {
        const auto v = s.top();
        s.pop();
        const auto r = s.top() + v;
        sum += r;
        s.push(v);
        s.push(r);
      } else {
        const auto r = std::stoi(v);
        sum += r;
        s.push(r);
      }
      // fmt::print("{} {} \n", v, sum);
    }
    // fmt::print("stack= {}\n", s.size());
    return sum;
  }
};

#endif /* LIB_HPP */