package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	employee1 := Employee{"1", "Employee1", 18}
	t.Log(employee1.ID, employee1.Name, employee1.Age)

	employee2 := Employee{ID: "2", Name: "Employee2", Age: 19}
	t.Log(employee2.ID, employee2.Name, employee2.Age)

	employee3 := new(Employee)
	employee3.ID = "3"
	employee3.Name = "employee3"
	employee3.Age = 20
	t.Log(employee3.ID, employee3.Name, employee3.Age)

	t.Logf("e1 is %T", employee1)
	t.Logf("e3 is %T", employee3)

}

func (e Employee) String() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

func (e *Employee) NewEmployeeString() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

func TestStructureOperations(t *testing.T)  {
	employee1 := Employee{"1", "Employee1", 18}
	t.Log(employee1.String())
	fmt.Printf("Address is %x", unsafe.Pointer(&employee1.Name))


	employee2 := &Employee{"2", "Employee2", 19}
	fmt.Printf("Address is %x", unsafe.Pointer(&employee2.Name))
	t.Log(employee2.NewEmployeeString())

}



