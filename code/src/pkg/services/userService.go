package services

import (
	"Forum-back/pkg/dtos"
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"Forum-back/pkg/utils"
	"fmt"

	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo        *repositories.UserRepository
	sessionRepo *repositories.SessionRepository
	roleRepo    *repositories.RoleRepository
}

// FindById retrieves a user by their ID.
func (service *UserService) FindById(id uuid.UUID) *models.User {
	user, err := service.repo.FindById(id)
	if err != nil {
		return nil
	}
	return user
}

// FindByEmail retrieves a user by their email address.
func (service *UserService) FindByEmail(email string) *models.User {
	user, err := service.repo.FindByEmail(&email)
	if err != nil {
		return nil
	}
	return user
}

// FindByUsername retrieves a user by their username.
func (service *UserService) FindByUsername(username string) *models.User {
	user, err := service.repo.FindByUsername(&username)
	if err != nil {
		return nil
	}
	return user
}

// FindMultipleByAny searches for users by any field (username, email, etc.) using a search string.
func (service *UserService) FindMultipleByAny(search string) *[]*models.User {
	query := "%" + search + "%"
	users, err := service.repo.FindMultipleByAny(query)
	if err != nil {
		log.Println("Error finding users by any:", err)
		return nil
	}
	return users
}

// FindByUsernameOrEmail retrieves a user by either their username or email address.
func (service *UserService) FindByUsernameOrEmail(pseudo *string, email *string) *models.User {
	if pseudo == nil && email == nil {
		return nil
	}
	user, err := service.repo.FindByUsernameOrEmail(pseudo, email)
	if err != nil {
		return nil
	}
	return user
}

// FindByIncompleteEmail searches for users with an email that contains the provided string.
func (service *UserService) FindByIncompleteEmail(email string) *models.User {
	if email == "" {
		return nil
	}
	query := "%" + email + "%"
	user, err := service.repo.FindByEmail(&query)
	if err != nil {
		return nil
	}
	return user
}

// FindByGoogleId retrieves a user by their Google ID.
func (service *UserService) FindByGoogleId(googleId string) *models.User {
	if googleId == "" {
		return nil
	}
	user, err := service.repo.FindByGoogleID(googleId)
	if err != nil {
		return nil
	}
	return user
}

// FindByGithubId retrieves a user by their GitHub ID.
func (service *UserService) FindByGithubId(githubId int64) *models.User {
	if githubId < 0 {
		return nil
	}
	user, err := service.repo.FindByGithubID(githubId)
	if err != nil {
		return nil
	}
	return user
}

// FindByIncompleteUsername searches for users with a username that contains the provided string.
func (service *UserService) FindByIncompleteUsername(username string) *models.User {
	if username == "" {
		return nil
	}
	query := "%" + username + "%"
	user, err := service.repo.FindByUsername(&query)
	if err != nil {
		return nil
	}
	return user
}

// GetSession retrieves the session associated with a user.
func (service *UserService) GetSession(user *models.User) *models.Session {
	if user == nil || user.ID == uuid.Nil {
		return nil
	}
	session, err := service.sessionRepo.FindByUserID(user.ID)
	if err != nil {
		return nil
	}
	user.Session = *session
	return session
}

// GetRole retrieves the role associated with a user.
func (service *UserService) GetRole(user *models.User) *models.Role {
	if user == nil || user.ID == uuid.Nil {
		return nil
	}
	role, err := service.roleRepo.FindById(user.Role_ID)
	if err != nil {
		return nil
	}
	user.Role = *role
	return role
}

// GetAllUsers retrieves all users from the repository.
func (service *UserService) GetAllUsers() ([]*models.User, error) {

	return service.repo.GetAllUsers()
}

// IsEmailAlreadyUse checks if the provided email is already in use by another user.
func (service *UserService) IsEmailAlreadyUse(email string) bool {
	if email == "" {
		return false
	}
	user, err := service.repo.FindByEmail(&email)
	if err != nil || user == nil {
		return false
	}
	return true
}

// IsUsernameAlreadyUse checks if the provided username is already in use by another user.
func (service *UserService) IsUsernameAlreadyUse(username string) bool {
	if username == "" {
		return false
	}
	user, err := service.repo.FindByUsername(&username)
	if err != nil || user == nil {
		return false
	}
	return true
}

// Update updates an existing user in the repository.
func (service *UserService) Update(user *models.User) bool {
	if user == nil || user.ID == uuid.Nil {
		return false
	}
	// If no role is set, set the default role ID from environment variable
	if user.Role_ID == uuid.Nil && user.Role.ID == uuid.Nil {
		if os.Getenv("DEFAULT_ROLE_ID") == "" {
			return false // Default role ID is not set in environment variables
		}
		// Set the default role ID from environment variable
		roleIdString := os.Getenv("DEFAULT_ROLE_ID")
		roleId, err := uuid.Parse(roleIdString)
		if err != nil {
			return false // Invalid UUID format in environment variable
		}
		user.Role_ID = roleId
	} else if user.Role_ID == uuid.Nil && user.Role.ID != uuid.Nil { // If Role_ID is not set but Role object is provided
		user.Role_ID = user.Role.ID // Use the role ID from the Role object if available
	}

	err := service.repo.Update(user)

	return err == nil
}

