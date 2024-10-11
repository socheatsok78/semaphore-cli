package semaphore

import "net/http"

type Semaphore struct {
	client *http.Client
}

func New(client *http.Client) *Semaphore {
	return &Semaphore{client: client}
}

func (s *Semaphore) Backup(projectID string) error {
	return nil
}

func (s *Semaphore) Restore(projectID string) error {
	return nil
}
