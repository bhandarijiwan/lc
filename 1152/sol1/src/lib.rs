// https://leetcode.com/problems/analyze-user-website-visit-pattern/
// https://play.rust-lang.org/?version=stable&mode=debug&edition=2018&gist=0d39b681f8c4fd8a7391a5ebc968d332
pub struct Solution;
use std::cmp::Ordering;
use std::collections::HashMap;
use std::collections::HashSet;

impl Solution {
    pub fn permuate<'a>(
        v: &[&'a str],
        i: usize,
        j: usize,
        n: usize,
        o: &mut Vec<&'a str>,
        output: &mut HashSet<Vec<&'a str>>,
    ) {
        if o.len() >= n {
            output.insert(o.to_owned());
            return;
        }
        for k in i..j + 1 {
            o.push(v[k]);
            Self::permuate(v, k + 1, j, n, o, output);
            o.pop();
        }
    }

    pub fn most_visited_pattern(
        username: Vec<String>,
        timestamp: Vec<i32>,
        website: Vec<String>,
    ) -> Vec<String> {
        let mut seq = Vec::with_capacity(username.len());
        for i in 0..username.len() {
            seq.push((username[i].as_str(), timestamp[i], website[i].as_str()));
        }
        seq.sort_unstable_by(|(u1, t1, _), (u2, t2, _)| match u1.cmp(u2) {
            Ordering::Equal => t1.cmp(t2),
            order => order,
        });
        let mut result = vec![];
        let mut max_c = 0;
        let mut three_map: HashMap<(&str, &str, &str), i32> = HashMap::new();
        let mut three_sequences = |i: usize, j: usize| {
            let c = j - i + 1;
            if c < 3 {
                return;
            }
            let v: Vec<&str> = (&seq[i..j + 1]).iter().map(|(_, _, s)| *s).collect();
            let mut d = vec![];
            let mut output = HashSet::new();
            Self::permuate(&v, 0, v.len() - 1, 3, &mut d, &mut output);
            for item in output {
                let key = (item[0], item[1], item[2]);
                *three_map.entry(key).or_insert(0) += 1;
                if let Some(i) = three_map.get(&key) {
                    if *i > max_c {
                        max_c = *i;
                        result = vec![item[0], item[1], item[2]];
                    }
                }
            }
        };
        let mut i = 0;
        for k in 1..seq.len() {
            if (seq[k]).0 != (seq[k - 1]).0 {
                three_sequences(i, k - 1);
                i = k
            }
        }

        three_sequences(i, seq.len() - 1);
        let mut tr: Vec<_> = three_map.iter().collect();
        tr.sort_unstable_by(|(w1, c1), (w2, c2)| match c1.cmp(c2) {
            Ordering::Equal => w1.cmp(w2),
            Ordering::Greater => Ordering::Less,
            Ordering::Less => Ordering::Greater,
        });
        if tr.len() > 0 {
            let ((a, b, c), _) = tr[0];
            return vec![a.to_string(), b.to_string(), c.to_string()];
        }
        vec![]
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case1() {
        let username = vec![
            String::from("joe"),
            String::from("joe"),
            String::from("joe"),
            String::from("james"),
            String::from("james"),
            String::from("james"),
            String::from("james"),
            String::from("mary"),
            String::from("mary"),
            String::from("mary"),
        ];
        let timestamp = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
        let website = vec![
            String::from("home"),
            String::from("about"),
            String::from("career"),
            String::from("home"),
            String::from("cart"),
            String::from("maps"),
            String::from("home"),
            String::from("home"),
            String::from("about"),
            String::from("career"),
        ];
        let actual = Solution::most_visited_pattern(username, timestamp, website);
        let expected = vec![
            String::from("home"),
            String::from("about"),
            String::from("career"),
        ];
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_case2() {
        let username = vec![
            String::from("u1"),
            String::from("u1"),
            String::from("u1"),
            String::from("u2"),
            String::from("u2"),
            String::from("u2"),
        ];
        let timestamp = vec![1, 2, 3, 4, 5, 6];
        let website = vec![
            String::from("a"),
            String::from("b"),
            String::from("a"),
            String::from("a"),
            String::from("b"),
            String::from("c"),
        ];
        let actual = Solution::most_visited_pattern(username, timestamp, website);
        let exepcted = ["a", "b", "a"];
        assert_eq!(actual, exepcted);
    }
    #[test]
    fn test_case3() {
        let username = vec![
            String::from("u1"),
            String::from("u1"),
            String::from("u1"),
            String::from("u2"),
            String::from("u2"),
            String::from("u2"),
        ];
        let timestamp = vec![1, 2, 3, 4, 5, 6];
        let website = vec![
            String::from("a"),
            String::from("b"),
            String::from("c"),
            String::from("a"),
            String::from("b"),
            String::from("a"),
        ];
        let actual = Solution::most_visited_pattern(username, timestamp, website);
        let exepcted = ["a", "b", "a"];
        assert_eq!(actual, exepcted);
    }
    #[test]
    fn test_case4() {
        let username = vec![
            String::from("h"),
            String::from("eiy"),
            String::from("cq"),
            String::from("h"),
            String::from("cq"),
            String::from("txldsscx"),
            String::from("cq"),
            String::from("txldsscx"),
            String::from("h"),
            String::from("cq"),
            String::from("cq"),
        ];
        let timestamp = vec![
            527896567, 334462937, 517687281, 134127993, 859112386, 159548699, 51100299, 444082139,
            926837079, 317455832, 411747930,
        ];
        let website = vec![
            String::from("hibympufi"),
            String::from("hibympufi"),
            String::from("hibympufi"),
            String::from("hibympufi"),
            String::from("hibympufi"),
            String::from("hibympufi"),
            String::from("hibympufi"),
            String::from("hibympufi"),
            String::from("yljmntrclw"),
            String::from("hibympufi"),
            String::from("yljmntrclw"),
        ];
        let actual = Solution::most_visited_pattern(username, timestamp, website);
        let exepcted = vec![
            String::from("hibympufi"),
            String::from("hibympufi"),
            String::from("yljmntrclw"),
        ];
        assert_eq!(actual, exepcted);
    }

    #[test]
    fn test_permutate() {
        let v = vec!["home", "cart", "maps", "home"];
        let mut t = vec![];
        let mut output = HashSet::new();
        let expected_output: HashSet<Vec<&str>> = (vec![
            vec!["home", "cart", "maps"],
            vec!["home", "cart", "home"],
            vec!["home", "maps", "home"],
            vec!["cart", "maps", "home"],
        ])
        .into_iter()
        .collect();
        Solution::permuate(&v, 0, v.len() - 1, 3, &mut t, &mut output);
        assert_eq!(expected_output, output);
    }
    #[test]
    fn test_permutate_dup() {
        let v = vec!["a", "a", "a", "a"];
        let mut t = vec![];
        let mut output = HashSet::new();
        let expected_output = (vec![vec!["a", "a", "a"]])
            .into_iter()
            .collect::<HashSet<_>>();
        Solution::permuate(&v, 0, v.len() - 1, 3, &mut t, &mut output);
        assert_eq!(expected_output, output);
    }

    #[test]
    fn test_permutate_dup1() {
        let v = vec![
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "yljmntrclw",
            "hibympufi",
            "yljmntrclw",
        ];

        let mut t = vec![];
        let mut output = HashSet::new();
        let expected = (vec![
            vec!["hibympufi", "yljmntrclw", "hibympufi"],
            vec!["hibympufi", "hibympufi", "yljmntrclw"],
            vec!["hibympufi", "hibympufi", "hibympufi"],
            vec!["hibympufi", "yljmntrclw", "yljmntrclw"],
            vec!["yljmntrclw", "hibympufi", "yljmntrclw"],
        ])
        .into_iter()
        .collect::<HashSet<_>>();
        Solution::permuate(&v, 0, v.len() - 1, 3, &mut t, &mut output);
        assert_eq!(expected, output);
    }
}
