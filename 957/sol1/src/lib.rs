pub struct Solution;

use std::collections::HashMap;
impl Solution {
    fn to_decimal(cells: &[i32]) -> i32 {
        let mut bitmap = 0x0;
        for cell in cells {
            bitmap <<= 1;
            bitmap = bitmap | cell;
        }
        bitmap
    }
    fn next_state(cells: &mut Vec<i32>) {
        let l = cells.len();
        let mut prev = cells[0];
        for i in 1..(l - 1) {
            let a = cells[i];
            cells[i] = 0;
            if prev == cells[i + 1] {
                cells[i] = 1
            }
            prev = a;
        }
        cells[0] = 0;
        cells[l - 1] = 0;
    }

    pub fn prison_after_n_days(cells: Vec<i32>, n: i32) -> Vec<i32> {
        let mut cells = cells;
        let mut map: HashMap<i32, i32> = HashMap::new();
        let mut repeat_start_index: i32 = -1;
        let mut step: i32 = 0;
        while step < n {
            let v = Self::to_decimal(&cells);
            if repeat_start_index < 0 && map.contains_key(&v) {
                repeat_start_index = map.get(&v).unwrap().clone();
                let c = step - repeat_start_index;
                step = n - ((n - step) % c);
                if step >= n {
                    break;
                }
            }
            if repeat_start_index < 0 {
                map.insert(v, step);
            }
            step = step + 1;
            Self::next_state(&mut cells);
        }
        cells
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case1() {
        let cells = vec![0, 1, 0, 1, 1, 0, 0, 1];
        let actual = Solution::prison_after_n_days(cells, 7);
        let expected = vec![0, 0, 1, 1, 0, 0, 0, 0];
        assert_eq!(expected, actual);
    }
    #[test]
    fn test_case2() {
        let cells = vec![1, 0, 0, 1, 0, 0, 1, 0];
        let actual = Solution::prison_after_n_days(cells, 1000000000);
        let expected = vec![0, 0, 1, 1, 1, 1, 1, 0];
        assert_eq!(expected, actual);
    }

    #[test]
    fn test_case3() {
        let cells = vec![1, 0, 0, 0, 1, 0, 0, 1];
        let actual = Solution::prison_after_n_days(cells, 99);
        let expected = vec![0, 0, 1, 0, 1, 0, 0, 0];
        assert_eq!(expected, actual);
    }
}
