// https://leetcode.com/problems/concatenated-words/
#ifndef LIB_HPP
#define LIB_HPP
#include "fmt/core.h"
#include "fmt/ranges.h"
#include <cstdint>
#include <unordered_set>
#include <vector>

class Trie {
public:
  explicit Trie(const std::vector<std::string> &words) { this->build(words); }
  bool exists(const std::string &w) const {
    TrieNode *next = root.get();
    for (const auto c : w) {
      const int i = index(c);
      if (next->nodes[i] == nullptr) {
        return false;
      }
      next = next->nodes[i].get();
    }
    return next->is_leaf;
  }
  // doesn't work for all cases
  bool consists_multiple_words(const std::string &w) const {
    TrieNode *rootnode = root.get();
    TrieNode *next = rootnode;
    size_t i = 0;
    uint32_t count = 0;
    while (i < w.size()) {
      const int idx = index(w[i]);
      if (next->nodes[idx] == nullptr) {
        return false;
      }
      next = next->nodes[idx].get();
      if (next->is_leaf) {
        auto inner_next = next;
        i++;
        const auto j = i;
        while (i < w.size()) {
          const auto j = index(w[i]);
          inner_next = inner_next->nodes[j].get();
          if (inner_next == nullptr || !inner_next->is_leaf) {
            break;
          }
          i++;
        }
        if (j < w.size()) {
          const std::string p = w.substr(j, i - j);
          if (p.size() > 0 && exists(p)) {
            count++;
          }
        }
        if (i < w.size()) {
          next = rootnode;
        }
        count++;
      } else {
        i++;
      }
    }
    return count > 1;
  }

private:
  struct TrieNode {
    TrieNode(bool leaf) : nodes(26), is_leaf(leaf) {}
    std::vector<std::unique_ptr<TrieNode>> nodes;
    bool is_leaf;
  };

  int index(char c) const { return (int)(c - 'a'); }

  void build(const std::vector<std::string> &words) {
    root = std::make_unique<TrieNode>(false);
    for (const auto &w : words) {
      TrieNode *next_node = root.get();
      for (const auto c : w) {
        const auto i = index(c);
        if (next_node->nodes[i] == nullptr) {
          next_node->nodes[i] = std::make_unique<TrieNode>(false);
        }
        next_node = next_node->nodes[i].get();
      }
      next_node->is_leaf = true;
    }
  }
  std::unique_ptr<TrieNode> root = nullptr;
};

class Solution {
private:
  bool contains_multiple(const std::string &s,
                         const std::unordered_set<std::string> &set) {
    if (s.size() < 2) {
      return false;
    }
    std::vector<int> m(s.size(), 0);
    m[0] = set.find(s.substr(0, 1)) != set.end() ? 1 : 0;
    for (size_t i = 1; i < s.size(); i++) {
      int result = set.find(s.substr(0, i + 1)) != set.end() ? 1 : 0;
      for (size_t j = i - 1; j < i; j--) {
        if (m[j] > 0) {
          if (set.find(s.substr(j + 1, i - j)) != set.end()) {
            result = m[j] + 1;
            break;
          }
        }
      }
      m[i] = result;
    }
    // fmt::print("m for {} = {}\n", s, m);
    return m.at(s.size() - 1) > 1;
  }

public:
  std::vector<std::string>
  findAllConcatenatedWordsInADict(std::vector<std::string> &words) {
    std::unordered_set<std::string> set{};
    for (const auto &w : words) {
      set.insert(w);
    }
    std::vector<std::string> result{};
    for (const auto &w : words) {
      if (contains_multiple(w, set)) {
        result.push_back(w);
      }
    }
    return result;
  }
};

#endif /* LIB_HPP */