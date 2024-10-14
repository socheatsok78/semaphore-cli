package semaphore

import (
	"errors"
	"net/http"

	"github.com/socheatsok78/semaphore-cli/types"
)

func (s *Semaphore) Authenticate(username string, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password are required")
	}
	authJson := &types.SemaphoreAuth{
		Username: username,
		Password: password,
	}
	resp, err := s.Write("/api/auth/login", authJson)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return errors.New("login failed")
	}
	return nil
}
