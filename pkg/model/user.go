package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type User struct {
	gorm.Model
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Role     Role   `gorm:"type:enum('USER', 'ADMIN');default:'USER';not null"`
}

type UserDetails interface {
	GetAuthorities() []string
	GetPassword() string
	GetUsername() string
	IsAccountNonExpired() bool
	IsAccountNonLocked() bool
	IsCredentialsNonExpired() bool
	IsEnabled() bool
}

func (u *User) IsAccountNonExpired() bool {
	return true
}

func (u *User) IsAccountNonLocked() bool {
	return true
}

func (u *User) IsCredentialsNonExpired() bool {
	return true
}

func (u *User) IsEnabled() bool {
	return true
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetUsername() string {
	return u.Email
}

func (u *User) GetAuthorities() []string {
	return []string{fmt.Sprintf("ROLE_%s", u.Role)}
}
