// failed code,
// not reeally working
// from chatgpt
package main

/*
 * Complete the 'counterGame' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts LONG_INTEGER n as parameter.
 */

func counterGame(n int64) string {
	// If the number is already 1, Richard wins because Louise cannot make a move.
	if n == 1 {
		return "Richard"
	}

	// Function to check if a number is a power of 2
	isPowerOfTwo := func(n int64) bool {
		return n > 0 && (n&(n-1)) == 0
	}

	// Function to find the closest lower power of 2
	closestPowerOfTwo := func(n int64) int64 {
		power := int64(1)
		for power <= n {
			power <<= 1
		}
		return power >> 1
	}

	// Keep playing the game until n becomes 1
	moves := 0
	for n > 1 {
		if isPowerOfTwo(n) {
			// If n is a power of 2, divide by 2
			n /= 2
		} else {
			// If n is not a power of 2, subtract the closest power of 2
			closestPower := closestPowerOfTwo(n)
			n -= closestPower
		}
		// Alternate moves between Louise and Richard
		moves++
	}

	// If the number of moves is odd, it means Richard made the last move, so Louise wins.
	if moves%2 == 0 {
		return "Louise"
	} else {
		return "Richard"
	}
}
