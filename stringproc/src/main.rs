use std::collections::HashMap;

fn main() {
    // let contents = std::fs::read_to_string("faust16.txt").unwrap();
    // eprintln!(" = {:?}", contents);
    let string = String::from("Hello");
    eprintln!(" = {:#?}", string.as_ptr());
    eprintln!(" = {:?}", string.len());
    eprintln!(" = {:?}", string.capacity());
}

fn process(string: &str) -> HashMap<&str, Vec<&str>> {
    let mut result: HashMap<&str, Vec<&str>> = HashMap::new();
    for line in string.lines() {
        for sentence in line.split_terminator(".") {
            for word in sentence.trim().split_whitespace() {
                if word.chars().next().unwrap().is_uppercase() {
                    result.entry(word).or_default().push(sentence);
                }
            }
        }
    }
    result
}
