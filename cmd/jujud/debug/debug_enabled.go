package debug

import (
	"fmt"
	"os"

	"github.com/go-delve/delve/cmd/dlv/cmds"
	"github.com/juju/errors"
)

const (
	Command           = "dlv"
	EnvJujuNoDebug    = "JUJU_NO_DEBUG"
	ErrShouldContinue = errors.ConstError("should exit")
)

func SetupJujuDebug(args []string) error {
	if _, exists := os.LookupEnv(EnvJujuNoDebug); exists {
		fmt.Printf("%s defined not starting dlv\n", EnvJujuNoDebug)
		return ErrShouldContinue
	}

	fmt.Printf("Setting os env %q\n", EnvJujuNoDebug)
	os.Setenv(EnvJujuNoDebug, "1")

	file, err := os.CreateTemp("", "delv-init")
	if err != nil {
		return fmt.Errorf("creating delv init file: %w", err)
	}
	fmt.Fprintln(file, "continue")
	fmt.Fprintln(file, "exit")
	file.Close()

	command := args[0]
	dlvArgs := []string{
		"--headless",
		"--listen", ":1122",
		"--accept-multiclient",
		"--init", file.Name(),
		"exec",
		command, "--",
	}
	dlvArgs = append(dlvArgs, args[1:]...)
	fmt.Printf("Starting dlv with %v\n", dlvArgs)

	dlvCmd := cmds.New(false)
	dlvCmd.SetArgs(dlvArgs)

	go func() {
		cCmd := cmds.New(false)
		cCmd.SetArgs([]string{
			"connect",
			"localhost:1122",
			"--init", file.Name(),
		})
		fmt.Println(cCmd.Execute())
	}()

	defer fmt.Println("dlv has stopped")
	if err := dlvCmd.Execute(); err != nil {
		return err
	}
	return fmt.Errorf("some error")
}
