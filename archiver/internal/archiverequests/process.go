package archiverequests

import (
	"os/exec"

	"github.com/SEAPUNK/horahora/archiver/internal/config"
)

func ProcessArchiveRequest(cfg *config.Config, archiveId int64) error {
	err := _processArchiveRequest(cfg, archiveId)
	if err != nil {
		// TODO(ivan): update db with error
		return err
	}

	return nil
}

func _processArchiveRequest(cfg *config.Config, archiveId int64) error {
	_, err := GetArchiveRequest(cfg, archiveId)
	if err != nil {
		return err
	}

	// TODO(ivan): get the query from the archive request
	

	err := extractQueryVideos(query)

	return nil
}

func extractQueryVideos(query string) error {
	// TODO(ivan): install youtube-dl
	cmd := exec.Command("youtube-dl")
}
