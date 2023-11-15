package main

// Given an array of integers gas and an array of integers cost
// Return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1.
// If there exists a solution, it is guaranteed to be unique

// This solution works by keeping track of the total gas and total cost.
// If the total gas is less than the total cost, then we know that there is no solution.
// Otherwise, we iterate through the gas array and keep track of the tank.
// If the tank is less than 0, then we know that we cannot start at the current index.
// So, we set the start index to the next index and reset the tank to 0.
// At the end, we return the start index.
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost, tank, start := 0, 0, 0, 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}
	return start
}
