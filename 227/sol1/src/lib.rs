pub struct Solution;

use std::cell::RefCell;

const PRIORITIES: [i32; 6] = [2, 1, -1, 1, -1, 2];

impl Solution {
    fn add(a: i32, b: i32) -> i32 {
        return a + b;
    }
    fn sub(a: i32, b: i32) -> i32 {
        return a - b;
    }
    fn div(a: i32, b: i32) -> i32 {
        return a / b;
    }
    fn mul(a: i32, b: i32) -> i32 {
        return a * b;
    }

    fn priority(op: &str) -> i32 {
        let c = match op {
            "*" => '*',
            "-" => '-',
            "+" => '+',
            "/" => '/',
            _ => panic!("Unexpected operator"),
        };
        let i = (c as u8 - '*' as u8) as usize;
        return PRIORITIES[i];
    }

    pub fn token(s: &str) -> Vec<&str> {
        let mut tokens = vec![];
        let bytes = s.chars();
        let mut next_token_index = 0;
        for (i, b) in bytes.enumerate() {
            match b {
                ' ' => {
                    if i - next_token_index >= 1 {
                        tokens.push(&s[next_token_index..i]);
                    }
                    next_token_index = i + 1;
                }
                '+' | '-' | '*' | '/' => {
                    if i - next_token_index >= 1 {
                        tokens.push(&s[next_token_index..i]);
                    }
                    tokens.push(&s[i..i + 1]);
                    next_token_index = i + 1;
                }
                _ => continue,
            }
        }
        if s.len() - next_token_index > 0 {
            tokens.push(&s[next_token_index..s.len()]);
        }
        tokens
    }

    pub fn postfix(s: &str) -> Vec<&str> {
        let tokens = Self::token(s);
        let mut exp = vec![];
        let mut operators: Vec<&str> = vec![];
        for token in tokens {
            match token {
                "+" | "-" | "/" | "*" => {
                    let p1 = Self::priority(token);
                    while let Some(&t) = operators.last() {
                        let p2 = Self::priority(t);
                        if p1 > p2 {
                            break;
                        }
                        exp.push(operators.pop().unwrap())
                    }
                    operators.push(token)
                }
                _ => exp.push(token),
            }
        }
        while let Some(operator) = operators.pop() {
            exp.push(operator);
        }
        exp
    }
    pub fn eval(postfix_exp: &[&str]) -> i32 {
        let stack = RefCell::new(vec![]);
        let do_op = |op_func: fn(i32, i32) -> i32| {
            let mut stack = stack.borrow_mut();
            let a = stack.pop().unwrap();
            let b = stack.pop().unwrap();
            let c = op_func(b, a);
            stack.push(c);
        };
        for &token in postfix_exp {
            match token {
                "+" => do_op(Self::add),
                "-" => do_op(Self::sub),
                "*" => do_op(Self::mul),
                "/" => do_op(Self::div),
                _ => {
                    if let Ok(v) = token.parse::<i32>() {
                        stack.borrow_mut().push(v);
                    }
                }
            }
        }
        let r = stack.borrow_mut().pop().unwrap();
        r
    }

    pub fn calculate(s: String) -> i32 {
        let postfix_exp = Self::postfix(&s);
        Self::eval(&postfix_exp)
    }
}

#[cfg(test)]
mod tests {
    mod tokens {
        use crate::Solution;
        #[test]
        fn test_case1() {
            let exp = "3+2*2";
            assert_eq!(Solution::token(exp), ["3", "+", "2", "*", "2"]);
        }
        #[test]
        fn test_case2() {
            let exp = " 3/2 ";
            assert_eq!(Solution::token(exp), ["3", "/", "2"]);
        }
        #[test]
        fn test_case3() {
            let exp = " 3+5 / 2 ";
            assert_eq!(Solution::token(exp), ["3", "+", "5", "/", "2"]);
        }
        #[test]
        fn test_case4() {
            let exp = "   333    +    51   /      2 ";
            assert_eq!(Solution::token(exp), ["333", "+", "51", "/", "2"]);
        }
    }

    mod postfix {
        use crate::Solution;
        #[test]
        fn test_case1() {
            // let exp = "3+2/4*5+6-9";
            let exp = "A+B/C*D+E-F";
            let actual = Solution::postfix(exp);
            let expected = vec!["A", "B", "C", "/", "D", "*", "+", "E", "+", "F", "-"];
            assert_eq!(actual, expected);
        }
        #[test]
        fn test_case2() {
            let exp = " 333    +    51   /      2";
            let actual = Solution::postfix(exp);
            let expected = vec!["333", "51", "2", "/", "+"];
            assert_eq!(actual, expected);
        }
    }

    mod calculate {
        use crate::*;
        #[test]
        fn test_case1() {
            let exp = String::from("3+2*2");
            let actual = Solution::calculate(exp);
            assert_eq!(actual, 7);
        }
        #[test]
        fn test_case2() {
            let exp = String::from("3/2");
            let actual = Solution::calculate(exp);
            assert_eq!(actual, 1);
        }
        #[test]
        fn test_case3() {
            let exp = String::from("3+5/2");
            let actual = Solution::calculate(exp);
            assert_eq!(actual, 5);
        }
        #[test]
        fn test_case4() {
            let exp = String::from("3+5/2*0");
            let actual = Solution::calculate(exp);
            assert_eq!(actual, 3);
        }
        #[test]
        fn test_case5() {
            let exp = String::from("0");
            let actual = Solution::calculate(exp);
            assert_eq!(actual, 0);
        }
        #[test]
        fn test_case6() {
            let exp = String::from("0001 * 2000");
            let actual = Solution::calculate(exp);
            assert_eq!(actual, 2000);
        }
    }
}
