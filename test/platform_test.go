package test

import (
	"fmt"
	"github.com/EricJSanchez/wecom/externalcontact"
	"github.com/EricJSanchez/wecom/platform"
	"github.com/spf13/cast"
	"testing"
	"time"
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

func TestBatchChangeUsername(t *testing.T) {
	i := 0
	for {
		Pr("新一轮开始了。。。。。。。")
		time.Sleep(1 * time.Second)
		weCom, err := Wework("ww******5305").GetPlatform()
		if err != nil {
			t.Error(err)
			return
		}
		memberList, err := weCom.SearchStaff("微信用户")
		if err != nil || len(memberList.Data) == 0 {
			return
		}
		for _, v := range memberList.Data {
			if v.Vid != "" && v.Acctid != "" && v.IsJoinQyh == true {
				i++
				fmt.Println(v.Name, v.Acctid, SaveMember(v.Vid), i)
				time.Sleep(1 * time.Second)
			}
		}
	}
}

// func TestSaveMember(t *testing.T) {
func SaveMember(vid string) (err error) {
	//a := true
	//b := false
	//a1 := cast.ToString(a)
	//b1 := cast.ToString(b)
	//fmt.Println(a1, b1)
	//return
	weCom, err := Wework("ww******5305").GetPlatform()
	if err != nil {
		return
	}
	memberInfoData, err := weCom.GetSingleMember(vid)
	if err != nil {
		fmt.Println("获取用户详情失败", err)
		return
	}
	memberInfo := memberInfoData.Data
	//Pr(memberInfo)
	fmt.Println(memberInfo)
	if memberInfo.Name != "微信用户" {
		return
	}
	if memberInfo.WxNickName == "" || memberInfo.WxNickName == "微信用户" {
		fmt.Println("微信昵称为空", memberInfo.WxNickName)
		return
	}

	var dispOrder []platform.SaveMemberDispOrder
	for _, item := range memberInfo.DispOrder {
		dispOrder = append(dispOrder, platform.SaveMemberDispOrder{
			DepartID:   item.DepartID,
			DispOrder:  cast.ToString(item.DispOrder),
			IsTop:      cast.ToString(item.IsTop),
			LeaderRank: cast.ToString(item.LeaderRank),
		})
	}

	var saveMemberData platform.SaveMemberSchema
	saveMemberData = platform.SaveMemberSchema{
		JoinStatus:                          memberInfo.JoinStatus,
		Account:                             memberInfo.Account,
		Acctid:                              memberInfo.Acctid,
		AcctidStat:                          cast.ToString(memberInfo.AcctidStat),
		ActiveBiz:                           cast.ToString(memberInfo.ActiveBiz),
		Alias:                               memberInfo.Alias,
		Avatar:                              memberInfo.Avatar,
		BIsQymailGray:                       "true",
		BindStat:                            cast.ToString(memberInfo.BindStat),
		BizMail:                             memberInfo.BizMail,
		CountryCode:                         memberInfo.CountryCode,
		DeleteStat:                          cast.ToString(memberInfo.DeleteStat),
		DisableBiz:                          cast.ToString(memberInfo.DisableBiz),
		DisableStat:                         cast.ToString(memberInfo.DisableStat),
		DispOrder:                           dispOrder,
		Domain:                              memberInfo.Domain,
		Domainid:                            cast.ToString(memberInfo.Domainid),
		Email:                               memberInfo.Email,
		EnglishName:                         memberInfo.EnglishName,
		ExtTel:                              memberInfo.ExtTel,
		ExternJobTitle:                      memberInfo.ExternJobTitle,
		ExternPosition:                      memberInfo.ExternPositionInfo.ExternPosition,
		ExternPositionInfoBSynInnerPosition: cast.ToString(memberInfo.ExternPositionInfo.BSynInnerPosition),
		ExternPositionInfoExternPosition:    memberInfo.ExternPositionInfo.ExternPosition,
		ExternalAttrs:                       memberInfo.ExternalAttrs,
		ExternalCorpInfo:                    memberInfo.ExternalCorpInfo,
		ExternalWxfinder:                    "1",
		Gender:                              cast.ToString(memberInfo.Gender),
		GenderStr:                           memberInfo.GenderStr,
		HideMobile:                          cast.ToString(memberInfo.HideMobile),
		IdentityStat:                        cast.ToString(memberInfo.IdentityStat),
		IgnoreAbnormalMobile:                "false",
		Imgid:                               memberInfo.Imgid,
		IsSearchListShow:                    "true",
		IsJoinQyh:                           cast.ToString(memberInfo.IsJoinQyh),
		IsQuit:                              cast.ToString(memberInfo.IsQuit),
		IsReadyJoinAgain:                    cast.ToString(memberInfo.IsReadyJoinAgain),
		IsWwBizmail:                         cast.ToString(memberInfo.IsWwBizmail),
		IsWwBizmailVip:                      cast.ToString(memberInfo.IsWwBizmailVip),
		JoinStat:                            cast.ToString(memberInfo.JoinStat),
		LoginStat:                           cast.ToString(memberInfo.LoginStat),
		MainpartyID:                         memberInfo.MainpartyID,
		ManageStat:                          cast.ToString(memberInfo.ManageStat),
		Mobile:                              memberInfo.Mobile,
		ModelType:                           "full",
		Name:                                memberInfo.Name,
		Nickname:                            memberInfo.Nickname,
		PartyList:                           memberInfo.PartyList,
		Partyid:                             memberInfo.MainpartyID,
		PersonalEmail:                       memberInfo.PersonalEmail,
		Pinyin:                              memberInfo.Pinyin,
		Position:                            memberInfo.Position,
		PstnExtensionNumber:                 memberInfo.PstnExtensionNumber,
		Realname:                            memberInfo.Realname,
		Uin:                                 memberInfo.Uin,
		UserQuitTime:                        cast.ToString(memberInfo.UserQuitTime),
		Username:                            memberInfo.WxNickName,
		Vid:                                 memberInfo.Vid,
		VidBindGid:                          cast.ToString(memberInfo.VidBindGid),
		Wechat:                              memberInfo.Wechat,
		WxIDHash:                            memberInfo.WxIDHash,
		WxNickName:                          memberInfo.WxNickName,
		XcxCorpAddress:                      memberInfo.XcxCorpAddress,
	}
	d, err := weCom.SaveMember(saveMemberData)
	if err != nil {
		fmt.Println("修改用户失败", err)
		return
	}
	fmt.Println(d)
	return
}
