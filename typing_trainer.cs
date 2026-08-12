// typing_trainer.cs — C# версия

using System;
using System.Collections.Generic;
using System.IO;
using System.Threading;

class TypingTrainer {
    static Random rand = new Random();

    static string GenerateSequence(int length) {
        string chars = "0123456789";
        char[] result = new char[length];
        for (int i = 0; i < length; i++) {
            result[i] = chars[rand.Next(chars.Length)];
        }
        return new string(result);
    }

    static void PrintHighlight(string target, string userInput) {
        Console.Write("\r\033[K");
        for (int i = 0; i < target.Length; i++) {
            string color = "\033[33m";
            if (i < userInput.Length) {
                if (userInput[i] == target[i]) {
                    color = "\033[32m";
                } else {
                    color = "\033[31m";
                }
            }
            Console.Write($"{color}{target[i]}\033[0m ");
        }
        Console.WriteLine();
    }

    static void Main() {
        Console.WriteLine("\033[36m⌨️  Тренажёр печати (цифры) (C#)\033[0m");
        Console.Write("Длина упражнения (по умолч. 20): ");
        string line = Console.ReadLine();
        int length = string.IsNullOrEmpty(line) ? 20 : int.Parse(line);

        Console.WriteLine("Нажмите Enter, когда будете готовы...");
        Console.ReadLine();

        string target = GenerateSequence(length);
        Console.WriteLine($"\033[33mВведите:\033[0m {string.Join(" ", target.ToCharArray())}");

        Console.WriteLine("Вводите цифры последовательно (Enter после каждой):");
        string userInput = "";
        DateTime start = DateTime.Now;

        while (userInput.Length < target.Length) {
            Console.Write("> ");
            string ch = Console.ReadLine();
            if (string.IsNullOrEmpty(ch)) continue;
            userInput += ch[0];
            PrintHighlight(target, userInput);
        }

        double elapsed = (DateTime.Now - start).TotalSeconds;
        int correct = 0;
        for (int i = 0; i < target.Length && i < userInput.Length; i++) {
            if (userInput[i] == target[i]) correct++;
        }
        double accuracy = (double)correct / target.Length * 100;
        double wpm = target.Length / (elapsed / 60);

        Console.WriteLine($"\n\033[36mСтатистика:\033[0m");
        Console.WriteLine($"  Время: {elapsed:F1} сек");
        Console.WriteLine($"  Скорость: {wpm:F1} зн/мин");
        Console.WriteLine($"  Точность: {accuracy:F1}%");
        Console.WriteLine($"  Ошибок: {target.Length - correct}");

        File.AppendAllText("typing_stats.txt", $"C#\t{elapsed:F1}\t{wpm:F1}\t{accuracy:F1}\t{target.Length - correct}\n");
    }
}
