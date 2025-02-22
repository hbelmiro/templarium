package runner

func FakeRunnerReturning(output string) Runner {
	return fakeRunner{output: output}
}

type fakeRunner struct {
	output string
}

func (r fakeRunner) RunCommand(string, ...string) ([]byte, error) {
	return []byte(r.output), nil
}
