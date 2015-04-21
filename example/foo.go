package example

import "fmt"

func yoMe() {
	_ = "waitgroupgen"
	a := foo(1, 3)
	d := Math.pow(1, 3)
	b := bar()

	c, err := car()
	_ = "end waitgroupgen"
	fmt.Println(a + b)
}

func yoMeWaitGroup() {
}

func foo(a, b int) int {
	time.Sleep(200 * time.Milliseconds)
	return 200
}

func bar() int {
	time.Sleep(300 * time.Milliseconds)
	return 300
}

func car() (int, error) {
	time.Sleep(300 * time.Milliseconds)
	return 300, nil
}
