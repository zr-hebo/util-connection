package connassist

import (
	"database/sql"

	"github.com/dearcode/crab/log"
)

// CloseConnection 关闭数据库连接
func CloseConnection(conn *sql.DB) (err error) {
	err = conn.Close()

	connInfo, queryErr := GetMySQLConnInfo(conn)
	if queryErr != nil {
		log.Errorf("获取连接描述信息失败 <-- %s", queryErr.Error())

	} else {
		log.Infof("关闭连接后，连接描述信息：%s", connInfo)
	}

	return
}
