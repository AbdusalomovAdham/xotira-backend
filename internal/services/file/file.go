package file

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error) {
	if file == nil {
		return "", errors.New("file not provided")
	}

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return "", err
	}

	dst := filepath.Join(folder, file.Filename)

	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err := io.Copy(out, f); err != nil {
		return "", err
	}

	return dst, nil
}

func (s Service) Delete(ctx context.Context, url string) error {
	err := os.Remove("./" + url)
	return err
}

func (s Service) MultipleUpload(ctx context.Context, file []*multipart.FileHeader, folder string) ([]string, error) {
	var links []string

	for _, f := range file {
		link, err := s.Upload(ctx, f, folder)

		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}
