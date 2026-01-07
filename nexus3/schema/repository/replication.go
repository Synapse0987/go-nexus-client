package repository

type Replication struct {
	enabled bool `json:"preemptivePullEnabled"`
	path_regex *string `json:"assetPathRegex,omitempty"`
}