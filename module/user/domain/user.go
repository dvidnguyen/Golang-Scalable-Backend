package domain

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

// `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
// `first_name` varchar(50) DEFAULT NULL,
// `last_name` varchar(50) DEFAULT NULL,
// `email` varchar(100) DEFAULT NULL,
// `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
// `salt` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `avatar` json DEFAULT NULL,
// `role` enum('user','admin') NOT NULL DEFAULT 'user',
// `status` enum('activated','banned') DEFAULT 'activated',
// `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

type User struct {
	id        uuid.UUID
	firstName string
	lastName  string
	email     string
	password  string
	salt      string
	role      Role
}

func NewUser(id uuid.UUID, firstName string, lastName string, email string, password string, salt string, role Role) (*User, error) {
	// Todo Validation
	// Check UUID rỗng
	if id == uuid.Nil {
		return nil, fmt.Errorf("id is required")
	}
	// Check firstName rỗng hoặc chỉ có space
	if strings.TrimSpace(firstName) == "" {
		return nil, fmt.Errorf("firstName is required")
	}
	// Check lastName rỗng hoặc chỉ có space
	if strings.TrimSpace(lastName) == "" {
		return nil, fmt.Errorf("lastName is required")
	}
	// Check email rỗng
	if strings.TrimSpace(email) == "" {
		return nil, fmt.Errorf("email is required")
	}
	// Validate email đúng định dạng
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, fmt.Errorf("invalid email")
	}
	// Check password rỗng
	if strings.TrimSpace(password) == "" {
		return nil, fmt.Errorf("password is required")
	}
	// Check salt rỗng
	if strings.TrimSpace(salt) == "" {
		return nil, fmt.Errorf("salt is required")
	}
	return &User{id: id, firstName: firstName, lastName: lastName, email: email, password: password, salt: salt, role: role}, nil
}

func (u User) Id() uuid.UUID {
	return u.id
}

func (u User) FirstName() string {
	return u.firstName
}

func (u User) LastName() string {
	return u.lastName
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) Salt() string {
	return u.salt
}

func (u User) Role() Role {
	return u.role
}

type Role int

const (
	RoleUser Role = iota
	RoleAdmin
)

func (r Role) String() string {
	return [2]string{"user", "admin"}[r]
}

func GetRole(s string) Role {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "admin":
		return RoleAdmin
	default:
		return RoleUser
	}
}
