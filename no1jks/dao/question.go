package dao

import (
	"github.com/jinzhu/gorm"
	"no1jks/no1jks/models"
	"time"
)
// 暂时不晓得如何join后全行扫描，先用这个东西
type qa struct {
	QuestionID         int
	QuestionTitle      string
	QuestionContent    string
	QuestionViewCount  int
	QuestionLikeCount  int
	QuestionCommCount  int
	QuestionUpdateTime time.Time
	QuestionUserID     int
	QuestionUserName   string
	QuestionUserAvatar string
	IsBlog             int
	Cover              string

	AnswerId         int
	AnswerContent    string
	AnswerViewCount  int
	AnswerLikeCount  int
	AnswerCommCount  int
	AnswerUpdateTime time.Time
	AnswerUserID     int
	AnswerUserName   string
	AnswerUserAvatar string

	AnswerScore      int
	QuestionIsLocked int
}
// 返回给首页的博客-自问自答的问题
type HomepageBlog struct {
	BlogID         int
	BlogTitle      string
	BlogContent    string
	BlogViewCount  int
	BlogLikeCount  int
	BlogCommCount  int
	BlogUpdateTime time.Time
	BlogCover      string
	BlogUserID     int
	BlogUserAvatar string
	BlogUserName   string
}
// 手动聚合问题答案
type QuestionSet struct {
	Question qa
	Answers []qa
}
// 返回给客户端的问题与答案
type QuestionHomepageDataSet struct {
	DataSet
	Questions *[]QuestionSet
}

// 问题答案查询语句。暂时不知道如何select * 然后scan
const _sql = "question.id as question_id, " +
	"question.title as question_title," +
	"question.content as question_content," +
	"question.view_count as question_view_count," +
	"question.like_count as question_like_count," +
	"question.comment_count as question_comm_count," +
	"question.update_at as question_update_time," +
	"question.is_blog as question_is_blog," +
	"question.thumb_img as question_cover," +
	"question.is_locked as question_is_locked," +
	"question_user.id as question_user_id," +
	"question_user.name as question_user_name," +
	"question_user.avatar as question_user_avatar," +
	"answer.id as answer_id, " +
	"answer.content as answer_content," +
	"answer.view_count as answer_view_count," +
	"answer.like_count as answer_like_count," +
	"answer.comment_count as answer_comm_count," +
	"answer.update_at as answer_update_time," +
	"answer.score as answer_score," +
	"answer_user.id as answer_user_id," +
	"answer_user.name as answer_user_name," +
	"answer_user.avatar as answer_user_avatar"

// 组装问题及答案
func assembleQuestions(rawData *[]qa) *[]QuestionSet{
	var set []QuestionSet
	container := make(map[int]QuestionSet)
	for _, question := range *rawData {
		q, ok := container[question.QuestionID]
		if ok {
			q.Answers = append(q.Answers, question)
		} else {
			var temp QuestionSet
			var answers []qa
			temp.Question = question
			temp.Answers = answers
			temp.Answers = append(temp.Answers, question)
			container[question.QuestionID] = temp
		}
	}
	for _, v := range container{
		set = append(set, v)
	}
	return &set
}

// 获取首页的博客
func (d *Dao) GetHomepageBlog(limit uint8) *[]*HomepageBlog {
	var result []*HomepageBlog
	db := d.Mysql.Table("question").
		Select("question.id as blog_id, "+
			"question.title as blog_title,"+
			"question.content as blog_content,"+
			"question.view_count as blog_view_count,"+
			"question.like_count as blog_like_count,"+
			"question.comment_count as blog_comm_count,"+
			"question.update_at as blog_update_time,"+
			"question.thumb_img as blog_cover,"+
			"blog_user.id as blog_user_id,"+
			"blog_user.avatar as blog_user_avatar,"+
			"blog_user.name as blog_user_name").
		Joins("left join user as blog_user on question.user_id = blog_user.id").
		Where("question.display_homepage = ? AND question.is_top = ? AND "+
			"question.is_blog = ?", models.True, models.True, models.True).
		Scan(&result)
	if err := db.Error; err != nil {
		panic(err)
	}
	return &result
}

// 获取首页的问答
func (d *Dao) GetHomepageQuestions(limit uint8) *QuestionHomepageDataSet {
	var rawQuestion []qa
	db := d.Mysql.Table("question").
		Select(_sql).
		Joins("left join user as question_user on question.user_id = question_user.id").
		Joins("left join answer on question.id = answer.question_id").
		Joins("left join user as answer_user on answer.user_id = question_user.id").
		Where("question.display_homepage = ? AND question.is_top = ? AND question.is_blog=0",
			models.True, models.True).Scan(&rawQuestion)
	if err := db.Error; err != nil {
		panic(err)
	}

	var ret QuestionHomepageDataSet
	ret.Questions = assembleQuestions(&rawQuestion)
	ret.Page = 0
	ret.TotalCount = 0
	return &ret
}


// 上Scope
func questionBaseFilter(db *gorm.DB) *gorm.DB {
	return db.Where("question.is_deleted = ?", models.False)
}

func answerBaseFilter(db *gorm.DB) *gorm.DB {
	return db.Where("answer.is_deleted = ?", models.False)
}

func blogFilter(db *gorm.DB) *gorm.DB {
	return db.Where("question.is_blog = ?", models.True)
}

func questionFilter(db *gorm.DB) *gorm.DB {
	return db.Where("question.is_blog = ?", models.False)
}

func (d *Dao) GetNewsHomepageNewsList(page int, onlyCount bool, filters *map[string]interface{}) interface{} {
	var rawQuestion []qa
	var totalCount int

	db := d.Mysql.Table("question").
		Select(_sql).
		Joins("left join user as question_user on question.user_id = question_user.id").
		Joins("left join answer on question.id = answer.question_id").
		Joins("left join user as answer_user on answer.user_id = answer_user.id").
		Scopes(questionBaseFilter, answerBaseFilter)
	db.Count(&totalCount)
	if onlyCount {
		return totalCount
	}

	err := db.Order("question.is_top asc, question.create_at desc, answer.score desc").
		Offset(getOffset(page)).
		Limit(models.Limit).
		Scan(&rawQuestion).Error
	if err != nil {
		panic(err)
	}

	var ret QuestionHomepageDataSet
	ret.Questions = assembleQuestions(&rawQuestion)
	ret.Page = page
	ret.TotalCount = totalCount
	return &ret
}

// 方案三 优化方案2的TODO
