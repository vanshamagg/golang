package panicrecover

import "fmt"

func CreatePanic() {
	defer HandlePanic()
	arr := []string{"Vansham", "Aggarwal"}
	third := arr[2]

	fmt.Println(third)
}

func HandlePanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered")
	}
}