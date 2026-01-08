package validator

type ValidationResult struct {
	Valid  bool
	Errors []string
}

func ValidateServiceName(name string) *ValidationResult {
	var errors []string

	if name == "" {
		errors = append(errors, "service name cannot be empty")
	}

	if len(name) > 100 {
		errors = append(errors, "service name exceeds maximum length of 100 characters")
	}

	return &ValidationResult{
		Valid:  len(errors) == 0,
		Errors: errors,
	}
}

func ValidateServiceData(data string) *ValidationResult {
	var errors []string

	if len(data) > 10000 {
		errors = append(errors, "service data exceeds maximum length of 10000 characters")
	}

	return &ValidationResult{
		Valid:  len(errors) == 0,
		Errors: errors,
	}
}
