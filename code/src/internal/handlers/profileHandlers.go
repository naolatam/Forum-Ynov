package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"Forum-back/pkg/utils"
	"database/sql"
	"html/template"
	"io"
	"net/http"

	"github.com/google/uuid"
)

// ProfileHandler handles the display of a user's profile page.
func ProfileHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {

	userService := services.NewUserService(db) // Init all services
	postService := services.NewPostService(db)
	commentService := services.NewCommentService(db)

	if session == nil { // If the user is not connected, session will be nil
		session = &models.Session{User_ID: uuid.Nil} // Create a session with a nil user ID to avoid nil pointer dereference
	}

	// Fetch user_id from URL query
	userUuid := getUserIdFromURL(w, r, header, &session.User_ID)
	if userUuid == uuid.Nil {
		return
	}

	user := userService.FindById(userUuid) // Find the user with the given ID
	if user == nil {                       // If no user is found, send a Not Found error page
		ShowCustomError404(w, header, "User not found.")
		return
	}

	user.Email = "private"                          // Set the email to private. Goal is to not share the email with anyone
	postCount := postService.GetUserPostCount(user) // Get user stats
	commentCount := commentService.GetUserCommentCount(user)
	data := createDto(header, user, nil, nil, false, postCount, commentCount) // Create a DTO from data received

	showProfilePage(w, data)
}

// MyProfileHandler handles the display of the connected user's profile page.
func MyProfileHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {

	userService := services.NewUserService(db)
	postService := services.NewPostService(db)
	commentService := services.NewCommentService(db)
	ras := services.NewRecentActivityService(db)

	user := userService.FindById(session.User_ID)
	if user == nil {
		ShowCustomError500(w, header, "Unable to retrieve your information from the database.")
		return
	}

	postCount := postService.GetUserPostCount(user) // Get user stats
	commentCount := commentService.GetUserCommentCount(user)
	recentActivity, _ := ras.FindByUser(user) // Get recent activity for the user
	if recentActivity == nil {
		recentActivity = &([]*models.RecentActivity{}) // If no recent activity, initialize to an empty slice
	}

	data := createDto(header, user, *recentActivity, nil, true, postCount, commentCount)
	showProfilePage(w, data)
}

// DeleteMyProfileHandler handles the deletion of the connected user's profile.
func DeleteMyProfileHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {

	userService := services.NewUserService(db)

	user := userService.FindById(session.User_ID)
	if user == nil {
		ShowCustomError500(w, header, "Unable to retrieve user information from the database.")
		return
	}

	if success, err := userService.Delete(user); err != nil && !success {
		ShowCustomError500(w, header, "Failed to delete user profile. Error: "+err.Error())
		return
	}

	deleteSessionCookie(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// EditMyProfileHandler handles the editing of the connected user's profile.
func EditMyProfileHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	userService := services.NewUserService(db)

	user := userService.FindById(session.User_ID)
	if user == nil {
		ShowCustomError500(w, header, "Unable to retrieve user information from the database.")
		return
	}
	if err := updateUserProfile(user, userService, r); err != nil {
		data := createDto(header, user, nil, err, true, 0, 0)
		showProfilePage(w, data)
		return
	}

	http.Redirect(w, r, "/me", http.StatusSeeOther)
}

// updateUserProfile updates the user's profile based on the form data from the request.
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

// showProfilePage renders the profile page template with the provided data.
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

// createDto constructs a ProfilePageDto from the provided parameters.
func createDto(header *dtos.HeaderDto, user *models.User, RecentActivity []*models.RecentActivity, err *dtos.ProfilePageErrorDto, isMine bool, postCount, commentCount int) *dtos.ProfilPageDto {
	if err == nil {
		err = &dtos.ProfilePageErrorDto{}
	}
	data := dtos.ProfilPageDto{
		Header:         *header,
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

// getUserIdFromURL retrieves the user ID from the URL query parameters.
func getUserIdFromURL(
	w http.ResponseWriter,
	r *http.Request,
	header *dtos.HeaderDto,
	connectedUserId *uuid.UUID) uuid.UUID {
	userId := r.URL.Query().Get("user_id")
	if userId == "" { // If userId is not set
		if header.IsConnected { // And user is connected, redirect him to his profile
			http.Redirect(w, r, "/me", http.StatusSeeOther)
			return uuid.Nil
		}
		ShowError400(w, header) // Else, send a bad request error
		return uuid.Nil
	}

	userUuid, err := uuid.Parse(userId)     // Parse the id from URL into a uuid
	if err != nil || userUuid == uuid.Nil { // If the parse fail
		ShowError400(w, header) // Send a bad request error due to invalid ID
		return uuid.Nil
	}
	if connectedUserId != nil && userUuid == *connectedUserId { // If the user_id from URL is corresponding to the id of the connected user
		http.Redirect(w, r, "/me", http.StatusSeeOther) // Redirect the user to his profile
		return uuid.Nil
	}

	return userUuid
}
