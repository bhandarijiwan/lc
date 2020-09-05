// https://leetcode.com/problems/partition-labels/

use std::default::Default;
use std::marker::Copy;

#[derive(Debug)]
pub struct CharMap<V> {
    chars: [V; 26],
}

impl<V: Default + Copy> CharMap<V> {
    pub fn new() -> Self {
        Self {
            chars: [V::default(); 26],
        }
    }
    pub fn new_with_default(v: V) -> Self {
        Self { chars: [v; 26] }
    }
    pub fn insert(&mut self, k: u8, v: V) {
        let i = (k - 'a' as u8) as usize;
        self.chars[i] = v;
    }
    pub fn get(&self, k: u8) -> V {
        let i = (k - 'a' as u8) as usize;
        self.chars[i]
    }
}

pub struct Solution {}

impl Solution {
    pub fn partition_labels(s: String) -> Vec<i32> {
        let mut left = CharMap::new_with_default(-1 as isize);
        let chars = s.as_bytes();
        let mut l_i = 0 as usize;
        for _ in 0..chars.len() {
            let c = chars[l_i];
            left.insert(c, l_i as isize);
            l_i += 1;
        }
        let mut res = Vec::new();
        let mut start = 0 as usize;
        let mut m: usize = 0;
        for i in 0..chars.len() {
            m = std::cmp::max(m, left.get(chars[i]) as usize);
            if m == i {
                res.push((i - start + 1) as i32);
                start = i + 1
            }
        }
        res
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn it_works() {
        assert_eq!(
            Solution::partition_labels(String::from("ababcbacadefegdehijhklij")),
            vec![9, 7, 8]
        )
    }
}
