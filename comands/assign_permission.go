package commands

import (
	"RBAC/models"
	"RBAC/storage"
	"fmt"
)

func AssignPermission(roleName, objectID, action string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	roleIdx := -1
	for i, role := range data.Roles {
		if role.Name == roleName {
			roleIdx = i
			break
		}
	}

	if roleIdx == -1 {
		fmt.Printf("Роль '%s' не найдена\n", roleName)
		return
	}

	for _, perm := range data.Roles[roleIdx].Permissions {
		if perm.ObjectID == objectID && perm.Action == action {
			fmt.Printf("Разрешение '%s:%s' уже существует для роли '%s'\n", objectID, action, roleName)
			return
		}
	}

	newPerm := models.Permission{
		ObjectID: objectID,
		Action:   action,
	}

	data.Roles[roleIdx].Permissions = append(data.Roles[roleIdx].Permissions, newPerm)

	err = storage.SaveData(data)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	fmt.Printf("Разрешение '%s:%s' добавлено для роли '%s'\n", objectID, action, roleName)
}
