package contact

const (
	// 增量更新成员
	batchSyncuserAddr = "https://qyapi.weixin.qq.com/cgi-bin/batch/syncuser?access_token=%s"
	// 全量覆盖成员
	batchReplaceuserAddr = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceuser?access_token=%s"
	// 全量覆盖部门
	batchReplacepartyAddr = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceparty?access_token=%s"
	// 获取异步任务结果
	batchGetresultAddr = "https://qyapi.weixin.qq.com/cgi-bin/batch/getresult?access_token=%s&jobid=%s"
)
