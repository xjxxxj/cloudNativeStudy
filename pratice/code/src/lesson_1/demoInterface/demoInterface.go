package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Connection interface {
	doQuery(sql string) string
}

type JDBC interface {
	getConnection() Connection
}

type MYSQL struct {
	url, username, password string `json:",string"`
}

func (mysql MYSQL) getConnection() Connection {
	fmt.Println("get mysql connection!")
	return MysqlConnection{
		url:      mysql.url,
		username: mysql.username,
		password: mysql.password,
	}
}

type ORACLE struct {
	url, username, password string
}

func (oracle ORACLE) getConnection() Connection {
	fmt.Println("get oracle connection!")
	return OracleConnection{
		url:      oracle.url,
		username: oracle.username,
		password: oracle.password,
	}
}

type MysqlConnection struct {
	url, username, password string
}

func (connection MysqlConnection) doQuery(sql string) string {
	fmt.Printf("do query in mysql, sql:%s\n", sql)
	return "mysql"
}

type OracleConnection struct {
	url, username, password string
}

func (connection OracleConnection) doQuery(sql string) string {
	fmt.Printf("do query in oracle, sql:%s\n", sql)
	return "oracle"
}

func main() {
	sql := "show parameters"
	var sqlEngines []JDBC
	mysql := MYSQL{url: "jdbc:mysql://localhost:3306", username: "xjx", password: "12345"}
	oracle := ORACLE{url: "jdbc:oracle://localhost:1521", username: "xjx", password: "12345"}
	sqlEngines = append(sqlEngines, mysql, oracle)
	for _, engine := range sqlEngines {
		engine.getConnection().doQuery(sql)
	}

	// do reflect
	fmt.Println("type", reflect.TypeOf(sqlEngines[0]))
	fmt.Println("value", reflect.ValueOf(sqlEngines[1]))

	// do json transfer
	jsonString, err := json.Marshal(&mysql)

	if err != nil {
		fmt.Println("Fail to trans to json, err:", err)
	}

	fmt.Println(string(jsonString))
}
