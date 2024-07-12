package google

import (
	"context"
	"encoding/json"

	"github.com/Grafiters/archive/app/payload"
	"github.com/Grafiters/archive/configs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig  *oauth2.Config
	googleClientID     = "311070042723-d2q0lhe9se560g1v7taikht1ucs9dqfj.apps.googleusercontent.com"
	googleClientSecret = "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"
)

type GoogleData struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	GoogleID  string `json:"gooogle_id"`
	GivenName string `json:"given_name"`
	Picture   string `json:"picture"`
}

func init() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleCallbackHandler(auth payload.SessionGooglePayload) (*GoogleData, error) {
	googleOauthConfig.RedirectURL = auth.RedirectURL
	token, err := googleOauthConfig.Exchange(context.Background(), auth.Code)

	if err != nil {
		configs.Logger.Error("Error exchanging code: %s\n", err)
		return nil, err
	}

	userInfo, err := getUserInfo(token)
	if err != nil {
		configs.Logger.Error("Error getting user info: %s\n", err)
		return nil, err
	}

	userInfoJson, err := json.Marshal(userInfo)
	if err != nil {
		configs.Logger.Error("Error getting user info: %s\n", err)
		return nil, err
	}

	var googleData GoogleData
	err = json.Unmarshal([]byte(userInfoJson), &googleData)
	if err != nil {
		configs.Logger.Error("Error decoding JSON:", err)
		return nil, err
	}

	return &googleData, err
}

func getUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := googleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	configs.Logger.Info(userInfo)

	return userInfo, nil
}
