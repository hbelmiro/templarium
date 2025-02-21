package runner

import "os/exec"

type Runner interface {
	RunCommand(name string, arg ...string) ([]byte, error)
}

func DefaultRunner() Runner {
	return defaultRunner
}

var defaultRunner Runner

type runner struct {
}

func (r runner) RunCommand(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return cmd.CombinedOutput()
}

func init() {
	defaultRunner = &runner{}
}
