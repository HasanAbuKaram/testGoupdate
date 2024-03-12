package main

import "fmt"

func main() {
	fmt.Printf("Hello, World! Running version %s\n", version)
}

// go build -ldflags="-s -w" -o myapp.exe github.com/HasanAbuKaram/testGoupdate
