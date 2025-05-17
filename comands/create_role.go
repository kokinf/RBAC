package commands

import (
	"RBAC/models"
	"RBAC/storage"
	"fmt"
)

func CreateRole(roleName string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	for _, role := range data.Roles {
		if role.Name == roleName {
			fmt.Printf("Роль '%s' уже существует\n", roleName)
			return
		}
	}

	newRole := models.Role{
		Name:            roleName,
		Permissions:     []models.Permission{},
		ParentRoleNames: []string{},
	}

	data.Roles = append(data.Roles, newRole)

	err = storage.SaveData(data)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	fmt.Printf("Роль '%s' успешно создана\n", roleName)
}
