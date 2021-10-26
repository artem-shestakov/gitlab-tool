package gitlab

type GitLab struct {
	Token string
}

type GitLabPipeline struct {
	ID             int                  `json:"id"`
	ProjectID      int                  `json:"project_id"`
	SHA            string               `json:"sha"`
	Ref            string               `json:"ref"`
	Status         string               `json:"status"`
	CretaedAt      string               `json:"created_at"`
	UpdatedAt      string               `json:"updated_at"`
	WebURL         string               `json:"web_url"`
	BeforeSHA      string               `json:"before_sha"`
	Tag            bool                 `json:"tag"`
	YAMLErrors     bool                 `json:"yaml_errors"`
	User           GitLabUser           `json:"user"`
	StartedAt      string               `json:"started_at"`
	FinishedAt     string               `json:"finished_at"`
	CommittedAt    string               `json:"committed_at"`
	Duration       int                  `json:"duration"`
	QueuedDuration int                  `json:"queued_duration"`
	Coverage       string               `json:"coverage"`
	DetailedStatus GitLabPipelineStatus `json:"detailed_status"`
}

type GitLabPipelineStatus struct {
	Icon        string `json:"icon"`
	Text        string `json:"text"`
	Label       string `json:"label"`
	Group       string `json:"group"`
	Tooltip     string `json:"tooltip"`
	HasDetails  bool   `json:"has_details"`
	DetailsPath string `json:"details_path"`
}

type GitLabUser struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	State        string `json:"state"`
	AvatarURL    string `json:"avatar_url"`
	WebURL       string `json:"web_url"`
	Illustration string `json:"illustration"`
	Favicon      string `json:"favicon"`
}
