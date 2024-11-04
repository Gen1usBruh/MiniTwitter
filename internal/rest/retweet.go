package rest

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Gen1usBruh/MiniTwitter/internal/logger/sl"
	"github.com/Gen1usBruh/MiniTwitter/internal/storage/models"
	postgresdb "github.com/Gen1usBruh/MiniTwitter/internal/storage/postgres/sqlc"
)

// createRetweet is the handler for creating a new retweet.
// @Summary Create a retweet
// @Description Create retweet to a tweet/retweet with quote
// @Tags Retweet
// @Accept json
// @Produce json
// @Param query_params query models.CreateRetweetStruct true "Query parameters"
// @Success 200 {object} postgresdb.tweet "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /retweets/ [post]
func (h *HandlerConfig) createRetweet(c *gin.Context) {
	var rq models.CreateRetweetStruct
	const op = "internal.rest.auth.createRetweet"
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error binding: ", op), sl.Err(err))
		c.JSON(400, models.BaseResponse{Error: "Wrong data", ErrorCode: 404})
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
	arg := postgresdb.CreateRetweetParams{}
	arg.UserID = userId
	arg.Quote.Scan(rq.Quote)
	arg.IsQuote = rq.IsQuote
	if rq.ParentRetweetID == 0 {
		arg.ParentTweetID.Scan(rq.ParentTweetID)
	} else if rq.ParentTweetID == 0 {
		arg.ParentRetweetID.Scan(rq.ParentRetweetID)
	} else {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error Parsing Request: ", op), sl.Err(err))
		c.JSON(400, models.BaseResponse{Error: err.Error()})
		return
	}
	retweet, err := h.Dep.Db.CreateRetweet(c.Request.Context(), arg)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s Error database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create Follow", ErrorCode: 500})
		return
	}
	c.JSON(200, gin.H{"retweet": retweet})
	return
}

// deleteRetweet is the handler for deleting an existing retweet.
// @Summary Delete a retweet
// @Description Delete a retweet by id
// @Tags Retweet
// @Accept json
// @Produce json
// @Param id path int true "ID of the retweet"
// @Success 200 {object} postgresdb.tweet "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /retweets/{id} [delete]
func (h *HandlerConfig) deleteRetweet(c *gin.Context) {
	const op = "internal.rest.auth.deleteRetweet"
	tweetIDStr := c.Param("retweet_id")
	tweetID, err := strconv.Atoi(tweetIDStr)
	if err != nil {
		c.JSON(400, models.BaseResponse{Error: err.Error()})
		return
	}
	err = h.Dep.Db.DeleteRetweet(c.Request.Context(), int32(tweetID))
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s Error database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create Follow", ErrorCode: 500})
		return
	}
	c.JSON(200, "retweet")
}

// getRetweet is the handler for selecting an existing retweet.
// @Summary Select a retweet
// @Description Select a retweet by id
// @Tags Retweet
// @Accept json
// @Produce json
// @Param id path int true "ID of the retweet"
// @Success 200 {object} postgresdb.tweet "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /retweets/{id} [get]
func (h *HandlerConfig) getRetweet(c *gin.Context) {
	const op = "internal.rest.auth.getRetweet"
	tweetIDStr := c.Param("retweet_id")
	tweetID, err := strconv.Atoi(tweetIDStr)
	if err != nil {
		c.JSON(400, models.BaseResponse{Error: err.Error()})
		return
	}
	retweet, err := h.Dep.Db.SelectRetweet(c.Request.Context(), int32(tweetID))
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s Error database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create Follow", ErrorCode: 500})
		return
	}
	c.JSON(200, gin.H{"retweet": retweet})
}

func (h *HandlerConfig) getRetweetRetweet(c *gin.Context) {

}

func (h *HandlerConfig) getRetweetTweet(c *gin.Context) {

}
