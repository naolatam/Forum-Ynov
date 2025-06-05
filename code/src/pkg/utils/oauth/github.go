package oauth

import (
	"Forum-back/pkg/dtos"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func GetGithubUserInfoFromCode(code string) (*dtos.GitHubUserInfo, error) {
	// Exchange the code received from Github for an access token
	githubOauthConfig := GetGithubOauthConfig()

	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	// Fetch user info from Github API
	client := githubOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Fetch user email since Github does not return it in the user info response if the user has not made it public
	respForEmail, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		return nil, err
	}
	defer respForEmail.Body.Close()

	if respForEmail.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user Email info from Github")
	}

	// Decode principal user info and emails into a GitHubUserInfo struct
	var userInfo dtos.GitHubUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	if err := json.NewDecoder(respForEmail.Body).Decode(&userInfo.Emails); err != nil {
		return nil, err
	}
	return &userInfo, nil
}
