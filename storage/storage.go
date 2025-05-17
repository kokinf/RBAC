package storage

import (
	"RBAC/models"
	"encoding/json"
	"os"
)

const dataFile = "rbac_data.json"

func LoadData() (*models.RBACData, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return &models.RBACData{
			Subjects: []models.Subject{},
			Objects:  []models.Object{},
			Roles:    []models.Role{},
		}, nil
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	var rbacData models.RBACData
	err = json.Unmarshal(data, &rbacData)
	return &rbacData, err
}

func SaveData(data *models.RBACData) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, file, 0644)
}

func InitStorage() error {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return SaveData(&models.RBACData{
			Subjects: []models.Subject{},
			Objects:  []models.Object{},
			Roles:    []models.Role{},
		})
	}
	return nil
}
