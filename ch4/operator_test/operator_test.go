package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	c := [...]int{1, 3, 4, 2}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b, a == c, a == d)
}

func TestBitClear(t *testing.T) {
	// &^ 按位清零操作符，若右边位的值为1，则最终结果为0; 若右边的位值为0，则结果等于左边的值
	// 例如：
	// 1 &^ 1 = 0
	// 1 &^ 0 = 1
	// 0 &^ 0 = 0
	// 0 &^ 1 = 0

	a := 7              // 0111
	a = a &^ Readable   // 去除可读
	a = a &^ Executable // 去除可执行
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
	t.Log(a)

}
