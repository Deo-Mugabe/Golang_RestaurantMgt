package utility

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Utility to parse ID from URL
func ParseID(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// Utility for JSON response
func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// WithTimeoutContext creates a context with a timeout for database operations.
func WithTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

// ValidateStruct validates a struct's fields, ensuring required fields are not empty
func ValidateStruct(model interface{}) error {
	v := reflect.ValueOf(model)
	// Ensure the input is a pointer to a struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("invalid input: expected a pointer to a struct")
	}

	// Loop through all fields of the struct and validate based on a basic check
	for i := 0; i < v.Elem().NumField(); i++ {
		field := v.Elem().Field(i)
		fieldName := v.Type().Field(i).Name

		// If the field is of type string and is empty, return an error
		if field.Kind() == reflect.String && field.String() == "" {
			return errors.New("field " + fieldName + " cannot be empty")
		}
		// Add more type-specific validations as needed
	}

	return nil
}
