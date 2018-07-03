package user

import (
	. "ApiServer/demo09/handler"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log/lager"
	"ApiServer/demo09/pkg/errno"
	"github.com/lexkong/log"
	"ApiServer/demo09/model"
	"ApiServer/demo09/util"
)


// 创建一个新的用户账户
//func Create(c *gin.Context) {
//	var r struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//	}
//	var err error
//	if err := c.Bind(&r); err != nil {
//		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
//		return
//	}
//	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
//	if r.Username == "" {
//		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
//		log.Errorf(err, "Get an error")
//	}
//	if errno.IsErrUserNotFound(err) {
//		log.Debug("err type is ErrUserNotFound")
//	}
//	if r.Password == "" {
//		err = fmt.Errorf("password is empty")
//	}
//	code, message := errno.DecodeErr(err)
//	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
//}

// 第二版 创建一个新的用户账户
//func Create(c *gin.Context) {
//	var r CreateRequest
//	if err := c.Bind(&r); err != nil {
//		SendResponse(c, errno.ErrBind, nil)
//		return
//	}
//
//	admin2 := c.Param("username")
//	log.Infof("URL username: %s", admin2)
//
//	desc := c.Query("desc")
//	log.Infof("URL key param desc: %s", desc)
//
//	contentType := c.GetHeader("Content-Type")
//	log.Infof("Header Content-Type: %s", contentType)
//
//	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
//	if r.Username == "" {
//		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
//		return
//	}
//
//	if r.Password == "" {
//		SendResponse(c, fmt.Errorf("password is empty"), nil)
//	}
//
//	rsp := CreateResponse{
//		Username: r.Username,
//	}
//
//	// Show the user information.
//	SendResponse(c, nil, rsp)
//}

// 从 HTTP 消息体获取参数（用户名和密码）
// 参数校验
// 加密密码
// 在数据库中添加数据记录
// 返回结果（这里是用户名）


// 新增用户

// Create creates a new user account.
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}


// 校验方法 -- 伪代码
// Create creates a new user account.
//func Create(c *gin.Context) {
//	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
//	var r CreateRequest
//	if err := c.Bind(&r); err != nil {
//		SendResponse(c, errno.ErrBind, nil)
//		return
//	}
//
//	if err := r.checkParam(); err != nil {
//		SendResponse(c, err, nil)
//		return
//	}
//	...
//}

//func (r *CreateRequest) checkParam() error {
//	if r.Username == "" {
//		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
//	}
//
//	if r.Password == "" {
//		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
//	}
//
//	return nil
//}
