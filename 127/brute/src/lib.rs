// https://leetcode.com/problems/word-ladder/
pub struct Solution;
use std::collections::LinkedList;

impl Solution {
    fn is_next(a: &str, b: &str) -> bool {
        let l = a.len();
        let a_bytes = a.as_bytes();
        let b_bytes = b.as_bytes();
        let mut d = 0;
        for i in 0..l {
            if a_bytes[i] != b_bytes[i] {
                d += 1
            }
            if d > 1 {
                return false;
            }
        }
        true
    }

    pub fn ladder_length(begin_word: String, end_word: String, word_list: Vec<String>) -> i32 {
        let mut visited = vec![false; word_list.len() + 1];
        let mut q = LinkedList::new();
        q.push_back((begin_word.as_str(), 1 as i32));
        visited[0] = true;
        while !q.is_empty() {
            let (str_a, level) = q.pop_front().unwrap();
            if str_a == &end_word {
                return level;
            }
            for i in 0..word_list.len() {
                if visited[i + 1] {
                    continue;
                }
                if Self::is_next(str_a, &word_list[i]) {
                    q.push_back((&word_list[i], level + 1));
                    visited[i + 1] = true;
                }
            }
        }
        0
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case1() {
        let begin_word = String::from("hit");
        let end_word = String::from("cog");
        let word_list = vec![
            String::from("hot"),
            String::from("dot"),
            String::from("dog"),
            String::from("lot"),
            String::from("log"),
            String::from("cog"),
        ];
        let res = Solution::ladder_length(begin_word, end_word, word_list);
        assert_eq!(res, 5)
    }

    #[test]
    fn test_case2() {
        let begin_word = String::from("hit");
        let end_word = String::from("cog");
        let word_list = vec![
            String::from("hot"),
            String::from("dot"),
            String::from("dog"),
            String::from("lot"),
            String::from("log"),
        ];
        let res = Solution::ladder_length(begin_word, end_word, word_list);
        assert_eq!(res, 0)
    }
    #[test]
    fn test_case3() {
        let begin_word = String::from("h");
        let end_word = String::from("c");
        let word_list = vec![String::from("c")];
        let res = Solution::ladder_length(begin_word, end_word, word_list);
        assert_eq!(res, 2)
    }

    #[test]
    fn test_case4() {
        let begin_word = String::from("hat");
        let end_word = String::from("ant");
        let word_list = vec![
            String::from("ant"),
            String::from("cat"),
            String::from("rat"),
            String::from("rnt"),
        ];
        let res = Solution::ladder_length(begin_word, end_word, word_list);
        assert_eq!(res, 4)
    }
    #[test]
    fn test_case5() {
        let begin_word = String::from("hit");
        let end_word: String = String::from("cog");
        let word_list = vec![
            String::from("hot"),
            String::from("dot"),
            String::from("dog"),
            String::from("lot"),
            String::from("log"),
        ];
        assert_eq!(Solution::ladder_length(begin_word, end_word, word_list), 0);
    }

    #[test]
    fn test_case6() {
        use std::fs::File;
        use std::io::{BufRead, BufReader};
        let input_file = File::open("./../data/input.txt").unwrap();
        let mut lines_iter = BufReader::new(input_file).lines().into_iter();
        let line1 = lines_iter.next().unwrap().unwrap();
        let line2 = lines_iter.next().unwrap().unwrap();
        let line3: String = lines_iter.next().unwrap().unwrap();
        let begin_word: String = line1.chars().skip(1).take(line1.len() - 2).collect();
        let end_word: String = line2.chars().skip(1).take(line2.len() - 2).collect();
        let word_list: Vec<String> = line3
            .split(",")
            .map(|s| s[1..s.len() - 1].to_owned())
            .collect();
        assert_eq!(Solution::ladder_length(begin_word, end_word, word_list), 20);
    }
}
