package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2, len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(s, len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	d := []string{
		"A", "B", "C",
		"D", "E", "F",
		"G", "H", "I",
	}
	s1 := d[3:5]
	t.Log(s1, len(s1), cap(s1))
	s2 := d[4:6]
	t.Log(s2, len(s2), cap(s2))

	s2[0] = "Z"
	t.Log(s1)
	t.Log(d)

}

//func TestSliceComparing(t *testing.T) {
//	a := []int{1, 2, 3, 4}
//	b := []int{1, 2, 3, 4}
//	if a == b {
//		t.Log("Equal!")
//	}
//}
