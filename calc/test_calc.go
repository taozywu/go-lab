package calc

// "fmt"

/*
var m = calc.Test11(1, "ww", 30)
	fmt.Println(*m)

	fmt.Println(m.Test22(), m.Age)
*/

type person struct {
	id   int
	name string
	Age  int
}

// person结构体只能在test_calc.go中使用，如果想外部使用则需要写一个func来调用。
func Test11(a int, b string, c int) *person {
	return &person{
		id:   a,
		name: b,
		Age:  c,
	}
}

// name是小写，则外部不能直接调用，同理想调用则需要一个func来处理。
func (m person) Test22() string {
	f := m.name
	return f
}
