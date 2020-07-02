package request

type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
	Search string `json:"search" form:"search"`
}

type TMachineUsersRequest struct {
	AdminId []int `json:"AdminId" form:"AdminId"`
	MachineId int `json:"machineId" form:"machineId"`
}
type RequestAdmin struct {
	Username     string `json:"username"  form:"username" `
	PasswordHash string `json:"password"  form:"password" `
	RoleId       int    `json:"roleId"  form:"roleId"`
	Status       int    `json:"status"  form:"status"`
	ResetPassword       string    `json:"ResetPassword"  form:"ResetPassword"`
	ID       int    `json:"id"  form:"id"`
}
func PageInfoData(pageInfo PageInfo) (info PageInfo ) {
	if pageInfo.Page==0 {
		info.Page=1
	}else {
		info.Page=pageInfo.Page
	}
	if pageInfo.PageSize==0 {
		info.PageSize=20
	}else {
		info.PageSize=pageInfo.PageSize
	}
	if len(pageInfo.Search)!=0{
		info.Search=pageInfo.Search
	}
	return info
}

