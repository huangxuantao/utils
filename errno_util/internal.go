package errno_util

import "errors"

var (
	EmptySQL          = errors.New("sql was empty")
	DBNotInit         = errors.New("DB connection not initialized")
	NoDataNeedUpdate  = errors.New("数据表中没有传入要更新的键值")
	InsertDataIllegal = errors.New("插入数据非法")
	UpdateDataIllegal = errors.New("更新数据非法")
	DeleteDataIllegal = errors.New("删除数据非法")
	NotFoundContext   = errors.New("not found Context")
	NotFoundUserName  = errors.New("not found username")

	HttpNoResponseContentReturn = errors.New("no response content return")
	HttpNoResultData            = errors.New("no result data found")

	AuthTokenNotFound       = errors.New("token not found")
	AuthSysNameInconsistent = errors.New("sys-name inconsistent")
	AuthIPInconsistent      = errors.New("client ip inconsistent")
	AuthIssuerInconsistent  = errors.New("issuer inconsistent")
	AuthSysNameNotFound     = errors.New("no sys-name found")
	AuthSysNameNotApplied   = errors.New("sys-name not applied")
)
