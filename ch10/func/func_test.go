package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ReturnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)

}

func timeSpent(inner func(int) int) func(int) int {
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
	a, b := ReturnMultiValues()
	t.Log(a, b)

	tsSF := timeSpent(slowFn)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	s := 0

	for _, op := range ops {
		s += op
	}
	return s

}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clean resources")
}

func TestDefer(t *testing.T) {
	defer Clear()
	t.Log("Started")
	panic("Fatal error")
}
