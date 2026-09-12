package main
import "fmt"
// import "time"
import "golang_tutorial/lib1"

const (
	BEIJING = 110
	SHANGHAI = 310
)

// func fn(city1 string, city2 string) (int, int) {
// 	if city1 == "Beijing" && city2 == "Shanghai" {
// 		return BEIJING, SHANGHAI
// 	}
// 	return 0, 0
// }

func main(){
	fmt.Println("Hello, Go!")
	// time.Sleep(1 * time.Second)
	// fmt.Println(fn("Beijing", "Shanghai"))
	lib1.Lib1Func()

}
