// typing_trainer.rs — Rust версия

use rand::Rng;
use std::io::{self, Write, stdout};
use std::time::Instant;
use termion::color;
use termion::raw::IntoRawMode;

fn generate_sequence(length: usize, _use_punctuation: bool) -> String {
    let chars = b"0123456789";
    let mut rng = rand::thread_rng();
    (0..length).map(|_| chars[rng.gen_range(0..chars.len())] as char).collect()
}

fn print_highlight(target: &str, user_input: &str) {
    print!("\r\x1b[K"); // очистка строки
    for (i, ch) in target.chars().enumerate() {
        let color_code = if i < user_input.len() {
            let u = user_input.chars().nth(i).unwrap();
            if u == ch {
                color::Fg(color::Green)
            } else {
                color::Fg(color::Red)
            }
        } else {
            color::Fg(color::Yellow)
        };
        print!("{}{}{} ", color_code, ch, color::Fg(color::Reset));
    }
    println!();
    io::stdout().flush().unwrap();
}

fn main() -> io::Result<()> {
    let mut rng = rand::thread_rng();
    println!("\x1b[36m⌨️  Тренажёр печати (цифры) (Rust)\x1b[0m");
    print!("Длина упражнения (по умолч. 20): ");
    io::stdout().flush()?;
    let mut input = String::new();
    io::stdin().read_line(&mut input)?;
    let length: usize = input.trim().parse().unwrap_or(20);

    println!("Нажмите Enter, когда будете готовы...");
    io::stdin().read_line(&mut input)?;

    let target = generate_sequence(length, false);
    println!("\x1b[33mВведите:\x1b[0m {}", target.chars().collect::<Vec<_>>().join(" "));

    let stdout = stdout();
    let mut stdout = stdout.lock().into_raw_mode()?;

    // Для чтения посимвольно используем termion::async_stdin
    // Но для простоты используем построчный ввод
    // Однако мы можем попробовать читать посимвольно с помощью termion::input::TermRead
    // В этой версии используем построчный ввод (Enter).
    println!("Вводите цифры последовательно (Enter после каждой):");

    let mut user_input = String::new();
    let start = Instant::now();

    while user_input.len() < target.len() {
        print!("> ");
        io::stdout().flush()?;
        let mut ch = String::new();
        io::stdin().read_line(&mut ch)?;
        ch = ch.trim().to_string();
        if ch.is_empty() {
            continue;
        }
        user_input.push_str(&ch[0..1]);
        print_highlight(&target, &user_input);
        io::stdout().flush()?;
    }

    let elapsed = start.elapsed().as_secs_f64();
    let correct = target.chars().zip(user_input.chars()).filter(|(t, u)| t == u).count();
    let total = target.len();
    let accuracy = (correct as f64 / total as f64) * 100.0;
    let wpm = total as f64 / (elapsed / 60.0);

    println!("\n\x1b[36mСтатистика:\x1b[0m");
    println!("  Время: {:.1} сек", elapsed);
    println!("  Скорость: {:.1} зн/мин", wpm);
    println!("  Точность: {:.1}%", accuracy);
    println!("  Ошибок: {}", total - correct);

    // Сохранение
    let mut f = std::fs::OpenOptions::new()
        .append(true)
        .create(true)
        .open("typing_stats.txt")?;
    writeln!(f, "Rust\t{:.1}\t{:.1}\t{:.1}\t{}", elapsed, wpm, accuracy, total - correct)?;

    Ok(())
}
