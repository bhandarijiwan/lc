type Bucket<K, V> = Vec<(K, V)>;
pub struct HashMap<K, V> {
    buckets: Vec<Bucket<K, V>>,
}

impl<K, V> HashMap<K, V> {
    pub fn iter(&self) -> Iter<K, V> {
        Iter {
            map: &self,
            entry_index: 0,
            bucket_index: 0,
        }
    }
}

pub struct Iter<'a, K, V> {
    map: &'a HashMap<K, V>,
    bucket_index: usize,
    entry_index: usize,
}

impl<'a, K, V> Iterator for Iter<'a, K, V> {
    type Item = (&'a K, &'a V);
    fn next(&mut self) -> Option<Self::Item> {
        while self.bucket_index < self.map.buckets.len() {
            if let Some(bucket) = self.map.buckets.get(self.bucket_index) {
                if self.entry_index < bucket.len() {
                    let (key, value) = &bucket[self.entry_index];
                    self.entry_index += 1;
                    return Some((key, value));
                }
                self.entry_index = 0
            }
            self.bucket_index += 1
        }
        return None;
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn it_works() {
        assert_eq!(true, true)
    }
}
