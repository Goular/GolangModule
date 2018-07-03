package user

import (
	"strconv"
	. "ApiServer/demo09/handler"
	"github.com/gin-gonic/gin"
	"ApiServer/demo09/pkg/errno"
	"ApiServer/demo09/model"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
