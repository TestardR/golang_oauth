package app

import "example.com/controllers/oauth"

func mapUrls() {
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
