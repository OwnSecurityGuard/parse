package monitor

import (
	"io"
)

type MessageEncoder[T Message] interface {
	Encode(writer io.Writer, msg T, dir Direction) error
}

type MessageDecoder[T Message] interface {
	Decode(reader io.Reader) (msg T, dir Direction, err error)
}

//type JSONEncoder[T Message] struct{}
//
//func (j *JSONEncoder[T]) Encode(w io.Writer, msg T, dir Direction) error {
//	data, err := msg.ToJSON()
//	if err != nil {
//		return err
//	}
//
//	return json.NewEncoder(w).Encode(data)
//}
//
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
