package main

import "fmt"

func main() {
	switch score := 50; score {

	case 90:
		fmt.Println("达标")
	case 80:
		fmt.Println("差10")
	case 70:
		fmt.Println("差20")
	case 60:
		fmt.Println("差30")
	default:
		fmt.Println("不达标")
	}
}
