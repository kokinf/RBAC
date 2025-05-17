package commands

import (
	"RBAC/models"
	"RBAC/storage"
	"fmt"
)

func AssignRole(subjectID, roleName string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	roleExists := false
	for _, role := range data.Roles {
		if role.Name == roleName {
			roleExists = true
			break
		}
	}

	if !roleExists {
		fmt.Printf("Роль '%s' не найдена\n", roleName)
		return
	}

	for i, subject := range data.Subjects {
		if subject.ID == subjectID {
			data.Subjects[i].RoleName = roleName
			err = storage.SaveData(data)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
				return
			}
			fmt.Printf("Роль '%s' назначена субъекту '%s'\n", roleName, subjectID)
			return
		}
	}

	newSubject := models.Subject{
		ID:       subjectID,
		RoleName: roleName,
	}
	data.Subjects = append(data.Subjects, newSubject)

	err = storage.SaveData(data)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Printf("Создан новый субъект '%s' с ролью '%s'\n", subjectID, roleName)
}
