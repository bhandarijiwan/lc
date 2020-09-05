// https://leetcode.com/problems/critical-connections-in-a-network/
// https://stackoverflow.com/questions/11218746/bridges-in-a-connected-graph
pub struct Solution;

impl Solution {
    fn create_adj_matrix(n: i32, connections: &Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut adj_matrix = vec![vec![]; n as usize];
        for v in connections {
            adj_matrix[v[0] as usize].push(v[1]);
            adj_matrix[v[1] as usize].push(v[0]);
        }
        adj_matrix
    }

    fn dfs(
        count: &mut i32,
        u: i32, // node
        v: i32, // parent
        adj_list: &Vec<Vec<i32>>,
        pre: &mut Vec<i32>,
        low: &mut Vec<i32>,
        result: &mut Vec<Vec<i32>>,
    ) {
        let u_i = u as usize;
        *count = *count + 1;
        pre[u_i] = *count + 1;
        low[u_i] = pre[u_i];

        for &w in adj_list[u_i].iter() {
            let w_i = w as usize;
            if pre[w_i] == -1 {
                Solution::dfs(count, w, u, adj_list, pre, low, result);
                low[u_i] = std::cmp::min(low[u_i], low[w_i]);
                if low[w_i] == pre[w_i] {
                    result.push(vec![u, w]);
                }
            } else {
                if w != v {
                    low[u_i] = std::cmp::min(low[u_i], pre[w_i])
                }
            }
        }

        // let mut stack = vec![u];
        // pre[u as usize] = u;
        // low[u as usize] = pre[u as usize];
        // while !stack.is_empty() {
        //     let &top = stack.last().unwrap();
        //     let t_i = top as usize;
        //     match adj_list[top as usize].next() {
        //         Some(&v) => {
        //             let v_i = v as usize;
        //             if pre[v_i] == -1 {
        //                 pre[v_i] = top + 1;
        //                 low[v_i] = pre[v_i];
        //                 stack.push(v);
        //             } else {
        //                 // println!("Cycle found {} -> {} ,", t_i, v_i);
        //                 low[t_i] = std::cmp::min(low[t_i], pre[v_i]);
        //             }
        //         }
        //         None => {
        //             // println!("popping {} from stack", t_i);
        //             stack.pop();
        //             if !stack.is_empty() {
        //                 let &k = stack.last().unwrap();
        //                 let k_i = k as usize;
        //                 low[k_i] = std::cmp::min(low[k_i], low[t_i]);
        //                 if low[t_i] == pre[t_i] {
        //                     result.push(vec![k, top])
        //                 }
        //             }
        //         }
        //     }
        //     // println!("stack {:?}", stack);
        //     // println!("pre = {:?}", pre);
        //     // println!("low = {:?}", low);
        // }
    }

    pub fn critical_connections(n: i32, connections: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let adj_matrix = Solution::create_adj_matrix(n, &connections);
        // let mut adj_list = adj_matrix.iter().map(|v| v.iter()).collect::<Vec<_>>();
        let mut result = vec![];
        let mut pre = vec![-1; n as usize];
        let mut low = vec![-1; n as usize];
        let mut count = 0;
        for i in 0..n {
            if pre[i as usize] == -1 {
                Solution::dfs(
                    &mut count,
                    i,
                    i,
                    &adj_matrix,
                    &mut pre,
                    &mut low,
                    &mut result,
                );
            }
        }
        result
    }
}

fn main() {
    let mut input = vec![vec![0, 1], vec![1, 2], vec![2, 0], vec![1, 3]];
    let mut s = Solution::critical_connections(4, input);
    println!("Bridge {:?}", s);
    input = vec![vec![0, 1], vec![1, 2], vec![2, 3], vec![3, 1]];
    s = Solution::critical_connections(4, input);
    println!("Bridge {:?}", s);

    input = vec![
        vec![1, 0],
        vec![2, 0],
        vec![3, 2],
        vec![4, 2],
        vec![4, 3],
        vec![3, 0],
        vec![4, 0],
    ];
    s = Solution::critical_connections(5, input);
    println!("Bridge {:?}", s)
}
