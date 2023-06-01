# mycd

## Problem statement
Write a program simulating a "cd" Unix command that changes a current directory in a file system. The simulated command takes two path strings from the command line and prints either a new path or an error.

 

The first path is a current directory. The second path is a new directory.  To make it simple let's assume that a directory name can only contain alphanumeric characters. A single dot (".") indicates a current directory, and the two dots ("..") indicate a step to a previous directory, up from the current one. A single forward slash "/" indicates a root directory. Multiple consecutive slashes are treated as equivalent to one. The program needs to check that the new directory path is valid without relying on any OS call or shell command for verification, to construct a new path, and print it out. This is a pure string manipulation exercise; no system calls are needed. Use any of the C++/Java/Python/Golang languages.

-----------------------------------------------------------------------

## How to compile and run

1) Clone repository
2) Ensure go environment is set up: https://go.dev/doc/install
3) The program can be run by one of the two options below:
    - Run from source `go run main.go path.go <current_path> <destination_path>`
    - Compile `mycd` executable: `go build -o mycd main.go path.go`. The mycd executable can be run as follows `./mycd <current_path> <destination_path>`. Optionally the `mycd` executable can be copied to the system path.

The usage is exactly the same as described in the problem statement

## Test results

The tests provided in the problem statement are all passing.

## How to run automated tests

Run `go test -v` from repository directory. An IDE such a VSCode with suitable plugin installed would facilite running and viewing test results.
