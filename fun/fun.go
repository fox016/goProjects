package main

import (
	"log"
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
	"math"
)

func init() {
	log.SetPrefix("fun: ")
	log.SetFlags(0)
}

func main() {
	runSqrt()
}

/*
 * Functions for square root calculation using Newton's method
 */
func runSqrt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a number: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading in input text: %v", err)
		}
		text = strings.Replace(text, "\n", "", -1)
		if text == "" {
			fmt.Println("Thank you, come again!")
			break
		}
		numIn, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Printf("You have to enter a number, silly. Error: %v", err)
		} else {
			sqrt := Sqrt(numIn)
			fmt.Printf("The square root of %v is %v\n", numIn, sqrt)
		}
	}
}

const min_diff float64 = 1E-10

func Sqrt(x float64) float64 {
	z := x/2.0
	for {
		fmt.Println(z)
		new_z := z - (z*z - x) / (2*z)
		if math.Abs(new_z - z) < min_diff {
			return new_z
		}
		z = new_z
	}
	return z
}

/*
 * Functions for Fahrenheit to Celcius converter
 */
func runFahrToCel() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter degrees fahrenheit: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading in input text: %v", err)
		}
		text = strings.Replace(text, "\n", "", -1)
		if text == "" {
			fmt.Println("Thank you, come again!")
			break
		}
		fahr, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Printf("You have to enter a number, silly. Error: %v", err)
		}
		cel := FahrToCel(fahr)
		fmt.Printf("%v degrees Fahrenheit => %v degrees Celsius\n", fahr, cel)
	}
}

func FahrToCel(fahr float64) float64 {
	return (fahr - 32) * 5.0/9.0
}
