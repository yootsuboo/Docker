use std::fs;

// fn main() {
//     let text = fs::read_to_string("hoge.txt").unwrap();
//     println!("{}", text);
// }

// エラー処理
// fn main() {
//     let text = fs::read_to_string("hoge.txt").unwrap_or("失敗した値".to_string());
//     println!("{}", text);
// }

// match文でのエラー処理
fn main() {
    let text = match fs::read_to_string("hoge.txt") {
        Ok(text) => text,
        // エラー処理の記載方法が不明
        Err(e) => {}
    };
    println!("{}", text);
}
