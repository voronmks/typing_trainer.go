// typing_trainer.go — Go версия

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateSequence(length int, usePunctuation bool) string {
	chars := "0123456789"
	if usePunctuation {
		chars += "!@#$%^&*()-_=+[]{};:,.<>/?`~"
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func printHighlight(target, userInput string) {
	fmt.Print("\r\033[K") // очистка строки
	for i := 0; i < len(target); i++ {
		var ch byte
		var color string
		if i < len(userInput) {
			ch = userInput[i]
			if ch == target[i] {
				color = "\033[32m"
			} else {
				color = "\033[31m"
			}
		} else {
			ch = target[i]
			color = "\033[33m"
		}
		fmt.Printf("%s%c\033[0m ", color, ch)
	}
	fmt.Print("\n")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\033[36m⌨️  Тренажёр печати (цифры) (Go)\033[0m")
	fmt.Print("Длина упражнения (по умолч. 20): ")
	reader := bufio.NewReader(os.Stdin)
	lengthStr, _ := reader.ReadString('\n')
	lengthStr = strings.TrimSpace(lengthStr)
	length := 20
	if lengthStr != "" {
		if l, err := strconv.Atoi(lengthStr); err == nil && l > 0 {
			length = l
		}
	}

	fmt.Println("Нажмите Enter, когда будете готовы...")
	reader.ReadString('\n')

	target := generateSequence(length, false)
	fmt.Printf("\033[33mВведите:\033[0m %s\n", strings.Join(strings.Split(target, ""), " "))

	fmt.Println("Вводите цифры (по одной, без Enter). Для выхода нажмите Ctrl+C.")
	userInput := ""

	// Для чтения посимвольно без Enter используем пакет "github.com/eiannone/keyboard"
	// Но для упрощения используем построчный ввод
	// Попробуем читать посимвольно с помощью bufio с буфером, но в Go сложно без внешних библиотек.
	// В этой версии используем построчный ввод (Enter).
	fmt.Println("Вводите цифры последовательно (Enter после каждой):")
	start := time.Now()
	for len(userInput) < len(target) {
		fmt.Print("> ")
		ch, _ := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)
		if len(ch) == 0 {
			continue
		}
		userInput += ch[:1]
		printHighlight(target, userInput)
	}
	elapsed := time.Since(start).Seconds()

	correct := 0
	for i := 0; i < len(target) && i < len(userInput); i++ {
		if userInput[i] == target[i] {
			correct++
		}
	}
	accuracy := float64(correct) / float64(len(target)) * 100
	wpm := float64(len(target)) / (elapsed / 60)

	fmt.Printf("\n\033[36mСтатистика:\033[0m\n")
	fmt.Printf("  Время: %.1f сек\n", elapsed)
	fmt.Printf("  Скорость: %.1f зн/мин\n", wpm)
	fmt.Printf("  Точность: %.1f%%\n", accuracy)
	fmt.Printf("  Ошибок: %d\n", len(target)-correct)

	// Сохранение
	f, _ := os.OpenFile("typing_stats.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	fmt.Fprintf(f, "Go\t%.1f\t%.1f\t%.1f\t%d\n", elapsed, wpm, accuracy, len(target)-correct)
}
