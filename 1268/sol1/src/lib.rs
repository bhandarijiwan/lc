// https://leetcode.com/problems/search-suggestions-system/

pub struct Solution;

const SENTINEL: char = ('z' as u8 + 1 as u8) as char;

impl Solution {
    pub fn suggested_products(products: Vec<String>, search_word: String) -> Vec<Vec<String>> {
        let mut products = products;
        products.sort_unstable();
        let mut result = Vec::with_capacity(search_word.len());
        for i in 1..search_word.len() + 1 {
            let mut p = String::with_capacity(i + 1);
            p.push_str(&search_word[..i]);
            let start_index = match products.binary_search(&p) {
                Ok(n) => n,
                Err(n) => n,
            };
            p.push(SENTINEL);
            let mut end_index = match products.binary_search(&p) {
                Ok(n) => n,
                Err(n) => n,
            };
            if end_index > start_index {
                end_index = std::cmp::min(end_index, start_index + 3)
            }
            let r = &products[start_index..end_index];
            result.push(r.to_owned())
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case1() {
        let products = vec![
            String::from("mobile"),
            String::from("mouse"),
            String::from("moneypot"),
            String::from("monitor"),
            String::from("mousepad"),
        ];
        let search_word = String::from("mouse");
        let expected_output = vec![
            vec![
                String::from("mobile"),
                String::from("moneypot"),
                String::from("monitor"),
            ],
            vec![
                String::from("mobile"),
                String::from("moneypot"),
                String::from("monitor"),
            ],
            vec![String::from("mouse"), String::from("mousepad")],
            vec![String::from("mouse"), String::from("mousepad")],
            vec![String::from("mouse"), String::from("mousepad")],
        ];
        let actual_output = Solution::suggested_products(products, search_word);
        assert_eq!(expected_output, actual_output);
    }
    #[test]
    fn test_case2() {
        let products = vec![String::from("havana")];
        let search_word = String::from("havana");
        let expected_output = vec![
            vec![String::from("havana")],
            vec![String::from("havana")],
            vec![String::from("havana")],
            vec![String::from("havana")],
            vec![String::from("havana")],
            vec![String::from("havana")],
        ];
        let actual_output = Solution::suggested_products(products, search_word);
        assert_eq!(expected_output, actual_output);
    }
    #[test]
    fn test_case3() {
        let products = vec![
            String::from("bags"),
            String::from("baggage"),
            String::from("banner"),
            String::from("box"),
            String::from("cloths"),
        ];
        let search_word = String::from("bags");
        let expected_output = vec![
            vec![
                String::from("baggage"),
                String::from("bags"),
                String::from("banner"),
            ],
            vec![
                String::from("baggage"),
                String::from("bags"),
                String::from("banner"),
            ],
            vec![String::from("baggage"), String::from("bags")],
            vec![String::from("bags")],
        ];
        let actual_output = Solution::suggested_products(products, search_word);
        assert_eq!(expected_output, actual_output);
    }

    #[test]
    fn test_case4() {
        let products = vec![String::from("havana")];
        let search_word = String::from("tatiana");
        let expected_output: Vec<Vec<String>> =
            vec![vec![], vec![], vec![], vec![], vec![], vec![], vec![]];
        let actual_output = Solution::suggested_products(products, search_word);
        assert_eq!(expected_output, actual_output);
    }
}
