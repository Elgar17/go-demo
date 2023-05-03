package servise

import "gin-gorm-mysql/models" //公共包含程序/组件/方法/变量等。这些程序

type WordServise struct {
}

// 添加word
func (ws *WordServise) AddWord(word *models.Word) (err error) {
	err = models.DB.Create(&word).Error
	return err
}

// 获取 word 列表
func (ws *WordServise) GetWords(size int, offset int) (words []*models.Word, err error) {
	err = models.DB.Limit(size).Offset(offset).Find(&words).Error
	if err != nil {
		return nil, err
	}
	return words, nil
}

// 删除
func (ws *WordServise) DeleteWordByID(id int) (err error) {
	return models.DB.Delete(&models.Word{}, id).Error
}

// 更新 word 内容或保存到数据库中的值或保存到新的 Word 对象中的值。
func (ws *WordServise) UpdateWord(word *models.Word) (err error) {
	return models.DB.Where("id = ?", word.ID).Updates(&word).Error
}
