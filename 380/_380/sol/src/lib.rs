use rand::distributions::{Distribution, Uniform};
use std::collections::HashMap;
#[derive(Debug)]
pub struct RandomizedSet {
    map: HashMap<i32, usize>,
    list: Vec<i32>,
}

impl RandomizedSet {
    /** Initialize your data structure here. */
    pub fn new() -> Self {
        Self {
            map: HashMap::new(),
            list: Vec::new(),
        }
    }

    /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
    pub fn insert(&mut self, val: i32) -> bool {
        match self.map.contains_key(&val) {
            true => false,
            false => {
                self.list.push(val);
                self.map.insert(val, self.list.len() - 1);
                true
            }
        }
    }

    /** Removes a value from the set. Returns true if the set contained the specified element. */
    pub fn remove(&mut self, val: i32) -> bool {
        match self.map.remove(&val) {
            Some(i1) => {
                let i2 = self.list.len() - 1;
                if i1 != i2 {
                    let v = self.list[i2];
                    self.list.swap(i1, i2);
                    self.map.insert(v, i1);
                }
                self.list.pop();
                true
            }
            None => false,
        }
    }

    /** Get a random element from the set. */
    pub fn get_random(&self) -> i32 {
        let between = Uniform::new(0, self.list.len());
        let mut rng = rand::thread_rng();
        let index = between.sample(&mut rng);
        self.list[index]
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
