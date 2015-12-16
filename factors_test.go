package main

import (
	"testing";
	"fmt";
	"math";
)

//Cylcles though a bunch of instances of StartFactor, wich is the only function call in func main()
//this function calls Factor then factor Calls Loop
func TestStartFactor(t *testing.T){
	fmt.Println("\n Testing StartFactor() now \n")
	for i := 2; i < 333; i+=7{
		StartFactor(i)

	}
}

//tests Loop() 
//most of this code was copy pasted from the Factor() function that calls Loop() in the first place
func TestLoop(t *testing.T){
	
	fmt.Println("\n Testing Loop() funcion now \n")
	factor := 0
	originalNum := 0
	for j := 120; j < 220; j+=3{
		originalNum = j
		for math.Remainder(float64(originalNum), float64(factor)) != 0{
			factor++
		}
		
		zed := 0
		tempNum := 0
		tempArray := make([]int, 0)
		tempNum = originalNum / factor //devides the original number by the factor, to find another factor

		//finds the factors of the passed in number
		for i := 0; i <= tempNum; i++ {

			//if you devide the temperary number by i and there is no remainder
			if math.Remainder(float64(tempNum), float64(i)) == 0 {
				//assigns all possible integer factors to a spot in the array
				tempArray = append(tempArray, i)
			}
		}
		Loop(factor, originalNum, tempArray, zed, len(tempArray)-1)
		factor++
		
	}
}

//Prime factorization, not actually used in factors.go
//but I wanted to test it anyhow
func TestIsPrime(t *testing.T){

	fmt.Println("\n Testing IsPrime() now \n")
	primes := []int{2,3,5,7,11,13,17,23,29,31,37}
	for i := 0; i < len(primes)-1; i++{
		pTest := IsPrime(primes[i])
		
		if pTest == false{
			t.Fail()
		} else {
			fmt.Println(primes[i], " is prime")
		}
	}
}
	


/*
//redundant because StartFactor() calls Factor and Factor calls Loop()
//
func TestFactor(t *testing.T){
	for i := 100; i < 300; i+=3{
		for j := 1; j < 20; j+=1{
			Factor(i, j)
		}
	}
}
*/
