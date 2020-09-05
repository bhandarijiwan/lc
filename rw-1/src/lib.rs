use std::collections::hash_map::DefaultHasher;
use std::hash::{Hash, Hasher};

const INITIAL_BUCKET_COUNT: usize = 16;
const MAX_LOAD_FACTOR: f64 = 0.9;

type Slot<K, V> = Option<((K, V), usize)>; // tuple of ((key, value), usize)
#[derive(Debug)]
pub struct HashMap<K: Hash + Eq, V> {
    slots: Vec<Slot<K, V>>, // vector of tuples
    slot_count: usize,
    item_count: usize,
}

impl<K: Hash + Eq, V> HashMap<K, V> {
    pub fn new() -> Self {
        Self {
            slots: Self::create_slots(INITIAL_BUCKET_COUNT),
            slot_count: INITIAL_BUCKET_COUNT,
            item_count: 0,
        }
    }

    pub fn insert(&mut self, key: K, value: V) -> Option<V> {
        let load_factor = self.item_count as f64 / self.slot_count as f64;
        if load_factor >= MAX_LOAD_FACTOR {
            self.resize();
        }
        let new_slot_index = self.slot_index(&key);
        let slot = self.slot_mut(new_slot_index, &key);
        match slot {
            Some(slot) => {
                let old = slot.replace(((key, value), new_slot_index));
                match old {
                    Some(((_, v), _)) => Some(v),
                    None => {
                        self.item_count += 1;
                        None
                    }
                }
            }
            None => {
                self.slots.push(Some(((key, value), new_slot_index)));
                None
            }
        }
    }
    fn slot_mut(&mut self, slot_index: usize, key: &K) -> Option<&mut Slot<K, V>> {
        for i in (slot_index + 1)..self.slots.len() {
            if let Some(((t, _), _)) = &self.slots[i] {
                if t == key {
                    let k = self.slots.get_mut(i);
                    return k;
                }
            } else {
                return self.slots.get_mut(i);
            }
        }
        None
    }

    pub fn get(&self, key: &K) -> Option<&V> {
        let slot_index = self.slot_index(key);
        let ((_, v), _v) = self.slot(slot_index, key)?.as_ref()?;
        Some(v)
    }

    pub fn remove(&mut self, key: &K) -> Option<V> {
        let slot_index = self.slot_index(key);
        let slot = self.slot_mut(slot_index, key)?;
        let ((_, v), _) = slot.take()?;
        Some(v)
    }

    fn resize(&mut self) {
        let new_slot_count = self.slot_count * 2;
        let new_slots = Self::create_slots(new_slot_count);
        let old_slots = std::mem::replace(&mut self.slots, new_slots);
        for slot in old_slots.into_iter() {
            self.slots.push(slot);
        }
    }

    fn slot_index(&self, key: &K) -> usize {
        let mut hasher = DefaultHasher::new();
        key.hash(&mut hasher);
        let hash = hasher.finish();
        (hash % self.slot_count as u64) as usize
    }

    fn slot(&self, slot_index: usize, key: &K) -> Option<&Slot<K, V>> {
        for i in (slot_index + 1)..self.slots.len() {
            if let Some(((k, _), _)) = &self.slots[i] {
                if k == key {
                    return self.slots.get(i);
                }
            } else {
                return self.slots.get(i);
            }
        }
        None
    }

    fn create_slots(slot_count: usize) -> Vec<Slot<K, V>> {
        let mut new_slots = Vec::with_capacity(slot_count);
        for _ in 0..slot_count {
            new_slots.push(None)
        }
        new_slots
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn map_works() {
        let mut map = HashMap::new();
        let key1 = "foo";
        assert_eq!(map.insert(key1, "bar"), None);
        assert_eq!("bar", map.insert(key1, "baz").unwrap());
        assert_eq!(map.get(&key1), Some(&"baz"));
        assert_eq!(map.insert("key2", "value2"), None);
        assert_eq!(map.get(&"key2"), Some(&"value2"));

        assert_eq!(map.get(&"key3"), None);
        map.insert("key3", "value3");
        assert_eq!(map.get(&"key3"), Some(&"value3"));

        assert_eq!(map.remove(&"key3"), Some("value3"));

        assert_eq!(map.remove(&"key3"), None);

        // assert_eq!(map.get(&"foo"), Some(&"lol"));
        // assert_eq!(map.get(&"qux"), None);

        // assert_eq!(map.remove(&"foo"), Some("lol"));
        // assert_eq!(map.get(&"foo"), None);
    }
}
