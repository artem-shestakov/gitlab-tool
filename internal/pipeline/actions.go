package pipeline

import (
	"errors"
	"fmt"
	"os"

	"github.com/artem-shestakov/gitlab-tool.git/internal/gitlab"
	"github.com/artem-shestakov/gitlab-tool.git/internal/telegram"
	"github.com/urfave/cli/v2"
)

func notify(c *cli.Context) error {
	GL_TOKEN := os.Getenv("GL_TOKEN")
	CI_SERVER_URL := os.Getenv("CI_SERVER_URL")
	CI_PROJECT_ID := os.Getenv("CI_PROJECT_ID")
	CI_PIPELINE_ID := os.Getenv("CI_PIPELINE_ID")

	gitlab := gitlab.GitLab{
		Token: GL_TOKEN,
	}

	_, err := gitlab.GetPipeline(CI_SERVER_URL, CI_PROJECT_ID, CI_PIPELINE_ID)
	if err != nil {
		return err
	}

	switch c.String("channel") {
	case "telegram":
		TG_TOKEN := os.Getenv("TG_TOKEN")
		TG_CHAT_ID := os.Getenv("TG_CHAT_ID")
		bot := telegram.NewBot(TG_TOKEN)
		fmt.Println(TG_TOKEN, TG_CHAT_ID)
		bot.SendMesage(TG_CHAT_ID, "Test")
	default:
		return errors.New("Channel unknown")
	}
	return nil
}
