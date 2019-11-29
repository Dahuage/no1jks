package dao

import (
	"no1jks/no1jks/models"
	"time"
)

// 看起来臭臭的！！！the map[interface{}]interface{} sucks
// 我们将在下个接口中尝试用struct 表述复杂的json

type qa struct {
	QuestionID int
	QuestionTitle string
	QuestionContent string
	QuestionViewCount int
	QuestionLikeCount int
	QuestionCommCount int
	QuestionUpdateTime time.Time
	QuestionUserID int
	QuestionUserName string
	QuestionUserAvatar string
	IsBlog int
	Cover string

	AnswerId int
	AnswerContent string
	AnswerViewCount int
	AnswerLikeCount int
	AnswerCommCount int
	AnswerUpdateTime time.Time
	AnswerUserID int
	AnswerUserName string
	AnswerUserAvatar string
}

func makeAnswer(q *qa) *map[string]interface{} {
	answer := make(map[string]interface{})
	answer["AnswerId"] = (*q).AnswerId
	answer["AnswerContent"] = (*q).AnswerContent
	answer["AnswerViewCount"] = (*q).AnswerViewCount
	answer["AnswerCommCount"] = (*q).AnswerCommCount
	answer["AnswerLikeCount"] = (*q).AnswerLikeCount
	answer["AnswerUpdateTime"] = (*q).AnswerUpdateTime
	answer["AnswerUserID"] = (*q).AnswerUserID
	answer["AnswerUserName"] = (*q).AnswerUserName
	answer["AnswerUserAvatar"] = (*q).AnswerUserAvatar
	return &answer
}

func AssembleQA(rows *[]*qa) *[]*map[string]interface{} {
	var Questions []*map[string]interface{}
	container := make(map[int]map[string]interface{})
	for _, value := range *rows {
		qid := (*value).QuestionID
		q := container[qid]
		answer := makeAnswer(value)
		if  q != nil {
			answers := q["Answers"].([]*map[string]interface{})
			answers = append(answers, answer)
		}else {
			q := make(map[string]interface{})
			q["QuestionID"] = (*value).QuestionID
			q["QuestionTitle"] = (*value).QuestionTitle
			q["QuestionContent"] = (*value).QuestionContent
			q["QuestionViewCount"] = (*value).QuestionViewCount
			q["QuestionLikeCount"] = (*value).QuestionLikeCount
			q["QuestionCommCount"] = (*value).QuestionCommCount
			q["QuestionUpdateTime"] = (*value).QuestionUpdateTime
			q["QuestionUserID"] = (*value).QuestionUserID
			q["QuestionUserName"] = (*value).QuestionUserName
			q["QuestionUserAvatar"] = (*value).QuestionUserAvatar
			q["IsBlog"] = (*value).IsBlog
			q["Cover"] = (*value).Cover
			var Answers []*map[string]interface{}
			Answers = append(Answers, answer)
			q["Answers"] = Answers
			container[(*value).QuestionID] = q
		}
	}
	for _, v:= range container{
		Questions = append(Questions, &v)
	}
	return &Questions
}

func (d *Dao) GetHomepageQuestions(limit uint8) *[]*map[string]interface{} {
	var result []*qa
	db := d.mysql.Table("question").
		Select("question.id as question_id, " +
		"question.title as question_title," +
		"question.content as question_content," +
		"question.view_count as question_view_count," +
		"question.like_count as question_like_count," +
		"question.comment_count as question_comm_count," +
		"question.update_at as question_update_time," +
		"question.is_blog as question_is_blog," +
		"question.thumb_img as question_cover," +
		"question_user.id as question_user_id," +
		"question_user.name as question_user_name," +
		"question_user.avatar as question_user_avatar," +
		"answer.id as answer_id, " +
		"answer.content as answer_content," +
		"answer.view_count as answer_view_count," +
		"answer.like_count as answer_like_count," +
		"answer.comment_count as answer_comm_count," +
		"answer.update_at as answer_update_time," +
		"answer_user.id as answer_user_id," +
		"answer_user.name as answer_user_name," +
		"answer_user.avatar as answer_user_avatar").
		Joins("left join user as question_user on question.user_id = question_user.id").
		Joins("left join answer on question.id = answer.question_id").
		Joins("left join user as answer_user on answer.user_id = question_user.id").
		Where("question.display_homepage = ? AND question.is_top = ?", models.True, models.True).
		Scan(&result)
	if err := db.Error; err != nil {
		panic(err)
	}
	return AssembleQA(&result)
}