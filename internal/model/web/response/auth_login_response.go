package response

type AuthLoginResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type AuthLoginTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
