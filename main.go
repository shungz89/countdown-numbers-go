package main

import (
	"fmt"
	"math"
	"strings"
)

// Define the operations and their string representations
var operations = []func(float64, float64) (float64, string){
	func(a, b float64) (float64, string) { return a + b, fmt.Sprintf("%.0f + %.0f", a, b) },
	func(a, b float64) (float64, string) { return a - b, fmt.Sprintf("%.0f - %.0f", a, b) },
	func(a, b float64) (float64, string) { return a * b, fmt.Sprintf("%.0f * %.0f", a, b) },
	func(a, b float64) (float64, string) {
		if b != 0 {
			return a / b, fmt.Sprintf("%.0f / %.0f", a, b)
		} else {
			return math.NaN(), ""
		}
	},
}

// countdownSolver is the main recursive function to solve the Countdown numbers game
func countdownSolver(numbers []float64, target float64) []string {
	// Base case: if the target is directly in the list of numbers, return it
	if contains(numbers, target) {
		return []string{fmt.Sprintf("%.0f", target)}
	}

	// Iterate over all pairs of numbers in the list
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j { // Ensure we are not using the same number twice
				for _, op := range operations { // Apply all four operations
					result, expr := op(numbers[i], numbers[j]) // Get the result and expression
					if result == target {                      // Check if the result matches the target
						return []string{expr, fmt.Sprintf("= %.0f", target)}
					}
					// Check if the result is a valid positive integer
					if !math.IsNaN(result) && result > 0 && math.Floor(result) == result {
						newNumbers := removePair(numbers, i, j)          // Remove the pair of numbers used
						newNumbers = append(newNumbers, result)          // Add the result to the new list of numbers
						subResult := countdownSolver(newNumbers, target) // Recursive call
						if subResult != nil {                            // If a solution is found, return it
							return append([]string{expr + ","}, subResult...)
						}
					}
				}
			}
		}
	}
	return nil // If no solution is found, return nil
}

// contains checks if the target is in the list of numbers
func contains(numbers []float64, target float64) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}

// removePair removes two numbers from the list by their indices
func removePair(numbers []float64, i, j int) []float64 {
	newNumbers := []float64{}
	for k := 0; k < len(numbers); k++ {
		if k != i && k != j {
			newNumbers = append(newNumbers, numbers[k])
		}
	}
	return newNumbers
}

func main() {
	// Define the list of numbers and the target number
	numbers := []float64{100, 75, 3, 10, 10, 3}
	target := 129.0

	// Call the solver function
	solution := countdownSolver(numbers, target)

	// Print the solution if found, ensuring the output is comma-separated
	if solution != nil {
		fmt.Println("Solution found:", strings.Join(solution, " "))
	} else {
		fmt.Println("No solution found.")
	}
}
