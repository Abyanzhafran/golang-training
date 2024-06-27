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

func (r *SubmissionRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Submission, error) {
	var submission *entity.Submission
	
	if err := r.DB.Where("id = ?", id).First(&submission).Error; err != nil {
		return nil, err
	}

	return submission, nil
}

func (r *SubmissionRepositoryImpl) Delete(ctx context.Context, id int64) error {
	return r.DB.Where("id = ?", id).Delete(&entity.Submission{}).Error
}
