package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	ids      = [5]int{1, 2, 3, 4, 5}
	titles   = [5]string{"сделать отчет", "посмотреть урок", "решить задачу", "проверить Pull Request", "прочитать главу книги"}
	statuses = [5]string{"новая", "в работе", "завершена", "новая", "в работе"}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
menu:
	for {
		fmt.Println("=== Менеджер задач ===")
		fmt.Println("1 - Показать все задачи")
		fmt.Println("2 - Фильтр по статусу")
		fmt.Println("3 - Выйти")

		fmt.Print("Выберите пункт: ")

		if !scanner.Scan() {
			break menu

		}

		command := scanner.Text()

		switch command {
		case "1":
			PrintAll()

		case "2":
			fmt.Print("Введите статус: ")

			if !scanner.Scan() {
				break menu
			}

			status := scanner.Text()

			if IsValidStatus(status) {
				PrintByStatus(status)
			} else {
				fmt.Println("Неверный статус. Допустимые: новая, в работе, завершена")
			}

		case "3":
			fmt.Println("Выход")
			break menu

		}

	}

	PrintAll()

}

func PrintAll() {
	for i := 0; i < len(ids); i++ {
		fmt.Printf("ID: %-4d | название: %-25s | статус: %s\n", ids[i], titles[i], statuses[i])

	}
}

func PrintByStatus(status string) {
	for i := 0; i < len(ids); i++ {
		if statuses[i] != status {
			continue
		}
		fmt.Printf("ID: %-4d | название: %-25s | статус: %s\n", ids[i], titles[i], statuses[i])
	}
}

func IsValidStatus(s string) bool {
	switch s {
	case "новая", "в работе", "завершена":
		return true
	}
	return false
}
