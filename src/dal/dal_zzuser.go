package dal

import (
	"base_gin/commons/dataBase"
	"database/sql"
	"fmt"
	"strings"
)

// DalInfoCreateOneZzUser ZzUser 添加一条数据
func DalCreateOneZzUser(args []any) (int, int) {
	valTemplate := dataBase.BuildInQuerySmp(len(args))
	res, err := dataBase.ExecuteSQL("INSERT INTO zzuser(name,account,age,email,count,create_time,update_time) VALUES"+valTemplate, args...)
	if err != nil {
		fmt.Println("插入一条数据出错！", err)
	}
	affectedRows, _ := res.RowsAffected()
	lastInsertId, _ := res.LastInsertId()
	return int(affectedRows), int(lastInsertId)

}

// DalCreateManyZzUser mmmm
func DalCreateManyZzUser(args [][]any) (int, int) {
	var valTemplateSlice []string
	var argsSlice []any
	for _, arg := range args {
		argsSlice = append(argsSlice, arg...)
		valTemplateSlice = append(valTemplateSlice, dataBase.BuildInQuerySmp(len(arg)))
	}
	valTemplate := strings.Join(valTemplateSlice, ",")
	res, _ := dataBase.ExecuteSQL("INSERT INTO zzuser(name,account,age,email,count,create_time,update_time) VALUES"+valTemplate, argsSlice...)
	affectedRows, _ := res.RowsAffected()
	lastInsertId, _ := res.LastInsertId()
	return int(affectedRows), int(lastInsertId)

}

// DalQueryAllZzUser 没有意义
func DalQueryAllZzUser() []map[string]any {

	toMap := dataBase.QueryToMap("SELECT * FROM zzuser ORDER BY age DESC")
	return toMap
}

func DalQueryWhereZzUser(ids []any, ageMax int) []map[string]any {
	valTemplate := dataBase.BuildInQuerySmp(len(ids))
	var argsSlice []any
	argsSlice = append(argsSlice, ageMax)
	argsSlice = append(argsSlice, ids...)
	queryString := fmt.Sprintf("SELECT * FROM zzuser WHERE age < ? AND id IN %s AND count IS NOT NULL ORDER BY id DESC", valTemplate)
	//queryString = fmt.Sprintf("SELECT * FROM zzuser WHERE age < 60 AND id IN (1,2,3,4,5,6,7,8,9) ORDER BY id DESC")
	fmt.Println(queryString, argsSlice)
	toMap := dataBase.QueryToMap(queryString, argsSlice...)
	fmt.Println(toMap)
	return toMap
}

func DalChangeWhereZzUser(ids []any, newEmail string) int {
	valTemplate := dataBase.BuildInQuerySmp(len(ids))
	var argsSlice []any
	argsSlice = append(argsSlice, newEmail)
	argsSlice = append(argsSlice, ids...)
	queryString := fmt.Sprintf("UPDATE zzuser SET email = ? WHERE id IN %s ", valTemplate)

	res, _ := dataBase.ExecuteSQL(queryString, argsSlice...)
	affectedRows, _ := res.RowsAffected()
	return int(affectedRows)

}

// DalDeleteWhereZzUser
func DalDeleteWhereZzUser(id int) int {
	queryString := fmt.Sprintf("DELETE FROM zzuser WHERE id = ?")
	res, err := dataBase.ExecuteSQL(queryString, id)
	fmt.Println(res, err)
	affectedRows, _ := res.RowsAffected()
	return int(affectedRows)
}

// DalQueryJoinUser  cc
func DalQueryJoinUser(ids []any, age int, count int) []map[string]any {
	valTemplate := dataBase.BuildInQuerySmp(len(ids))
	var argsSlice []any
	argsSlice = append(argsSlice, count, age)
	argsSlice = append(argsSlice, ids...)

	queryString := "SELECT u.*,i.tel,i.likes FROM zzuser as u JOIN zzinfo as i ON u.id = i.u_id WHERE i.count < ? AND age < ? AND u.id IN" + valTemplate + " ORDER BY age DESC"
	data := dataBase.QueryToMap(queryString, argsSlice...)
	return data
}

func DalTestExecTx() error {
	return dataBase.ExecTx(
		func(tx *sql.Tx) error {
			// 业务处理 在这里写事物的逻辑
			// likes:="鸡哥来了"
			// id:=10
			//queryString := "UPDATE zzuser SET likes = ? WHERE id = ?"
			// afr_, err := tx.Exec(queryString, likes,id)
			// return nil
			return fmt.Errorf("这里有问题了%s", "Hello world!")
		})
}
