package controllers

type AdminNewsController struct {
	baseController
}


type AdminNewsDetailController struct {
	baseController
}

func (c *AdminNewsController) Get(){
	var resp adminJsonView
	var data = []map[string]interface{} {{
		"id": 1,
		"timestamp": 1577361111,
		"author": "原创",
		"importance": 1,
		"status": 1,
		"title": "考试时间发布",
	},
	}

	resp.Code = 0
	resp.Data = map[string]interface{}{
		"News": data,
		"Total": 100,
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AdminNewsDetailController) Get(){
	var resp adminJsonView
	var data = map[string]interface{}  {
		"status": "draft",
		"title": "文章题目",
		"content": "文章题目neirong", // 文章内容
		"content_short": "文章摘要", // 文章摘要
		// source_uri: '', // 文章外链
		//image_uri: '', // 文章图片
		//display_time: undefined, // 前台展示时间
		"id": 1,
		//platforms: ['a-platform'],
		//comment_disabled: false,
		"display_homepage": false,
		"importance": 0,
		"source": "原创",
		"order": 9999,
	}

	resp.Code = 0
	resp.Data = data
	c.Data["json"] = resp
	c.ServeJSON()
}
