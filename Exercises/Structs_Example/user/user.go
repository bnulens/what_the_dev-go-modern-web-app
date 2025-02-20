package user

import (
	"errors"
	"fmt"
	"time"
)

// Any
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// A struct can be embedded without a field name (anonymous embedding)
// This still includes the User into the Admin struct but can be omitted when calling
// upon any methods of the Admin struct
//	type Admin struct {
//		email    string
//		password string
//		User
//	}

// User struct has to made available globally inside project by uppercasing 'user' field
type Admin struct {
	email    string
	password string
	User     User
}

func GetUserData(message string) string {
	fmt.Println(message)
	var value string
	fmt.Scan(&value)
	return value
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name or date of birth are empty. Please fill in correctly")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// Returns a value instead of a pointer - this is fine because the Struct is not too complex yet
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Brecht",
			lastName:  "Nulens",
			birthDate: "--",
			createdAt: time.Now(),
		},
	}
}
func (u *User) OutputUserDetails() {
	fmt.Printf("%s, %s, %s", u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
