package handlers

import (
	"net/http"
	auth "projects/chatterbox/server/pkgs/auth"
	"projects/chatterbox/server/pkgs/dao"
	utils "projects/chatterbox/server/pkgs/utilities"

	"github.com/gin-gonic/gin"
)

type Handles struct {
	Dao dao.DAO
}

// register service
func (h Handles) Register(c *gin.Context) {
	var registerData dao.User

	err := c.BindJSON(&registerData)
	if err != nil {
		utils.FailureResponse(c, http.StatusBadRequest, utils.BadRequestMsg, err.Error())
		return
	}

	if registerData.Username == "" || registerData.Email == "" || registerData.Password == "" {
		utils.FailureResponse(c, http.StatusBadRequest, utils.BadRequestMsg, nil)
		return
	}

	isExistingUser, _, err := h.Dao.CheckExistingUser(registerData)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	if isExistingUser {
		utils.FailureResponse(c, http.StatusUnprocessableEntity, utils.UserExists, nil)
		return
	}

	err = h.Dao.AddUserToDatabase(registerData)
	if err != nil {
		utils.FailureResponse(c, http.StatusBadRequest, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, utils.UserRegistered, nil)
}

// login service
func (h Handles) Login(c *gin.Context) {
	var user dao.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.FailureResponse(c, http.StatusBadRequest, utils.BadRequestMsg, err.Error())
		return
	}

	isExisting, existingUser, err := h.Dao.CheckExistingUser(user)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	if !isExisting {
		utils.FailureResponse(c, http.StatusForbidden, utils.UserNotFound, nil)
		return
	}

	if match := auth.CompareHashPassword(user.Password, existingUser.Password); !match {
		utils.FailureResponse(c, http.StatusForbidden, utils.InvalidPassword, nil)
		return
	}

	sessionId, signedString, ttl, err := auth.SignToken(user)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, "generate token failed", nil)
		return
	}

	startSession(c, sessionId, signedString, ttl)
	utils.SuccessResponse(c, http.StatusOK, utils.UserLoggedIn, map[string]interface{}{"token": signedString})
}

func (h Handles) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	utils.SuccessResponse(c, http.StatusOK, utils.UserLoggedOut, nil)
}

func startSession(c *gin.Context, sessionId, signedString string, ttl int) {
	c.SetCookie("session_id", sessionId, int(ttl), "/", "localhost", false, true)
	c.SetCookie("auth_session_id", signedString, int(ttl), "/", "localhost", false, true)
	// c.SetCookie(sessionId, "session data", int(ttl), "/", "localhost", false, true)
}
