package errors

import "fmt"

func InvalidType() error {
	return fmt.Errorf("invalid type")
}

func InvalidParam() error {
	return fmt.Errorf("invalid param")
}

func NotFound() error {
	return fmt.Errorf("not found")
}
