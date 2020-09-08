package process

import "errors"

var (
	usrM *usrMgr
)

type usrMgr struct {
	onlineUsrs map[int]*UsrProcess
}

func init() {
	usrM = &usrMgr{
		onlineUsrs: make(map[int]*UsrProcess, 1024),
	}
}

func (usrM *usrMgr) addOnlineUsr(up *UsrProcess) {
	usrM.onlineUsrs[up.UsrID] = up
}

func (usrM *usrMgr) deleteOnlineUsr(usrID int) {
	delete(usrM.onlineUsrs, usrID)
}

func (usrM *usrMgr) getOnlineUsrs() map[int]*UsrProcess {
	return usrM.onlineUsrs
}

func (usrM *usrMgr) getConnByUsrID(usrID int) (up *UsrProcess, err error) {
	up, ok := usrM.onlineUsrs[usrID]
	if !ok {
		err = errors.New("用户不在线！")
	}
	return
}
