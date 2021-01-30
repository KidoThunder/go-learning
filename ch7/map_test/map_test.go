package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[string]int{
		"a": 1, "b": 2,
	}
	t.Log(m1["a"])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestTAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])

	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not existing.")
	}

	m1[3] = 0

	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not existing.")
	}

}

func TestMapTravel(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 3, 3: 4}

	for k, v := range m1 {
		t.Log(k, v)
	}

}
