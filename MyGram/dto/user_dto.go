package dto

import "time"


type Response struct{
	Success 	bool `json:"success"`
	Message 	interface{} `json:"message"`
	Data 	interface{} `json:"data"`
}


type RegisterRequest struct{
	Age int `validate:"required,min=8" json:"age"`
	Email string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6" json:"password"`
	UserName string `validate:"required" json:"user_name"`
	
}

type RegisterResponse struct{
	Age int `json:"age"`
	Email string `json:"email"`
	ID string `json:"id"`
	UserName string `json:"user_name"`
}

type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct{
	Token string `json:"token"`
}

type UpdateUserRequest struct{
	Email string `validate:"required,email" json:"email"`
	UserName string `validate:"required" json:"user_name"`
}

type UpdateUserResponse struct{
	ID string `json:"id"`
	Email string `json:"email"`
	UserName string `json:"user_name"`
	Age int `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`

}

type DeleleUserResponse struct{
	Message string `json:"message"`
}