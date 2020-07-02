package request

type Permissions struct {
	Permissions []int  `json:"permissionIds" form:"permissionIds"`
	roleId        int  `json:"roleId" form:"roleId"`
}
