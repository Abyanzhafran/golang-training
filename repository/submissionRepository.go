package repository

import (
	"context"

	"golang-assignment/entity"
)

type SubmissionRepository interface {
	GetAll(ctx context.Context) ([]*entity.Submission, error)
	GetById(ctx context.Context, id int) (*entity.Submission, error)
	Create(ctx context.Context, submission *entity.Submission) error
	Delete(ctx context.Context, id int) error
}
