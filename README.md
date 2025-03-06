# Go bcrypt Cost Benchamrk

This is a very basic Go program that checks the bcrypt cost needed to reach the target time, starting with the Default Cost.
It has a default target cost of 250 millisecond and it can be changed with the target parameter (`--target <time in millisecond>`).

## Usage
* `git clone https://github.com/mtzfederico/bcrypt-cost-benchmark.git`
* `cd bcrypt-cost-benchmark`
* `go get .`
* `go run .` or `go run . --target <time in millisecond>`