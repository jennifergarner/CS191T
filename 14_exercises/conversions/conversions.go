package main
import "fmt"

func main () {
	//int to float64
	myInt := 2
	fmt.Println(int64(myInt))

	//float64 to int
	myFloat := 3.14
	fmt.Println(int(myFloat))

	//[]byte to string
	myBytes := []byte{'h','e','l','l','o'}
	fmt.Println(string(myBytes))

	//string to []byte
	myString := "hello world"
	fmt.Println([]byte(myString))

}
