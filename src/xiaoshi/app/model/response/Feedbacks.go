package response

import "xiaoshi/app/model"

type RespModel struct {
	Success string `json:"success"`
	Message string `json:"message"`
}

/**
返回对象
 */
type RespFeedback struct {
	RespModel
	Data respData `json:"data"`
}

type respData struct {
	Feedback model.Feedbacks `json:"feedback"`
}

/**
返回数组
 */
type RespFeedbacks struct {
	RespModel
	Data respDatas `json:"data"`
}

type respDatas struct {
	Feedbacks []model.Feedbacks `json:"feedback"`
}
