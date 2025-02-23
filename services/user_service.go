package services

import (
	"errors"
	"time"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Secret key for signing JWTs - This should be securely stored (e.g., in an environment variable)
var jwtSecret = []byte("your-secret-key")

// GetUsers returns all users
func GetUsers() ([]models.User, error) {
	var users []models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUser fetches a user by ID
func GetUser(id uint) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail fetches a user by email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		// Check if the error is due to the record not being found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // No user found with the given email
		}
		// Return any other errors that occurred
		return nil, result.Error
	}
	return &user, nil
}

// CreateUser creates a new user after validating and hashing the password
func CreateUser(user *models.User) error {
	// Validate user data
	if err := validateUserData(user); err != nil {
		return err
	}

	// Hash the user's password
	if err := HashPassword(user); err != nil {
		return err
	}

	// Set timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Generate a unique UserID
	user.UserID = uuid.New().String()

	// Save user to database
	result := db.DB.Create(user)
	return result.Error
}

// UpdateUser updates an existing user's details
func UpdateUser(id uint, updatedUser *models.User) error {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	// Update fields only if they are provided
	if updatedUser.FirstName != "" {
		user.FirstName = updatedUser.FirstName
	}
	if updatedUser.LastName != "" {
		user.LastName = updatedUser.LastName
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Avatar != "" {
		user.Avatar = updatedUser.Avatar
	}
	if updatedUser.Phone != "" {
		user.Phone = updatedUser.Phone
	}
	if updatedUser.Password != "" {
		// Hash new password before saving
		if err := HashPassword(&user); err != nil {
			return err
		}
	}

	// Update timestamp
	user.UpdatedAt = time.Now()

	// Save the updated user
	result = db.DB.Save(&user)
	return result.Error
}

// DeleteUser removes a user by ID
func DeleteUser(id uint) error {
	result := db.DB.Delete(&models.User{}, id)
	return result.Error
}

// HashPassword hashes the password for secure storage
func HashPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// ValidateUserData validates that the user's data is correct (email, phone, etc.)
func validateUserData(user *models.User) error {
	// Ensure required fields are present
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing required fields")
	}

	// Check if email or phone already exists
	if err := checkEmailAndPhone(user); err != nil {
		return err
	}

	return nil
}

// Check if email or phone already exists in the database
func checkEmailAndPhone(user *models.User) error {
	var existingUser models.User

	// Check if email already exists
	if err := db.DB.First(&existingUser, "email = ?", user.Email).Error; err == nil {
		return errors.New("email already in use")
	}

	// Check if phone number already exists
	if err := db.DB.First(&existingUser, "phone = ?", user.Phone).Error; err == nil {
		return errors.New("phone number already in use")
	}

	return nil
}

// GenerateTokens generates both the access token and the refresh token
func GenerateTokens(user models.User) (string, string, error) {
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
	accessTokenString, err := accessToken.SignedString([]byte("your-secret-key"))
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
	refreshTokenString, err := refreshToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", "", err
	}

	// Return both tokens
	return accessTokenString, refreshTokenString, nil
}

// VerifyPassword compares the provided password with the stored hashed password
func VerifyPassword(storedPassword, providedPassword string) (bool, error) {
	// Compare the provided password with the stored password (hashed)
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	if err != nil {
		// If passwords don't match, return false and the error
		return false, errors.New("invalid password")
	}
	// If passwords match, return true
	return true, nil
}
