package task

import (
	"fmt"
	"github.com/uma-co82/Shupple-api/src/api/domain/user"
	"github.com/uma-co82/Shupple-api/src/api/infrastructure/db"
	"os"
	"time"
)

func UserCombinationCheckCreatedAtTask() {
	fmt.Println("開始＊＊＊＊＊＊＊＊＊＊＊＊＊＊＊")
	db := db.Init()
	tx := db.Begin()
	defer db.Close()

	var (
		userCombinations []user.UserCombination
		targetUserIds    []string
		updateTarget     = map[string]interface{}{"is_combination": false, "opponent_uid": nil}
	)
	now := time.Now()
	endTargetHours := 48

	file, err := os.OpenFile(now.String()+"-UserCombinationTask.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("************************%v scheduler file open&create&update Failed ********************************", now)
		return
	}

	defer file.Close()

	if err := tx.Where("created_at + INTERVAL ? HOUR < ?", endTargetHours, now).Find(&userCombinations).Error; err != nil {
		tx.Rollback()
		return
	}

	for _, userCombination := range userCombinations {
		targetUserIds = append(targetUserIds, userCombination.UID)
		targetUserIds = append(targetUserIds, userCombination.OpponentUID)
	}

	if err := db.Where("uid IN (?)", targetUserIds).Updates(updateTarget).Error; err != nil {
		tx.Rollback()
		return
	}

	fmt.Fprintln(file, now.String()+"から"+string(endTargetHours)+"のUserCombinationを確認しUserのマッチングを調整しました。")
}
