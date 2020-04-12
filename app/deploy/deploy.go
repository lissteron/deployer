package deploy

import (
	"log"
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
}
