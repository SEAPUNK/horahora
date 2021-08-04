package archiverequests

import (
	"database/sql"

	"github.com/SEAPUNK/horahora/archiver/internal/config"
)

type ArchiveRequest struct {
	Id    int64
	Query string
	Error sql.NullString
}

type UserArchiveRequest struct {
	ArchiveRequest
	UserId int64
}

func CreateArchiveRequest(cfg *config.Config, uar *UserArchiveRequest) (int64, error) {
	db := cfg.PostgresConn
	// TODO(ivan): security implications of processing with raw URL

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	var arId int64
	err = tx.QueryRow(`
		INSERT INTO archive_requests (query) VALUES ($1) RETURNING id
	`, uar.Query).Scan(&arId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(`
		INSERT INTO user_archive_requests (user_id, archive_id) VALUES ($1, $2)
	`, uar.UserId, arId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return arId, nil
}

func GetArchiveRequest(cfg *config.Config, archiveId int64) (*ArchiveRequest, error) {
	db := cfg.PostgresConn

	var ar *ArchiveRequest
	err := db.Get(&ar, `
		SELECT id, query, error FROM archive_requests WHERE id = $1
	`, archiveId)

	return ar, err
}

func ListArchiveRequestsForUser(cfg *config.Config, userId int64) ([]*ArchiveRequest, error) {
	db := cfg.PostgresConn

	var entries []*ArchiveRequest

	err := db.Select(&entries, `
		SELECT
			a.id as id,
			a.query as query,
			a.error as error
		FROM archive_requests a
		LEFT JOIN user_archive_requests u
			ON u.archive_id = a.id
		WHERE u.user_id = $1
	`, userId)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func QueueArchiveRequest(cfg *config.Config, archiveId int64) error {
	// TODO(ivan): implement this
	return nil
}
