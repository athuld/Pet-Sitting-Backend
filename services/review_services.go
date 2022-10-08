package services

import (
	"pet-sitting-backend/domain/review"
	"pet-sitting-backend/utils/errors"
)

func SaveReview(review review.Reviews) *errors.RestErr {
	if err := review.InsertReviewToDB(); err != nil {
		return err
	}
	return nil
}

func FetchReviewForOwner(reviewReq review.Reviews) (*review.ReviewsWithSitter, *errors.RestErr) {
	result, err := reviewReq.GetReviewsForOwnerFromDB()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FetchAllReviewsForSitter(sitterId int64) (*[]review.ReviewsWithSitter, *errors.RestErr) {
	review := &review.Reviews{SitterId: sitterId}
	result, err := review.GetAllReviewsBySitter()
	if err != nil {
		return nil, err
	}
	return result, nil
}
