package form

type OAuthRequest struct {
	Provider          string `json:"provider" binding:"required"`
	Email             string `json:"email" binding:"required"`
	ProviderAccountId string `json:"provider_account_id" binding:"required"`
	Name              string `json:"name"`
	Image             string `json:"image"`
}