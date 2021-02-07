package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(int) int

func timeSpent(inner IntConv) IntConv {
	return func(i int) int {
		start := time.Now()
		ret := inner(i)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFn(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFn)
	t.Log(tsSF(10))
}