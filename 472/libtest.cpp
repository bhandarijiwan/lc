#include "lib.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"
#include <fstream>
#include <sstream>
#include <string>

using namespace testing;
using namespace std;

TEST(LibTest, TestCase1) {
  vector<string> input{"cat", "cats",        "catsdogcats",
                       "dog", "dogcatsdog",  "hippopotamuses",
                       "rat", "ratcatdogcat"};
  const vector<string> expected{"catsdogcats", "dogcatsdog", "ratcatdogcat"};
  Solution sol{};
  const auto actual = sol.findAllConcatenatedWordsInADict(input);
  ASSERT_THAT(actual, ContainerEq(expected));
}
TEST(LibTest, TestCase2) {
  vector<string> input{"a", "b", "ab", "abc"};
  const vector<string> expected{"ab"};
  Solution sol{};
  const auto actual = sol.findAllConcatenatedWordsInADict(input);
  ASSERT_THAT(actual, ContainerEq(expected));
}
TEST(LibTest, TestCase3) {
  const auto read_file = [](const auto filename, auto &output) {
    ifstream file_stream(filename);
    ASSERT_TRUE(file_stream.is_open()) << "Failed to open " << filename << '\n';
    std::string line;
    while (std::getline(file_stream, line)) {
      istringstream iss(line.substr(1, line.size() - 2));
      std::string word;
      while (std::getline(iss, word, ',')) {
        output.emplace_back(word.substr(1, word.size() - 2));
      }
    }
    file_stream.close();
  };
  vector<string> input{};
  read_file("../input.txt", input);
  vector<string> expected{};
  read_file("../output.txt", expected);
  Solution sol{};
  const auto actual = sol.findAllConcatenatedWordsInADict(input);
  ASSERT_THAT(actual, ContainerEq(expected));
}
TEST(LibTest, TestCase4) {
  vector<string> input{"c", "w", "cw", "k", "u", "ku", "uz",          "y",
                       "a", "m", "e",  "n", "r", "p",  "cwkuzyamnerp"};
  vector<string> expected{"cw", "ku", "cwkuzyamnerp"};
  Solution sol{};
  const auto actual = sol.findAllConcatenatedWordsInADict(input);
  ASSERT_THAT(actual, ContainerEq(expected));
}
