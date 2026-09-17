package main
import "fmt"
// import "time"
// import "golang_tutorial/lib1"

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
	defer fmt.Println("main end")
	c := make(chan int, 3)

	go func(){
		for i:=0; i<5; i++ {
			c <- i
			fmt.Println("send data:", i)
			fmt.Println("channel length:", len(c))
			fmt.Println("channel capacity:", cap(c))
		}
		close(c)
	}()

	// fmt.Println("Hello, Go!")
	// time.Sleep(1 * time.Second)
	// fmt.Println(fn("Beijing", "Shanghai"))
	// lib1.Lib1Func()
	
	for{
		if data, ok := <-c; ok {
			fmt.Println(data)
		}else{
			break
		}
	}


}
