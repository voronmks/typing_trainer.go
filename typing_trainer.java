// typing_trainer.java — Java версия

import java.util.*;
import java.io.*;
import java.time.*;

public class typing_trainer {
    private static final Scanner scanner = new Scanner(System.in);

    public static String generateSequence(int length) {
        String chars = "0123456789";
        Random rand = new Random();
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < length; i++) {
            sb.append(chars.charAt(rand.nextInt(chars.length())));
        }
        return sb.toString();
    }

    public static void printHighlight(String target, String userInput) {
        System.out.print("\r\033[K");
        for (int i = 0; i < target.length(); i++) {
            String color = "\033[33m";
            if (i < userInput.length()) {
                char u = userInput.charAt(i);
                if (u == target.charAt(i)) {
                    color = "\033[32m";
                } else {
                    color = "\033[31m";
                }
            }
            System.out.print(color + target.charAt(i) + "\033[0m ");
        }
        System.out.println();
    }

    public static void main(String[] args) throws Exception {
        System.out.println("\033[36m⌨️  Тренажёр печати (цифры) (Java)\033[0m");
        System.out.print("Длина упражнения (по умолч. 20): ");
        String line = scanner.nextLine();
        int length = line.isEmpty() ? 20 : Integer.parseInt(line);

        System.out.println("Нажмите Enter, когда будете готовы...");
        scanner.nextLine();

        String target = generateSequence(length);
        System.out.println("\033[33mВведите:\033[0m " + String.join(" ", target.split("")));

        System.out.println("Вводите цифры последовательно (Enter после каждой):");
        String userInput = "";
        long start = System.currentTimeMillis();

        while (userInput.length() < target.length()) {
            System.out.print("> ");
            String ch = scanner.nextLine();
            if (ch.isEmpty()) continue;
            userInput += ch.charAt(0);
            printHighlight(target, userInput);
        }

        double elapsed = (System.currentTimeMillis() - start) / 1000.0;
        int correct = 0;
        for (int i = 0; i < target.length() && i < userInput.length(); i++) {
            if (userInput.charAt(i) == target.charAt(i)) correct++;
        }
        double accuracy = (double) correct / target.length() * 100;
        double wpm = target.length() / (elapsed / 60);

        System.out.println("\n\033[36mСтатистика:\033[0m");
        System.out.printf("  Время: %.1f сек\n", elapsed);
        System.out.printf("  Скорость: %.1f зн/мин\n", wpm);
        System.out.printf("  Точность: %.1f%%\n", accuracy);
        System.out.printf("  Ошибок: %d\n", target.length() - correct);

        // Сохранение
        try (FileWriter fw = new FileWriter("typing_stats.txt", true)) {
            fw.write(String.format("Java\t%.1f\t%.1f\t%.1f\t%d\n", elapsed, wpm, accuracy, target.length() - correct));
        }
    }
}
