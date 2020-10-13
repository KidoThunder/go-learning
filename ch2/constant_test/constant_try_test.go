package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestDayConstant(t *testing.T)  {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}

func TestReadableConstant(t *testing.T)  {
	t.Log(Readable, Writeable, Executable)
	a := 7 // 0111
	t.Log(a & Readable == Readable, a & Writeable == Writeable, a & Executable == Executable)

	b := 1 // 0001
	t.Log(b & Readable == Readable, b & Writeable == Writeable, b & Executable == Executable)

}
