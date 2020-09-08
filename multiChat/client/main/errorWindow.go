package main

import (
	"github.com/lxn/walk"
)

func errorTip(err error) {
	walk.MsgBox(nil, "错误", err.Error(), walk.MsgBoxIconInformation)
}
