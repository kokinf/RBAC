package commands

import (
	"RBAC/models"
	"RBAC/storage"
	"fmt"
)

func InheritRole(childRole, parentRole string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	var childIdx, parentIdx = -1, -1
	for i, role := range data.Roles {
		if role.Name == childRole {
			childIdx = i
		}
		if role.Name == parentRole {
			parentIdx = i
		}
	}

	if childIdx == -1 || parentIdx == -1 {
		fmt.Println("Одна из указанных ролей не найдена")
		return
	}

	if createsCycle(data.Roles, childRole, parentRole) {
		fmt.Println("Ошибка: создание цикла в иерархии ролей")
		return
	}

	for _, parent := range data.Roles[childIdx].ParentRoleNames {
		if parent == parentRole {
			fmt.Printf("Роль '%s' уже наследует от '%s'\n", childRole, parentRole)
			return
		}
	}

	data.Roles[childIdx].ParentRoleNames = append(data.Roles[childIdx].ParentRoleNames, parentRole)

	err = storage.SaveData(data)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	fmt.Printf("Роль '%s' теперь наследует от '%s'\n", childRole, parentRole)
}

func createsCycle(roles []models.Role, child, parent string) bool {
	if child == parent {
		return true
	}

	for _, role := range roles {
		if role.Name == parent {
			for _, grandparent := range role.ParentRoleNames {
				if createsCycle(roles, child, grandparent) {
					return true
				}
			}
		}
	}

	return false
}
