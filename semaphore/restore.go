package semaphore

import "github.com/socheatsok78/semaphore-cli/types"

func (s *Semaphore) Restore(projectID string, backup *types.SemaphoreBackup) error {
	resp, err := s.Write("/api/project", backup.Meta)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
