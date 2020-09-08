package model

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"
)

//MyUsrDao 唯一持久层管理
var (
	MyUsrDao *usrDao
)

type usrDao struct {
	pool *redis.Pool
}

//InitUsrDao 初始化
func InitUsrDao(pool *redis.Pool) {
	MyUsrDao = &usrDao{
		pool: pool,
	}
}

func (usrDao *usrDao) GetRedisConn() (conn redis.Conn) {
	conn = usrDao.pool.Get()
	return
}

func (usrDao *usrDao) GetUserByID(conn redis.Conn, id int) (usr usr, err error) {
	res, err := redis.String(conn.Do("HGet", "usrs", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrUsrNotExist
			return
		}
	}
	err = json.Unmarshal([]byte(res), &usr)
	if err != nil {
		return
	}
	return
}

func (usrDao *usrDao) Login(usrID int, usrPwd string) (usr usr, err error) {
	conn := usrDao.pool.Get()
	defer conn.Close()

	usr, err = usrDao.GetUserByID(conn, usrID)
	if err != nil {
		return
	}
	if usr.UsrPwd != usrPwd {
		err = ErrUsrPwdErr
		return
	}
	return

}

func (usrDao *usrDao) Regist(usr usr) (err error) {
	conn := usrDao.pool.Get()
	defer conn.Close()
	_, err = usrDao.GetUserByID(conn, usr.UsrID)
	if err == nil {
		err = ErrUsrExisted
		return
	}

	data, err := json.Marshal(usr)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "usrs", usr.UsrID, string(data))
	if err != nil {
		return
	}
	return
}
