package main 
 
import "fmt"

import "rsc.io/quote"

func main () { 
	fmt.Print(quote.Glass())
	fmt.Println(Mygame())
	fmt.Println(Kmsaio())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(minecraft())
}

func Mygame() string {
	return "I love lemon rice."
}

func Kmsaio() string {
	return "This is easier than playing minecraft"
}

func minecraft() string {
	return "minecraft is easy since I watch Dream"
}

