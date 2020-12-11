package models

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ImageService is a set of methods used to manipulate and
// work with the images model.
type ImageService interface {
	Create(galleryID uint, r io.Reader, filename string) error
}

type imageService struct{}

// NewImageService is used to create a new image service.
func NewImageService() ImageService {
	return &imageService{}
}

func (is *imageService) Create(galleryID uint, r io.Reader, filename string) error {
	path, err := is.mkImagePath(galleryID)
	if err != nil {
		return err
	}

	// Create Destination File
	dst, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy Reader Data to Destination File
	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}

	return nil
}

func (is *imageService) mkImagePath(galleryID uint) (string, error) {
	galleryPath := filepath.Join("images", "galleries", fmt.Sprintf("%v", galleryID))
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", err
	}
	return galleryPath, nil
}
