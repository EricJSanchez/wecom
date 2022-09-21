package test

import (
	"fmt"
	"github.com/EricJSanchez/wecom/contact"
	"testing"
)

func TestStaffUserId(t *testing.T) {
	weCom, err := Wework("*****").GetContact()
	if err != nil {
		t.Error(err)
		return
	}
	userList, err := weCom.UserListId(contact.UserListIdOptions{
		Cursor: "",
		Limit:  100,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(userList)
}
