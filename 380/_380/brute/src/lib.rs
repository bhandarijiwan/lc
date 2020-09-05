// https://leetcode.com/problems/insert-delete-getrandom-o1/

use rand::distributions::{Distribution, Uniform};
use std::collections::hash_map::DefaultHasher;
use std::hash::Hash;
use std::hash::Hasher; // 0.6.5

const INITIAL_SIZE: usize = 16;

type Bucket = Option<(i32, bool)>;
#[derive(Debug)]
pub struct RandomizedSet {
    items: Vec<Bucket>,
    items_count: usize,
    buckets_count: usize,
}

impl RandomizedSet {
    pub fn new() -> Self {
        Self {
            items: Self::create_buckets(INITIAL_SIZE),
            buckets_count: INITIAL_SIZE,
            items_count: 0,
        }
    }

    pub fn insert(&mut self, val: i32) -> bool {
        let load_factor = self.items_count as f64 / self.buckets_count as f64;
        if load_factor >= 0.9 {
            self.resize();
        }
        let item_index = self.item_index(val);
        // eprintln!("key = {:#?} item_index = {:#?}", val, item_index);
        let bucket = self.bucket_mut(item_index, val);
        // eprintln!("bucket = {:#?}", bucket);
        if let Some(b) = bucket {
            if let Some((_, del)) = b {
                if !*del {
                    return false;
                }
            }
            b.replace((val, false));
            self.items_count += 1;
        } else {
            self.items_count += 1;
            self.items.push(Some((val, false)));
        }
        true
    }

    fn resize(&mut self) {
        self.buckets_count *= 2;
        let new_items = Self::create_buckets(self.buckets_count);
        let old_items = std::mem::replace(&mut self.items, new_items);
        for old_item in old_items.into_iter() {
            if let Some((key, del)) = old_item {
                let bucket_index = self.item_index(key);
                let bucket = self.bucket_mut(bucket_index, key);
                if let Some(b) = bucket {
                    b.replace((key, del));
                } else {
                    self.items.push(Some((key, del)));
                }
            }
        }
    }

    pub fn remove(&mut self, val: i32) -> bool {
        let item_index = self.item_index(val);
        let bucket = self.bucket_mut(item_index, val);
        if let Some(b) = bucket {
            if let Some((_, del)) = b {
                if !*del {
                    *del = true;
                    self.items_count -= 1;
                    return true;
                }
            }
        }
        false
    }

    fn bucket_mut(&mut self, index: usize, key: i32) -> Option<&mut Bucket> {
        let mut first_del_i: isize = -1;
        for i in (index)..self.items.len() {
            if let Some((k, del)) = self.items[i] {
                if del {
                    first_del_i = i as isize;
                }
                if k == key {
                    // eprintln!("goes inside here = {:?}", key);
                    return self.items.get_mut(i);
                }
            } else {
                if first_del_i >= 0 {
                    return self.items.get_mut(first_del_i as usize);
                }
                // eprintln!("goes inside here = {:?}", key);
                return self.items.get_mut(i);
            }
        }
        // eprintln!("size of map = {:#?}", self.items.len());
        // eprintln!("goes inside here  = {:#?} index={:#?}", key, index);
        None
    }

    pub fn get_random(&self) -> i32 {
        let between = Uniform::from(0..self.buckets_count);
        let mut rng = rand::thread_rng();
        while let Some(b) = self.items.get(between.sample(&mut rng)) {
            if let Some((v, del)) = b {
                if !del {
                    return *v;
                }
            }
        }
        panic!()
    }

    fn item_index(&self, key: i32) -> usize {
        let mut hasher = DefaultHasher::new();
        key.hash(&mut hasher);
        let hash = hasher.finish();
        (hash % self.buckets_count as u64) as usize
    }
    fn create_buckets(cap: usize) -> Vec<Bucket> {
        let mut items = Vec::with_capacity(cap);
        for _ in 0..cap {
            items.push(None);
        }
        items
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn it_works() {
        let mut s = RandomizedSet::new();
        let v = s.insert(0);
        assert_eq!(v, true);
        for i in 1..32 {
            let v = s.insert(i);
            assert_eq!(v, true)
        }
        for i in 0..32 {
            let v = s.insert(i);
            assert_eq!(v, false)
        }

        for i in 0..10 {
            let v = s.remove(i);
            assert_eq!(v, true)
        }

        for i in 0..10 {
            let v = s.insert(i);
            assert_eq!(v, true)
        }

        let v = s.get_random();
        eprintln!("v = {:#?}", v);
    }

    #[test]
    fn test_case1() {
        let mut v = RandomizedSet::new();
        assert_eq!(v.insert(-1), true);
        assert_eq!(v.remove(-2), false);
        assert_eq!(v.insert(-2), true);

        let r = v.get_random();
        eprintln!("r = {:#?}", r);

        assert_eq!(v.remove(-1), true);
        assert_eq!(v.remove(-1), false);
        v.insert(-2);
        eprintln!("v = {:?}", v);
        assert_eq!(v.insert(-2), false);
        assert_eq!(v.get_random(), -2);
    }

    #[test]
    fn test_case2() {
        use std::fs::File;
        use std::io::{self, BufRead};
        let input_file = File::open("../data/input.txt").unwrap();
        let mut lines_iter = io::BufReader::new(input_file).lines().into_iter();
        let ins_lines = lines_iter.next().unwrap().unwrap();
        let val_lines = lines_iter.next().unwrap().unwrap();
        let output_lines = lines_iter.next().unwrap().unwrap();
        let ins_array: Vec<&str> = ins_lines.split(",").collect();
        let val_arry: Vec<&str> = val_lines.split(",").collect();
        let output_arr: Vec<&str> = output_lines.split(",").collect();

        let mut map = RandomizedSet::new();
        for i in 0..ins_array.len() {
            match ins_array[i] {
                "\"insert\"" => {
                    let v = val_arry[i];
                    let expected_str = output_arr[i];
                    let mut expected = false;
                    if expected_str == "true" {
                        expected = true;
                    }
                    let k = &v[1..v.len() - 1].parse::<i32>().unwrap();
                    let actual = map.insert(*k);
                    assert_eq!(actual, expected);
                    // eprintln!(
                    //     "inserting = {:#?} actual = {:#?}, expected={:#?}",
                    //     k, actual, expected
                    // );
                }
                "\"remove\"" => {
                    let v = val_arry[i];
                    let expected_str = output_arr[i];
                    let mut expected = false;
                    if expected_str == "true" {
                        expected = true;
                    }
                    let k = &v[1..v.len() - 1].parse::<i32>().unwrap();
                    let actual = map.remove(*k);
                    assert_eq!(actual, expected);

                    // eprintln!(
                    //     "removing = {:#?} actual = {:#?}, expected={:#?}",
                    //     k, actual, expected
                    // );
                    // if actual != expected {
                    //     eprintln!("map = {:?}", map);
                    //     panic!()
                    // }
                }
                "\"getRandom\"" => {
                    // eprintln!("Random = {:#?}", map.get_random());
                }
                _ => (),
            }
        }
    }
}
