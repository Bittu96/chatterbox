package handlers

import (
	"fmt"
	"net/http"
	utils "projects/chatterbox/server/pkgs/utilities"

	"github.com/gin-gonic/gin"
)

// api to get all followers
func (h Handles) GetFollowers(c *gin.Context) {
	userId, found := c.Params.Get("auth_user_id")
	if !found || userId == "" {
		utils.FailureResponse(c, http.StatusForbidden, utils.UserNotFound, nil)
		return
	}

	followers, err := h.Dao.GetFollowers(userId)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "followers fetch success", followers)
}

func (h Handles) GetFollowing(c *gin.Context) {
	userId, found := c.Params.Get("auth_user_id")
	if !found || userId == "" {
		utils.FailureResponse(c, http.StatusForbidden, utils.UserNotFound, nil)
		return
	}

	following, err := h.Dao.GetFollowing(userId)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "following fetch success", following)
}

func (h Handles) FollowUser(c *gin.Context) {
	followerId, found := c.Params.Get("auth_user_id")
	if !found || followerId == "" {
		utils.FailureResponse(c, http.StatusForbidden, utils.UserNotFound, nil)
		return
	}

	followingId := c.Query("following_id")
	if followingId == "" {
		utils.FailureResponse(c, http.StatusUnprocessableEntity, utils.FollowingUserNotFound, nil)
		return
	}

	if followerId == followingId {
		utils.FailureResponse(c, http.StatusUnprocessableEntity, utils.FollowingUserNotFound, nil)
		return
	}

	err := h.Dao.FollowUser(followingId, followerId)
	if err != nil {
		fmt.Println(err)
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "follow success", nil)
}

func (h Handles) UnfollowUser(c *gin.Context) {
	followerId, found := c.Params.Get("auth_user_id")
	if !found || followerId == "" {
		utils.FailureResponse(c, http.StatusForbidden, utils.UserNotFound, nil)
		return
	}

	followingId := c.Query("following_id")
	if followingId == "" {
		utils.FailureResponse(c, http.StatusUnprocessableEntity, utils.FollowingUserNotFound, nil)
		return
	}

	err := h.Dao.UnfollowUser(followingId, followerId)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "unfollow success", nil)
}
