package semaphore

type SemaphoreProject struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name"`
	Alert            bool   `json:"alert,omitempty"`
	AlertChat        string `json:"alert_chat,omitempty"`
	MaxParallelTasks int    `json:"max_parallel_tasks,omitempty"`
	Type             string `json:"type,omitempty"`
	Created          string `json:"created,omitempty"`
}

func NewProject(name string, maxParallelTasks int) *SemaphoreProject {
	return &SemaphoreProject{
		Name:             name,
		MaxParallelTasks: maxParallelTasks,
	}
}
