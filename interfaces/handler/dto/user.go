package dto

type (
	GetUserListReq struct {
		Limit  int `validate:"gte=0" form:"limit" json:"limit"`
		Offset int `form:"offset" json:"offset"`
	}

	AddUserReq struct {
		Name string `validate:"required" json:"name"`
		Sex  int    `validate:"required" json:"sex"`
		Age  int    `validate:"required" json:"age"`
		Home string `validate:"required" json:"home"`
	}

	UpdateUserReq struct {
		Name string `json:"name"`
		Sex  int    `json:"sex"`
		Age  int    `json:"age"`
		Home string `json:"home"`
	}
)

type (
	GetUserByIdRsp struct {
		Id   uint64 `json:"id"`
		Name string `json:"name,omitempty"`
		Sex  int    `json:"sex"`
		Age  int    `json:"age"`
		Home string `json:"home"`
	}

	GetUserByQueryRsp struct {
		Id   uint64 `json:"id"`
		Name string `json:"name,omitempty"`
		Sex  int    `json:"sex"`
		Age  int    `json:"age"`
		Home string `json:"home"`
	}

	GetUserListRsp struct {
		Id   uint64 `json:"id"`
		Name string `json:"name,omitempty"`
		Sex  int    `json:"sex"`
		Age  int    `json:"age"`
		Home string `json:"home"`
	}
)
