package main
import "fmt"

func main (){

	userch := make(chan string)

	userch <- "Bob"

	user := <- userch 

	fmt.Println(user)
}