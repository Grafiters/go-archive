package services

import (
	"archive/actions/utils/interfaces"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig  *oauth2.Config
	googleClientID     = "311070042723-d2q0lhe9se560g1v7taikht1ucs9dqfj.apps.googleusercontent.com"
	googleClientSecret = "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"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleCallbackHandler(auth interfaces.LoginGoogle) (*interfaces.GoogleData, error) {
	googleOauthConfig.RedirectURL = auth.RedirectUrl
	token, err := googleOauthConfig.Exchange(context.Background(), auth.Code)

	if err != nil {
		fmt.Println("==============================")
		log.Printf("Error exchanging code: %s\n", err.Error())
		return nil, err
	}

	userInfo, err := getUserInfo(token)
	if err != nil {
		log.Printf("Error getting user info: %s\n", err.Error())
		return nil, err
	}

	userInfoJson, err := json.Marshal(userInfo)
	if err != nil {
		log.Printf("Error getting user info: %s\n", err.Error())
		return nil, err
	}

	var googleData interfaces.GoogleData
	err = json.Unmarshal([]byte(userInfoJson), &googleData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
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

	return userInfo, nil
}
