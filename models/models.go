package models

// Субъект (юзер)
type Subject struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	RoleName string `json:"role_name"`
}

// Объект (ресурс, к которому регулируется доступ)
type Object struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Разрешение (действие над объектом)
type Permission struct {
	ObjectID string `json:"object_id"` // Используется ObjectID вместо ResourceID
	Action   string `json:"action"`
}

// Роль (набор разрешений + наследование)
type Role struct {
	Name            string       `json:"name"` // Используется как идентификатор
	Permissions     []Permission `json:"permissions"`
	ParentRoleNames []string     `json:"parent_role_names"` // Используется ParentRoleNames
}

// Основная структура данных
type RBACData struct {
	Subjects []Subject `json:"subjects"`
	Objects  []Object  `json:"objects"`
	Roles    []Role    `json:"roles"`
}
