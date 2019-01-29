package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func capturePrimes(max int, nums []bool) ([]int, int) {
	var primes []int
	var sum = 0

	// Traverses array and anything that is still marked 'false'
	// is a prime number based on the sieve of Eratosthenes.
	for i := 2; i <= max; i++ {
		if nums[i] == false {
			sum += i
			primes = append(primes, i)
		}
	}

	return primes, sum
}

func sieveHelper(a int, max int, nums[]bool) {
	// Checks all multiples of the value 'a'
	for j := a+a; j <= max; j += a {
		if nums[j] != true {
			nums[j] = true
		}
	}
	// Subtract 1 from wait group.
	wg.Done()
}

func sieveOfEratosthenes(max int) ([]int, int) {

	// Defines number of threads we're allowed to spawn.
	runtime.GOMAXPROCS(8)

	// Array containing all numbers in consideration as boolean.
	allNumbers := make([]bool, max+1)

	for i := 2; i*i <= max; i++ {
		if allNumbers[i] == false {
			// Add 1 to wait group
			wg.Add(1)
			go sieveHelper(i, max, allNumbers)
		}
	}

	// Waits for all threads to finish before checking for primes.
	wg.Wait()
	return capturePrimes(max, allNumbers)
}

func inputFail () {
	fmt.Println()
	fmt.Println("Please provide a valid command line argument. [100 < n < ∞]")
	fmt.Println()
	fmt.Println("EXAMPLE: go run soe.go <n>")
	fmt.Println()
}

func main() {

	// Input Validation
	if len(os.Args) < 2 {
		inputFail()
		return
	}

	// Grabs number of philosophers from command line argument.
	arg := os.Args[1]
	num, _ := strconv.Atoi(arg)

	// Input Validation
	if num < 100 {
		inputFail()
		return
	}

	start := time.Now()
	primes, sum := sieveOfEratosthenes(num)
	end := time.Now()
	elapsed := end.Sub(start)

	// Creates output file and closes upon function completion.
	fout, _ := os.Create("primes.txt")
	defer fout.Close()

	// Print required information to "primes.txt" file.
	fmt.Fprintln(fout,"Execution Time: ", elapsed)
	fmt.Fprintln(fout,"Total Number of Primes Found: ", len(primes))
	fmt.Fprintln(fout,"Sum of all Primes Found: ", sum)
	fmt.Fprintln(fout,"Top Ten Max Primes:")

	for i := len(primes) - 10; i < len(primes); i++ {
		fmt.Fprintln(fout,primes[i])
	}
}