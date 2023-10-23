package models

import (
	"errors"
	"time"

	"github.com/azkazkazka/task-todo/auth"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

type UserService struct {
	Service IUserService
}

type GormUserService struct {
	DB *gorm.DB
}

func (us *GormUserService) AuthResponseToUserResponse(authResponse *AuthResponse) *UserResponse {
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

func (us *GormUserService) Register(userRequest *UserRequest) (interface{}, error) {
	var err error

	user := &UserResponse{
		Fullname: userRequest.Fullname,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Username: userRequest.Username,
	}

	user.ID = uuid.New().String()

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	if err := us.DB.Table("users").Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (us *GormUserService) AuthenticateUser(loginRequest *LoginRequest) (*UserResponse, error) {
	user, err := us.findUserByUsernameOrEmail(loginRequest.Username, loginRequest.Email)
	if err != nil {
		return nil, err
	}

	if user == nil || !auth.CheckPasswordHash(loginRequest.Password, user.Password) {
		return nil, err
	}

	return us.AuthResponseToUserResponse(user), nil
}

func (us *GormUserService) findUserByUsernameOrEmail(username, email string) (*AuthResponse, error) {
	user := &AuthResponse{}

	if err := us.DB.Table("users").Where("username = ?", username).First(user).Error; err == nil {
		return user, nil
	}

	if err := us.DB.Table("users").Where("email = ?", email).First(user).Error; err == nil {
		return user, nil
	}

	return nil, errors.New("user not found")
}

func (us *GormUserService) FetchAllUsers() (interface{}, error) {
	var users []UserResponse

	if err := us.DB.Table("users").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (us *GormUserService) GetUser(userID string) (interface{}, error) {
	user := &UserResponse{}

	if err := us.DB.Table("users").Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (us *GormUserService) DeleteUser(userID string) (interface{}, error) {
	existingUser := &UserResponse{}

	if err := us.DB.Table("users").First(&existingUser).Where("id = ?", userID).Error; err != nil {
		return nil, err
	}

	if err := us.DB.Table("users").Delete(existingUser).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *GormUserService) UpdateUser(user *UpdateRequest) (interface{}, error) {

	existingUser := &UserResponse{}
	if err := us.DB.Table("users").First(&existingUser).Where("id = ?", user.ID).Error; err != nil {
		return nil, err
	}

	user.UpdatedAt = time.Now()

	if err := us.DB.Table("users").Model(&UserResponse{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
