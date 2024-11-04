package rest

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Gen1usBruh/MiniTwitter/internal/logger/sl"
	"github.com/Gen1usBruh/MiniTwitter/internal/storage/models"
	postgresdb "github.com/Gen1usBruh/MiniTwitter/internal/storage/postgres/sqlc"
)

// follow is the handler for following a user.
// @Summary Follow a user
// @Description Follow a user by id
// @Tags Follow
// @Accept json
// @Produce json
// @Param query_params query models.FollowStruct true "Query parameters"
// @Success 200 {object} postgresdb.follow "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /follow/ [post]
func (h *HandlerConfig) follow(c *gin.Context) {
	var rq models.FollowStruct
	const op = "internal.rest.follow"
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error binding: ", op), sl.Err(err))
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
	followerID := int32(tok_num)
	arg := postgresdb.CreateFollowerParams{}
	arg.FollowerID = followerID
	arg.FollowingID = rq.FollowingID
	follow, err := h.Dep.Db.CreateFollower(c.Request.Context(), arg)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create Follow", ErrorCode: 500})
		return
	}
	c.JSON(200, gin.H{"follow": follow})
}

// unfollow is the handler for unfollowing a user.
// @Summary Unfollow a user
// @Description Unfollow a user by id
// @Tags Follow
// @Accept json
// @Produce json
// @Param query_params query models.FollowStruct true "Query parameters"
// @Success 200 {string} string "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /follow/ [delete]
func (h *HandlerConfig) unfollow(c *gin.Context) {
	var rq models.FollowStruct
	const op = "internal.rest.unfollow"
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error binding: ", op), sl.Err(err))
		return
	}
	token := c.Request.Header.Get("Authorization")
	tok_id, err := h.ParseToken(token[7:])
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error parse token: ", op), sl.Err(err))
		c.JSON(404, models.BaseResponse{Error: err.Error()})
		return
	}
	tok_num, err := strconv.Atoi(tok_id)
	followerID := int32(tok_num)
	arg := postgresdb.DeleteFollowerParams{}
	arg.FollowerID = followerID
	arg.FollowingID = rq.FollowingID
	err = h.Dep.Db.DeleteFollower(c.Request.Context(), arg)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create Follow", ErrorCode: 500})
		return
	}
	c.JSON(200, "unfollow")
}
