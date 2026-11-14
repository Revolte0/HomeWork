package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"todo/internal/storage"
	"todo/internal/task"
)

var tasks []*task.Task

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t1, _ := task.NewTask("Купить хлеб", "Описание")
	t2, _ := task.NewTask("Сделать ДЗ", "Описание ДЗ")
	tasks = append(tasks, t1, t2)
	store := storage.NewMemoryStorage()

menu:
	for {
		fmt.Println("=== Менеджер задач ===")
		fmt.Println("1 - Показать все задачи")
		fmt.Println("2 - Фильтр по статусу")
		fmt.Println("3 - Выйти")
		fmt.Println("4 - Добавить задачу")
		fmt.Println("5 - Изменить статус задачи")
		fmt.Println("6 - Удалить задачу")
		fmt.Println("7 - Показать задачу по ID")

		fmt.Print("Выберите пункт: ")
		if !scanner.Scan() {
			break menu
		}
		command := scanner.Text()

		switch command {
		case "1":
			PrintAll(store)

		case "2":
			fmt.Print("Введите статус: ")
			if !scanner.Scan() {
				break menu
			}
			status := scanner.Text()
			PrintByStatus(store, status)

		case "3":
			fmt.Println("Выход")
			break menu

		case "4":
			AddTask(scanner, store)

		case "5":
			ChangeStatus(scanner, store)

		case "6":
			fmt.Print("Введите ID задачи для удаления: ")
			if !scanner.Scan() {
				break menu
			}
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Ошибка: ID должен быть числом")
				break
			}
			err = store.Delete(id)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Задача успешно удалена!")
			}
		case "7":
			fmt.Print("Введите ID задачи: ")
			if !scanner.Scan() {
				break menu
			}
			idStr := scanner.Text()

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Ошибка: ID должен быть числом")
				break
			}

			t, err := store.GetByID(id)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println(t)
			}

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func AddTask(scanner *bufio.Scanner, store *storage.MemoryStorage) {
	for {
		fmt.Print("Введите название: ")
		if !scanner.Scan() {
			return
		}
		title := scanner.Text()

		fmt.Print("Введите описание: ")
		if !scanner.Scan() {
			return
		}
		description := scanner.Text()

		newTask, err := task.NewTask(title, description)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		err = store.Add(newTask)
		if err != nil {
			fmt.Println("Ошибка при добавлении:", err)
			continue
		}
		fmt.Println("Задача успешно добавлена!")
		break
	}
}

func ChangeStatus(scanner *bufio.Scanner, store *storage.MemoryStorage) {
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

	if newStatus != "new" && newStatus != "in_progress" && newStatus != "done" {
		fmt.Println("Ошибка: недопустимый статус. Допустимые: new, in_progress, done")
		return
	}

	t, err := store.GetByID(id)
	if err != nil {
		fmt.Println("Задача не найдена")
		return
	}
	t.SetStatus(task.Status(newStatus))
	fmt.Println("Статус обновлён!")
}

func PrintAll(store *storage.MemoryStorage) {
	for _, t := range store.GetAll() {
		fmt.Println(t)
	}

}

func PrintByStatus(store *storage.MemoryStorage, status string) {
	filtered := store.FilterByStatus(task.Status(status))
	for _, t := range filtered {
		fmt.Println(t)
	}
}
