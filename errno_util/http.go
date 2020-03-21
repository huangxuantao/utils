package errno_util

var (
	OK = &Errno{Code: 200, Message: "操作正常"}

	InternalServer = &Errno{Code: 10001, Message: "内部服务器错误"}
	Bind           = &Errno{Code: 10002, Message: "将请求内容绑定到结构体时发生错误"}
	RequestMethod  = &Errno{Code: 10003, Message: "请求方法错误"}
	RequestParams  = &Errno{Code: 10004, Message: "请求参数错误"}

	DatabaseInsert = &Errno{Code: 20001, Message: "数据库插入失败"}
	DatabaseQuery  = &Errno{Code: 20002, Message: "数据库查询失败"}
	DatabaseUpdate = &Errno{Code: 20003, Message: "数据库更新失败"}
	DatabaseDelete = &Errno{Code: 20004, Message: "数据库删除失败"}

	AuthNotAllowed       = &Errno{Code: 403, Message: "用户无权限进行此操作"}
	AuthTokenCanNotGet   = &Errno{Code: 403, Message: "您无权限进行此操作"}
	AuthTokenNotValidYet = &Errno{Code: 403, Message: "您无权限进行此操作"}
	AuthTokenTimeout     = &Errno{Code: 403, Message: "您无权限进行此操作"}
	AuthTokenIllegal     = &Errno{Code: 403, Message: "您无权限进行此操作"}
	AuthNotAllowedLogin  = &Errno{Code: 403, Message: "您无权限登录此系统"}

	AuthCanNotGetToken = &Errno{Code: 20102, Message: "获取Token失败"}

	JsonMarshal = &Errno{Code: 20201, Message: "Json Marshal失败"}

	WebsocketUpgrade = &Errno{Code: 20301, Message: "Http update to WebSocket Failed"}

	UserLoginFailed     = &Errno{Code: 30001, Message: "用户名或密码错误"}
	UsernameNotExisted  = &Errno{Code: 30002, Message: "用户名不存在"}
	OriginPassword      = &Errno{Code: 30003, Message: "原密码错误"}
	NewPassword         = &Errno{Code: 30004, Message: "新密码错误"}
	UserNotExisted      = &Errno{Code: 30005, Message: "用户不存在"}
	UsernameHaveExisted = &Errno{Code: 30006, Message: "用户名已存在"}
	UsernameError       = &Errno{Code: 30007, Message: "不允许修改其他人密码"}
	PasswordUpdate      = &Errno{Code: 30008, Message: "更新密码失败"}

	RoleNotExisted  = &Errno{Code: 32001, Message: "角色不存在"}
	RoleHaveExisted = &Errno{Code: 32002, Message: "角色已存在"}

	UserRoleBind   = &Errno{Code: 33001, Message: "用户与角色绑定错误"}
	UserRoleUnBind = &Errno{Code: 33002, Message: "用户与角色解绑错误"}
	UserOrgBind    = &Errno{Code: 33003, Message: "用户与组织绑定错误"}
	UserOrgUnBind  = &Errno{Code: 33004, Message: "用户与组织解绑错误"}

	ChipOrgNotExisted     = &Errno{Code: 34001, Message: "码房不存在"}
	ChipOrgHaveExisted    = &Errno{Code: 34002, Message: "码房已存在"}
	ChipOrgNotAllowDelete = &Errno{Code: 34003, Message: "码房不允许删除"}
	ChipOrgTopNotAll      = &Errno{Code: 34004, Message: "总码房只能有一个"}
	ChipOrgTopAdd         = &Errno{Code: 34006, Message: "总码房入库操作失败"}
	ChipHistory           = &Errno{Code: 34007, Message: "写入历史记录失败"}
	ChipReserveLess       = &Errno{Code: 34008, Message: "库存不足"}
	ChipOrgMudLess        = &Errno{Code: 34009, Message: "码房公司码不足"}
	ChipOrgCashLess       = &Errno{Code: 34010, Message: "码房现金码不足"}
	ChipShuffleData       = &Errno{Code: 34011, Message: "码房洗码码数不对等"}

	CardOrgNotExisted     = &Errno{Code: 35001, Message: "牌房不存在"}
	CardOrgHaveExisted    = &Errno{Code: 35002, Message: "牌房已存在"}
	CardOrgNotAllowDelete = &Errno{Code: 35003, Message: "牌房不允许删除"}
	CardOrgTopNotAll      = &Errno{Code: 35004, Message: "总牌房只能有一个"}
	CardOrgTopAllow       = &Errno{Code: 35005, Message: "只有总牌房允许入库操作"}
	CardOrgTopAdd         = &Errno{Code: 35006, Message: "总牌房入库操作失败"}
	CardHistory           = &Errno{Code: 35007, Message: "写入历史记录失败"}
	CardUseError          = &Errno{Code: 35008, Message: "牌房使用参数有误"}

	OrgNotExisted     = &Errno{Code: 36001, Message: "组织不存在"}
	OrgHaveExisted    = &Errno{Code: 36002, Message: "组织已存在"}
	OrgNotAllowDelete = &Errno{Code: 36003, Message: "组织不允许删除"}
	OrgTopNotAll      = &Errno{Code: 36004, Message: "顶级组织只能有一个"}

	RequestNotExisted    = &Errno{Code: 37001, Message: "请求单不存在"}
	RequestCreate        = &Errno{Code: 37002, Message: "请求单创建错误"}
	RequestCategory      = &Errno{Code: 37003, Message: "请求单类型错误"}
	RequestCreateParams  = &Errno{Code: 37004, Message: "请求单创建参数错误"}
	RequestNotPermission = &Errno{Code: 37005, Message: "没有权限进行此操作"}
	RequestDataError     = &Errno{Code: 37006, Message: "请求单数据有误"}

	MemberNotExisted  = &Errno{Code: 38001, Message: "会员不存在"}
	MemberHaveExisted = &Errno{Code: 38002, Message: "会员已存在"}
	MemberCommission  = &Errno{Code: 38003, Message: "结算佣金错误"}
	MemberHistory     = &Errno{Code: 38004, Message: "写入历史记录失败"}
	MemberSaveError   = &Errno{Code: 38005, Message: "会员存码数据有误"}
	MemberMudLess     = &Errno{Code: 38006, Message: "会员公司码不足"}
	MemberCashLess    = &Errno{Code: 38007, Message: "会员现金码不足"}

	ChipCurrencyHaveExisted   = &Errno{Code: 39001, Message: "筹码货币中文名重复"}
	ChipCurrencyEnHaveExisted = &Errno{Code: 39002, Message: "筹码货币英文名重复"}
)
