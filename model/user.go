package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID int32 `gorm:"primaryKey"json:"id"`
	CreatedAt time.Time `gorm:"column:add_time"json:"add_time"`
	UpdateAt time.Time `gorm:"column:update_time" json:"update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool `json:"is_deleted"`
}


type User struct {
	BaseModel
	Phone string `gorm:"index:idx_phone;unique;type:varchar(11);not null"json:"phone"`
	Password string `gorm:"type:varchar(100);not null"json:"password"`
	NickName string `gorm:"type:varchar(20);" json:"nick_name"`
	Birthday *time.Time `gorm:"gorm:type:datetime;" json:"birthday"`
	Gender string `gorm:"column:gender;default:男;type:varchar(6)" json:"gender"`
	Role int`gorm:"column:role;default:1;type:int comment '1表示普通用户,2表示管理员用户'"json:"role"`
}