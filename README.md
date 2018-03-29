# connection reflector

use reflect get a mysql connection info, to get not closed connection. Help resolve mysql connection leak problem.
mysql connection structs in go1.9 and go1.10 have a bit differences, we all support.

Example ignore error deal:
```
connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/", userName, passwd, ip, port)
conn, _ := sql.Open("mysql", connStr)
connInfo, _ := GetMySQLConnInfo(conn)
println(connInfo)
```
