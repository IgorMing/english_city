package utils

import "fmt"

func HandleError(message string, err error) error {
	if message != "" {
		return fmt.Errorf("%s Details: %s", message, err.Error())
	}

	return fmt.Errorf("Unexpected error. Details: %s", err.Error())
}
