package system

import (
	"server/global"
	"server/models/system"
)

//GetUserListService 获取用户列表
func (*SysUserService) GetUserListService(query *system.SysUserRequest) (list interface{}, total int64, mg string, err error) {
	limit := query.PageSize
	offset := query.PageSize * (query.Page - 1)
	var userList []system.User
	db := global.DB.Model(system.User{})
	if query.Username != "" {
		db = db.Where("user_name LIKE ?", "%"+query.Username+"%")
	}
	if query.Nickname != "" {
		db = db.Where("nick_name LIKE ?", "%"+query.Nickname+"%")
	}
	//查询总数
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, "查询用户总数失败", err
	}
	err = db.Omit("password").Limit(limit).Offset(offset).Find(&userList).Error
	if err != nil {
		return nil, 0, "查询用户列表失败", err
	}
	return userList, total, "查询用户列表成功", nil
}
