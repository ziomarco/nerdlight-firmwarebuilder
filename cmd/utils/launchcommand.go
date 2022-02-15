package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func LaunchCommand(commandDir string, commandToLaunch string, args ...string) error {
	cmd := exec.Command(commandToLaunch, args...)

	if len(commandDir) > 0 {
		cmd.Dir = commandDir
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	res, _ := out.ReadString('\n')

	if err != nil {
		log.Fatalln(fmt.Sprintf("Error in %s: %s %s", commandToLaunch, err.Error(), res))
		panic(err)
	}
	if len(res) < 1 {
		fmt.Println("Empty res for " + commandToLaunch)
		fmt.Println(res)
	}
	return nil
}

func LaunchCommandBool(commandToLaunch string, args ...string) (bool, error) {
	cmd := exec.Command(commandToLaunch, args...)
	err := cmd.Run()
	return cmd.ProcessState.ExitCode() != 0, err
}

func LaunchCommandWithEnv(commandDir string, env []string, commandToLaunch string, args ...string) error {
	cmd := exec.Command(commandToLaunch, args...)
	cmd.Env = append(cmd.Env, env...)

	if len(commandDir) > 0 {
		cmd.Dir = commandDir
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	res, _ := out.ReadString('\n')

	if err != nil {
		log.Fatalln(fmt.Sprintf("Error in %s: %s %s", commandToLaunch, err.Error(), res))
		panic(err)
	}
	if res != "" {
		fmt.Println(res)
	}
	return nil
}
