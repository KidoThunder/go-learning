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

func TestXx(t *testing.T){
	s1 := []int{1, 2, 3, 4}
	s2 := s1[1:2]
	fmt.Println(s2)
	s2 = append(s2, 5)
	fmt.Println(s2)
	for _, v := range s2 {
		v += 10
	}
	fmt.Println(s2)
	fmt.Println(s1)
	for i := range s2 {
		s2[i] += 10
	}
	fmt.Println(s1)
	fmt.Println(s2)
}