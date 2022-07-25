package img

import (
	"fmt"
	"os"
	"path/filepath"
)

type Storage struct {
	DirPath string
}

func NewStorage(path string) *Storage {
	return &Storage{
		DirPath: path,
	}
}

func (s *Storage) Init() error {
	return os.MkdirAll(s.DirPath, 0750)
}

func (s *Storage) Delete(filename string) error {
	if err := os.Remove(s.FilePath(filename)); err != nil {
		return fmt.Errorf("Failed to remove a file: %v", err)
	}

	return nil
}

func (s *Storage) FilePath(filename string) string {
	return filepath.Join(s.DirPath, filename)
}
