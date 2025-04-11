package dto

import (
	response2 "github.com/adisetiawanx/novel-app/internal/dto/response"
)

type APIResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type UserResponseWrapper struct {
	User response2.AuthRegisterResponse `json:"user"`
}

type LoginResponseWrapper struct {
	User  response2.AuthLoginResponse      `json:"user"`
	Token response2.AuthLoginTokenResponse `json:"token"`
}
