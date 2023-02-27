package main

import "servicesManager/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
