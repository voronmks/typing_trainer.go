

### 1. `typing_trainer.py` (Python)

```python
# typing_trainer.py — Python версия

import random
import time
import sys
from colorama import init, Fore, Style
try:
    import keyboard  # для чтения клавиш без Enter
except ImportError:
    keyboard = None

init(autoreset=True)

def generate_sequence(length=20, use_punctuation=False):
    chars = '0123456789'
    if use_punctuation:
        chars += '!@#$%^&*()-_=+[]{};:,.<>/?'
    return ''.join(random.choice(chars) for _ in range(length))

def print_highlight(target, user_input):
    """Печатает строку, выделяя правильные/неправильные символы."""
    result = []
    for i, (t, u) in enumerate(zip(target, user_input)):
        if u == t:
            result.append(Fore.GREEN + u + Style.RESET_ALL)
        else:
            result.append(Fore.RED + u + Style.RESET_ALL)
    # Дописываем остаток target, если ввод короче
    if len(user_input) < len(target):
        for ch in target[len(user_input):]:
            result.append(Fore.YELLOW + ch + Style.RESET_ALL)
    print(' '.join(result))

def main():
    print(f"{Fore.CYAN}⌨️  Тренажёр печати (цифры) (Python)")
    length = input("Длина упражнения (по умолч. 20): ").strip()
    length = int(length) if length else 20

    print("Нажмите Enter, когда будете готовы...")
    input()

    target = generate_sequence(length)
    print(f"{Fore.YELLOW}Введите:{Style.RESET_ALL} " + ' '.join(target))

    start = time.time()
    user_input = []

    # Используем keyboard для посимвольного ввода (если доступно)
    if keyboard:
        print("Вводите цифры (нажмите Esc для выхода):")
        while len(user_input) < len(target):
            event = keyboard.read_event(suppress=True)
            if event.event_type == 'down':
                if event.name == 'esc':
                    print("\nПрервано.")
                    return
                if event.name.isdigit():
                    user_input.append(event.name)
                    sys.stdout.write('\r' + ' ' * 80 + '\r')
                    print_highlight(target, ''.join(user_input))
                    sys.stdout.flush()
                elif event.name == 'backspace' and user_input:
                    user_input.pop()
                    sys.stdout.write('\r' + ' ' * 80 + '\r')
                    print_highlight(target, ''.join(user_input))
                    sys.stdout.flush()
    else:
        # fallback: построчный ввод
        print("Вводите цифры последовательно (Enter после каждой):")
        while len(user_input) < len(target):
            ch = input("> ").strip()
            if not ch:
                continue
            user_input.append(ch[0])
            sys.stdout.write('\r' + ' ' * 80 + '\r')
            print_highlight(target, ''.join(user_input))
            sys.stdout.flush()

    elapsed = time.time() - start
    user_str = ''.join(user_input)

    # Статистика
    correct = sum(1 for t, u in zip(target, user_str) if t == u)
    total = len(target)
    accuracy = (correct / total) * 100 if total > 0 else 0
    wpm = (total / (elapsed / 60)) if elapsed > 0 else 0

    print(f"\n{Fore.CYAN}Статистика:")
    print(f"  Время: {elapsed:.1f} сек")
    print(f"  Скорость: {wpm:.1f} зн/мин")
    print(f"  Точность: {accuracy:.1f}%")
    print(f"  Ошибок: {total - correct}")

    # Сохраняем результат в файл
    with open("typing_stats.txt", "a", encoding="utf-8") as f:
        f.write(f"Python\t{elapsed:.1f}\t{wpm:.1f}\t{accuracy:.1f}\t{total - correct}\n")

if __name__ == "__main__":
    main()
