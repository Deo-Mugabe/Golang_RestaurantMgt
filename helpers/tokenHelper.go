package helpers

import (
	"errors"
	"time"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/dgrijalva/jwt-go"
)

// Secret key for signing JWTs - This should be securely stored (e.g., in an environment variable)
var jwtSecret = []byte("your-secret-key")

// GenerateAllTokens generates both the access token and the refresh token
func GenerateAllTokens(user models.User) (string, string, error) {
	// Set expiration times for access token and refresh token
	accessTokenExpiry := time.Now().Add(15 * time.Minute).Unix()    // 15 minutes for access token
	refreshTokenExpiry := time.Now().Add(7 * 24 * time.Hour).Unix() // 7 days for refresh token

	// Create the JWT claims for access token
	accessClaims := jwt.MapClaims{
		"email": user.Email,
		"exp":   accessTokenExpiry,
	}
	// Create access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Create the JWT claims for refresh token
	refreshClaims := jwt.MapClaims{
		"email": user.Email,
		"exp":   refreshTokenExpiry,
	}
	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Return both tokens
	return accessTokenString, refreshTokenString, nil
}

// UpdateAllTokens generates a new access and refresh token when the user requests a refresh
func UpdateAllTokens(user models.User) (string, string, error) {
	// Set new expiration times
	accessTokenExpiry := time.Now().Add(15 * time.Minute).Unix()    // 15 minutes for access token
	refreshTokenExpiry := time.Now().Add(7 * 24 * time.Hour).Unix() // 7 days for refresh token

	// Generate the new access token
	accessClaims := jwt.MapClaims{
		"email": user.Email,
		"exp":   accessTokenExpiry,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Generate the new refresh token
	refreshClaims := jwt.MapClaims{
		"email": user.Email,
		"exp":   refreshTokenExpiry,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// ValidateToken validates the provided JWT access token
func ValidateToken(tokenString string) (*models.User, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method matches what we used for signing
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Validate the token claims (e.g., expiration)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, errors.New("Token expired")
			}
		}
		// Return the user from the claims (you may want to extract more fields)
		user := models.User{
			Email: claims["email"].(string),
		}
		return &user, nil
	}
	return nil, errors.New("Invalid token")
}
