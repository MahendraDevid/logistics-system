package service

import (
	"auth-service/internal/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo domain.UserRepository
}

func NewAuthService(r domain.UserRepository) domain.AuthService {
	return &authService{repo: r}
}

func (s *authService) Register(name, email, password, role string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}
	return s.repo.Create(user)
}

func (s *authService) Login(email, password string) (string, string, *domain.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", nil, errors.New("user tidak ditemukan")
	}

	// Cek Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", nil, errors.New("password salah")
	}

	// Generate JWT (Stateless sesuai spek)
	accessToken, _ := s.generateToken(user, time.Minute*15)
	refreshToken, _ := s.generateToken(user, time.Hour*24*7)

	return accessToken, refreshToken, user, nil
}

func (s *authService) generateToken(user *domain.User, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("UPI_SECRET_KEY"))
}

func (s *authService) ValidateToken(token string) (*domain.User, error) {
	// Logic validasi token untuk middleware API Gateway nanti
	return nil, nil
}
