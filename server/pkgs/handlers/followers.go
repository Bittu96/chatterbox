package handlers

import (
	"net/http"
	utils "projects/chatterbox/server/pkgs/utilities"

	"github.com/gin-gonic/gin"
)

// api to get all followers
func (h Handles) GetFollowers(c *gin.Context) {
	authUserId, found := c.Params.Get("auth_user_id")
	if !found || authUserId == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	followers, err := h.Dao.GetFollowers(authUserId)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "followers fetch success", followers)
}

func (h Handles) GetFollowing(c *gin.Context) {
	authUserId, found := c.Params.Get("auth_user_id")
	if !found || authUserId == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	following, err := h.Dao.GetFollowing(authUserId)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "following fetch success", following)
}

func (h Handles) FollowUser(c *gin.Context) {
	authUserId, found := c.Params.Get("auth_user_id")
	if !found || authUserId == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	followId := c.Query("follow_id")
	if followId == "" || authUserId == followId {
		utils.FailureResponse(c, http.StatusBadRequest, utils.BadRequestMsg, nil)
		return
	}

	if err := h.Dao.FollowUser(followId, authUserId); err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "follow success", nil)
}

func (h Handles) UnfollowUser(c *gin.Context) {
	authUserId, found := c.Params.Get("auth_user_id")
	if !found || authUserId == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	followId := c.Query("follow_id")
	if followId == "" || authUserId == followId {
		utils.FailureResponse(c, http.StatusBadRequest, utils.BadRequestMsg, nil)
		return
	}

	if err := h.Dao.UnfollowUser(followId, authUserId); err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "unfollow success", nil)
}
