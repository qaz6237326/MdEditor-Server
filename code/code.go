package code

import "server/global"

const (
	SUCCESS                        = 200
	ERROR                          = 500
	ErrorNotLogin                  = 50000
	ErrorAuthorizationExpires      = 50001
	ErrorImageNotFound             = 1001  //图片不存在
	UploadImageSuccess             = 2001  //上传图片成功
	ErrorCreateDir                 = 1002  //创建文件夹失败
	SaveFileSuccess                = 2002  //保存文件成功
	ErrorSaveFile                  = 1003  //保存文件失败
	ErrorSaveFileData              = 1004  //保存文件信息到数据库失败
	ErrorUploadQiNiu               = 1005  //七牛云上传失败
	ErrorOpenFile                  = 1006  //打开文件失败
	ErrorUploadQiNiuSuccess        = 3001  //七牛云上传成功
	ErrorMissingId                 = 4001  //缺少id
	ErrorFileNotFound              = 404   //文件不存在
	ErrorDeleteFile                = 5001  //删除文件失败
	ErrorDeleteFileData            = 5002  //删除文件数据库失败
	ErrorListFile                  = 6001  //获取文件列表失败
	ErrorCreateTheme               = 7001  //创建主题失败
	ErrorGetTheme                  = 8001  //获取主题失败
	ErrorDeleteTheme               = 9001  //删除主题失败
	ErrorThemeNotFound             = 10001 //主题不存在
	ErrorUpdateTheme               = 10002 //主题更新失败
	ErrorRegisterMissingParam      = 10003 //缺少参数
	ErrorUserExist                 = 10011 //用户已存在
	ErrorUserExistAccount          = 10012 //账户已存在
	ErrorUserExistNickname         = 10013 //昵称已存在
	ErrorUserExistEmail            = 10014 //邮箱已存在
	ErrorUserExistPhone            = 10015 //手机号已存在
	ErrorCreateUser                = 10016 //用户创建失败
	ErrorGetQueryParam             = 10017 //获取查询参数失败
	ErrorNotFound                  = 10018 //未找到记录
	ErrorEditorTheme               = 10019 //编辑主题失败
	ErrorFindArticle               = 10020 //查询文章失败
	ErrorFindUser                  = 10021 //	查询用户信息失败
	ErrorUpYunPut                  = 10022 //又拍云上传失败
	ErrorSetUserConfigMissingParam = 10023 //缺少参数
	ErrorSetUserConfig             = 10024 //设置用户配置失败
	ErrorFindUserConfig            = 10025 //查询用户配置失败
)

func Text(code int) string {
	lang := global.Config.System.Language
	if lang == "en-us" {
		return enUSText[code]
	}
	if lang == "zh-cn" {
		return zhCNText[code]
	}
	return zhCNText[code]
}
