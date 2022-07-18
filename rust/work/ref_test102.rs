fn main() {
    // 変数vを10にする
    let mut v = 10;

    // 関数を呼び出す
    set_value(&mut v);

    println!("v={}", v);
}

fn set_value(arg: &mut u32) {
    *arg = 100;
}
