use futures::join;
use futures::Future;
use futures::{executor, StreamExt};
use std::{thread, time};

async fn print_temp() {
    let task1 = async {
        let ten_millis = time::Duration::from_millis(3000);
        thread::sleep(ten_millis);
        println!("{}", 11);
    };
    let task2 = async {
        println!("{}", 12);
    };

    let a = 10;
    println!("{}", a);

    println!("{}", a);
}

fn main() {
    println!("Hello, world!");
    executor::block_on(print_temp());
    println!("Hello, world!2");
}
