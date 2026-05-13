package storage

import (
	"io"
	"mime/multipart"
	"os"
)

type LocalStorage struct {
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (s *LocalStorage) SaveFile(file multipart.File, path string) error {

	dst, err := os.Create(path)

	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)

	return err
}