package pqtype

import (
	"encoding/binary"
	"github.com/jackc/pgio"
)

type Int2 int16

const (
	Int2OID  = 21
	int2Size = 2
)

func (v *Int2) DecodeType(src []byte) ([]byte, error) {
	err := LenCheck(src, int4Size)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, Int4OID)
	if err != nil {
		return nil, err
	}

	return v.DecodeValue(src)
}

func (v *Int2) DecodeValue(src []byte) ([]byte, error) {
	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	return v.Read(src)
}

func (v *Int2) Read(src []byte) ([]byte, error) {
	*v = Int2(binary.BigEndian.Uint16(src))
	return src[int2Size:], nil
}

func (v Int2) EncodeType(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, Int2OID)
	return v.EncodeValue(buf)
}

func (v Int2) EncodeValue(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, int2Size)
	return v.Write(buf)
}

func (v Int2) Write(buf []byte) []byte {
	return pgio.AppendUint16(buf, uint16(v))
}

func DecodeInt2(src []byte) (*Int2, []byte, error) {
	err := LenCheck(src, 0)
	if err != nil {
		return nil, nil, err
	}

	src, err = TypeCheck(src, Int2OID)
	if err != nil {
		return nil, nil, err
	}

	size, src := ValueSize(src)
	if size == -1 {
		return nil, src, nil
	}

	v := new(Int2)
	src, _ = v.Read(src)
	return v, src, err
}
