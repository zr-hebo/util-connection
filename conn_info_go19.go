// +build !go1.10

package connassist

import (
	"database/sql"
	"reflect"
)


func getMySQLConnDSN(conn *sql.DB) string {
	return reflect.ValueOf(conn).Elem().FieldByName("dsn").String()
}

