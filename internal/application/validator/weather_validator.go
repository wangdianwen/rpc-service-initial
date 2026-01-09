package validator

import (
	"strings"
)

func ValidateGetCurrentWeatherRequest(city string, lat, lon float64) *ValidationResult {
	var errors []string

	if city == "" && (lat == 0 || lon == 0) {
		errors = append(errors, "either city name or coordinates (lat, lon) must be provided")
	}

	if lat < -90 || lat > 90 {
		errors = append(errors, "latitude must be between -90 and 90")
	}

	if lon < -180 || lon > 180 {
		errors = append(errors, "longitude must be between -180 and 180")
	}

	if city != "" && len(city) > 100 {
		errors = append(errors, "city name exceeds maximum length of 100 characters")
	}

	return &ValidationResult{
		Valid:  len(errors) == 0,
		Errors: errors,
	}
}

func ValidateGetForecastRequest(city string, lat, lon float64, days int) *ValidationResult {
	var errors []string

	result := ValidateGetCurrentWeatherRequest(city, lat, lon)
	if !result.Valid {
		return result
	}

	if days < 1 {
		days = 5
	} else if days > 7 {
		errors = append(errors, "days must be between 1 and 7")
	}

	return &ValidationResult{
		Valid:  len(errors) == 0,
		Errors: errors,
	}
}

func ValidateCityName(city string) *ValidationResult {
	var errors []string

	if city == "" {
		errors = append(errors, "city name cannot be empty")
	}

	city = strings.TrimSpace(city)
	if len(city) < 2 {
		errors = append(errors, "city name must be at least 2 characters")
	}

	if len(city) > 100 {
		errors = append(errors, "city name exceeds maximum length of 100 characters")
	}

	return &ValidationResult{
		Valid:  len(errors) == 0,
		Errors: errors,
	}
}
