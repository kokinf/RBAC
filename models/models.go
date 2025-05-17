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
	ObjectID string `json:"object_id"`
	Action   string `json:"action"`
}

// Роль (набор разрешений + наследование)
type Role struct {
	Name            string       `json:"name"`
	Permissions     []Permission `json:"permissions"`
	ParentRoleNames []string     `json:"parent_roles"`
}

// Основная структура данных
type RBACData struct {
	Subjects []Subject `json:"subjects"`
	Objects  []Object  `json:"objects"`
	Roles    []Role    `json:"roles"`
}
