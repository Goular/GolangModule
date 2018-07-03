package user

import (
	. "ApiServer/demo07/handler"
	"github.com/gin-gonic/gin"
	"ApiServer/demo07/pkg/errno"
	"ApiServer/demo07/service"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
