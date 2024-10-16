package rowbinary

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

var UInt32 Type[uint32] = &typeUInt32{}

type typeUInt32 struct {
}

func (t *typeUInt32) String() string {
	return "UInt32"
}

func (t *typeUInt32) Write(w Writer, v uint32) error {
	binary.LittleEndian.PutUint32(w.buffer(), v)
	_, err := w.Write(w.buffer()[:4])
	return err
}

func (t *typeUInt32) Read(r Reader) (uint32, error) {
	_, err := io.ReadAtLeast(r, r.buffer()[:4], 4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(r.buffer()[:4]), nil
}

func (t *typeUInt32) WriteAny(w Writer, v any) error {
	value, ok := v.(uint32)
	if !ok {
		return errors.New("unexpected type")
	}
	return t.Write(w, value)
}

func (t *typeUInt32) ReadAny(r Reader) (any, error) {
	return t.Read(r)
}
