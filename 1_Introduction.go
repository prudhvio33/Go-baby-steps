package main

import "fmt"

func main() { 	// Program exec starts here
	fmt.Println("hello")
}

/*
	Running Go code:
1) go run <file_name>.go
2) go install workspacepath/bin/<file_name>, When you type go install hello, the go tool searches for the hello package (hello is called as package, we will look into packages in more detail later) inside the workspace. Then it creates a binary named hello inside the bin directory of the workspace.
*/