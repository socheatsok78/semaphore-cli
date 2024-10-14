package semaphore

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-kit/log/level"
	"github.com/socheatsok78/semaphore-cli/internals"
)

type SemaphoreBackup struct {
	Meta *SemaphoreProject `json:"meta"`
}

func (s *Semaphore) Backup(projectID string, backupFile string) error {
	level.Info(internals.Logger).Log("msg", "Creating backup", "project", projectID)
	resp, err := s.Read(fmt.Sprintf("/api/project/%s/backup", projectID))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Backup failed")
	} else {
		if backupFile == "stdout" {
			// Write to stdout
			level.Info(internals.Logger).Log("msg", "Writing backup to console", "project", projectID)
			io.Copy(os.Stdout, resp.Body)
		} else {
			// Write to file
			level.Info(internals.Logger).Log("msg", "Writing backup to file", "project", projectID, "file", backupFile)
			file, err := os.Create(backupFile)
			if err != nil {
				return err
			}
			defer file.Close()
			io.Copy(file, resp.Body)
		}
	}
	return nil
}

func (s *Semaphore) Restore(projectID string, backupFile string) error {
	return nil
}
