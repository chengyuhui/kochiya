package main

import (
	"os"
	"os/exec"
)

func encode(conf config, avsPath, binPath string) error {
	cmd := exec.Command("encode.bat", avsPath, conf.Output.Path, binPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
