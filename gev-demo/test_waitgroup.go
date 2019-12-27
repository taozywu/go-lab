package main

// import (
// 	"fmt"
// 	"sync"
// 	// "time"
// )

// type WaitGroupWrapper struct {
// 	sync.WaitGroup
// }

// func (w *WaitGroupWrapper) Wrap(cb func(argvs ...interface{}), argvs ...interface{}) {
// 	w.Add(1)
// 	go func() {
// 		cb(argvs...)
// 		w.Done()
// 	}()
// }

// type MyStruct struct {
// 	Age  int
// 	Name string
// }

// func GetAge(argvs ...interface{}) {
// 	my := argvs[1].(*MyStruct)
// 	my.Age = argvs[0].(int)
// }

// func main() {
// 	var wg WaitGroupWrapper
// 	var my MyStruct
// 	wg.Wrap(GetAge, 12, &my)
// 	wg.Wait()
// 	fmt.Println(my)
// }
import (
	"fmt"

	"github.com/Allenxuxu/toolkit/sync"
)

func main() {

	fmt.Println("start")
	// 如下初始化两种都可以。
	// sw := sync.WaitGroupWrapper{}
	sw := &sync.WaitGroupWrapper{}
	for i := 0; i < 10; i++ {
		sw.AddAndRun(func() {
			fmt.Println("ok")
		})
	}
	sw.Wait()
	fmt.Println("end")
}
