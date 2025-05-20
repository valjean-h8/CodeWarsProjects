/*
https://www.codewars.com/kata/514b92a657cdc65150000006/train/go
*/
package kata

func Multiple3And5(number int) int {
	if number < 0 { // Check for negative number
		return 0
	}
	sum := 0
	for i := 1; i < number; i++ { // Cycle of numbers from 1 to number - 1
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
