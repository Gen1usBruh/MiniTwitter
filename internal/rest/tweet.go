package rest

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Gen1usBruh/MiniTwitter/internal/logger/sl"
	"github.com/Gen1usBruh/MiniTwitter/internal/storage/models"
	postgresdb "github.com/Gen1usBruh/MiniTwitter/internal/storage/postgres/sqlc"
)

// createTweet is the handler for creating a new tweet.
// @Summary Create a tweet
// @Description Create tweet with content and media
// @Tags Tweet
// @Accept json
// @Produce json
// @Param query_params query models.CreateTweetStruct true "Query parameters"
// @Success 200 {object} postgresdb.tweet "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /tweets/ [post]
func (h *HandlerConfig) createTweet(c *gin.Context) {
	var rq models.CreateTweetStruct
	const op = "internal.rest.auth.createTweet"
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error binding: ", op), sl.Err(err))
		c.JSON(400, models.BaseResponse{Error: "Wrong data", ErrorCode: 400})
		return
	}
	token := c.Request.Header.Get("Authorization")
	tok_id, err := h.ParseToken(token[7:])
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error parse token: ", op), sl.Err(err))
		c.JSON(400, models.BaseResponse{Error: err.Error()})
		return
	}
	tok_num, err := strconv.Atoi(tok_id)
	userId := int32(tok_num)
	arg := postgresdb.CreateTweetParams{}
	arg.UserID = userId
	arg.Content = rq.Content
	arg.Media = rq.Media
	tweet, err := h.Dep.Db.CreateTweet(c.Request.Context(), arg)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s Error database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create Follow", ErrorCode: 500})
		return
	}
	c.JSON(200, gin.H{"tweet": tweet})
}

// deleteTweet is the handler for deleting a tweet.
// @Summary Deleting a tweet
// @Description Delete tweet by ID
// @Tags Tweet
// @Accept json
// @Produce json
// @Param tweet_id path int true "ID of the tweet"
// @Success 200 {string} string "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /tweets/{tweet_id} [delete]
func (h *HandlerConfig) deleteTweet(c *gin.Context) {
	const op = "internal.rest.auth.deleteTweet"
	tweetIDStr := c.Param("tweet_id")
	tweetID, err := strconv.Atoi(tweetIDStr)
	if err != nil {
		c.JSON(400, models.BaseResponse{Error: err.Error()})
		return
	}
	err = h.Dep.Db.DeleteTweet(c.Request.Context(), int32(tweetID))
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s Error database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create Follow", ErrorCode: 500})
		return
	}
	c.JSON(200, "tweet")
}

// getTweet is the handler for selecting a tweet.
// @Summary Selecting a tweet
// @Description Select tweet by ID
// @Tags Tweet
// @Accept json
// @Produce json
// @Param tweet_id path int true "ID of the tweet"
// @Success 200 {object} postgresdb.tweet "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /tweets/{tweet_id} [get]
func (h *HandlerConfig) getTweet(c *gin.Context) {
	const op = "internal.rest.auth.getTweet"
	tweetIDStr := c.Param("tweet_id")
	tweetID, err := strconv.Atoi(tweetIDStr)
	if err != nil {
		c.JSON(400, models.BaseResponse{Error: err.Error()})
		return
	}
	tweet, err := h.Dep.Db.SelectTweet(c.Request.Context(), int32(tweetID))
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s Error database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create Follow", ErrorCode: 500})
		return
	}
	c.JSON(200, gin.H{"tweet": tweet})
}
