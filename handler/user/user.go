package user

import "apiserver-study/model"

//创建用户信息-请求
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//创建用户信息-返回
type CreateResponse struct {
	Username string `json:"username"`
}

//获取用户列表-请求
type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

//获取用户列表-返回
type ListResponse struct {
	TotalCount uint64            `json:"total_count"`
	UserList   []*model.UserInfo `json:"user_list"`
}
