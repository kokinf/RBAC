package commands

import (
	"RBAC/models"
	"RBAC/storage"
	"fmt"
)

func CreateObject(id, name string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	// Проверка на существование объекта
	for _, obj := range data.Objects {
		if obj.ID == id {
			fmt.Printf("Объект с ID '%s' уже существует\n", id)
			return
		}
	}

	data.Objects = append(data.Objects, models.Object{
		ID:   id,
		Name: name,
	})

	if err := storage.SaveData(data); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	fmt.Printf("Объект '%s' (ID: %s) создан\n", name, id)
}
