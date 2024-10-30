package models

import (
	"fmt"
	"github.com/catseeker/EasyGoLib/db"
	"github.com/catseeker/EasyGoLib/utils"
)

func Init() (err error) {
	err = db.Init()
	if err != nil {
		return
	}
	db.SQLite.AutoMigrate(User{}, Stream{})
	count := 0
	sec := utils.Conf().Section("http")
	defUser := sec.Key("default_username").MustString("admin")
	defPass := sec.Key("default_password").MustString("admin")
	db.SQLite.Model(User{}).Where("username = ?", defUser).Count(&count)
	if count == 0 {
		db.SQLite.Create(&User{
			Username: defUser,
			Password: defPass,
			//Password: utils.MD5(defPass),
		})
	}
	fmt.Printf("defUser=$s defPass=%s\n", defUser, defPass)
	return
}

func Close() {
	db.Close()
}
