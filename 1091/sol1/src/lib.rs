// https://leetcode.com/problems/shortest-path-in-binary-matrix/

pub struct Solution;

impl Solution {
    pub fn shortest_path_binary_matrix(grid: Vec<Vec<i32>>) -> i32 {
        let (m, n) = (grid.len(), grid[0].len());
        if grid[0][0] != 0 || grid[m - 1][n - 1] != 0 {
            return -1;
        }
        let mut visited: Vec<Vec<i32>> = Vec::with_capacity(m);
        for _ in 0..m {
            visited.push(vec![-1; m]);
        }
        let mut v = vec![0; 4];
        for i in (0..m).rev() {
            for j in (0..n).rev() {
                if grid[i][j] != 0 {
                    continue;
                }
                // right
                if j + 1 < n && visited[i][j + 1] > 0 {
                    v.push(visited[i][j + 1])
                }
                // down
                if i + 1 < m && visited[i + 1][j] > 0 {
                    v.push(visited[i + 1][j]);
                }
                // down left
                if i + 1 < m && j > 0 && visited[i + 1][j - 1] > 0 {
                    v.push(visited[i + 1][j - 1]);
                }
                // down right
                if i + 1 < m && j + 1 < n && visited[i + 1][j + 1] > 0 {
                    v.push(visited[i + 1][j + 1])
                }
                if v.len() > 0 {
                    visited[i][j] = v.iter().min().unwrap().to_owned() + 1;
                }
                v.clear()
            }
        }
        println!("visited = {:?}", visited);
        visited[0][0]
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case1() {
        let input = vec![vec![0, 1], vec![1, 0]];
        let actual = Solution::shortest_path_binary_matrix(input);
        assert_eq!(actual, 2)
    }
    #[test]
    fn test_case2() {
        let input = vec![vec![0, 0, 0], vec![1, 1, 0], vec![1, 1, 0]];
        let actual = Solution::shortest_path_binary_matrix(input);
        assert_eq!(actual, 4)
    }
    #[test]
    fn test_case3() {
        let input = vec![vec![0, 0, 0], vec![1, 1, 0], vec![1, 1, 1]];
        let actual = Solution::shortest_path_binary_matrix(input);
        assert_eq!(actual, -1)
    }
    #[test]
    fn test_case4() {
        let input = vec![
            vec![0, 0, 1, 0, 0, 0, 0],
            vec![0, 1, 0, 0, 0, 0, 1],
            vec![0, 0, 1, 0, 1, 0, 0],
            vec![0, 0, 0, 1, 1, 1, 0],
            vec![1, 0, 0, 1, 1, 0, 0],
            vec![1, 1, 1, 1, 1, 0, 1],
            vec![0, 0, 1, 0, 0, 0, 0],
        ];

        let actual = Solution::shortest_path_binary_matrix(input);
        assert_eq!(actual, 10)
    }

    #[test]
    fn test_case5() {
        let input = vec![vec![0, 0, 0], vec![1, 0, 1], vec![0, 0, 0]];

        let actual = Solution::shortest_path_binary_matrix(input);
        assert_eq!(actual, 3)
    }
    #[test]
    fn test_case6() {
        let input = vec![
            vec![0, 0, 0, 1],
            vec![0, 0, 1, 0],
            vec![0, 1, 0, 0],
            vec![1, 0, 0, 0],
        ];

        let actual = Solution::shortest_path_binary_matrix(input);
        assert_eq!(actual, 4)
    }
    #[test]
    fn test_case7() {
        let input = vec![
            vec![0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 0, 1, 0, 0, 0, 0, 1],
            vec![1, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 1, 1, 0],
            vec![0, 0, 1, 0, 1, 0, 1, 1],
            vec![0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 1, 1, 1, 0, 0],
            vec![1, 0, 1, 1, 1, 0, 0, 0],
        ];

        let actual = Solution::shortest_path_binary_matrix(input);
        assert_eq!(actual, 9)
    }
    #[test]
    fn test_case8() {
        let input = vec![vec![0, 0, 0], vec![1, 1, 1], vec![1, 0, 0]];

        let actual = Solution::shortest_path_binary_matrix(input);
        assert_eq!(actual, -1)
    }
    #[test]
    fn test_case9() {
        let input = vec![
            vec![0, 1, 0, 0, 0],
            vec![0, 1, 0, 1, 0],
            vec![0, 1, 0, 1, 0],
            vec![0, 1, 0, 1, 0],
            vec![0, 0, 0, 1, 0],
        ];
        let visited = [
            [-1, -1, 6, 5, 5],
            [-1, -1, -1, -1, 4],
            [-1, -1, -1, -1, 3],
            [-1, -1, -1, -1, 2],
            [-1, -1, -1, -1, 1],
        ];

        let actual = Solution::shortest_path_binary_matrix(input);
        assert_eq!(actual, 13)
    }
}
