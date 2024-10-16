package rowbinary

import (
	"time"

	"github.com/pkg/errors"
)

var DateTime Type[time.Time] = &typeDateTime{}

type typeDateTime struct {
}

func (t *typeDateTime) String() string {
	return "DateTime"
}

func (t *typeDateTime) Write(w Writer, value time.Time) error {
	if value.Year() < 1970 {
		return UInt32.Write(w, 0)
	}
	return UInt32.Write(w, uint32(value.Unix()))
}

func (t *typeDateTime) Read(r Reader) (time.Time, error) {
	n, err := UInt32.Read(r)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(n), 0).UTC(), nil
}

func (t *typeDateTime) WriteAny(w Writer, v any) error {
	value, ok := v.(time.Time)
	if !ok {
		return errors.New("unexpected type")
	}
	return t.Write(w, value)
}

func (t *typeDateTime) ReadAny(r Reader) (any, error) {
	return t.Read(r)
}
