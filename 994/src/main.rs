// https://leetcode.com/problems/rotting-oranges/

pub struct Solution();

type Grid = Vec<Vec<i32>>;

impl Solution {
    fn has_neighbor_with(grid: &[Vec<i32>], i: usize, j: usize) -> bool {
        let m = grid.len();
        let n = grid[0].len();
        if j > 0 && grid[i][j - 1] == 2 {
            return true;
        }
        if j < n - 1 && grid[i][j + 1] == 2 {
            return true;
        }
        if i > 0 && grid[i - 1][j] == 2 {
            return true;
        }
        return i < m - 1 && grid[i + 1][j] == 2;
    }
    fn process_minute(grid: &[Vec<i32>]) -> Option<Grid> {
        let mut new_grid: Grid = grid.iter().cloned().collect();
        let m = new_grid.len();
        let n = new_grid[0].len();
        let mut changed = false;
        for i in 0..m {
            for j in 0..n {
                if new_grid[i][j] == 1 && Solution::has_neighbor_with(grid, i, j) {
                    changed = true;
                    new_grid[i][j] = 2
                }
            }
        }
        if changed {
            return Some(new_grid);
        }
        return None;
    }

    fn check_impossible(grid: &[Vec<i32>]) -> bool {
        let m = grid.len();
        let n = grid[0].len();
        for i in 0..m {
            for j in 0..n {
                if grid[i][j] == 1 {
                    return true;
                }
            }
        }
        false
    }

    pub fn oranges_rotting(grid: Grid) -> i32 {
        let mut c = 0;
        let mut g = grid;
        while let Some(grid) = Solution::process_minute(&g) {
            g = grid;
            c += 1
        }
        if Self::check_impossible(&g) {
            return -1;
        }
        c
    }
}

fn main() {
    let input: Grid = vec![vec![2, 1, 1], vec![1, 1, 0], vec![0, 1, 1]];
    // eprintln!("output = {:#?}", Solution::oranges_rotting(input));
    assert_eq!(Solution::oranges_rotting(input), 4);

    let input: Grid = vec![vec![2, 1, 1], vec![0, 1, 1], vec![1, 0, 1]];
    debug_assert_eq!(Solution::oranges_rotting(input), -1);

    let input: Grid = vec![vec![0, 2]];
    assert_eq!(Solution::oranges_rotting(input), 0);

    let input: Grid = vec![vec![0]];
    assert_eq!(Solution::oranges_rotting(input), 0);

    let input: Grid = vec![vec![0, 1], vec![2, 0]];
    assert_eq!(Solution::oranges_rotting(input), -1);

    let input: Grid = vec![vec![2], vec![2], vec![1], vec![0], vec![1], vec![1]];
    debug_assert_eq!(Solution::oranges_rotting(input), -1);

    let input: Vec<Vec<i32>> = vec![vec![1, 2]];
    debug_assert_eq!(Solution::oranges_rotting(input), 1);

    let input: Vec<Vec<i32>> = vec![vec![2, 1, 0, 2]];
    debug_assert_eq!(Solution::oranges_rotting(input), 1);
}
