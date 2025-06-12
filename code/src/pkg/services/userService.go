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

func (service *UserService) FindById(id uuid.UUID) *models.User {
	user, err := service.repo.FindById(id)
	if err != nil {
		return nil
	}
	return user
}

func (service *UserService) FindByEmail(email string) *models.User {
	user, err := service.repo.FindByEmail(&email)
	if err != nil {
		return nil
	}
	return user
}

func (service *UserService) FindByUsername(username string) *models.User {
	user, err := service.repo.FindByUsername(&username)
	if err != nil {
		return nil
	}
	return user
}

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

func (service *UserService) GetUserCount() (int, error) {
	count := service.repo.GetUserCount()
	if count == -1 {
		return -1, fmt.Errorf("failed to get user count")
	}
	return count, nil
}

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

func (service *UserService) IsAdminOrModerator(user *models.User) bool {
	if user == nil || user.ID == uuid.Nil {
		return false
	}
	return service.IsAdmin(user) || service.IsModerator(user)
}
