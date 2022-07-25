package storage

type Storage interface {
	Init() error
	Delete(filename string) error
	FilePath(filename string) string
}
