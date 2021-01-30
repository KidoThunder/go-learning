package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}

	t.Log(arr1)
	t.Log(arr[1], arr[2])
	t.Log(arr2)
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	for _, e := range arr2 {
		t.Log(e)
	}

}

func TestMutiArray(t *testing.T) {
	arr1 := [2][3]int{{1, 2, 3}, {3, 4, 5}}
	t.Log(arr1)
}

func TestArraySection(t *testing.T) {
	arr1 := [2][3]int{{1, 2, 3}, {3, 4, 5}}
	t.Log(arr1[:2])

	sec1 := arr1[1:]
	t.Log(sec1)

	t.Log(arr1[:])

}
