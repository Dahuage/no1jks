package routers

import (
	"no1jks/no1jks/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/demo", &controllers.MainController{})

	beego.Router("/", &controllers.HomepageController{})
	beego.Router("/news", &controllers.NewsHomeController{})
	beego.Router("/question", &controllers.QuestionHomeController{})
	beego.Router("/examination", &controllers.ExaminationHomeController{})
	beego.Router("/material", &controllers.GoodHomeController{})
	beego.Router("/train", &controllers.TrainHomeController{})

	beego.Router("/user/login", &controllers.UserLoginController{})
	beego.Router("/user/signup", &controllers.UserSignupController{})
	beego.Router("/user/terms", &controllers.UserTermController{})
}
