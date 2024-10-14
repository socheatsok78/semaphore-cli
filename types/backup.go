package types

type SemaphoreBackup struct {
	Meta         SemaphoreBackupProject
	Templates    []SemaphoreBackupProjectTemplate    `json:"templates"`
	Repositories []SemaphoreBackupProjectRepository  `json:"repositories"`
	Keys         []SemaphoreBackupProjectKey         `json:"keys"`
	Views        []SemaphoreBackupProjectView        `json:"views"`
	Inventories  []SemaphoreBackupProjectInventory   `json:"inventories"`
	Environments []SemaphoreBackupProjectEnvironment `json:"environments"`
}

type SemaphoreBackupProject struct {
	Name             string `json:"name"`
	Alert            bool   `json:"alert"`
	AlertChat        string `json:"alert_chat"`
	MaxParallelTasks int    `json:"max_parallel_tasks"`
}

type SemaphoreBackupProjectTemplate = map[string]interface{}

type SemaphoreBackupProjectRepository struct {
	Name   string `json:"name"`
	Url    string `json:"git_url"`
	Branch string `json:"git_branch"`
	SSHKey string `json:"ssh_key"`
}

type SemaphoreBackupProjectKey struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SemaphoreBackupProjectView struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
}

type SemaphoreBackupProjectInventory struct {
	Name      string `json:"name"`
	Inventory string `json:"inventory"`
	SSHKey    string `json:"ssh_key"`
	BecomeKey string `json:"become_key"`
	Type      string `json:"type"`
}

type SemaphoreBackupProjectEnvironment struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Json     string `json:"json"`
	Env      string `json:"env"`
}
