package deploy

import (
	"fmt"
	"log"
	"os/exec"
)

func pull(logger *log.Logger, path string) (err error) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s && git pull", path))

	b, err := cmd.CombinedOutput()
	if err != nil {
		logger.Println("[error]", err)
		logger.Println(string(b))

		return
	}

	return
}
