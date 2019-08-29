package structs

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
 * DBとのやり取りを担うUser構造体
 * フロントへ返却する
 */
type User struct {
	gorm.Model
	UID              string            `json:"uid" sql:"index"`
	NickName         string            `json:"nickName"`
	Sex              int               `json:"sex"`
	BirthDay         time.Time         `json:"birthDay"`
	Age              int               `json:"age"`
	ImageURL         string            `json:"imageUrl"`
	UserInformation  UserInformation   `gorm:"foreignkey:uid;association_foreignkey:uid"`
	UserCombinations []UserCombination `gorm:"foreignkey:uid;association_foreignkey:uid"`
}

/*
 * DBとのやり取りを担うUserInformation構造体
 */
type UserInformation struct {
	gorm.Model
	UID              string `json:"uid" sql:"index"`
	OpponentAgeLow   int    `json:"opponentAgeLow"`
	OpponentAgeUpper int    `json:"opponentAgeUpper"`
	Hobby            string `json:"hobby"`
	Residence        int    `json:"residence"`
	Job              int    `json:"job"`
	Personality      int    `json:"personality"`
}

/**
 * 任意のUIDの組み合わせを表す構造体
 */
type UserCombination struct {
	gorm.Model
	UID         string `json:"uid" sql:"index"`
	OpponentUID string `json:"otherID"`
}

/**
 * 相性が良いUserInformationを記録していくUserCompatible構造体
 * TODO: 複合主キーワンチャン
 */
type InfoCompatible struct {
	gorm.Model
	InfoID  string `json:"infoID"`
	OtherID string `json:"otherID"`
}
