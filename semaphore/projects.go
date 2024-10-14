package semaphore

import "github.com/socheatsok78/semaphore-cli/types"

func NewProject(name string, maxParallelTasks int) *types.SemaphoreProject {
	return &types.SemaphoreProject{
		Name:             name,
		MaxParallelTasks: maxParallelTasks,
	}
}
