package model

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 公共model
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {

	addr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.PassWord,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(databaseSetting.DBType, addr)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		// 是否打印log
		db.LogMode(true)
	}
	// 全局禁用表名复数 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	db.SingularTable(true)
	// 空闲最大连接数
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	// 数据库最大连接数
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
