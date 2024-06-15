package main
import (
	"fmt"
)

func expo(x int, n int) int{
	if x == 0{
		return 1
	} 
	 return x * expo(x, n-1)
}

func main1 (){
	var x,y int
	fmt.Scan(&x,&y)
	fmt.Print(expo (x,y))
}