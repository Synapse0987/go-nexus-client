package repository

type Replication struct {
	PreemptivePullEnabled bool    `json:"preemptivePullEnabled"`
	AssetPathRegex        *string `json:"assetPathRegex,omitempty"`
}
