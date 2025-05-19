package commands

import (
	"RBAC/models"
	"RBAC/storage"
	"fmt"
)

// ShowPermissions отображает все разрешения роли (свои и унаследованные)
func ShowPermissions(roleName string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	// Находим роль
	var targetRole *models.Role
	for i, role := range data.Roles {
		if role.Name == roleName {
			targetRole = &data.Roles[i]
			break
		}
	}

	if targetRole == nil {
		fmt.Printf("Роль '%s' не найдена\n", roleName)
		return
	}

	// Получаем все разрешения (рекурсивно)
	allPerms := getAllPermissions(targetRole, data.Roles)

	// Отображаем результат
	fmt.Printf("Все разрешения роли '%s':\n", roleName)
	for _, perm := range allPerms {
		fmt.Printf("- %s:%s\n", perm.ObjectID, perm.Action)
	}
}

// Рекурсивно собираем все разрешения
func getAllPermissions(role *models.Role, allRoles []models.Role) []models.Permission {
	permissions := make([]models.Permission, len(role.Permissions))
	copy(permissions, role.Permissions)

	// Добавляем разрешения от родительских ролей
	for _, parentName := range role.ParentRoleNames {
		for i, parentRole := range allRoles {
			if parentRole.Name == parentName {
				parentPerms := getAllPermissions(&allRoles[i], allRoles)
				permissions = append(permissions, parentPerms...)
				break
			}
		}
	}

	return permissions
}
