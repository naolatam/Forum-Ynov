package oauth

import (
	"Forum-back/pkg/dtos"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// GetGoogleUserInfoFromCode exchanges an authorization code for a Google user info DTO.
func GetGoogleUserInfoFromCode(code string) (*dtos.GoogleUserInfo, error) {
	// 1. Échanger le code contre un token
	googleOauthConfig := GetGoogleOauthConfig()

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	// 2. Récupérer les infos de l'utilisateur
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user info from Google")
	}

	// 3. Décode dans le DTO
	var userInfo dtos.GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
