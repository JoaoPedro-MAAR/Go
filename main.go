package main

import (
	"slices"
	"errors"
	"fmt"
	"sort"
	s "strings"
)


func main() {

	numbersList := make([]int, 0)
	var option int
	var a int
	var b int
	var types string
	var err error
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
			"7 - Order List\n"+
			"8 - Show Evens\n"+
			"0 - Exit(Bye Bye)\n"+
			"Option: ",
			numbersList,
		)
		fmt.Scan(&option)
		switch option{
			case 1:
				numbersList,err = addNumbersInterface(numbersList)
				if (err != nil){
					fmt.Println(err)
				}
			case 2:
				listNumbers(numbersList)
			case 3:
				numbersList, err = rmByIndexInterface(numbersList)
				if (err != nil){
					fmt.Println(err)
				}
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
			case 7:
				fmt.Print("Select the option Ascending [A] or Descending [D]: ")
				fmt.Scan(&types)
				numbersList = orderList(numbersList, types)
			case 8:
			    showEven(numbersList)
			case 0:
				break outerloop
			default:
				fmt.Print("Not a known option try again")
		}

	}




}

func showEven(list []int) {
	fmt.Print("[")
	for i := 0; i < len(list); i++ {
		if list[i] % 2 == 0{
		fmt.Print(list[i])
		if i < len(list) -1{
		fmt.Print(",")
		}
		}
	}
	fmt.Print("]\n")

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

func rmByIndexInterface(list []int) ([]int, error){
	var position int
	fmt.Print("Digit an integer value to be the index: ")
	fmt.Scan(&position)
	return rmByIndex(list, position)}

func addNumbersInterface(list []int) ([]int, error){
	var newNumber int
	fmt.Print("Digit an integer value to be added: ")
	_, err :=fmt.Scan(&newNumber)
    if err != nil{
		return list,err
	}
	if newNumber < 0{
		return list, errors.New("The list don't support negative numbers")
	}
	return addNumbers(list, newNumber), nil

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

func rmByIndex(list []int, index int) ([]int, error) {
	if index < 0 || index > len(list){
		return list, errors.New("Not valid index")
	}
	return append(list[:index], list[index+1:]...), nil
}


func staticstics(list []int) (int, int, float64, error){
	listCopy := make([]int, len(list))
	if len(list) == 0{
		return -1,-1,-1,errors.New("No data in the List")
	}
	listCopy = orderList(list,"A")

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

func orderList(list []int,types string )([]int){
	if (s.ToLower(types) == "a"){
		slices.Sort(list)
	}else if s.ToLower(types)== "d"{
		sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
	}
	return list
}