package service

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
)

func CreateApp(app model.TApp) (err error) {
	db := common.GetDB()
	err = db.Create(app).Error
	return err
}

func DeleteApp(id int) (err error) {
	db := common.GetDB()
	err = db.Where("id =?", id).Delete(model.TApp{}).Error
	return err
}

func UpdateApp(app *model.TApp) (err error) {
	db := common.GetDB()
	err = db.Model(&app).Updates(model.TApp{Name: app.Name, Status: app.Status, Type: app.Type, Desc: app.Desc}).Error
	return err
}

func GetApp(id uint) (err error, app model.TApp) {
	db := common.GetDB()
	err = db.Where("id = ?", id).First(&app).Error
	return err, app
}

func GetAppInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := common.GetDB()
	var apps []model.TApp
	err = db.Find(&apps).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&apps).Error
	return err, apps, total
}
