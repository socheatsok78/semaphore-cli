package semaphore

// /api/projects
// {name: "Demo", alert: true, alert_chat: "tfchatid", max_parallel_tasks: 0}
type SemaphoreProject struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Alert            bool   `json:"alert"`
	AlertChat        string `json:"alert_chat"`
	MaxParallelTasks int    `json:"max_parallel_tasks"`
	Type             string `json:"type"`
	Created          string `json:"created"`
}
