package handlers

import (
	"net/http"
	utils "projects/chatterbox/server/pkgs/utilities"

	"github.com/gin-gonic/gin"
)

func (h Handles) GetUsers(c *gin.Context) {
	// userId, found := c.Params.Get("auth_user_id")
	// if !found || userId == "" {
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"message": "user_id not found",
	// 	})
	// 	return
	// }

	users, err := h.Dao.GetAllUsersFromDatabase()
	if err != nil {
		utils.FailureResponse(c, http.StatusBadRequest, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "user fetched", users)
}

func (h Handles) Home(c *gin.Context) {
	auth_user_id, found := c.Params.Get("auth_user_id")
	if !found || auth_user_id == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	users, err := h.Dao.GetAllUserProfilesFromDatabase(auth_user_id)
	if err != nil {
		utils.FailureResponse(c, http.StatusBadRequest, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "user fetched", users)
}

func (h Handles) Premium(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "premium page", "role": ""})
}
