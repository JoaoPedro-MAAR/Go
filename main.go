package main

import (
	"errors"
	"fmt"
	"sort"
)


func main() {

	numbersList := make([]int, 0)
	var option int
	var a int
	var b int
	fmt.Println("Hello, World!")
	outerloop: for{
		fmt.Printf(
			"\nManagement GO numbers\n"+
			"List: %v\n"+
			"1 - Add Number\n"+
			"2 - Show List\n"+
			"3 - Remove by position\n"+
			"4 - Show Metrics\n"+
			"5 - Safe division (Integer, Integer)\n"+
			"6 - Clear List\n"+
			"0 - Exit(Bye Bye)\n"+
			"Option: ",
			numbersList,
		)
		fmt.Scan(&option)
		switch option{
			case 1:
				numbersList = addNumbersInterface(numbersList)
			case 2:
				listNumbers(numbersList)
			case 3:
				numbersList = rmByIndexInterface(numbersList)
			case 4:
				showMetrics(numbersList)
			case 5:
				fmt.Print("First number: ")
				fmt.Scan(&a)
				fmt.Print("Second Number: ")
				fmt.Scan(&b)
				result, err :=divide(a,b)
				if err == nil{
					fmt.Printf("The result is %v", result)
				} else {
					fmt.Print(err)
				}
				fmt.Println()
			case 6:
				numbersList = cleanList(numbersList)
			case 0:
				break outerloop
			default:
				fmt.Print("Not a known option try again")
		}

	}




}

func showMetrics(list []int){
	minimun, maximum, mean, err := staticstics(list)
	if err == nil{
	fmt.Printf("Minimun: %v\n", minimun)
	fmt.Printf("Maximum: %v\n", maximum)
	fmt.Printf("Mean: %v\n", mean)
	} else{
		fmt.Print(err)
	}

}

func rmByIndexInterface(list []int) []int{
	var position int
	fmt.Print("Digit an integer value to be added: ")
	fmt.Scan(&position)
	return rmByIndex(list, position)}

func addNumbersInterface(list []int) []int{
	var newNumber int
	fmt.Print("Digit an integer value to be added: ")
	fmt.Scan(&newNumber)
	return addNumbers(list, newNumber)

}



func addNumbers(list []int, num int) []int {
	return append(list, num)
}


func listNumbers(list []int){
	fmt.Print("[")
	for i := 0; i < len(list); i++ {
		fmt.Print(list[i])
		if i < len(list) -1{
		fmt.Print(",")
		}
	}
	fmt.Print("]\n")

}

func rmByIndex(list []int, index int) []int {
	return append(list[:index], list[index+1:]...)
}


func staticstics(list []int) (int, int, float64, error){
	listCopy := make([]int, len(list))
	if len(list) == 0{
		return -1,-1,-1,errors.New("No data in the List")
	}
	copy(listCopy, list)
	sort.Slice(listCopy, func(i, j int) bool {
		return listCopy[i] < listCopy[j]
	})

	minimun := listCopy[0]
	maximum := listCopy[len(listCopy)-1]
	medium := calculateAverage(listCopy)
	return minimun, maximum, medium, nil
}



func calculateAverage(list []int) float64{
	total := 0
	for _,number := range list {
		total = total + number
		
	}

	return float64(total) / float64(len(list))
}


func divide(a int, b int) (float64, error) {
	if b == 0{
		return -1, errors.New("Divison by zero is not allowed")
	}
	return float64(a) / float64(b), nil

}

func cleanList(list []int) []int{
	list = make([]int, 0)
	return list
}