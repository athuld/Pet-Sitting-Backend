package review

import (
	"context"
	"pet-sitting-backend/datasource"
	"pet-sitting-backend/utils/errors"
	"pet-sitting-backend/utils/logger"

	"github.com/randallmlough/pgxscan"
)

var(
    queryInsertReview = "insert into reviews(sitter_id,user_id,rating,review) values ($1,$2,$3,$4)"
    queryGetReviewsForOwner= "select r.*,name,address,pincode,phone,avatar_img from reviews r inner join userdetails ud on r.sitter_id=ud.user_id where r.user_id=$1 and sitter_id=$2"
    queryGetAllReviewsBySitter= "select r.*,name,address,pincode,phone,avatar_img from reviews r inner join userdetails ud on r.user_id=ud.user_id where sitter_id=$1"
    queryGetAllReviews="select r.sitter_id,name,address,pincode,phone,avatar_img,cast((avg(rating)) as decimal(10,1)) as avg_rating from reviews r inner join userdetails ud on r.sitter_id=ud.user_id group by r.sitter_id,name,address,phone,avatar_img,pincode"
)

func (review *Reviews) InsertReviewToDB() (*errors.RestErr){
    _,err:= datasource.Client.Query(context.Background(),queryInsertReview,review.SitterId,review.UserId,review.Rating,review.Review)
    if err!=nil{
        return errors.NewBadRequestError("Database error")
    }
    return nil
}

func (review *Reviews) GetReviewsForOwnerFromDB() (*ReviewsWithSitter,*errors.RestErr){
    result,err:=datasource.Client.Query(context.Background(),queryGetReviewsForOwner,review.UserId,review.SitterId)
    if err!=nil{
        return nil,errors.NewBadRequestError("Database error")
    }
    var reviewData ReviewsWithSitter
    if err:= pgxscan.NewScanner(result).Scan(&reviewData);err!=nil{
        return nil,errors.NewBadRequestError("Failed to scan")
    }
    return &reviewData,nil
}

func (review *Reviews) GetAllReviewsBySitter() (*[]ReviewsWithSitter,*errors.RestErr){
    result,err:= datasource.Client.Query(context.Background(),queryGetAllReviewsBySitter,review.SitterId)
    if err!=nil{
        logger.Error.Println(err)
        return nil,errors.NewBadRequestError("Database error")
    }
    var reviewData []ReviewsWithSitter
    if err:= pgxscan.NewScanner(result).Scan(&reviewData);err!=nil{
        return nil,errors.NewBadRequestError("Failed to scan")
    }
    return &reviewData,nil
}

func (review *Reviews) GetAllReviewsGroup() (*[]ReviewsWithSitter,*errors.RestErr){
    result,err:= datasource.Client.Query(context.Background(),queryGetAllReviews)
    if err!=nil{
        logger.Error.Println(err)
        return nil,errors.NewBadRequestError("Database error")
    }
    var reviewData []ReviewsWithSitter
    if err:= pgxscan.NewScanner(result).Scan(&reviewData);err!=nil{
        logger.Error.Println(err)
        return nil,errors.NewBadRequestError("Failed to scan")
    }
    return &reviewData,nil
}
