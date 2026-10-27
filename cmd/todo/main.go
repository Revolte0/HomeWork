package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	ids      = []int{1, 2, 3, 4, 5}
	titles   = []string{"сделать отчет", "посмотреть урок", "решить задачу", "проверить Pull Request", "прочитать главу книги"}
	statuses = []string{"новая", "в работе", "завершена", "новая", "в работе"}
)

var ErrEmptyTitle = errors.New("название не может быть пустым")
var ErrTitleTooLong = errors.New("название слишком длинное")
var ErrInvalidStatus = errors.New("недопустимый статус")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
menu:
	for {
		fmt.Println("=== Менеджер задач ===")
		fmt.Println("1 - Показать все задачи")
		fmt.Println("2 - Фильтр по статусу")
		fmt.Println("3 - Выйти")
		fmt.Println("4 - Добавить задачу")

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

		case "4":
			fmt.Print("Введите название: ")
			if !scanner.Scan() {
				break menu
			}
			title := scanner.Text()
			fmt.Print("Введите статус: ")
			if !scanner.Scan() {
				break menu
			}
			status := scanner.Text()

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


fmt.Print


func ValidateTask(title string, status string) error {
	if title == "" {
		return fmt.Errorf("Валидация не пройдена(%s, %s): %w", title, status, ErrEmptyTitle)
	}
	if len(title) > 100 {
		return fmt.Errorf("Слишком большая длина(%s): %w", title, ErrTitleTooLong)
	}
	if !IsValidStatus(status) {
		return fmt.Errorf("Невалидный статус(%s): %w", status, ErrInvalidStatus)
	}
	return nil

}

func AddTask(scanner *bufio.Scanner) {
	for {
		fmt.Print("Введите данные:")
		err := ValidateTask(title, status)
		if errors.Is(err, ErrEmptyTitle) {
			fmt.Println("Ошибка: название пустое. Введите заново")
			continue
		} 
		if errors.Is(err, ErrTitleTooLong) {
			fmt.Println("Ошибка: название слишком длинное")
			continue
			
		}
		if errors.Is(err, ErrInvalidStatus) {
			fmt.Println("Ошибка: недопустимый статус")
			
		continue
		}
		if errors.Is(err, ErrInvalidStatus) {
			fmt.Println("Ошибка: недопустимый статус")
			continue

		}
		if err != nil {
		continue
		} else if err == nil {
			
		}
	}
}
