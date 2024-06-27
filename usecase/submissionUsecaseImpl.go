package usecase

import (
	"net/http"

	"golang-assignment/entity"
	"golang-assignment/repository"

	"github.com/gin-gonic/gin"
)

type SubmissionUsecaseImpl struct {
	SubmissionRepo repository.SubmissionRepository
	UserRepo       repository.UserRepository
}

func NewSubmissionUsecase(SubmissionRepo repository.SubmissionRepository, UserRepo repository.UserRepository) submissionUsecase {
	return &SubmissionUsecaseImpl{SubmissionRepo: SubmissionRepo, UserRepo: UserRepo}
}

func (usecase *SubmissionUsecaseImpl) FindAll(ctx *gin.Context) {
	submissions, err := usecase.SubmissionRepo.GetAll(ctx)
	if err != nil {
		// Handle the error and return an Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"error":  "",
		"data":   submissions,
	})
}

func (usecase *SubmissionUsecaseImpl) Create(ctx *gin.Context) {
	var submission entity.Submission

	if err := ctx.ShouldBindJSON(&submission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	// // Marshal the user struct to JSON
	// jsonData, err := json.Marshal(submission)
	// if err != nil {
	// 	log.Fatalf("Failed to marshal data to JSON: %v", err)
	// }

	// // Log the JSON data as a string
	// fmt.Println("LOGGING DATA:", string(jsonData))

	if _, err := usecase.UserRepo.GetById(ctx, submission.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "User id not found",
		})
		return
	}

	if err := usecase.SubmissionRepo.Create(ctx, &submission); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"error":  "",
		"data":   submission,
	})
}
