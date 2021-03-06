// Copyright 2019 The SQLFlow Authors. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseGetTestingDBSingleton(t *testing.T) {
	db := GetTestingDBSingleton()
	a := assert.New(t)

	switch dbms := getEnv("SQLFLOW_TEST_DB", "mysql"); dbms {
	case "mysql":
		a.Equal(testingMySQLURL(), db.URL())
	case "hive":
		a.Equal(testingHiveURL(), db.URL())
	case "maxcompute":
		a.Equal(testingMaxComputeURL(), db.URL())
	default:
		a.Fail("Unrecognized environment variable SQLFLOW_TEST_DB %s", dbms)
	}
}

func TestDatabaseTestingMySQLURL(t *testing.T) {
	a := assert.New(t)
	a.Equal("mysql://root:root@tcp(127.0.0.1:3306)/?maxAllowedPacket=0", testingMySQLURL())
	if db := GetTestingDBSingleton(); db.DriverName == "mysql" {
		a.Equal(testingMySQLURL(), db.URL())
	}
}

func TestDatabaseTestingHiveURL(t *testing.T) {
	a := assert.New(t)
	a.Equal("hive://root:root@localhost:10000/churn", testingHiveURL())
	if db := GetTestingDBSingleton(); db.DriverName == "hive" {
		a.Equal(testingHiveURL(), db.URL())
	}
}

func TestDatabaseTestingMaxComputeURL(t *testing.T) {
	a := assert.New(t)
	a.Equal("maxcompute://test:test@service-maxcompute.com/api?curr_project=test&scheme=http", testingMaxComputeURL())
	if db := GetTestingDBSingleton(); db.DriverName == "maxcompute" {
		a.Equal(testingMaxComputeURL(), db.URL())
	}
}
