# Advent of Code 2021
This repo contains my solutions to [Advent of Code 2021](https://adventofcode.com/2021)'s puzzles.

On my machine this runs all solutions entirely single threaded in less then a second.

Puzzle inputs go into their respective folders as `input.txt`.
Run all available days with with `go run .`.

Each package contains its own test cases, a global test for all solutions is available in the main package, ran via `go test` and requires a `solutions.json` file, see the example file given.

A VSCode task is provided to run the benchmark, alternatively run `go test -benchmem -run =^$ -bench ^BenchmarkSolutions > benchmark.txt`

Solutions can be output to file in json format via the -o flag, eg `go run . -o solutions.json`

Inputs and solutions are not available in this repo