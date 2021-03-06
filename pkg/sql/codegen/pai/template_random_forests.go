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

package pai

type randomForestsTrainFiller struct {
	DataSource     string
	TmpTrainTable  string
	FeatureColumns []string
	LabelColumn    string
	Save           string
	TreeNum        int
}

type randomForestsPredictFiller struct {
	DataSource      string
	TmpPredictTable string
	FeatureColumns  []string
	Save            string
	ResultTable     string
}

type randomForestsExplainFiller struct {
	DataSource      string
	TmpExplainTable string
	FeatureColumns  []string
	LabelColumn     string
	Save            string
	ResultTable     string
}

const randomForestsTrainTemplate = `
import os
import subprocess
import sqlflow_submitter.db

driver, dsn = "{{.DataSource}}".split("://")
assert driver == "maxcompute"
user, passwd, address, database = sqlflow_submitter.db.parseMaxComputeDSN(dsn)

column_names = []
{{ range $colname := .FeatureColumns }}
column_names.append("{{$colname}}")
{{ end }}
pai_cmd = 'pai -name randomforests -project algo_public -DinputTableName="{{.TmpTrainTable}}" -DmodelName="{{.Save}}" -DlabelColName="{{.LabelColumn}}" -DfeatureColNames="%s" -DtreeNum="{{.TreeNum}}" ' % (
    ",".join(column_names)
)

# Submit the tarball to PAI
subprocess.run(["odpscmd", "-u", user,
                           "-p", passwd,
                           "--project", database,
                           "--endpoint", address,
                           "-e", pai_cmd],
               check=True)
`

const randomForestsPredictTemplate = `
import os
import subprocess
import sqlflow_submitter.db

driver, dsn = "{{.DataSource}}".split("://")
assert driver == "maxcompute"
user, passwd, address, database = sqlflow_submitter.db.parseMaxComputeDSN(dsn)

column_names = []
{{ range $colname := .FeatureColumns }}
column_names.append("{{$colname}}")
{{ end }}
pai_cmd = 'pai -name prediction -project algo_public -DmodelName="{{.Save}}" -DinputTableName="{{.TmpPredictTable}}"  -DoutputTableName="{{.ResultTable}}" -DfeatureColNames="%s" ' % (
    ",".join(column_names)
)

# Submit the tarball to PAI
subprocess.run(["odpscmd", "-u", user,
                           "-p", passwd,
                           "--project", database,
                           "--endpoint", address,
                           "-e", pai_cmd],
               check=True)
`

const randomForestsExplainTemplate = `
import os
import subprocess
import sqlflow_submitter.db

driver, dsn = "{{.DataSource}}".split("://")
assert driver == "maxcompute"
user, passwd, address, database = sqlflow_submitter.db.parseMaxComputeDSN(dsn)

column_names = []
{{ range $colname := .FeatureColumns }}
column_names.append("{{$colname}}")
{{ end }}
pai_cmd = 'pai -name feature_importance -project algo_public -DmodelName="{{.Save}}" -DinputTableName="{{.TmpExplainTable}}"  -DoutputTableName="{{.ResultTable}}" -DlabelColName="{{.LabelColumn}}" -DfeatureColNames="%s" ' % (
    ",".join(column_names)
)

# Submit the tarball to PAI
subprocess.run(["odpscmd", "-u", user,
                           "-p", passwd,
                           "--project", database,
                           "--endpoint", address,
                           "-e", pai_cmd],
               check=True)
`
