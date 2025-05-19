package commands

import (
	"RBAC/models"
	"RBAC/storage"
	"fmt"
	"strings"
)

func AssignPermission(roleName, objectID, action string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	// Находим роль
	roleIndex := -1
	for i, role := range data.Roles {
		if role.Name == roleName {
			roleIndex = i
			break
		}
	}

	if roleIndex == -1 {
		fmt.Printf("Роль '%s' не найдена\n", roleName)
		return
	}

	// Проверяем существующие разрешения для этого объекта
	existingIndex := -1
	for i, perm := range data.Roles[roleIndex].Permissions {
		if perm.ObjectID == objectID {
			existingIndex = i
			break
		}
	}

	// Если разрешение для этого объекта уже существует - объединяем действия
	if existingIndex >= 0 {
		existingPerm := &data.Roles[roleIndex].Permissions[existingIndex]

		// Проверяем, не существует ли уже такое действие
		if strings.Contains(existingPerm.Action, action) {
			fmt.Printf("Действие '%s' уже существует для объекта '%s' в роли '%s'\n",
				action, objectID, roleName)
			return
		}

		// Объединяем действия через запятую
		existingPerm.Action = strings.Join([]string{existingPerm.Action, action}, ",")
		fmt.Printf("Действие '%s' добавлено к существующим разрешениям для объекта '%s' в роли '%s'\n",
			action, objectID, roleName)
	} else {
		// Создаем новое разрешение
		newPermission := models.Permission{
			ObjectID: objectID,
			Action:   action,
		}
		data.Roles[roleIndex].Permissions = append(data.Roles[roleIndex].Permissions, newPermission)
		fmt.Printf("Новое разрешение '%s:%s' создано для роли '%s'\n",
			objectID, action, roleName)
	}

	err = storage.SaveData(data)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
	}
}
