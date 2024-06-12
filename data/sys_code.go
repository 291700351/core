package data

type sysCode struct {
	Success    *ErrCode
	ErrSystem  *ErrCode
	ErrUnknown *ErrCode

	ErrMissParam       *ErrCode
	ErrNotFount        *ErrCode
	ErrUnSupportMethod *ErrCode
	ErrInvalidParam    *ErrCode

	//数据库
	ErrSave   *ErrCode
	ErrDelete *ErrCode
	ErrUpdate *ErrCode
	ErrQuery  *ErrCode

	//token
	ErrGenToken     *ErrCode
	ErrToken        *ErrCode
	ErrTokenInvalid *ErrCode
}

var SysCode = sysCode{
	Success:    NewErrCode(0, ""),
	ErrSystem:  NewErrCode(100, "系统错误,请稍候重试"),
	ErrUnknown: NewErrCode(101, "系统发生未知错误,请稍候重试"),
	// 请求相关
	ErrMissParam:       NewErrCode(200, "缺少必要参数,请检查后重试"),
	ErrNotFount:        NewErrCode(201, "未找到相关资源,请稍后重试"),
	ErrUnSupportMethod: NewErrCode(202, "不支持的请求方法,请检查后重试"),
	ErrInvalidParam:    NewErrCode(203, "请求参数无效,请检查后重试"),
	// 数据库 CRUD
	ErrSave:   NewErrCode(300, "保存数据发生错误,请稍后重试"),
	ErrDelete: NewErrCode(301, "删除数据发生错误,请稍后重试"),
	ErrUpdate: NewErrCode(302, "修改数据发生错误,请稍后重试"),
	ErrQuery:  NewErrCode(303, "查询数据发生错误,请稍后重试"),
	// Token
	ErrGenToken:     NewErrCode(400, "创建token失败,请稍候重试"),
	ErrToken:        NewErrCode(401, "请重新登陆"),
	ErrTokenInvalid: NewErrCode(402, "请重新登陆"),
}
