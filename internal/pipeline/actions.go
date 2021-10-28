package pipeline

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/artem-shestakov/gitlab-tool.git/internal/gitlab"
	"github.com/artem-shestakov/gitlab-tool.git/internal/telegram"
	"github.com/urfave/cli/v2"
)

//
func notify(c *cli.Context) error {

	msg, err := pipilineInfo()
	if err != nil {
		return err
	}

	switch c.String("channel") {
	case "telegram":
		TG_TOKEN := os.Getenv("TG_TOKEN")
		TG_CHAT_ID := os.Getenv("TG_CHAT_ID")
		bot := telegram.NewBot(TG_TOKEN)
		bot.SendMesage(TG_CHAT_ID, msg)
	default:
		return errors.New("Channel unknown")
	}
	return nil
}

// Get pipiline info from GitLab
func pipilineInfo() (string, error) {
	var MSG = "%s <b>The pipeline completed</b>\n\n" +
		"<b>Status:</b> %s\n" +
		"<b>Pipeline:</b> %s\n" +
		"<b>Branch:</b> %s\n" +
		"<b>User:</b> %s\n"

	GL_TOKEN := os.Getenv("GL_TOKEN")
	if GL_TOKEN == "" {
		return "", errors.New("Can't get GitLab token")
	}

	CI_SERVER_URL := os.Getenv("CI_SERVER_URL")
	CI_PROJECT_ID := os.Getenv("CI_PROJECT_ID")
	CI_PIPELINE_ID := os.Getenv("CI_PIPELINE_ID")

	gitlab := gitlab.GitLab{
		Token: GL_TOKEN,
	}

	pipeline, err := gitlab.GetPipeline(CI_SERVER_URL, CI_PROJECT_ID, CI_PIPELINE_ID)
	if err != nil {
		return "", err
	}

	var icon string
	switch pipeline.Status {
	case "running":
		icon = "✅"
	case "failed":
		icon = "‼️"
	default:
		icon = "⚠️"
	}

	MSG = fmt.Sprintf(MSG, icon, pipeline.Status, fmt.Sprintf("<a href=\"%s\">%s</a>", pipeline.WebURL, strconv.Itoa(pipeline.ID)), pipeline.Ref, pipeline.User.Username)
	return MSG, nil
}
