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
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://git.digital.rt.ru/api/v4/projects/%s/pipelines/%s", projectId, pipelineId), nil)
	req.Header.Add("PRIVATE-TOKEN", c.Token)
	resp, _ := client.Do(req)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &pipeline)
	return pipeline, nil
}
