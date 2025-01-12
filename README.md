# Unexpected Behavior When Accessing Elements in a Nil Map in Go

This repository demonstrates an uncommon error in Go:  accessing elements in a nil map.  Unlike some languages that would raise an exception, Go silently returns the zero value for the map's value type if you try to access a key in a nil map.

## The Bug
The `bug.go` file shows a simple example:  a nil map is accessed before it is initialized.  The program will compile and run without errors, but the output might be unexpected.

## The Solution
The `bugSolution.go` file presents the corrected code.  It initializes the map before accessing its elements.  This is generally the recommended way to prevent such issues and makes code behaviour much more explicit.

## How to Run
1. Clone this repository.
2. Navigate to the directory.
3. Run the Go program using `go run bug.go` and `go run bugSolution.go`.
Compare the output to understand the difference in behavior between the two programs.
