package dto

type (
	GetUserListReq struct {
		Limit  int `validate:"gte=0" form:"limit" json:"limit"`
		Offset int `form:"offset" json:"offset"`
	}

	// CreateUserReq 创建User 请求对象
	CreateUserReq struct {
		Name string `validate:"required" json:"name"`
		Sex  int    `validate:"required" json:"sex"`
		Age  int    `validate:"required" json:"age"`
		Home string `validate:"required" json:"home"`
	}

	// UpdateUserReq 修改User 请求对象
	UpdateUserReq struct {
		Name string `json:"name"`
		Sex  int    `json:"sex"`
		Age  int    `json:"age"`
		Home string `json:"home"`
	}

	// DelUsersReq 删除 请求对象
	DelUsersReq struct {
		Id int64 ` validate:"required" uri:"id" json:"id"` // '用户id'
	}

	// FindUserByIdReq 查询 请求对象
	FindUserByIdReq struct {
		Id int64 ` validate:"required" uri:"id" json:"id"` // '用户id'
	}

	// UpdateUserByIdReq 更新 请求对象
	UpdateUserByIdReq struct {
		Id int64 ` validate:"required" uri:"id" json:"id"` // '用户id'
	}
)

type (
	GetUserByIdRsp struct {
		Id   int64  `json:"id"`
		Name string `json:"name,omitempty"`
		Sex  int    `json:"sex"`
		Age  int    `json:"age"`
		Home string `json:"home"`
	}

	GetUserByQueryRsp struct {
		Id   int64  `json:"id"`
		Name string `json:"name,omitempty"`
		Sex  int    `json:"sex"`
		Age  int    `json:"age"`
		Home string `json:"home"`
	}

	GetUserListRsp struct {
		Id   int64  `json:"id"`
		Name string `json:"name,omitempty"`
		Sex  int    `json:"sex"`
		Age  int    `json:"age"`
		Home string `json:"home"`
	}
)
