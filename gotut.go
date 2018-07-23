package main

import "fmt"


func main() {
	grades := make(map[string]float32)

	grades["tim"] = 42
	grades["monty"] = 92
	grades["sam"] = 67


	TimsGrade := grades["tim"]
	
	fmt.Println(TimsGrade)

	delete(grades, "tim")
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k,":",v)
	}



}
