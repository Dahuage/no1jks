package dao

import (
	"github.com/astaxie/beego/logs"
	"no1jks/no1jks/models"
	"time"
)

// 创建问题
func (d *Dao) AnswerCreate(uid, questionId int, conclusion, content string) bool {
	logs.Info("===============", questionId, conclusion, content)
	var answer models.Answer
	answer.UserID = uid
	answer.QuestionID = questionId
	answer.Conclusion = conclusion
	answer.Content = content
	answer.CreateAt = int(time.Now().Unix())
	answer.UpdateAt = int(time.Now().Unix())
	db := d.Mysql.Create(&answer)
	if err := db.Error;  err != nil {
		logs.Error("Create question err", err, uid, questionId, content, conclusion)
		return false
	}
	return true
}