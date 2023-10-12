package models

import (
	"errors"
	"time"

	"github.com/azkazkazka/task-todo/auth"
	"github.com/azkazkazka/task-todo/db"
	"github.com/google/uuid"
)

type UserRequest struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UpdateRequest struct {
	ID        string    `json:"id"`
	Fullname  string    `json:"fullname,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	Fullname  string    `json:"fullname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthResponse struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	Fullname  string    `json:"fullname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"-"`
	UpdatedAt time.Time `json:"updated_at" gorm:"-"`
}

func AuthResponseToUserResponse(authResponse *AuthResponse) *UserResponse {
	if authResponse == nil {
		return nil
	}

	userResponse := &UserResponse{
		ID:        authResponse.ID,
		Fullname:  authResponse.Fullname,
		Username:  authResponse.Username,
		Email:     authResponse.Email,
		CreatedAt: authResponse.CreatedAt,
		UpdatedAt: authResponse.UpdatedAt,
	}
	return userResponse
}

func Register(userRequest *UserRequest) (interface{}, error) {
	var err error
	con := db.CreateCon()

	user := &UserResponse{
		Fullname: userRequest.Fullname,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Username: userRequest.Username,
	}

	user.ID = uuid.New().String()

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	if err := con.Table("users").Create(user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func AuthenticateUser(loginRequest *LoginRequest) (*UserResponse, error) {
	user, err := findUserByUsernameOrEmail(loginRequest.Username, loginRequest.Email)
	if err != nil {
		return nil, errors.New("could not find user by username or email")
	}

	if user == nil || !auth.CheckPasswordHash(loginRequest.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return AuthResponseToUserResponse(user), nil
}

func findUserByUsernameOrEmail(username, email string) (*AuthResponse, error) {
	con := db.CreateCon()
	user := &AuthResponse{}

	if err := con.Table("users").Where("username = ?", username).First(user).Error; err == nil {
		return user, nil
	}

	if err := con.Table("users").Where("email = ?", email).First(user).Error; err == nil {
		return user, nil
	}

	return nil, errors.New("user not found")
}

func FetchAllUsers() (interface{}, error) {
	var users []UserResponse
	con := db.CreateCon()

	if err := con.Table("users").Find(&users).Error; err != nil {
		return nil, errors.New("no users found")
	}

	return users, nil
}

func GetUser(userID string) (interface{}, error) {
	user := &UserResponse{}
	con := db.CreateCon()

	if err := con.Table("users").Where("id = ?", userID).First(user).Error; err == nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(userID string) (interface{}, error) {
	con := db.CreateCon()
	existingUser := &UserResponse{}

	if err := con.Table("users").First(existingUser, userID).Error; err != nil {
		return nil, errors.New("user does not exist")
	}

	if err := con.Table("users").Delete(existingUser).Error; err != nil {
		return nil, errors.New("failed to delete user")
	}

	return nil, nil
}

func UpdateUser(user *UpdateRequest) (interface{}, error) {
	con := db.CreateCon()

	existingUser := &UserResponse{}
	if err := con.Table("users").First(existingUser, user.ID).Error; err != nil {
		return nil, errors.New("user does not exist")
	}

	user.UpdatedAt = time.Now()

	if err := con.Table("users").Model(&UserResponse{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, errors.New("failed to update user")
	}

	return user, nil
}
