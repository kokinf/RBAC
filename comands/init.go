package commands

import (
	"RBAC/storage"
	"fmt"
)

func Init() {
	err := storage.InitStorage()
	if err != nil {
		fmt.Println("Ошибка инициализации:", err)
		return
	}
	fmt.Println("Система успешно инициализирована")
}
