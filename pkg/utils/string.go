package utils

import uuid "github.com/satori/go.uuid"

func UUID() string {
	return uuid.NewV4().String()
}

func Str2Byte(str string) []byte {
	return []byte(str)
}

func Byte2Str(bytes []byte) string {
	return string(bytes)
}
