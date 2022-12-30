package pkg

import "os/exec"

func PbCopy(content string) error {
	cmd := exec.Command("pbcopy")
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(content)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return cmd.Wait()
}

func PbPaste() (content string, err error) {
	cmd := exec.Command("pbpaste")
	out, err := cmd.Output()
	return string(out), err
}
