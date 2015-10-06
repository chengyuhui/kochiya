package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func encode(conf config, avsPath, binPath string) error {
	cmd := exec.Command(filepath.Join(binPath, "encode.bat"), avsPath, conf.Output.Path, binPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
