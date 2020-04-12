package deploy

import (
	"log"
	"os/exec"
	"strings"

	"github.com/lissteron/deployer/pkg/github"
	"github.com/spf13/viper"
)

func ProcessPush(logger *log.Logger, event *github.PushEvent) {
	if !strings.Contains(viper.GetString("ACCEPTED_REPOSITORY"), getBranch(event.Ref)) {
		return
	}

	if err := pull(logger, viper.GetString("REPOSITORY_PATH")); err != nil {
		logger.Panicln("[error]", err)
		return
	}

	if viper.GetString("GO_TOUCH") != "" {
		var needRestart bool

		for _, commit := range event.Commits {
			if needRestart {
				break
			}

			files := append([]string{}, commit.Added...)
			files = append(files, commit.Removed...)
			files = append(files, commit.Modified...)

			for _, fn := range files {
				if strings.Contains(fn, "go") {
					needRestart = true
					break
				}
			}
		}

		if needRestart {
			b, err := exec.Command("/bin/sh", "-c", viper.GetString("GO_TOUCH")).CombinedOutput()
			if err != nil {
				logger.Println("[error]", err)
				logger.Println(string(b))

				return
			}
		}
	}
}
