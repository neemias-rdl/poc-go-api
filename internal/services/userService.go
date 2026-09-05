package services

import (
	"api/internal/data/entities"
	"api/internal/data/repositories"
	"api/internal/dto"

	"golang.org/x/crypto/bcrypt"

	"context"
	"net/http"
	"unicode"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

func (s *UserService) Register(ctx context.Context, registerDto dto.Register) dto.Response {
	var response dto.Response
	existingUser, err := s.userRepository.FindByUsername(ctx, registerDto.Username)

	if err != nil {
		response.Error = ("Error fetching user data")
		response.ErrorCode = http.StatusBadRequest
		return response
	}

	if existingUser != nil {
		response.Error = "Username already in use"
		response.ErrorCode = http.StatusBadRequest
		return response
	}

	validPass := validatePassword(registerDto.Password)

	if !validPass {
		response.Error = "Invalid password"
		response.ErrorCode = http.StatusBadRequest
		return response
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)

	s.userRepository.CreateUser(
		ctx,
		*entities.NewUser(
			registerDto.Username,
			string(hashedPass),
			"",
		),
	)

	return response
}

// Verify if password passes needed checks
// 1. 8 characters minimum
// 2. at least one special character
// 3. at least one number
// 4. at least one capitalized letter
func validatePassword(pass string) bool {
	if len(pass) < 8 {
		return false
	}

	hasSpecialCharacter := false
	hasNumber := false
	hasCapitalChar := false

	for _, c := range pass {
		switch {
		case unicode.IsUpper(c):
			hasCapitalChar = true
		case unicode.IsNumber(c):
			hasNumber = true
		case !unicode.IsLetter(c):
			hasSpecialCharacter = true
		}
	}

	return hasCapitalChar && hasNumber && hasSpecialCharacter
}
