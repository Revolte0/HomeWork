package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"todo/internal/task"
)

var tasks []*task.Task

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t1, _ := task.NewTask("Купить хлеб", "Описание")
	t2, _ := task.NewTask("Сделать ДЗ", "Описание ДЗ")
	tasks = append(tasks, t1, t2)

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
			PrintByStatus(status)

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

func AddTask(scanner *bufio.Scanner) {
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

		tasks = append(tasks, newTask)
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

	if newStatus != "new" && newStatus != "in_progress" && newStatus != "done" {
		fmt.Println("Ошибка: недопустимый статус. Допустимые: new, in_progress, done")
		return
	}

	found := false

	for _, t := range tasks {
		if t.ID() == id {

			t.SetStatus(task.Status(newStatus))
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
	for _, t := range tasks {
		fmt.Println(t)
	}

}

func PrintByStatus(status string) {
	for _, t := range tasks {
		if string(t.Status()) != status {
			continue
		}
		fmt.Println(t)
	}
}
