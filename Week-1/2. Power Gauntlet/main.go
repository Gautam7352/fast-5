package main

import (
	"fmt"
)

func canPowerup(n int) bool{
	
	if(n<=0){
		return false
	}

	for i := 1; i<=n;i*=2{
		if(n%i!=0){
		return false
		}
	}

	return true
}

func main() {
	var x int
	fmt.Print("Please enter the number of crystals: ")
	fmt.Scanln(&x)

	result := canPowerup(x)
	if(result){
	fmt.Println("Yes")
	} else {
	fmt.Println("No")
	}

}
