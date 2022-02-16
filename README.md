# Solution for Codility sample test in Go

Test located in [Codility sample test site](https://app.codility.com/demo/take-sample-test/)
(accessed the 2022-02-16)

Just a normal problem where you get an array of integers and need to look for the
positive smallest positive integer not in the array.

## Explanation of the Solution

1. Exclude all the negative (`<0`) numbers from the array. If size of array is
zero then return `1`.
2. Ascending sort and remove duplicates from the array
3. Get the first value of the array, compare it with initialized major number in
`1`, if it's the same, set new major number with `+1`
4. If not, go to the next value in the array. It will `break` the for before
checking all the numbers in the array making the program efficient.

There are some examples in the main.go

## Build

```sh
bo build .
```

## Run

```sh
go run .
```

or if it was builded

```sh
./codility
```
