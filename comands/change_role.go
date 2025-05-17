package commands

import (
	"RBAC/storage"
	"fmt"
)

func ChangeRole(subjectID, newRoleName string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	roleExists := false
	for _, role := range data.Roles {
		if role.Name == newRoleName {
			roleExists = true
			break
		}
	}

	if !roleExists {
		fmt.Printf("Роль '%s' не найдена\n", newRoleName)
		return
	}

	subjectFound := false
	for i, subject := range data.Subjects {
		if subject.ID == subjectID {
			oldRole := subject.RoleName
			data.Subjects[i].RoleName = newRoleName
			subjectFound = true

			err = storage.SaveData(data)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
				return
			}

			fmt.Printf("Роль субъекта '%s' изменена с '%s' на '%s'\n",
				subjectID, oldRole, newRoleName)
			break
		}
	}

	if !subjectFound {
		fmt.Printf("Субъект '%s' не найден\n", subjectID)
	}
}
