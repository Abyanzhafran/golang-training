package repository

import (
	"context"

	"golang-assignment/entity"
)

type SubmissionRepository interface {
	GetAll(ctx context.Context) ([]*entity.Submission, error)
	Create(ctx context.Context, submission *entity.Submission) error
}
