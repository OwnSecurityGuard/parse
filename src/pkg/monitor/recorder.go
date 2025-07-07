package monitor

import (
	"encoding/gob"
	"os"
)

type MessageEncoder[T Message] interface {
	Encode(msg T) error
}

type MessageDecoder[T Message] interface {
	Decode() (msg T, err error)
}

type JSONCoder[T Message] struct {
	e *gob.Encoder
	d *gob.Decoder
}

func NewJSONCoder[T Message](path string) *JSONCoder[T] {
	f, _ := os.Open(path)
	jc := &JSONCoder[T]{}
	jc.d = gob.NewDecoder(f)
	return jc
}
func (j *JSONCoder[T]) Encode(msg T) error {
	data, err := msg.ToJSON()
	if err != nil {
		return err
	}

	return j.e.Encode(data)
}
func (j *JSONCoder[T]) Decode() (msg T, err error) {
	var d []byte
	err = j.d.Decode(&d)
	msg.FromJSON(d)

	return msg, err
}

//type JSONDecoder[T Message] struct {
//	newInstance func() T // 提供新 message 实例的构造函数
//}
//
//func NewJSONDecoder[T Message](ctor func() T) *JSONDecoder[T] {
//	return &JSONDecoder[T]{newInstance: ctor}
//}
//
//func (j *JSONDecoder[T]) Decode(r io.Reader) (T, Direction, error) {
//	var record EncodedRecord
//	err := json.NewDecoder(r).Decode(&record)
//	if err != nil {
//		var zero T
//		return zero, 0, err
//	}
//
//	msg := j.newInstance()
//	if err := msg.FromJSON(record.Payload); err != nil {
//		var zero T
//		return zero, 0, err
//	}
//
//	var dir Direction
//	switch record.Direction {
//	case "c2s":
//		dir = ClientToServer
//	case "s2c":
//		dir = ServerToClient
//	default:
//		return msg, 0, fmt.Errorf("unknown direction: %s", record.Direction)
//	}
//
//	return msg, dir, nil
//}
