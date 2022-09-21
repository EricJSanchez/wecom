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
