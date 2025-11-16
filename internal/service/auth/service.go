package auth

import (
	"context"
	"encoding/json"
	"errors"
	"exptracker/internal/cache"
	"exptracker/internal/config"
	"exptracker/internal/domain/dto/auth"
	models "exptracker/internal/domain/model"
	"exptracker/internal/logger"
	repository "exptracker/internal/repository/auth"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessTokenTTL  = time.Hour * 1
	refreshTokenTTL = time.Hour * 24 * 7
)

type Service interface {
	Register(req auth.RegisterRequest, ip, userAgent string) (*auth.AuthResponse, error)
	Login(req auth.LoginRequest, ip, userAgent string) (*auth.AuthResponse, error)
	Refresh(req auth.RefreshRequest, ip, userAgent string) (*auth.AuthResponse, error)
}

type authService struct {
	repo repository.Repository
	cfg  config.Config
}

type RefreshSession struct {
	UserID    string `json:"user_id"`
	TokenID   string `json:"token_id"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	ExpiresAt int64  `json:"expires_at"`
}

func NewAuthService(repo repository.Repository, cfg config.Config) Service {
	return &authService{repo: repo, cfg: cfg}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (s *authService) generateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(accessTokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.App.JWTSecret))
}

func (s *authService) generateRefreshToken(userID, tokenID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"jti":     tokenID, // token ID
		"exp":     time.Now().Add(refreshTokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.App.JWTSecret))
}

func (s *authService) storeRefreshSession(ctx context.Context, sess RefreshSession) error {
	if cache.Client == nil {
		return errors.New("redis not initialized")
	}

	key := "refresh:" + sess.UserID + ":" + sess.TokenID
	data, err := json.Marshal(sess)
	if err != nil {
		return err
	}

	ttl := time.Until(time.Unix(sess.ExpiresAt, 0))
	return cache.Client.Set(ctx, key, data, ttl).Err()
}

func (s *authService) deleteRefreshSession(ctx context.Context, userID, tokenID string) {
	if cache.Client == nil {
		return
	}
	key := "refresh:" + userID + ":" + tokenID
	if err := cache.Client.Del(ctx, key).Err(); err != nil {
		logger.L().Error().Err(err).Msg("Failed to delete refresh session")
	}
}

func (s *authService) buildAuthResponse(ctx context.Context, userID, ip, userAgent string) (*auth.AuthResponse, error) {
	accessToken, err := s.generateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	tokenID := uuid.NewString()
	refreshToken, err := s.generateRefreshToken(userID, tokenID)
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(refreshTokenTTL).Unix()

	sess := RefreshSession{
		UserID:    userID,
		TokenID:   tokenID,
		IP:        ip,
		UserAgent: userAgent,
		ExpiresAt: expiresAt,
	}

	if err := s.storeRefreshSession(ctx, sess); err != nil {
		return nil, err
	}

	return &auth.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) Register(req auth.RegisterRequest, ip, userAgent string) (*auth.AuthResponse, error) {
	_, err := s.repo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	hashed, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    req.Email,
		Password: hashed,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.buildAuthResponse(ctx, user.ID, ip, userAgent)
}

func (s *authService) Login(req auth.LoginRequest, ip, userAgent string) (*auth.AuthResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !checkPassword(user.Password, req.Password) {
		return nil, errors.New("invalid credentials")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.buildAuthResponse(ctx, user.ID, ip, userAgent)
}

func (s *authService) Refresh(req auth.RefreshRequest, ip, userAgent string) (*auth.AuthResponse, error) {
	// 1. Parse refresh token
	token, err := jwt.Parse(req.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.cfg.App.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid refresh token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("invalid refresh token user")
	}

	tokenID, ok := claims["jti"].(string)
	if !ok {
		return nil, errors.New("invalid refresh token id")
	}

	// 2. Lookup in Redis
	if cache.Client == nil {
		return nil, errors.New("redis not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := "refresh:" + userID + ":" + tokenID
	val, err := cache.Client.Get(ctx, key).Result()
	if err != nil {
		// Token not found => reused / revoked / expired
		// At this point you might choose to revoke all sessions for this user
		// but for now just treat as invalid.
		return nil, errors.New("refresh token expired or revoked")
	}

	var sess RefreshSession
	if err := json.Unmarshal([]byte(val), &sess); err != nil {
		return nil, errors.New("invalid session data")
	}

	// 3. Check IP + User-Agent match
	if sess.IP != ip || sess.UserAgent != userAgent {
		// Optionally: delete all sessions for userID to force logout everywhere
		// scanKey := "refresh:" + userID + ":*"
		// (left as future enhancement)
		return nil, errors.New("device mismatch")
	}

	// 4. Rotate token: delete old, issue new
	s.deleteRefreshSession(ctx, userID, tokenID)

	return s.buildAuthResponse(ctx, userID, ip, userAgent)
}
