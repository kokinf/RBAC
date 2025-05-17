package commands

import (
	"RBAC/models"
	"RBAC/storage"
	"fmt"
)

func CreateUser(userID, userName string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	// Проверяем существует ли уже пользователь
	for _, subject := range data.Subjects {
		if subject.ID == userID {
			fmt.Printf("Субъект с ID '%s' уже существует\n", userID)
			return
		}
	}

	// Создаем нового пользователя
	newUser := models.Subject{
		ID:   userID,
		Name: userName,
	}

	data.Subjects = append(data.Subjects, newUser)

	err = storage.SaveData(data)
	if err != nil {
		fmt.Println("Ошибка сохранения данных:", err)
		return
	}

	fmt.Printf("Субъект '%s' (ID: %s) успешно создан\n", userName, userID)
}
