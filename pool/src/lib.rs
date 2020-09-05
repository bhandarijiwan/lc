use std::cell::RefCell;
pub struct Pool<T> {
    items: RefCell<Vec<T>>,
}

impl<T> Pool<T>
where
    T: PoolItem,
{
    pub fn new() -> Self {
        Self {
            items: RefCell::new(Vec::new()),
        }
    }
    pub fn get(&self) -> PoolGuard<T> {
        let item = match self.items.borrow_mut().pop() {
            Some(item) => item,
            None => T::new(),
        };
        PoolGuard {
            inner: Some(item),
            items: &self.items,
        }
    }
}

pub trait PoolItem {
    fn new() -> Self;
    fn reset()
}

pub struct PoolGuard<'a, T> {
    inner: Option<T>,
    items: &'a RefCell<Vec<T>>,
}

impl<T> Drop for PoolGuard<'_, T> {
    fn drop(&mut self) {
        let item = self.inner.take().unwrap();
        // somehow return the item to the pool
        self.items.borrow_mut().push(item);
    }
}
impl<T> std::ops::Deref for PoolGuard<'_, T> {
    type Target = T;
    fn deref(&self) -> &Self::Target {
        self.inner.as_ref().unwrap()
    }
}

impl<T> std::ops::DerefMut for PoolGuard<'_, T> {
    fn deref_mut(&mut self) -> &mut Self::Target {
        self.inner.as_mut().unwrap()
    }
}

#[cfg(test)]
mod test {
    use super::*;
    #[test]
    fn it_works() {
        struct Awesome(usize);
        impl Awesome {
            fn do_thing(&self) {
                eprintln!("Yay");
            }
            pub fn inc(&mut self) {
                self.0 += 1;
            }
        }
        impl PoolItem for Awesome {
            fn new() -> Self {
                Self(0)
            }
        };
        let pool: Pool<Awesome> = Pool::new();
        let mut item1 = pool.get();
        item1.inc();
        assert_eq!(item1.0, 1);
        drop(item1);
        pool.get();
    }
}
