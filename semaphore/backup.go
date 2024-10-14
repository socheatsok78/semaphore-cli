package semaphore

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (s *Semaphore) Backup(projectID string, backupFile string) (io.ReadCloser, error) {
	resp, err := s.Read(fmt.Sprintf("/api/project/%s/backup", projectID))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Backup failed")
	}
	return resp.Body, nil
}
