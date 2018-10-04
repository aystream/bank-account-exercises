package main

import "github.com/aystream/bank-account-exercises/src/app"

func main() {
	newApp := &app.App{}
	newApp.Initialize()
	newApp.Run(":8080")
}
