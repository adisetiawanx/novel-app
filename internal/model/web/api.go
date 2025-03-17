package web

import (
	"github.com/adisetiawanx/novel-app/internal/model/web/response"
)

type APIResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type UserResponseWrapper struct {
	User response.AuthRegisterResponse `json:"user"`
}

type LoginResponseWrapper struct {
	User  response.AuthLoginResponse      `json:"user"`
	Token response.AuthLoginTokenResponse `json:"token"`
}
