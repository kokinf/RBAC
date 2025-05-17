package commands

import (
	"RBAC/models"
	"RBAC/storage"
	"fmt"
)

func MergeRoles(newRoleName string, roleNames []string, strategy string) {
	data, err := storage.LoadData()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	for _, role := range data.Roles {
		if role.Name == newRoleName {
			fmt.Printf("Роль '%s' уже существует\n", newRoleName)
			return
		}
	}

	var rolesToMerge []*models.Role
	for _, roleName := range roleNames {
		found := false
		for i := range data.Roles {
			if data.Roles[i].Name == roleName {
				rolesToMerge = append(rolesToMerge, &data.Roles[i])
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Роль '%s' не найдена\n", roleName)
			return
		}
	}

	var mergedPermissions []models.Permission
	switch strategy {
	case "union":
		mergedPermissions = mergeUnion(rolesToMerge)
	case "intersection":
		mergedPermissions = mergeIntersection(rolesToMerge)
	default:
		fmt.Println("Неизвестная стратегия. Используйте 'union' или 'intersection'")
		return
	}

	newRole := models.Role{
		Name:            newRoleName,
		Permissions:     mergedPermissions,
		ParentRoleNames: roleNames,
	}

	data.Roles = append(data.Roles, newRole)

	err = storage.SaveData(data)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	fmt.Printf("Создана новая роль '%s' путем объединения %v с стратегией '%s'\n",
		newRoleName, roleNames, strategy)
}

func mergeUnion(roles []*models.Role) []models.Permission {
	uniquePerms := make(map[models.Permission]bool)
	var result []models.Permission

	for _, role := range roles {
		for _, perm := range role.Permissions {
			if !uniquePerms[perm] {
				uniquePerms[perm] = true
				result = append(result, perm)
			}
		}
	}

	return result
}

func mergeIntersection(roles []*models.Role) []models.Permission {
	if len(roles) == 0 {
		return []models.Permission{}
	}

	permCounts := make(map[models.Permission]int)
	for _, role := range roles {
		seen := make(map[models.Permission]bool)
		for _, perm := range role.Permissions {
			if !seen[perm] {
				seen[perm] = true
				permCounts[perm]++
			}
		}
	}

	var result []models.Permission
	required := len(roles)
	for perm, count := range permCounts {
		if count == required {
			result = append(result, perm)
		}
	}

	return result
}
