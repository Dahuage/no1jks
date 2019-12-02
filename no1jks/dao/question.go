package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"no1jks/no1jks/models"
	"time"
)

// 方案一 摸石头过河
// 看起来臭臭的！！！the map[interface{}]interface{} sucks
// 我们将在下个接口中尝试用struct 表述复杂的json

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
		if q != nil {
			answers := q["Answers"].([]*map[string]interface{})
			answers = append(answers, answer)
		} else {
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
	for _, v := range container {
		Questions = append(Questions, &v)
	}
	return &Questions
}

// 我服了， 我怂了。。。。
// 尼玛 这么长。
// 第一是我的问题 第二是那个gorm那个文档呀，写的跟💩没两样。
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

func (d *Dao) GetHomepageQuestions(limit uint8) *[]*map[string]interface{} {
	var result []*qa
	db := d.mysql.Table("question").
		Select(_sql).
		Joins("left join user as question_user on question.user_id = question_user.id").
		Joins("left join answer on question.id = answer.question_id").
		Joins("left join user as answer_user on answer.user_id = question_user.id").
		Where("question.display_homepage = ? AND question.is_top = ? AND question.is_blog=0",
			models.True, models.True).Scan(&result)
	if err := db.Error; err != nil {
		panic(err)
	}
	return AssembleQA(&result)
}

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

func (d *Dao) GetHomepageBlog(limit uint8) *[]*HomepageBlog {
	var result []*HomepageBlog
	db := d.mysql.Table("question").
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

// 方案2 让代码看起来好一点
// 上面的尝试很沙雕，指针套娃🐸没必要。
// 试试struct会不会香一点
// TODO 仍然不完美，重复定义了struct，struct也不香

/* **
orm连表查询后的结果结构
TODO 看下是不是真的不能select *
直接来个
type qa struct {
	q models.Question
	qu models.User
	a models.Answer
	au models.User
} 进行结果扫描*/

// ain't work. fk
//type qa_ struct {
//	q  models.Question
//	qu models.User
//	a  models.Answer
//	au models.User
//}
//
//type _qa struct {
//	qa
//	AnswerScore     int
//	QuestionIsLocked int
//}

// 为了把_qa手动聚合，很shit，就这样吧
// TODO 再研究下不sum如何group
type QuestionSet struct {
	Question qa
	Answers []qa
}

// 最终返回给前端的数据
type QuestionHomepageDataSet struct {
	DataSet
	Questions *[]QuestionSet
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

func AssembleQuestions(rawData *[]qa) *[]QuestionSet{
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
		logs.Info("container===================", container)
	}
	for _, v := range container{
		set = append(set, v)
		logs.Info("sett============", set)
	}
	return &set
}

func (d *Dao) GetNewsHomepageNewsList(page int, onlyCount bool, filters *map[string]interface{}) interface{} {
	var rawQuestion []qa
	var totalCount int

	db := d.mysql.Table("question").
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
	ret.Questions = AssembleQuestions(&rawQuestion)
	ret.Page = page
	ret.TotalCount = totalCount
	return &ret
}

// 方案三 优化方案2的TODO
