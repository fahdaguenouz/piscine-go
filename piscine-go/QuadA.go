package main

import "fmt"
func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 && j == 0  && x > 1{
					fmt.Print("o")//
				}	else if i == 0 && j == 0  && x == 1{
					fmt.Print("o\n")//
				}	else if i == 0 && j >= 1 && j <= x-2 && x >1 {
					fmt.Print("-")//
				}	else if i == 0 && j == x-1 {
					fmt.Print("o\n")//
				}	else if i >= 1 && i <= y-2 && j == 0 && x >1 {
					fmt.Print("|")//
				}	else if i >= 1 && i <= y-2 && j == x-1 {
					fmt.Print("|\n")//
				}	else if i == y-1 && j == 0 && x >1 {
					fmt.Print("o")//
				}	else if i == y-1 && j >= 1 && j <= x-2 && x >1{
					fmt.Print("-")//
				}	else if i == y-1 && j == x-1 {
					fmt.Print("o\n")//
				}	else if i >= 1 && i <= y-2 && j >=1 && j <= x-2 && x >1{
					fmt.Print(" ")//
				}
			}
		}
	}
}
func main() {
	QuadA(5, -4)
	fmt.Print("\n")
	QuadA(5, 3)
	fmt.Print("\n")
	QuadA(5, 1)
	fmt.Print("\n")
	QuadA(1, 1)
	fmt.Print("\n")
	QuadA(1, 5)
	fmt.Print("\n")
	QuadA(6, 5)
	fmt.Print("\n")
	
}