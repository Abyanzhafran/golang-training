package repository

import (
	"context"

	"golang-assignment/entity"

	"gorm.io/gorm"
)

type SubmissionRepositoryImpl struct {
	DB *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &SubmissionRepositoryImpl{DB: db}
}

func (r *SubmissionRepositoryImpl) Create(ctx context.Context, submission *entity.Submission) error {
	return r.DB.Create(submission).Error
}

func (r *SubmissionRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Submission, error) {
	var submissions []*entity.Submission

	if err := r.DB.Find(&submissions).Error; err != nil {
		return nil, err
	}

	return submissions, nil
}
