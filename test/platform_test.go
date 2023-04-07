package test

import (
	"github.com/EricJSanchez/wecom/externalcontact"
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

func TestAddMomentTask(t *testing.T) {
	weCom, err := Wework("ww******5305").GetExternalContact()
	if err != nil {
		t.Error(err)
		return
	}
	adminList, err := weCom.AddMomentTask(externalcontact.AddMomentTask{
		Text: externalcontact.Text{
			Content: "abc1",
		},
		Attachments: nil,
		VisibleRange: externalcontact.AddMomentTaskVisibleRange{
			SenderList: externalcontact.AddMomentTaskSenderList{
				UserList:       []string{"DengHui"},
				DepartmentList: nil,
			},
			ExternalContactList: externalcontact.AddMomentTaskExternalContactList{
				TagList: nil,
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	Pr(adminList)
}

func TestGetMomentTaskResult(t *testing.T) {
	weCom, err := Wework("ww******5305").GetExternalContact()
	if err != nil {
		t.Error(err)
		return
	}
	adminList, err := weCom.GetMomentTaskResult("aof_Y*****j-UjjFnA")
	if err != nil {
		t.Error(err)
		return
	}
	Pr(adminList)
}

func TestGetMomentComments(t *testing.T) {
	weCom, err := Wework("ww******5305").GetExternalContact()
	if err != nil {
		t.Error(err)
		return
	}
	adminList, err := weCom.GetMomentComments(externalcontact.MomentCommentsOption{
		MomentId: "mom0**********5ehNXHQ",
		Userid:   "De**i1",
	})
	if err != nil {
		t.Error(err)
		return
	}
	Pr(adminList)
}

func TestGetRoleList(t *testing.T) {
	weCom, err := Wework("ww******5305").GetPlatform()
	if err != nil {
		t.Error(err)
		return
	}
	adminList, err := weCom.GetRoleList()
	if err != nil {
		t.Error(err)
		return
	}
	Pr(adminList.Data.RoleList.Item)
	for _, item := range adminList.Data.RoleList.Item {
		if item.RoleName == "超级管理组" {

		} else {
			for _, item1 := range item.AdminList.Item {

				ss, _ := weCom.SearchStaff(item1.Name)
				for _, staff := range ss.Data {
					if staff.Vid == item1.ID {

					}
				}
			}
		}
	}
}
