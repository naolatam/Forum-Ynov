/* package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "34247152246-btfgsp7evifdtl9ads1lb7hpeefclv3h.apps.googleusercontent.com", //os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: "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",                                     //os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "random" // peut être plus sécurisé (csrf)
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Println("invalid oauth state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("code exchange failed:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("failed getting user info:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&userInfo)
	fmt.Fprintf(w, "User Info: %+v\n", userInfo)
}

func main() {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	fmt.Println("Started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/

/*
	package main

import (

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

)

var (

	githubOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/auth/callback/github",
		ClientID:     "Ov23liKWJXf8KLKApgQc",                     //os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: "5d2f715717b53d2bc8ecb486eabcffd99580545a", //os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
	}
	oauthStateString = "randomstate" // Utiliser une valeur plus sécurisée en prod

)

	func handleLogin(w http.ResponseWriter, r *http.Request) {
		url := githubOauthConfig.AuthCodeURL(oauthStateString)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}

	func handleCallback(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		if state != oauthStateString {
			http.Error(w, "Invalid OAuth state", http.StatusBadRequest)
			return
		}

		code := r.FormValue("code")
		token, err := githubOauthConfig.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, "Code exchange failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Récupérer infos utilisateur depuis l'API GitHub
		client := githubOauthConfig.Client(context.Background(), token)
		resp, err := client.Get("https://api.github.com/user")
		if err != nil {
			http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var userInfo map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&userInfo)
		fmt.Fprintf(w, "User Info: %+v\n", userInfo)
	}

	func main() {
		err := godotenv.Load(".env") // Chemin relatif vers ton .env
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		http.HandleFunc("/login", handleLogin)
		http.HandleFunc("/api/auth/callback/github", handleCallback)
		fmt.Println("Server started at http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
*/
package main
