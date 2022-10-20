package test

import (
	"testing"
)

func TestDeptCache(t *testing.T) {
	weCom, err := Wework("d80").GetPlatform()
	if err != nil {
		t.Error(err)
		return
	}
	deptList, err := weCom.GetDeptCache()
	if err != nil {
		t.Error(err)
		return
	}
	Pr(deptList)
}

func TestDeptStaff(t *testing.T) {
	weCom, err := Wework("d80").GetPlatform()
	if err != nil {
		t.Error(err)
		return
	}
	// 注意 partyId为后台给的部门 ID
	deptStaffList, err := weCom.GetDeptStaff("16*************1", 0, 2)
	if err != nil {
		t.Error(err)
		return
	}
	Pr(deptStaffList)
}

func TestSearchStaff(t *testing.T) {
	weCom, err := Wework("d80").GetPlatform()
	if err != nil {
		t.Error(err)
		return
	}
	// 注意 partyId为后台给的部门 ID
	deptStaffList, err := weCom.SearchStaff("Eric")
	if err != nil {
		t.Error(err)
		return
	}
	Pr(deptStaffList)
}

// 获取corp列表
func TestGetCorpList(t *testing.T) {
	weCom, err := Wework("305").GetPlatform()
	if err != nil {
		t.Error(err)
		return
	}
	appList, err := weCom.GetCorpApplication()
	if err != nil {
		t.Error(err)
		return
	}
	Pr(appList)
}

// 设置 API 接收
func TestSaveOpenApiApp(t *testing.T) {
	weCom, err := Wework("305").GetPlatform()
	if err != nil {
		t.Error(err)
		return
	}
	appList, err := weCom.SettingApiCallback("56***", "https://****", "T2d93***", "aesKey******")
	if err != nil {
		t.Error(err)
		return
	}
	Pr(appList)
}

// 设置 IP 名名单，设置前请先 设置可信域名 或 设置接收消息服务器URL,否则会报服务器异常
func TestSaveIpWhiteList(t *testing.T) {
	weCom, err := Wework("305").GetPlatform()
	if err != nil {
		t.Error(err)
		return
	}
	appList, err := weCom.SaveIpWhiteList("781", []string{"58.*.*.*"})
	if err != nil {
		t.Error(err)
		return
	}
	Pr(appList)
}

// 获取应用管理员列表
func TestGetAppAdminInfo(t *testing.T) {
	weCom, err := Wework("305").GetPlatform()
	if err != nil {
		t.Error(err)
		return
	}
	adminList, err := weCom.GetAppAdminInfo("305")
	if err != nil {
		t.Error(err)
		return
	}
	Pr(adminList)
}
