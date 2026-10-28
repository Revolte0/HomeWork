package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ids      = []int{1, 2, 3, 4, 5}
	titles   = []string{"сделать отчет", "посмотреть урок", "решить задачу", "проверить Pull Request", "прочитать главу книги"}
	statuses = []string{"новая", "в работе", "завершена", "новая", "в работе"}
)

var (
	ErrEmptyTitle    = errors.New("название не может быть пустым")
	ErrTitleTooLong  = errors.New("название слишком длинное")
	ErrInvalidStatus = errors.New("недопустимый статус")
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

menu:
	for {
		fmt.Println("=== Менеджер задач ===")
		fmt.Println("1 - Показать все задачи")
		fmt.Println("2 - Фильтр по статусу")
		fmt.Println("3 - Выйти")
		fmt.Println("4 - Добавить задачу")
		fmt.Println("5 - Изменить статус задачи")

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
			AddTask(scanner)

		case "5":
			ChangeStatus(scanner)

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func ValidateTask(title, status string) error {
	if title == "" {
		return fmt.Errorf("валидация не пройдена: %w", ErrEmptyTitle)
	}
	if len(title) > 100 {
		return fmt.Errorf("валидация не пройдена: %w", ErrTitleTooLong)
	}
	if !IsValidStatus(status) {
		return fmt.Errorf("валидация не пройдена: %w", ErrInvalidStatus)
	}
	return nil
}

func AddTask(scanner *bufio.Scanner) {
	for {
		fmt.Print("Введите название: ")
		if !scanner.Scan() {
			return
		}
		title := scanner.Text()

		fmt.Print("Введите статус: ")
		if !scanner.Scan() {
			return
		}
		status := scanner.Text()

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
		if err != nil {
			fmt.Println("Неизвестная ошибка:", err)
			continue
		}

		ids = append(ids, len(ids)+1)
		titles = append(titles, title)
		statuses = append(statuses, status)
		fmt.Println("Задача успешно добавлена!")
		break
	}
}

func ChangeStatus(scanner *bufio.Scanner) {
	fmt.Print("Введите ID задачи: ")
	if !scanner.Scan() {
		return
	}
	idStr := scanner.Text()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Ошибка: ID должен быть числом")
		return
	}

	fmt.Print("Введите новый статус: ")
	if !scanner.Scan() {
		return
	}
	newStatus := scanner.Text()

	if !IsValidStatus(newStatus) {
		fmt.Println("Ошибка: недопустимый статус")
		return
	}

	found := false
	for i := 0; i < len(ids); i++ {
		if ids[i] == id {
			statuses[i] = newStatus
			fmt.Println("Статус задачи успешно обновлен!")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Задача с таким ID не найдена")
	}
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
