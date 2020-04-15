package main

import "fmt"

func main() {

	var choice int
	fmt.Print("Enter the choice : ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		var size int
		var searchNumber int
		fmt.Print("Enter Size of an Array (Max 100): ")
		fmt.Scan(&size)
		LinearSearch(size, searchNumber)

	case 2:
		ar := []int{4, 2, 1, 5, 3}
		var sorted []int
		sorted = BubbleSort(ar)
		fmt.Println(sorted)

	// swap(1,2)
	default:
		fmt.Print("Please choose valid option\n")
	}

}

//****Linear Search Function****//
func LinearSearch(size, search int) {
	var number [100]int
	for i := 0; i < size; i++ {
		fmt.Print("\nEnter Number ", i, ": ")
		fmt.Scan(&number[i])
	}
	fmt.Print("Enter the value want to search : ")
	fmt.Scan(&search)
	for k := 0; k < 5; k++ {

		if number[k] == search {
			fmt.Println("Value Searched is : ", number[k], "at ", k)
		}

	}
}
//****Bubble Sort Function****//
func BubbleSort(numbers []int) []int {
	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				intermediate := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = intermediate
			}
		}
	}
	return numbers
}
