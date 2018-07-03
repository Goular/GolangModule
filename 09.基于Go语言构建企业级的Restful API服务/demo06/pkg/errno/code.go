package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// user errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
)

// 错误码设计
// 服务级别错误：1 为系统级错误；2 为普通错误，通常是由用户非法操作引起的
// 服务模块为两位数：一个大型系统的服务模块通常不超过两位数，如果超过，说明这个系统该拆分了
// 错误码为两位数：防止一个模块定制过多的错误码，后期不好维护
// code = 0 说明是正确返回，code > 0 说明是错误返回
// 错误通常包括系统级错误码和服务级错误码
// 建议代码中按服务模块将错误分类
// 错误码均为 >= 0 的数
// 在 apiserver 中 HTTP Code 固定为 http.StatusOK，错误码通过 code 来表示。

// 错误代码说明:
// 第一位 :		  1/2 		-----  服务级错误
// 第二三位:		 00-99		-----  服务模块代码
// 第四五位:       00-99      -----  具体错误代码


//如果 API 是对外的，错误信息数量有限，则制定错误码非常容易，强烈建议使用错误码。
// 如果是内部系统，特别是庞大的系统，内部错误会非常多，这时候没必要为每一个错误制定错误码，
// 而只需为常见的错误制定错误码，对于普通的错误，系统在处理时会统一作为 InternalServerError 处理。