// Create adds a new user to the repository.
func (service *UserService) Create(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, nil
	}
	user.ID = uuid.New()
	// If no role is set, set the default role ID from environment variable
	if user.Role_ID == uuid.Nil && user.Role.ID == uuid.Nil {
		if os.Getenv("DEFAULT_ROLE_ID") == "" {
			return nil, nil // Default role ID is not set in environment variables
		}
		// Set the default role ID from environment variable
		roleIdString := os.Getenv("DEFAULT_ROLE_ID")
		roleId, err := uuid.Parse(roleIdString)
		if err != nil {
			return nil, err // Invalid UUID format in environment variable
		}
		user.Role_ID = roleId
	} else if user.Role_ID == uuid.Nil && user.Role.ID != uuid.Nil { // If Role_ID is not set but Role object is provided
		user.Role_ID = user.Role.ID // Use the role ID from the Role object if available
	}

	err := service.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateFromGoogle creates a new user from Google user information.
func (service *UserService) CreateFromGoogle(userInfo *dtos.GoogleUserInfo) (*models.User, error) {
	profilePictureBlob, err := utils.FetchImage(userInfo.Picture)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Pseudo:    userInfo.Name,
		Email:     userInfo.Email,
		Google_ID: &userInfo.ID,
		Avatar:    profilePictureBlob,
		CreatedAt: time.Now(),
	}

	user, err = service.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateFromGithub creates a new user from GitHub user information.
func (service *UserService) CreateFromGithub(userInfo *dtos.GitHubUserInfo) (*models.User, error) {
	profilePictureBlob, err := utils.FetchImage(userInfo.AvatarURL)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Pseudo:    userInfo.Login,
		Email:     userInfo.Emails[0].Email, // Assuming the first email is the primary one
		Github_ID: &userInfo.ID,
		Avatar:    profilePictureBlob,
		Bio:       userInfo.Bio,
		CreatedAt: time.Now(),
	}

	user, err = service.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CheckUserWithSameGithubEmails checks if there are existing users with the same GitHub emails.
func (service *UserService) CheckUserWithSameGithubEmails(userInfo *dtos.GitHubUserInfo) (user *models.User, _ *string) {

	if userInfo == nil {
		return nil, nil
	}

	// Filter userInfo.Emails to make the primary one first
	for i, email := range userInfo.Emails {
		if email.Primary {
			// Move the primary email to the first position
			userInfo.Emails[0], userInfo.Emails[i] = userInfo.Emails[i], userInfo.Emails[0]
			break
		}
	}

	for _, e := range userInfo.Emails {
		if e.Email == "" {
			continue
		}
		user, err := service.repo.FindByEmail(&e.Email)
		if err != nil {
			continue
		}
		if user != nil {
			user.Github_ID = &userInfo.ID
			if !service.Update(user) {
				return user, nil
			}
			return user, &e.Email
		}
	}
	return nil, nil
}

// GetUserCount retrieves the total number of users in the repository.
func (service *UserService) GetUserCount() (int, error) {
	count := service.repo.GetUserCount()
	if count == -1 {
		return -1, fmt.Errorf("failed to get user count")
	}
	return count, nil
}

// Delete removes a user from the repository.
func (service *UserService) Delete(user *models.User) (bool, error) {
	if user == nil || user.ID == uuid.Nil {
		return false, fmt.Errorf("user cannot be nil or have an empty ID")
	}
	err := service.repo.Delete(&user.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsAdmin checks if a user has admin privileges.
func (service *UserService) IsAdmin(user *models.User) bool {
	if user == nil || user.ID == uuid.Nil {
		return false
	}
	role, err := service.roleRepo.FindById(user.Role_ID)
	if err != nil || role == nil {
		log.Println(err)
		return false
	}
	if r, err := service.roleRepo.FindHighestPermRole(); err == nil && r.ID == role.ID {
		return true
	}
	return false
}

// IsModerator checks if a user has moderator privileges.
func (service *UserService) IsModerator(user *models.User) bool {
	if user == nil || user.ID == uuid.Nil {
		return false
	}
	role, err := service.roleRepo.FindById(user.Role_ID)
	if err != nil || role == nil {
		return false
	}
	if r, err := service.roleRepo.FindMidPermRole(); err == nil && r.ID == role.ID {
		return true
	}
	return false
}

// IsAdminOrModerator checks if a user is either an admin or a moderator.
func (service *UserService) IsAdminOrModerator(user *models.User) bool {
	if user == nil || user.ID == uuid.Nil {
		return false
	}
	return service.IsAdmin(user) || service.IsModerator(user)
}
