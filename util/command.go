package util

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunCmd(name string, args ...string) (error, *bytes.Buffer) {
	var out bytes.Buffer
	c := exec.Command(name, args...)
	c.Stdout = &out
	err := c.Start()
	if err != nil {
		fmt.Println(out.String())
		return err, nil
	}
	err = c.Wait()
	if err != nil {
		fmt.Println(out.String())
		return err, nil
	}

	return nil, &out
}

func RunCmdIn(in *bytes.Buffer, name string, args ...string) (error, *bytes.Buffer) {
	var out bytes.Buffer
	c := exec.Command(name, args...)
	c.Stdout = &out
	c.Stdin = in
	err := c.Start()
	if err != nil {
		fmt.Println(out.String())
		return err, nil
	}
	err = c.Wait()
	if err != nil {
		fmt.Println(out.String())
		return err, nil
	}

	return nil, &out
}
