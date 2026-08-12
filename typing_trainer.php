<?php
// typing_trainer.php — PHP версия

function generateSequence($length) {
    $chars = '0123456789';
    $result = '';
    for ($i = 0; $i < $length; $i++) {
        $result .= $chars[rand(0, strlen($chars)-1)];
    }
    return $result;
}

function printHighlight($target, $userInput) {
    echo "\r\033[K";
    for ($i = 0; $i < strlen($target); $i++) {
        $color = "\033[33m";
        if ($i < strlen($userInput)) {
            if ($userInput[$i] == $target[$i]) {
                $color = "\033[32m";
            } else {
                $color = "\033[31m";
            }
        }
        echo $color . $target[$i] . "\033[0m ";
    }
    echo "\n";
}

echo "\033[36m⌨️  Тренажёр печати (цифры) (PHP)\033[0m\n";
echo "Длина упражнения (по умолч. 20): ";
$lengthInput = trim(fgets(STDIN));
$length = empty($lengthInput) ? 20 : (int)$lengthInput;

echo "Нажмите Enter, когда будете готовы...\n";
fgets(STDIN);

$target = generateSequence($length);
echo "\033[33mВведите:\033[0m " . implode(' ', str_split($target)) . "\n";

echo "Вводите цифры последовательно (Enter после каждой):\n";
$userInput = '';
$start = microtime(true);

while (strlen($userInput) < strlen($target)) {
    echo "> ";
    $ch = trim(fgets(STDIN));
    if ($ch === '') continue;
    $userInput .= $ch[0];
    printHighlight($target, $userInput);
}

$elapsed = microtime(true) - $start;
$correct = 0;
for ($i = 0; $i < strlen($target) && $i < strlen($userInput); $i++) {
    if ($userInput[$i] == $target[$i]) $correct++;
}
$accuracy = ($correct / strlen($target)) * 100;
$wpm = strlen($target) / ($elapsed / 60);

echo "\n\033[36mСтатистика:\033[0m\n";
echo "  Время: " . number_format($elapsed, 1) . " сек\n";
echo "  Скорость: " . number_format($wpm, 1) . " зн/мин\n";
echo "  Точность: " . number_format($accuracy, 1) . "%\n";
echo "  Ошибок: " . (strlen($target) - $correct) . "\n";

file_put_contents('typing_stats.txt', 
    "PHP\t" . number_format($elapsed, 1) . "\t" . number_format($wpm, 1) . "\t" . number_format($accuracy, 1) . "\t" . (strlen($target) - $correct) . "\n", 
    FILE_APPEND);
?>
