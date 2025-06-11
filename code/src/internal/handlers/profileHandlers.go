package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"Forum-back/pkg/utils"
	"html/template"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{})
		return
	}
	defer db.Close()
	userService := services.NewUserService(db) // Init all services
	postService := services.NewPostService(db)
	commentService := services.NewCommentService(db)
	sessionService := services.NewSessionService(db)
	isConnected, session := sessionService.IsAuthenticated(r)

	// Fetch user_id from URL query
	userUuid := getUserIdFromURL(w, r, isConnected, &session.User_ID)
	if userUuid == uuid.Nil {
		return
	}

	user := userService.FindById(userUuid) // Find the user with the given ID
	if user == nil {                       // If no user is found, send a Not Found error page
		ShowCustomError404(w, &dtos.HeaderDto{IsConnected: isConnected}, "User not found.")
		return
	}

	user.Email = "private"                          // Set the email to private. Goal is to not share the email with anyone
	postCount := postService.GetUserPostCount(user) // Get user stats
	commentCount := commentService.GetUserCommentCount(user)
	data := createDto(isConnected, user, nil, nil, false, postCount, commentCount) // Create a DTO from data received

	showProfilePage(w, data)
}

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{})
		return
	}
	defer db.Close()

	sessionService := services.NewSessionService(db)
	userService := services.NewUserService(db)

	isConnected, session := sessionService.IsAuthenticated(r)
	if !isConnected {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	user := userService.FindById(session.User_ID)
	if user == nil {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Unable to retrieve user information from the database.")
		return
	}
	data := createDto(isConnected, user, nil, nil, true, 0, 0)
	showProfilePage(w, data)
}

func DeleteMyProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ShowError405(w, &dtos.HeaderDto{})
		return
	}

	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{})
		return
	}
	defer db.Close()

	sessionService := services.NewSessionService(db)
	userService := services.NewUserService(db)
	isConnected, session := sessionService.IsAuthenticated(r)
	if !isConnected {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	user := userService.FindById(session.User_ID)
	if user == nil {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Unable to retrieve user information from the database.")
		return
	}

	if success, err := userService.Delete(user); err != nil && !success {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Failed to delete user profile. Error: "+err.Error())
		return
	}

	deleteSessionCookie(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func EditMyProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ShowError405(w, &dtos.HeaderDto{})
		return
	}

	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{})
		return
	}
	defer db.Close()

	sessionService := services.NewSessionService(db)
	userService := services.NewUserService(db)
	isConnected, session := sessionService.IsAuthenticated(r)
	if !isConnected {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	user := userService.FindById(session.User_ID)
	if user == nil {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Unable to retrieve user information from the database.")
		return
	}
	if err := updateUserProfile(user, userService, r); err != nil {
		data := createDto(isConnected, user, nil, err, true, 0, 0)
		showProfilePage(w, data)
		return
	}

	http.Redirect(w, r, "/me", http.StatusSeeOther)
}

func updateUserProfile(user *models.User, userService *services.UserService, r *http.Request) *dtos.ProfilePageErrorDto {
	if pseudo := r.FormValue("username"); pseudo != "" {
		searchedUser := userService.FindByUsername(pseudo)
		if searchedUser != nil && searchedUser.ID != uuid.Nil && searchedUser.ID != user.ID {
			return &dtos.ProfilePageErrorDto{
				ErrorMessage: "Username already exists, please choose another one.",
				ErrorTitle:   "Username already taken",
			}
		}
		user.Pseudo = pseudo
	}

	user.Bio = r.FormValue("bio")

	if r.FormValue("new-password") != "" {
		hashedPassword, err := utils.CheckForNewPassword(r.FormValue("new-password"), r.FormValue("confirm-password"))
		if err != nil {
			return &dtos.ProfilePageErrorDto{
				ErrorMessage: err.Error(),
				ErrorTitle:   "Password Error",
			}
		}
		user.Password = hashedPassword
	}

	if avatarFile, _, err := r.FormFile("avatar-upload"); err == nil {
		defer avatarFile.Close()
		if avatarBytes, err := io.ReadAll(avatarFile); err == nil {
			user.Avatar = avatarBytes
		} else {
			return &dtos.ProfilePageErrorDto{
				ErrorMessage: "Failed to read avatar file: " + err.Error(),
				ErrorTitle:   "Avatar Upload Error",
			}
		}
	}

	if success := userService.Update(user); !success {
		return &dtos.ProfilePageErrorDto{
			ErrorMessage: "Failed to update user profile in the database.",
			ErrorTitle:   "Database Update Error",
		}
	}

	return nil

}

func showProfilePage(w http.ResponseWriter, data *dtos.ProfilPageDto) error {
	tmpl, templateError := templates.GetTemplateWithLayout(&data.Header, "myProfile", "internal/templates/profile.gohtml")
	if templateError != nil {
		ShowTemplateError500(w, &data.Header)
		return templateError
	}
	if templateError = tmpl.Execute(w, data); templateError != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return nil
}

func createDto(isConnected bool, user *models.User, RecentActivity []*dtos.RecentActivityDto, err *dtos.ProfilePageErrorDto, isMine bool, postCount, commentCount int) *dtos.ProfilPageDto {
	if err == nil {
		err = &dtos.ProfilePageErrorDto{}
	}
	data := dtos.ProfilPageDto{
		Header: dtos.HeaderDto{
			IsConnected: isConnected,
		},
		Username:       user.Pseudo,
		Email:          user.Email,
		JoinedAt:       user.CreatedAt.Format("January 02, 2006"),
		Avatar:         template.URL(utils.ConvertBytesToBase64(user.Avatar, "image/png")),
		Bio:            user.Bio,
		PostsCount:     postCount,
		CommentsCount:  commentCount,
		RecentActivity: RecentActivity, // This should be populated with actual recent activity data
		Error:          *err,
		IsMine:         isMine,
	}
	return &data
}

func getUserIdFromURL(
	w http.ResponseWriter,
	r *http.Request,
	isConnected bool,
	connectedUserId *uuid.UUID) uuid.UUID {
	userId := r.URL.Query().Get("user_id")
	if userId == "" { // If userId is not set
		if isConnected { // And user is connected, redirect him to his profile
			http.Redirect(w, r, "/me", http.StatusSeeOther)
			return uuid.Nil
		}
		ShowError400(w, &dtos.HeaderDto{IsConnected: isConnected}) // Else, send a bad request error
		return uuid.Nil
	}

	userUuid, err := uuid.Parse(userId)     // Parse the id from URL into a uuid
	if err != nil || userUuid == uuid.Nil { // If the parse fail
		ShowError400(w, &dtos.HeaderDto{IsConnected: isConnected}) // Send a bad request error due to invalid ID
		return uuid.Nil
	}
	if connectedUserId != nil && userUuid == *connectedUserId { // If the user_id from URL is corresponding to the id of the connected user
		http.Redirect(w, r, "/me", http.StatusSeeOther) // Redirect the user to his profile
		return uuid.Nil
	}

	return userUuid
}
