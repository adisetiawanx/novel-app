package response

type LoginResponseWrapper struct {
	Token AuthLoginTokenResponse `json:"token"`
}

type AuthLoginTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
