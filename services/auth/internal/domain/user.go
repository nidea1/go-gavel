package domain

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint64     `db:"id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Password  string     `db:"password"`
	Email     string     `db:"email"`
	IsAdmin   bool       `db:"is_admin"`
	IsActive  bool       `db:"is_active"`
	LastLogin *time.Time `db:"last_login"`
	CreatedAt time.Time  `db:"created_at"`
}

func NewUser(firstName, lastName, password, email string, isAdmin bool) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
		Email:     email,
		IsAdmin:   isAdmin,
		IsActive:  false,
		CreatedAt: time.Now(),
	}
}

func (u *User) Validate() error {
	if u.FirstName == "" {
		return errors.New("First name is required")
	}
	if u.LastName == "" {
		return errors.New("Last name is required")
	}
	if u.Email == "" {
		return errors.New("Email is required")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("Invalid email address")
	}

	return nil
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Activate() {
	u.IsActive = true
}

func (u *User) Deactivate() {
	u.IsActive = false
}

func (u *User) SetLastLogin(t time.Time) {
	u.LastLogin = &t
}

func (u *User) ChangePassword(newPassword string) error {
	if newPassword == "" {
		return errors.New("Password is required")
	}
	if len(newPassword) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(newPassword)); err == nil {
		return errors.New("New password must be different from the current password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
