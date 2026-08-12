// typing_trainer.js — JavaScript версия

const readline = require('readline');
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function generateSequence(length) {
    const chars = '0123456789';
    let result = '';
    for (let i = 0; i < length; i++) {
        result += chars[Math.floor(Math.random() * chars.length)];
    }
    return result;
}

function printHighlight(target, userInput) {
    process.stdout.write('\r\x1b[K');
    for (let i = 0; i < target.length; i++) {
        let color = '\x1b[33m';
        if (i < userInput.length) {
            if (userInput[i] === target[i]) {
                color = '\x1b[32m';
            } else {
                color = '\x1b[31m';
            }
        }
        process.stdout.write(`${color}${target[i]}\x1b[0m `);
    }
    process.stdout.write('\n');
}

function prompt(question) {
    return new Promise(resolve => rl.question(question, resolve));
}

async function main() {
    console.log('\x1b[36m⌨️  Тренажёр печати (цифры) (JavaScript)\x1b[0m');
    let length = await prompt('Длина упражнения (по умолч. 20): ');
    length = parseInt(length) || 20;

    console.log('Нажмите Enter, когда будете готовы...');
    await prompt('');

    const target = generateSequence(length);
    console.log(`\x1b[33mВведите:\x1b[0m ${target.split('').join(' ')}`);

    console.log('Вводите цифры (без Enter, используйте Ctrl+C для выхода).');
    // В Node.js посимвольно читать без Enter сложно, используем построчный ввод.
    console.log('Вводите цифры последовательно (Enter после каждой):');

    let userInput = '';
    const start = Date.now();

    while (userInput.length < target.length) {
        const ch = await prompt('> ');
        if (ch.length === 0) continue;
        userInput += ch[0];
        printHighlight(target, userInput);
    }

    const elapsed = (Date.now() - start) / 1000;
    let correct = 0;
    for (let i = 0; i < target.length && i < userInput.length; i++) {
        if (userInput[i] === target[i]) correct++;
    }
    const accuracy = (correct / target.length) * 100;
    const wpm = target.length / (elapsed / 60);

    console.log(`\n\x1b[36mСтатистика:\x1b[0m`);
    console.log(`  Время: ${elapsed.toFixed(1)} сек`);
    console.log(`  Скорость: ${wpm.toFixed(1)} зн/мин`);
    console.log(`  Точность: ${accuracy.toFixed(1)}%`);
    console.log(`  Ошибок: ${target.length - correct}`);

    // Сохранение
    const fs = require('fs');
    fs.appendFileSync('typing_stats.txt', `JavaScript\t${elapsed.toFixed(1)}\t${wpm.toFixed(1)}\t${accuracy.toFixed(1)}\t${target.length - correct}\n`);
    rl.close();
}

main().catch(console.error);
