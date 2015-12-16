package main

import (
	"fmt"
	"math"
	"strconv"
	
)

var zed = 0 //a counter so numbers do not repeat

/////////////////////////////////////////////////////////////////////////////////////////
//takes in the original number and the factor to be used
//in this version StartFactor is not fully necassary, and could be incopreated here
///////////////////////////////////////////////////////////////////////////////////

func Factor(originalNum int, factor int) {
	zed = 0
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
	/*	
	////////////////////////////////////////////////////////
	//My early itterative solution
	////////////////////////////////////////////////////////
	
		//iterates over the array finding the correct multiples of the original number
		// i = zed so numbers dont repeat
		//i < lenght of tempArray/2 to split the array and stop repeats
		
		for i := zed; i < len(tempArray)/2; i++ {
		
			//j = half of tempArray length, so it starts where i leaves off
			for j := len(tempArray) - 1; j > len(tempArray)/2; j-- {
				if factor*tempArray[i]*tempArray[j] == originalNum {

					//displays the factors that equal the original number in the desired format
					fmt.Println(factor, " x ", tempArray[i], " x ", tempArray[j], " = ", originalNum)
				}
			}
		}
	*/
	//////////////////////////////////////////////////////////////////////////
	//Loop is recursive, but nearly the same as the itterative solution.
	////////////////////////////////////////////////////////////////////////////
	
	Loop(factor, originalNum, tempArray, zed, len(tempArray)-1)
	
	//increments zed to stop repeats
	zed++
	return
}
////////////////////////////////////////////////////////////////////////////////
//Not as important as it was in previous versions
//it could be implemented in the Factor function, but it works fine the way it is
/////////////////////////////////////////////////////////////////////////////////
func StartFactor(originalNum int) {
	zed = 0 //resets zed, zed is global
	factors := make([]int, 0)

	//finds all factors of the original number
	//'i' can stop at the square root of the originalNum
	for i := 0; i*i <= originalNum; i++ {
		if math.Remainder(float64(originalNum), float64(i)) == 0 {
		
			//adds the factors to the array
			factors = append(factors, i)

		}
	}
	//calls the Factor function for each element in the array of factors
	for i := 0; i < len(factors); i++ {
		Factor(originalNum, factors[i])
	}

	return
}
////////////////////////////////////////////////////////////////////////////////
//to check if a number is prime
//I made a few attemps to use prime factors with recursion for the bonus, but they all failed
//this code worked nicely though
/////////////////////////////////////////////////////////////////////////////////
//I had this block commented out, but then I couldn't test it
func IsPrime(x int) bool {
	//2 is prime but is a bit of a special case
	if x == 2 {
		return true
	}
	//0 and 1 are not prime, but I think thats up for debate
	if x <= 1 {
		return false
	}
	if math.Remainder(float64(x), float64(2)) == 0 {
		return false
	}
	//checks every possible factor up to the square root of the given number
	for i := 3; i*i <= x; i++ {
		if math.Remainder(float64(x), float64(i)) == 0 {
			return false
		}
	}
	//if it meets none of the above criteria it is prime
	return true
}

//////////////////////////////////////////////////////////////////////////////////
//This is the function in use.  It was an early attempt at the bonus. 
//In this form it is barely differnt from embedded for loops, maybe slower.
//I hear there is no tail call optimization in Go yet.
//////////////////////////////////////////////////////////////////////////////////

//function that replaces the embedded for loops in the Factor function
func Loop(factor int, originalNum int, tempArray []int, top int, bottom int) {
	temp := ""
	
	//a slice to hold strings of numbers
	
	
	//if the 'top' index is less than half the arra of factors
	if top <= len(tempArray)/2 {
		//if the 'bottom' index isgreater than half the facors array length
		if bottom >= len(tempArray)/2 {
		
			//if the passed in factor x the factors at index top and bottom = the original number
			if factor*tempArray[top]*tempArray[bottom] == originalNum {
				
				//copy the string to temp and then print to screen
				temp = (strconv.Itoa(factor) + " x " + strconv.Itoa(tempArray[top]) + " x " + strconv.Itoa(tempArray[bottom]) + " = " + strconv.Itoa(originalNum))
				fmt.Println(temp)
				
			}//decrement bottom and call the next Loop function
			bottom--
			Loop(factor, originalNum, tempArray, top, bottom)
			return
		}//increment top and call the next loop function
		top++
		Loop(factor, originalNum, tempArray, top, len(tempArray)-1)
		return
	}
	top++
}
/*
////////////////////////////////////////////////////////////////////////////////
//I think I keep making things too complicated, also channels are a bit different
//but go routines and channels are very interesting tools
///////////////////////////////////////////////////////////////////////////////
func ReCurve(originalNum int, factor int, factorsArr []int, recursionCount int, topLevel int, index chan int, total chan int) string {
	//currentTemp := 0
	totalTemp := 0
	indexTemp := 0
	string := ""
	fmt.Println("this happens")
	
	//goes to the bottom first
	if recursionCount > 1 {
		//adds the string returned by ReCurve to the collection
		string += (ReCurve(originalNum, factor, factorsArr, recursionCount-1, topLevel, total, index))
		if recursionCount < topLevel {
			//when we are not at the top level of the recursive calls
			totalTemp = <-total
			indexTemp = <-index
			for i := indexTemp; i < len(factorsArr); i++ {
				if math.Remainder(float64(totalTemp*factorsArr[i]), float64(originalNum)) == 0 && totalTemp*factorsArr[indexTemp] < originalNum {
					return (string + strconv.Itoa(factorsArr[i]) + " x ")
				}
			}
			if totalTemp*factorsArr[indexTemp] >= originalNum {
				index <- indexTemp - 1
				//ReCurve()originalNum, factor, factorsArr, recursionCount - 1, topLevel, index, total, current)
			}
		}
	} else if recursionCount == topLevel {
		fmt.Println("zzzzzzzzzzzz")
		for i := indexTemp; i < len(factorsArr); i++ {
			if totalTemp*factorsArr[i] == originalNum {
				return (string + (strconv.Itoa(factorsArr[i]) + " = " + strconv.Itoa(originalNum)))
			}
		}
	} else {
		//current <- originalNum/factor
		total <- factor
		index <- indexTemp
		return (string + (strconv.Itoa(factor)))
	}
	return string
}
*//////////////////////////////////////////////////////////////

func main() {
	for i := 120; i < 230; i+=101{
	StartFactor(i)
	}
}