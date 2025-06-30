package monitor

import (
	"os"
	"sync"
)

type FileOutput[T Message] struct {
	mu   sync.Mutex
	file *os.File

	dir Direction

	//recorder Recorder[T]
}

func NewFileOutput[T Message](path string) (*FileOutput[T], error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &FileOutput[T]{
		file: f,
	}, nil
}
