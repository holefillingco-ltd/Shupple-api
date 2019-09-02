package structs

import "time"

/**
 * POSTされた値を受け取る為の構造体
 */
type PostUser struct {
	UID              string    `json:"uid"`
	NickName         string    `json:"nickName" validate:"required,gte=1,lt=10"`
	Sex              int       `json:"sex" validate:"required"`
	BirthDay         time.Time `json:"birthDay" validate:"required"`
	OpponentAgeLow   int       `json:"opponentAgeLow" validate:"required"`
	OpponentAgeUpper int       `json:"opponentAgeUpper" validate:"required"`
	Hobby            string    `json:"hobby" validate:"required,gte=0,lt=10"`
	Residence        int       `json:"residence" validate:"required"`
	Job              int       `json:"job" validate:"required"`
	Personality      int       `json:"personality" validate:"required"`
}

/**
 * エラーが発生した場合にフロントへ返却するError構造体
 */
type Error struct {
	Code              int      `json:"code"`
	Message           string   `json:"message"`
	ValidationMessage []string `json:"validationMessage"`
}