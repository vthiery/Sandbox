use std::fmt;

#[derive(Debug)]
struct Complex{
    x: i32,
    y: i32,
}

impl fmt::Display for Complex {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self.y {
            0 => write!(f, "{}", self.x),
            1 => write!(f, "{} + i", self.x),
            _ => write!(f, "{} + {}i", self.x, self.y)
        }
    }
}

fn main() {
    println!("{}", Complex{ x: 3, y: 0 });
    println!("{}", Complex{ x: 3, y: 1 });
    println!("{}", Complex{ x: 3, y: 2 });
}
