package say_hello

import (
	"errors"
	"fmt"
)

var (
	VALIDATION_ERROR = func(msg string) error {
		return errors.New("Validation Error: " + msg)
	}
)

func SayHello(name string) (string, error) {
	if name == "" {
		return "", VALIDATION_ERROR(("Name cannot be nil"))
	}
	return fmt.Sprintf("Hello There %s", name), nil
}
