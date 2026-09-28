package workspace

type Workspace interface {
	Walk(dir string, fn func(path string, size int64) error) error

	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error

	Mkdir(path string) error
	Remove(path string) error

	Copy(src, dst string) error

	Close() error
}
