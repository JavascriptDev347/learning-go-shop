package services

import (
	"fmt"
	"mime/multipart"
	"path"
	"strings"

	"github.com/JavascriptDev347/learning-go-shop/internal/interfaces"
)

type LocalUploadService struct {
	provider interfaces.UploadProvider
}

func NewLocalUploadService(provider interfaces.UploadProvider) *LocalUploadService {
	return &LocalUploadService{
		provider: provider,
	}
}

func (s *LocalUploadService) UploadProductImage(productID uint, file *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(path.Ext(file.Filename))
	if !isValidImageExt(ext) {
		return "", fmt.Errorf("invalid file extension: %s", ext)
	}

	path := fmt.Sprintf("products/%d/%s", productID, file.Filename)
	return s.provider.UploadFile(file, path)
}

func isValidImageExt(ext string) bool {
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			return true
		}
	}
	return false
}
