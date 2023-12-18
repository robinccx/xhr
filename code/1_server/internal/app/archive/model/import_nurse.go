package model

import (
	"github.com/google/uuid"
	"github.com/tiger1103/gfast/v3/internal/utils"
	"strconv"
	"strings"
)

type ImportNurse struct {
	ImportBase
}

func (t *ImportNurse) Init(insertFlag bool) {

	t.AppendColumn("科室", "dept_code", "string", true, 2)
	t.AppendColumn("姓名", "nurse_name", "string", true, 2)
	t.AppendColumn("性别", "sex", "string", insertFlag, 0)
	t.AppendColumn("出生年月", "birthday", "date", insertFlag, 0)
	t.AppendColumn("来院工作时间", "start_date", "date", insertFlag, 0)
	t.AppendColumn("学历", "education_code", "string", insertFlag, 0)
	t.AppendColumn("职称类别", "title_code", "string", insertFlag, 0)
	t.AppendColumn("离职日期", "end_date", "date", false, 0)
	t.AppendColumn("执业证书号", "cert_no", "string", false, 2)
	t.AppendColumn("护士执业注册有效期", "cert_end_date", "date", false, 0)
	t.AppendColumn("执业注册时间", "reg_date", "date", false, 0)
	t.AppendColumn("身份证号", "id_no", "string", false, 2)
	t.AppendColumn("人员类别", "staff_type_code", "string", false, 2)

	return
}

func (t *ImportNurse) LoadData(data [][]string) (err ImportErrorInfo) {
	var (
		count      int
		value      string
		isError    bool
		firstIndex int
		isNone     bool
	)
	firstIndex = -1
	for i, record := range data {
		if len(record) < t.ColumnCount {
			continue
		}
		t.SetColumn(record)
		firstIndex = i
		break
	}
	if firstIndex < 0 {
		panic("导入文件格式错误")
	}

	err = ImportErrorInfo{}

	t.SessionId = uuid.New().String()
	for i, record := range data {
		if i <= firstIndex {
			continue
		}

		count = len(record)
		if count < t.ColumnCount-3 {
			continue
		}
		isNone = true
		for _, v := range record {
			if v != "" {
				isNone = false
				break
			}
		}
		if isNone {
			continue
		}
		isError = false
		newRecord := map[string]interface{}{}
		newRecord["status"] = 1
		for _, column := range t.Columns {
			if column.ColumnIndex < 0 {
				continue
			}
			value = ""
			if column.ColumnIndex < count {
				value = record[column.ColumnIndex]
			}

			if column.Required && value == "" {
				isError = true
				err.Push(i, column.ColumnName+"没有数据")
				break
			}
			if value == "" {
				continue
			}
			switch column.TrimType {
			case 1:
				value = strings.TrimSpace(value)
			case 2:
				value = strings.TrimSpace(value)
				value = strings.Replace(value, " ", "", -1)
				value = strings.Replace(value, "　", "", -1)
			}

			if column.FieldType == "date" {
				value = strings.Replace(value, "/", "-", -1)
				value = strings.Replace(value, ".", "-", -1)
				tempList := strings.Split(value, "-")

				if len(tempList) < 2 || len(tempList[0]) == 5 {
					if len(tempList[0]) == 5 {
						dateInt, err1 := strconv.Atoi(tempList[0])
						if err1 != nil {
							err.Push(i, column.ColumnName+"数据错误["+value+"]")
							isError = true
							break
						}

						value = utils.ExcelDateToDateString(dateInt)
						if value < "1949-01-01" {
							err.Push(i, column.ColumnName+"数据错误["+value+"]")
							isError = true
							break
						}
						if value > "2200-01-01" {
							err.Push(i, column.ColumnName+"数据错误["+value+"]")
							isError = true
							break
						}
					} else {
						err.Push(i, column.ColumnName+"数据错误["+value+"]")
						isError = true
						break
					}
				} else if len(tempList) == 2 {
					if len(tempList[1]) == 2 {
						value = value + "-01"

					} else if len(tempList[1]) == 1 {
						value = tempList[0] + "-0" + tempList[1] + "-01"
					} else {
						err.Push(i, column.ColumnName+"数据错误["+value+"]")
						isError = true
						break
					}
				}

				if value != "" && value < "1949-01-01" {
					err.Push(i, column.ColumnName+"数据错误["+value+"]")
					isError = true
					break
				}

				if value > "2200-01-01" {
					err.Push(i, column.ColumnName+"数据错误["+value+"]")
					isError = true
					break
				}

				if column.FieldName == "end_date" && value != "" {
					newRecord["status"] = 2
				}
			}

			newRecord[column.FieldName] = value
			newRecord["seq"] = i - firstIndex
		}
		if !isError {
			newRecord["session_id"] = t.SessionId
			t.Data = append(t.Data, newRecord)
		}
	}

	return
}
