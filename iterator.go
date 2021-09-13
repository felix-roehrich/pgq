package pgq

import (
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgtyp"
	"github.com/Alisaien/pgq/pgval"
	"github.com/jackc/pgtype"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type Iterator pgetc.Iterator

func NewIterator(src []byte) *Iterator {
	return (*Iterator)(pgetc.NewIterator(src))
}

func (iter *Iterator) Iterator() *pgetc.Iterator {
	return (*pgetc.Iterator)(iter)
}

func (iter *Iterator) Err() error {
	return iter.Iterator().Err()
}

func (iter *Iterator) ReportError(err error) {
	iter.Iterator().ReportError(err)
}

func (iter *Iterator) ReadBool() bool {
	return pgtyp.Bool.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadBoolPtr() *bool {
	return pgtyp.BoolPtr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadCompositeTypeHeader() uint32 {
	if err := (*pgetc.Iterator)(iter).Next4(); err != nil {
		return 0
	}
	return pgbin.Uint32.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadEnum(oid pgetc.OID) string {
	id := iter.Iterator().ReadUint32()
	if id == 0 {
		return ""
	} else if id != uint32(oid) {
		iter.ReportError(pgetc.ErrUnexpectedType)
	}

	return pgval.String.Read(iter.Iterator())
}

func (iter *Iterator) ReadInt() int {
	return pgtyp.Int.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadIntPtr() *int {
	return pgtyp.IntPtr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadInt16() int16 {
	return pgtyp.Int16.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadInt16Ptr() *int16 {
	return pgtyp.Int16Ptr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadJSONB(v interface{}) {
	oid := iter.Iterator().ReadUint32()
	if oid != pgtype.JSONBOID {
		return
	}
	size := int32(iter.Iterator().ReadUint32())

	var (
		err error
		data []byte
	)
	if size != -1 {
		if err = iter.Iterator().Next(int(size)); err != nil {
			return
		}
		data = iter.Iterator().Read()
	}

	if err = jsoniter.Unmarshal(data, v); err != nil {
		iter.ReportError(err)
	}
}

func (iter *Iterator) ReadString() string {
	return pgtyp.String.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadStringPtr() *string {
	return pgtyp.StringPtr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadTimestamptz() time.Time {
	return pgtyp.Timestamptz.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadUUID() pgetc.UUID {
	return pgtyp.UUID.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadUUIDPtr() *pgetc.UUID {
	return pgtyp.UUIDPtr.Read((*pgetc.Iterator)(iter))
}
