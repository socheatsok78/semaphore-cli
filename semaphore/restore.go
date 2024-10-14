package semaphore

func (s *Semaphore) Restore(projectID string, backup *SemaphoreBackup) error {
	resp, err := s.Write("/api/project", backup.Meta)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
