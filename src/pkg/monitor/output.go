package monitor

import (
	"context"
	"encoding/gob"
	"log"
	"os"
	"parse/src/pkg/tmp"
	"sync"
)

type FileOutput[T Message] struct {
	mu   sync.Mutex
	file *os.File

	dir tmp.Direction
	wb  *gob.Encoder
	//recorder Recorder[T]
}

func (f *FileOutput[T]) Write(ctx context.Context, msg T, dir tmp.Direction) error {
	data, err := msg.ToJSON()
	if err != nil {
		log.Println(err)
		return err
	}
	err = f.wb.Encode(data)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (f *FileOutput[T]) Close() error {
	//TODO implement me
	panic("implement me")
}

func NewFileOutput[T Message](path string) (*FileOutput[T], error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &FileOutput[T]{
		file: f,
		wb:   gob.NewEncoder(f),
	}, nil
}
