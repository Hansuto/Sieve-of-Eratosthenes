# Sieve-of-Eratosthenes
Implementing the Sieve of Eratosthenes to determine all prime numbers less than ***n*** using concurrent programming practices in Go.

    go run soe.go <n> 
   *where 100 < n < inf*

## Explanation
I make an array of Boolean values 0 through n, all initialized to false. Starting with the
number 2, I spawn a thread to flip each of its multiples to true. I repeat this
process for all values still initialized to false up to sqrt(n). As I get closer to sqrt(n) I don’t have to spawn as many threads due to
previous threads marking many multiples as true. I limit the amount of threads available to 8. After all threads have finished, I traverse the Boolean array and append every value that is still set to false to a separate primes array. Once completed I have a list of all primes, in-order, less than n. 

## Efficiency

Typically,in the Sieve of Eratosthenes, each thread completes **n/i** steps, where **i** is prime. The whole complexity is **summation(n/i) = n * summation(1/i)**. According to prime harmonic series, the **summation(1/i)** where **i** is prime is **log(log n)**. In total, **O(n * log(log n))**. I optimized my solution by running the outer loop **sqrt(n)** times, because once the outer loop reaches **sqrt(n)**, I would have already checked all multiples due to **n*n**. Therefore, the time complexity of my solution is **O(sqrt(n) * log(log n)**.

My solution is effective, although I do see one drawback in my approach that I was unable to resolve.  The thread checking all multiples of 2 takes much longer than threads checking the higher multiples.
