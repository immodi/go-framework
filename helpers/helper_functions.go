package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckForStaticFiles() bool {
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		return false
	}

	return true
}

func GenerateTmplScripts() {
	cmd := exec.Command("bash", "-c", "templ generate")

	// Get the output of the command
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output
	fmt.Println(string(output))
}
