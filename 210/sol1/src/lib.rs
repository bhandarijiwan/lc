pub struct Solution;

use std::collections::LinkedList;
impl Solution {
    fn make_adj_list(n: i32, edges: &[Vec<i32>]) -> Vec<Vec<usize>> {
        let mut adj_list: Vec<Vec<usize>> = Vec::with_capacity(n as usize);
        for _ in 0..n {
            adj_list.push(vec![])
        }
        for edge in edges {
            let u = edge[0] as usize;
            let v = edge[1] as usize;
            adj_list[v].push(u)
        }
        adj_list
    }

    fn dfs(
        u: usize,
        adj_list: &[Vec<usize>],
        visited: &mut Vec<u8>,
        list: &mut LinkedList<usize>,
    ) -> bool {
        visited[u] = 1;
        let mut cycle = false;
        for &v in adj_list[u].iter() {
            if visited[v] == 1 {
                cycle = true;
                break;
            }
            if !cycle && visited[v] == 0 {
                cycle = Self::dfs(v, adj_list, visited, list);
                if cycle {
                    break;
                }
            }
        }
        visited[u] = 2;
        list.push_front(u);
        cycle
    }

    pub fn find_order(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> Vec<i32> {
        let adj_list = Self::make_adj_list(num_courses, &prerequisites);
        let mut list = LinkedList::new();
        let mut visited = vec![0 as u8; num_courses as usize];
        for i in 0..visited.len() {
            let u = visited[i];
            if u == 0 {
                if Solution::dfs(i, &adj_list, &mut visited, &mut list) {
                    return vec![];
                }
            }
        }
        list.iter().map(|v| *v as i32).collect::<Vec<_>>()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case1() {
        let num_courses = 2;
        let prerequisites = vec![vec![1, 0]];
        let actual = Solution::find_order(num_courses, prerequisites);
        let exepcted = vec![0, 1];
        assert_eq!(actual, exepcted);
    }
    #[test]
    fn test_case2() {
        let num_courses = 4;
        let prerequisites = vec![vec![1, 0], vec![2, 0], vec![3, 1], vec![3, 2]];
        let actual = Solution::find_order(num_courses, prerequisites);
        let expected = vec![0, 2, 1, 3];
        assert_eq!(actual, expected);
    }
    #[test]
    fn test_case3() {
        let num_courses = 1;
        let prerequisites = vec![];
        let actual = Solution::find_order(num_courses, prerequisites);
        let expected = vec![0];
        assert_eq!(actual, expected);
    }
    #[test]
    fn test_case4() {
        let num_courses = 2;
        let prerequisites = vec![vec![0, 1]];
        let actual = Solution::find_order(num_courses, prerequisites);
        let expected = vec![1, 0];
        assert_eq!(actual, expected);
    }
    #[test]
    fn test_case5() {
        let num_courses = 2;
        let prerequisites = vec![vec![0, 1], vec![1, 0]];
        let actual = Solution::find_order(num_courses, prerequisites);
        let expected = vec![];
        assert_eq!(actual, expected);
    }
}
