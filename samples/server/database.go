// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

func init() {
	sql.Register("test", &testDriver{})
}

type testDriver struct{}

func (*testDriver) Open(string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

type testConn struct{}

func (*testConn) Prepare(string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

func (*testConn) Close() error { _ = "STUB: not implemented"; return nil }

func (*testConn) Begin() (driver.Tx, error) { _ = "STUB: not implemented"; return *new(driver.Tx), nil }

type testConnector struct{}

func (*testConnector) Connect(context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (*testConnector) Driver() driver.Driver { _ = "STUB: not implemented"; return *new(driver.Driver) }

func openDatabase() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func openDatabase2() *sql.DB { _ = "STUB: not implemented"; return nil }
