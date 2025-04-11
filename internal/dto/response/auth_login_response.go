package response

type AuthLoginResponse struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
}

type AuthLoginTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
