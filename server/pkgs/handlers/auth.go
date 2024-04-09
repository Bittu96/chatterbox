package handlers

import (
	"fmt"
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

	if err := c.BindJSON(&registerData); err != nil {
		utils.FailureResponse(c, http.StatusBadRequest, utils.BadRequestMsg, err.Error())
		return
	}

	// validate := validator.New(validator.WithRequiredStructEnabled())

	if registerData.Username == "" || registerData.Email == "" || registerData.Password == "" {
		utils.FailureResponse(c, http.StatusBadRequest, utils.BadRequestMsg, nil)
		return
	}

	if isExisting, _, err := h.Dao.CheckExistingUser(registerData); err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	} else if isExisting {
		utils.FailureResponse(c, http.StatusBadRequest, utils.UserExists, nil)
		return
	}

	if len(registerData.Password) < 8 {
		utils.FailureResponse(c, http.StatusBadRequest, utils.MinPasswordLengthConflict, nil)
		return
	}

	if err := h.Dao.AddUserToDatabase(registerData); err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, utils.UserRegistered, nil)
}

// login service
func (h Handles) Login(c *gin.Context) {
	var user dao.User
	if err := c.ShouldBindJSON(&user); err != nil || user.Username == "" || user.Password == "" {
		fmt.Println("error:", err, user.Username == "", user.Password == "")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	isExisting, existingUser, err := h.Dao.CheckExistingUser(user)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, utils.DaoFailureMsg, err.Error())
		return
	} else if !isExisting {
		utils.FailureResponse(c, http.StatusBadRequest, utils.UnregisteredUser, nil)
		return
	}

	if match := auth.CompareHashPassword(user.Password, existingUser.Password); !match {
		utils.FailureResponse(c, http.StatusBadRequest, utils.InvalidLoginRequest, nil)
		return
	}

	sessionId, signedString, ttl, err := auth.SignToken(existingUser)
	if err != nil {
		utils.FailureResponse(c, http.StatusInternalServerError, "generate token failed", nil)
		return
	}

	startSession(c, sessionId, signedString, ttl)
	utils.SuccessResponse(c, http.StatusOK, utils.UserLoggedIn, map[string]interface{}{"token": signedString})
}

func (h Handles) Logout(c *gin.Context) {
	c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
	c.SetCookie("session_token", "", -1, "/", "localhost", false, true)
	utils.SuccessResponse(c, http.StatusOK, utils.UserLoggedOut, nil)
}

func startSession(c *gin.Context, sessionId, signedString string, ttl int) {
	c.SetCookie("session_id", sessionId, int(ttl), "/", "localhost", false, true)
	c.SetCookie("session_token", signedString, int(ttl), "/", "localhost", false, true)
	// c.SetCookie(sessionId, "session data", int(ttl), "/", "localhost", false, true)
}
