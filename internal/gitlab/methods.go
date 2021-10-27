package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *GitLab) GetPipeline(serverURL string, projectId string, pipelineId string) (GitLabPipeline, error) {
	pipeline := GitLabPipeline{}
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://git.digital.rt.ru/api/v4/projects/%s/pipelines/%s", projectId, pipelineId), nil)
	if err != nil {
		return pipeline, err
	}
	req.Header.Add("PRIVATE-TOKEN", c.Token)
	resp, err := client.Do(req)
	if err != nil {
		return pipeline, err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &pipeline)
	if err != nil {
		return pipeline, err
	}
	return pipeline, nil
}
