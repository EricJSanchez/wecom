package test

import (
	"testing"
)

func TestGetDimissionUserInfoListV2(t *testing.T) {
	weCom, err := Wework("305").GetPlatform()
	if err != nil {
		t.Error(err)
		return
	}
	deptList, err := weCom.GetDimissionUserInfoListV2()
	if err != nil {
		t.Error(err)
		return
	}
	Pr(deptList)
}
