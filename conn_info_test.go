package connassist

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_GetConnInfo(t *testing.T) {
	userName := "test"
	passwd := "test"
	ip := "localhost"
	port := 3306
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/", userName, passwd, ip, port)

	conn, err := sql.Open("mysql", connStr)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	conn.SetMaxOpenConns(0)
	if err := conn.Ping(); err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	connInfo, err := GetMySQLConnInfo(conn)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log(connInfo)

	conn.Query("show databases")
	conn.Query("show processlist")
	conn.Query("show databases")
	conn.Query("show databases")
	connInfo = GetMySQLConnInfoIgnoreErr(conn)
	t.Log(connInfo)

}
