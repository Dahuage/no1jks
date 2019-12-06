package routers

import (
	"no1jks/no1jks/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/demo", &controllers.MainController{})

	beego.Router("/", &controllers.HomepageController{})
	beego.Router("/news", &controllers.NewsHomeController{})
	beego.Router("/news/:id", &controllers.NewsDetailController{})
	beego.Router("/news/like/:id", &controllers.NewsHomeController{})
	beego.Router("/news/comm/:id", &controllers.NewsHomeController{})

	beego.Router("/question", &controllers.QuestionHomeController{})
	beego.Router("/question/:id", &controllers.QuestionDetailController{})
	beego.Router("/question/answer/:id", &controllers.QuestionDetailController{})
	beego.Router("/answer/like/:id", &controllers.QuestionHomeController{})
	beego.Router("/answer/comm/:id", &controllers.QuestionHomeController{})

	beego.Router("/examination", &controllers.ExaminationHomeController{})
	beego.Router("/examination/download/:id", &controllers.ExaminationHomeController{})

	beego.Router("/material", &controllers.GoodHomeController{})
	beego.Router("/material/:id", &controllers.GoodHomeController{})
	beego.Router("/material/download/:id", &controllers.GoodHomeController{})

	beego.Router("/train", &controllers.TrainHomeController{})

	beego.Router("/user/terms", &controllers.UserTermController{})
	beego.Router("/user/signup", &controllers.UserSignupController{})
	beego.Router("/user/login", &controllers.UserLoginController{})
	beego.Router("/user/logout", &controllers.UserLoginController{})
	beego.Router("/user/home/:id", &controllers.UserLoginController{})
	beego.Router("/user/set/:id", &controllers.UserLoginController{})

}
