package types

type SemaphoreProject struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name"`
	Alert            bool   `json:"alert,omitempty"`
	AlertChat        string `json:"alert_chat,omitempty"`
	MaxParallelTasks int    `json:"max_parallel_tasks,omitempty"`
	Type             string `json:"type,omitempty"`
	Created          string `json:"created,omitempty"`
}

type SemaphoreProjectTemplate map[string]interface{}

func (t SemaphoreProjectTemplate) Get(key string) interface{} {
	return t[key]
}

func (t SemaphoreProjectTemplate) Has(key string) bool {
	_, ok := t[key]
	return ok
}

func (t SemaphoreProjectTemplate) Set(key string, value interface{}) {
	t[key] = value
}

func (t SemaphoreProjectTemplate) Del(key string) {
	delete(t, key)
}

type SemaphoreProjectRepository struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"git_url"`
	Branch string `json:"git_branch"`
	SSHKey string `json:"ssh_key"`
}

type SemaphoreProjectKey struct {
	ID             int                              `json:"id"`
	ProjectID      int                              `json:"project_id"`
	Name           string                           `json:"name"`
	Type           string                           `json:"type"`
	OverrideSecret bool                             `json:"override_secret"`
	LoginPassword  SemaphoreProjectKeyLoginPassword `json:"login_password"`
	SSHKey         SemaphoreProjectKeySSH           `json:"ssh"`
}

type SemaphoreProjectKeyLoginPassword struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SemaphoreProjectKeySSH struct {
	Login      string `json:"login"`
	Passphrase string `json:"passphrase"`
	PrivateKey string `json:"private_key"`
}

type SemaphoreProjectView struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

type SemaphoreProjectInventory struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Inventory string `json:"inventory"`
	SSHKey    string `json:"ssh_key"`
	BecomeKey string `json:"become_key"`
	Type      string `json:"type"`
}

type SemaphoreProjectEnvironment struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Json     string `json:"json"`
	Env      string `json:"env"`
}
