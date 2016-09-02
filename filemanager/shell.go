package filemanager

import "os/exec"

func run(dir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	c.Stderr = Stderr
	c.Stdout = Stdout

	return c.Run()
}
