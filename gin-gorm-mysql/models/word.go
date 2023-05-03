package models

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	CN string `json:"cn" gorm:"type:varchar(50);not null"`
	EN string `json:"en" gorm:"type:varchar(50);not null"`
}

// 获取列表
func GetWords(pageNum int, pageSize int, maps interface{}) (words []Word, err error) {

	if pageSize > 0 && pageNum > 0 {
		err = DB.Where(maps).Find(&words).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = DB.Where(maps).Find(&words).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return words, nil
}

// 添加单词
func AddWord(word Word) (err error) {
	return DB.Create(&word).Error
}
