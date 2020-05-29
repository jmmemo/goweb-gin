package src

type Topic struct {
	TopicID    int    `json:"id"`
	TopicTitle string `json:"title"`
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}

func Create(id int, title string) Topic {
	return Topic{
		TopicID:    id,
		TopicTitle: title,
	}
}
