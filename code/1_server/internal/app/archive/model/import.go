package model

import "errors"

type ImportErrorItem struct {
	Index    int    `json:"index"`    //
	ErrorMsg string `json:"errorMsg"` //
}

type ImportErrorInfo struct {
	ErrorCount int                `json:"Error_count"` //
	ErrorData  []*ImportErrorItem `json:"error_data"`  //
}

func (e *ImportErrorInfo) Push(index int, errorMsg string) {
	e.ErrorData = append(e.ErrorData, &ImportErrorItem{
		Index:    index,
		ErrorMsg: errorMsg,
	})
	e.ErrorCount = len(e.ErrorData)
}

type ImportFieldAlias struct {
	ColumnName  string
	FieldName   string
	FieldType   string
	ColumnIndex int
	Required    bool
	TrimType    int
}

type ImportBase struct {
	SessionId   string
	ColumnCount int
	Columns     []*ImportFieldAlias
	Data        []map[string]interface{}
}

/*
trimType:
0--空格不处理
1--去除头尾空格
2--去除所有空格
*/
func (t *ImportBase) AppendColumn(columnName string, fieldName string, fieldType string, req bool, trimType int) {
	c := &ImportFieldAlias{
		ColumnName:  columnName,
		FieldName:   fieldName,
		FieldType:   fieldType,
		ColumnIndex: -1,
		Required:    req,
		TrimType:    trimType,
	}
	if req {
		t.ColumnCount = t.ColumnCount + 1
	}

	t.Columns = append(t.Columns, c)
	return
}

func (t *ImportBase) SetColumn(data []string) (err error) {
	for index, d := range data {
		for i, column := range t.Columns {
			if column.ColumnName == d {
				t.Columns[i].ColumnIndex = index
				break
			}
		}
	}
	return
}

func (t *ImportBase) VoidColumn() (err error) {
	for _, column := range t.Columns {
		if column.Required && column.ColumnIndex < 0 {
			err = errors.New(column.ColumnName + "没有输入")
			return
		}
	}

	return
}
