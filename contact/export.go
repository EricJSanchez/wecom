package contact

const (
	// 导出成员
	exportSimpleUserAddr = "https://qyapi.weixin.qq.com/cgi-bin/export/simple_user?access_token=%s"
	// 导出成员详情
	exportUserAddr = "https://qyapi.weixin.qq.com/cgi-bin/export/user?access_token=%s"
	// 导出部门
	exportDepartmentAddr = "https://qyapi.weixin.qq.com/cgi-bin/export/department?access_token=%s"
	// 导出标签成员
	exportTaguserAddr = "https://qyapi.weixin.qq.com/cgi-bin/export/taguser?access_token=%s"
	// 获取导出结果
	exportGetResultAddr = "https://qyapi.weixin.qq.com/cgi-bin/export/get_result?access_token=%s&jobid=%s"
)
