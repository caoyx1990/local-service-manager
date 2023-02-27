package internal

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ipconfig")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
