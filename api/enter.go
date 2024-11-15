package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/article_api"
	"gvb_server/api/big_model_api"
	"gvb_server/api/chat_api"
	"gvb_server/api/comment_api"
	"gvb_server/api/data_api"
	"gvb_server/api/feedback_api"
	"gvb_server/api/gaode_api"
	"gvb_server/api/images_api"
	"gvb_server/api/log_api"
	"gvb_server/api/log_v2_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/message_api"
	"gvb_server/api/new_api"
	"gvb_server/api/role_api"
	"gvb_server/api/settings_api"
	"gvb_server/api/tag_api"
	"gvb_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	TagApi      tag_api.TagApi
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
	CommentApi  comment_api.CommentApi
	NewsApi     new_api.NewApi
	ChatApi     chat_api.ChatApi
	LogApi      log_api.LogApi
	DataApi     data_api.DataApi
	LogV2Api    log_v2_api.LogApi
	RoleApi     role_api.RoleApi
	GaodeApi    gaode_api.GaodeApi
	FeedbackApi feedback_api.FeedbackApi
	BigModelApi big_model_api.BigModelApi
}

var ApiGroupApp = new(ApiGroup)
