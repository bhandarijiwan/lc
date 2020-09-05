// https://leetcode.com/problems/alien-dictionary/

struct Solution;

use std::collections::LinkedList;

impl Solution {
    fn get_ordering_rules(a: &str, b: &str, v: &mut Vec<(u8, u8)>, byte_map: &mut Vec<u8>) -> bool {
        let bytes_a = a.bytes().collect::<Vec<_>>();
        let bytes_b = b.bytes().collect::<Vec<_>>();
        let mut found = false;
        let mut len_s = bytes_a.len();
        let l_s;
        if bytes_b.len() < bytes_a.len() {
            len_s = bytes_b.len();
            l_s = &bytes_a;
        } else {
            l_s = &bytes_b;
        }
        let index = |byte| (byte as usize - 'a' as usize);
        for i in 0..len_s {
            let byte_a_index = index(bytes_a[i]);
            let byte_b_index = index(bytes_b[i]);
            if byte_map[byte_a_index] == 0 {
                byte_map[byte_a_index] = 1
            }
            if byte_map[byte_b_index] == 0 {
                byte_map[byte_b_index] = 1;
            }
            if !found && byte_a_index != byte_b_index {
                v.push((byte_a_index as u8, byte_b_index as u8));
                byte_map[byte_a_index] = 2;
                byte_map[byte_b_index] = 2;
                found = true;
            }
        }
        for i in len_s..l_s.len() {
            let byte_index = index(l_s[i]);
            if byte_map[byte_index] == 0 {
                byte_map[byte_index] = 1
            }
        }
        if !found && a.len() > b.len() {
            return true;
        }
        false
    }

    fn make_adj_list(e: &Vec<(u8, u8)>) -> Vec<Vec<u8>> {
        let mut adj_list = vec![vec![]; 26];
        for t in e {
            adj_list[t.0 as usize].push(t.1);
        }
        adj_list
    }

    fn dfs(u: u8, adj_list: &[Vec<u8>], ll: &mut LinkedList<u8>, visited: &mut Vec<u8>) -> bool {
        visited[u as usize] = 1;
        for &w in adj_list[u as usize].iter() {
            if visited[w as usize] == 0 {
                let cycle = Self::dfs(w, adj_list, ll, visited);
                if cycle {
                    return cycle;
                }
            } else if visited[w as usize] == 1 {
                return true;
            }
        }
        visited[u as usize] = 2;
        ll.push_front(u);
        false
    }

    pub fn alien_order(words: Vec<String>) -> String {
        if words.len() <= 1 {
            let v = words.last().unwrap();
            return v.to_owned();
        }
        let mut byte_map = vec![0 as u8; 26];
        let words_refs = words.iter().collect::<Vec<&String>>();
        let windows = words_refs.windows(2);
        let mut rules = vec![];

        for win in windows {
            let prefix = Solution::get_ordering_rules(win[0], win[1], &mut rules, &mut byte_map);
            if prefix {
                return String::from("");
            }
        }
        let adj_list = Solution::make_adj_list(&rules);
        let mut visited: Vec<u8> = vec![0; adj_list.len()];
        let mut order_ll = LinkedList::new();

        for i in 0..adj_list.len() {
            if adj_list[i].len() > 0 && visited[i] == 0 {
                let cycle = Solution::dfs(i as u8, &adj_list, &mut order_ll, &mut visited);
                if cycle {
                    // eprintln!("Cycle Found ");
                    return String::from("");
                };
            }
        }
        let mut result = String::from("");
        for i in order_ll.into_iter() {
            result.push((i + 'a' as u8) as char)
        }
        for i in 0..byte_map.len() {
            if byte_map[i] == 1 {
                result.push((i as u8 + 'a' as u8) as char);
            }
        }
        result
    }
}

fn main() {
    let input = vec![
        "wrt".to_string(),
        "wrf".to_string(),
        "er".to_string(),
        "ett".to_string(),
        "rftt".to_string(),
    ];
    let output = Solution::alien_order(input);
    assert_eq!(output, "wertf");

    let input = vec!["xa".to_string(), "by".to_string()];
    let output = Solution::alien_order(input);
    assert_eq!(output, "xbay");

    let input = vec!["z".to_string(), "x".to_string(), "z".to_string()];
    let output = Solution::alien_order(input);
    assert_eq!(output, "");

    let input = vec!["hello".to_string(), "".to_string(), "".to_string()];
    let output = Solution::alien_order(input);
    assert_eq!(output, "");

    let input = vec!["abc".to_string(), "ab".to_string()];
    let output = Solution::alien_order(input);
    assert_eq!(output, "");

    let input = vec!["ab".to_string(), "abc".to_string(), "c".to_string()];
    let output = Solution::alien_order(input);
    assert_eq!(output, "acb");

    let input = vec!["wnlb".to_string()];
    let output = Solution::alien_order(input);
    assert_eq!(output, "wnlb")
}
