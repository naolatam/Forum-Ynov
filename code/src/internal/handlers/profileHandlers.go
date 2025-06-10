package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/services"
	"Forum-back/pkg/utils"
	"html/template"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{})
		return
	}
	defer db.Close()

	data := map[string]interface{}{
		/* "id":     UserID,
		"pseudo": Pseudo,
		"email":  Email, */
	}

	tmpl, err := template.ParseFiles("internal/templates/profile.gohtml", "internal/templates/components/headerComponent.gohtml")
	if err != nil {
		ShowTemplateError500(w, &dtos.HeaderDto{})
		return
	}

	tmpl.Execute(w, data)
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

	data := dtos.ProfilPageDto{
		Header: dtos.HeaderDto{
			IsConnected: isConnected,
		},
		Username:       user.Pseudo,
		Email:          user.Email,
		JoinedAt:       user.CreatedAt.Format("January 02, 2006"),
		Avatar:         template.URL(utils.ConvertBytesToBase64(user.Avatar, "image/png")),
		Bio:            user.Bio,
		PostsCount:     0,
		CommentsCount:  0,
		RecentActivity: nil, // This should be populated with actual recent activity data
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "myProfile", "internal/templates/profile.gohtml")
	if err != nil {
		ShowTemplateError500(w, &data.Header)
		return
	}

	tmpl.Execute(w, data)
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
