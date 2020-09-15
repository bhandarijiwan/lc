#ifndef LIB_HPP
#define LIB_HPP

#include "fmt/core.h"
#include "fmt/ranges.h"
#include <limits>
#include <string>
#include <vector>

class Solution {
public:
  int shortestWordDistance(std::vector<std::string> &words, std::string word1,
                           std::string word2) {
    int word1_last_index = std::numeric_limits<int>::max();
    int word2_last_index = std::numeric_limits<int>::max();
    const auto alternate_match = word1 == word2;
    bool skip1 = false;
    bool skip2 = alternate_match && true;
    const auto l = static_cast<int>(words.size());
    int length = std::numeric_limits<int>::max();
    // fmt::print("{} {} {}\n", skip1, skip2, alternate_match);
    for (int i = 0; i < l; i++) {
      if (words[i] == word1 && !skip1) {
        word1_last_index = i;
        skip1 = alternate_match;
        skip2 = alternate_match && !skip1;
        // fmt::print(" matches {} at index {}\n", word1, i);
      } else if (!skip2 && words[i] == word2) {
        word2_last_index = i;
        skip2 = alternate_match;
        skip1 = alternate_match && !skip2;
        // fmt::print(" matches {} at index {} {} {}\n", word2, i,
        //            word2_last_index, word1_last_index);
      }
      const auto r = std::abs(word1_last_index - word2_last_index);
      if (r > 0) {
        length = std::min(length, r);
      }
    }
    return length;
  }
};

#endif /* LIB_HPP */