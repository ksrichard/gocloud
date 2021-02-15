package util

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func RunCmdInteractive(name string, args ...string) error {
	var out bytes.Buffer
	c := exec.Command(name, args...)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Println(out.String())
		return err
	}
	return nil
}

func RunCmdInteractiveInDir(dir string, name string, args ...string) error {
	var out bytes.Buffer
	c := exec.Command(name, args...)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Println(out.String())
		return err
	}
	return nil
}

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

func RunCmdInDir(dir string, name string, args ...string) (error, *bytes.Buffer) {
	var out bytes.Buffer
	c := exec.Command(name, args...)
	c.Dir = dir
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
