package models

type IUserService interface {
	AuthResponseToUserResponse(authResponse *AuthResponse) *UserResponse
	Register(userRequest *UserRequest) (interface{}, error)
	AuthenticateUser(loginRequest *LoginRequest) (*UserResponse, error)
	FetchAllUsers() (interface{}, error)
	GetUser(userID string) (interface{}, error)
	DeleteUser(userID string) (interface{}, error)
	UpdateUser(user *UpdateRequest) (interface{}, error)
}