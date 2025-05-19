package main

import (
	"RBAC/comands"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Printhelp() {
	fmt.Println("Доступные команды:")
	fmt.Println("  init                           - Инициализация системы")
	fmt.Println("  create_role <имя_роли>         - Создать новую роль")
	fmt.Println("  assign_role <субъект> <роль>   - Назначить роль пользователю")
	fmt.Println("  inherit_role <роль1> <роль2>   - Сделать роль1 наследником роли2")
	fmt.Println("  assign_permission <роль> <объект> <действие> - Добавить разрешение")
	fmt.Println("  merge_roles <новая_роль> <стратегия> <роли...> - Объединить роли")
	fmt.Println("  change_role <субъект> <новая_роль> - Изменить роль субъекта")
	fmt.Println("  show_permissions <роль>       - Показать все разрешения роли")
	fmt.Println("  exit                           - Выйти из программы")
}
func main() {
	fmt.Println("=== Система управления доступом на основе ролей (RBAC) ===")
	Printhelp()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]
		args = args[1:]

		switch command {
		case "help":
			Printhelp()
		case "init":
			commands.Init()
		case "create_role":
			if len(args) != 1 {
				fmt.Println("Ошибка: требуется 1 аргумент - имя роли")
				continue
			}
			commands.CreateRole(args[0])
		case "assign_role":
			if len(args) != 2 {
				fmt.Println("Ошибка: требуется 2 аргумента - ID пользователя и имя роли")
				continue
			}
			commands.AssignRole(args[0], args[1])
		case "inherit_role":
			if len(args) != 2 {
				fmt.Println("Ошибка: требуется 2 аргумента - дочерняя и родительская роли")
				continue
			}
			commands.InheritRole(args[0], args[1])
		case "assign_permission":
			if len(args) != 3 {
				fmt.Println("Ошибка: требуется 3 аргумента - роль, объект и действие")
				continue
			}
			commands.AssignPermission(args[0], args[1], args[2])
		case "merge_roles":
			if len(args) < 3 {
				fmt.Println("Ошибка: требуется минимум 3 аргумента - новая роль, стратегия и список ролей")
				continue
			}
			strategy := args[1]
			if strategy != "union" && strategy != "intersection" {
				fmt.Println("Ошибка: стратегия должна быть 'union' или 'intersection'")
				continue
			}
			commands.MergeRoles(args[0], args[2:], strategy)
		case "change_role":
			if len(args) != 2 {
				fmt.Println("Ошибка: требуется 2 аргумента - ID пользователя и новая роль")
				continue
			}
			commands.ChangeRole(args[0], args[1])
		case "show_permissions":
			if len(args) != 1 {
				fmt.Println("Ошибка: требуется 1 аргумент - имя роли")
				continue
			}
			commands.ShowPermissions(args[0])
		default:
			fmt.Println("Неизвестная команда. Введите 'help' для списка команд.")
		}
	}

	fmt.Println("Завершение работы системы RBAC")
}
