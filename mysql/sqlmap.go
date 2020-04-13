package mysql

import "fmt"

// 使用原生 SQL 查询单条记录并封装成 map[string]interface{}
func QueryForMap(sql string, args ...interface{}) (map[string]interface{}, error) {
	rows, _ := Conn.DB().Query(sql, args...)
	cols, _ := rows.Columns()
	colTypes, _ := rows.ColumnTypes()
	fmt.Println(colTypes)
	defer rows.Close()
	if rows.Next() {

		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}
		m := make(map[string]interface{})
		for i, colName := range cols {
			// TODO 类型映射
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		return m, nil

	} else {
		//return nil, errors.New("未获取到返回信息");
		return map[string]interface{}{}, nil
	}
}
