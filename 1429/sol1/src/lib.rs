use std::cmp::Ordering;
use std::collections::HashMap;
use std::collections::LinkedList;

#[derive(Debug, Eq, PartialEq)]
struct Item(usize, i32);

impl Ord for Item {
    fn cmp(&self, other: &Self) -> Ordering {
        if (*self).0 <= (*other).0 {
            return Ordering::Less;
        }
        Ordering::Greater
    }
}

impl PartialOrd for Item {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

#[derive(Debug)]
pub struct FirstUnique {
    pq: LinkedList<Item>,
    map: HashMap<i32, usize>,
}

impl FirstUnique {
    pub fn new(nums: Vec<i32>) -> Self {
        let mut map = HashMap::new();
        let mut pq = LinkedList::new();
        for &v in &nums {
            Self::add_impl(&mut pq, &mut map, v);
        }
        Self { map, pq }
    }

    pub fn show_first_unique(&self) -> i32 {
        if let Some(Item(_, v)) = self.pq.front() {
            return *v;
        }
        -1
    }
    fn add_impl(pq: &mut LinkedList<Item>, map: &mut HashMap<i32, usize>, v: i32) {
        let entry = map.entry(v).or_insert(0);
        *entry += 1;
        pq.push_back(Item(*entry, v));
        while let Some(Item(_, b)) = pq.front() {
            if *map.get(b).unwrap() == 1 {
                break;
            }
            pq.pop_front();
        }
    }

    pub fn add(&mut self, value: i32) {
        Self::add_impl(&mut self.pq, &mut self.map, value);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case1() {
        let v = vec![2, 3, 5];
        let mut f = FirstUnique::new(v);
        assert_eq!(f.show_first_unique(), 2);
        f.add(5);
        assert_eq!(f.show_first_unique(), 2);
        f.add(2);
        assert_eq!(f.show_first_unique(), 3);
        f.add(5);
        assert_eq!(f.show_first_unique(), 3);
        f.add(2);
        assert_eq!(f.show_first_unique(), 3);
        f.add(3);
        assert_eq!(f.show_first_unique(), -1);
    }

    #[test]
    fn test_case2() {
        let v = vec![7, 7, 7, 7, 7, 7];
        let mut f = FirstUnique::new(v);
        assert_eq!(f.show_first_unique(), -1);
        f.add(7);
        f.add(3);
        f.add(3);
        f.add(7);
        f.add(17);
        assert_eq!(f.show_first_unique(), 17);
    }

    #[test]
    fn test_case3() {
        let v = vec![809];
        let mut f = FirstUnique::new(v);
        assert_eq!(f.show_first_unique(), 809);
        f.add(809);
        assert_eq!(f.show_first_unique(), -1);
    }

    #[test]
    fn test_case4() {
        let v = vec![233];
        let mut f = FirstUnique::new(v);
        assert_eq!(f.show_first_unique(), 233);
        f.add(11);
        assert_eq!(f.show_first_unique(), 233);
    }

    #[test]
    fn test_case5() {
        // ["FirstUnique","showFirstUnique","add","showFirstUnique","add","showFirstUnique","add","showFirstUnique","add","add","add","add","showFirstUnique"]
        // [[[2,3,5]],[],[5],[],[2],[],[3],[],[6],[7],[8],[6],[]]
        let v = vec![2, 3, 5];
        let mut f = FirstUnique::new(v);
        assert_eq!(f.show_first_unique(), 2);
        f.add(5);
        assert_eq!(f.show_first_unique(), 2);
        f.add(2);
        assert_eq!(f.show_first_unique(), 3);
        f.add(3);
        assert_eq!(f.show_first_unique(), -1);
        f.add(6);
        f.add(7);
        f.add(8);
        f.add(6);
        assert_eq!(f.show_first_unique(), 7);
    }
}
