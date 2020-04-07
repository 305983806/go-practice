package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	//lp -o media=Custom.15.2×10.2cm filename
	//lp -o media=Custom.7×3cm max.pdf -d gk888t
	//lp -o media=Custom.WIDTHxLENGTHmm filename
	cmd = exec.Command("/bin/bash", "-c", "lp -o media=Custom.4x7cm 4-7.pdf -d gt888k")
	//cmd.CombinedOutput()

	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
