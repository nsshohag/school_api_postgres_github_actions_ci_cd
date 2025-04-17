package validation

import (
	"errors"
	"school_api_postgres/models"
)

func ValidateStudent(s models.Student) error {

	if s.Name == "" {
		return errors.New("name is required")
	}

	if s.Age == 0 {
		return errors.New("provide age")
	}

	if s.Age < 0 || s.Age > 120 {

		if s.Age < 0 {
			return errors.New("age must be a positive number")
		}
		if s.Age > 120 {
			return errors.New("age is too high")
		}
	}

	if s.Class == 0 {
		return errors.New("provide class")
	}

	if s.Class < 1 || s.Class > 10 {
		if s.Class < 0 {
			return errors.New("class must be a positive number")
		}
		if s.Class > 10 {
			return errors.New("class is too high")
		}
	}

	return nil
}

// validateAge checks if the age is valid
func ValidateAge(age int) error {
	if age < 0 {
		return errors.New("age must be a positive number")
	}
	if age > 120 {
		return errors.New("age is too high")
	}
	return nil
}

// validateClass checks if the class is valid
func ValidateClass(class int) error {
	if class < 0 {
		return errors.New("class must be a positive number")
	}
	if class > 10 {
		return errors.New("class is too high")
	}
	return nil
}
