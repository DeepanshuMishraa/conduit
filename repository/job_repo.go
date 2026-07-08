package repository

import (
	"context"
	"database/sql"
	"github.com/DeepanshuMishraa/conduit.git/models"
	"time"
)

func CreateJob(job models.Job, db *sql.DB) (*models.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO jobs (title) VALUES ($1) RETURNING id, title, created_at, updated_at`

	var createdJob models.Job

	err := db.QueryRowContext(ctx, query, job.Title).Scan(&createdJob.ID, &createdJob.Title, &createdJob.CreatedAt, &createdJob.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &createdJob, nil
}
