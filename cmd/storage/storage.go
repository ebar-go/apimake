package storage

import (
	"encoding/gob"
	"os"
)

func Write(filename string, obj interface{}) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	enc := gob.NewEncoder(file)
	if err := enc.Encode(obj); err != nil {
		return err
	}

	return nil
}

func Read(filename string, obj interface{}) error {
	file, _ := os.Open(filename)
	return gob.NewDecoder(file).Decode(obj)
}
