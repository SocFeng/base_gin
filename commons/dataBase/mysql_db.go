package dataBase

import (
	"base_gin/commons/config"
	"base_gin/commons/logs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

var GlobalDB *sql.DB

func createDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		config.GlobalConfig.DataBase.User,
		config.GlobalConfig.DataBase.Password,
		config.GlobalConfig.DataBase.Host,
		config.GlobalConfig.DataBase.Port,
		config.GlobalConfig.DataBase.DbName)
	var err error
	db, err := sql.Open("mysql", dsn)
	fmt.Println(dsn)
	if err != nil {
		logs.AppFatal("数据库连接失败:", err)
	}

	// 配置连接池
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err = db.Ping(); err != nil {
		logs.AppFatal("数据库 Ping 失败:", err)
	}
	logs.AppInfo("✅✅✅ 数据库连接成功！")

	return db
}
func InitDB() {
	GlobalDB = createDB()

}

func formatResultData(rows *sql.Rows) ([]map[string]any, error) {
	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("获取列名失败: %w", err)
	}
	var results []map[string]any

	// 遍历每一行数据
	for rows.Next() {
		// 创建值的容器（必须用 []any 接收）
		values := make([]any, len(columns))
		pointers := make([]any, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}

		// 扫描数据到容器
		if err := rows.Scan(pointers...); err != nil {
			return nil, fmt.Errorf("扫描行失败: %w", err)
		}

		// 将值转换为 map
		rowMap := make(map[string]any)
		for i, colName := range columns {
			val := values[i]

			// 处理特殊类型（如 time.Time 转字符串）
			switch v := val.(type) {
			case time.Time:
				rowMap[colName] = v.Format("2006-01-02 15:04:05")
			case []byte:
				rowMap[colName] = string(v) // 处理 BLOB/TEXT 等二进制字段
			default:
				rowMap[colName] = val
			}
		}

		results = append(results, rowMap)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历行失败: %w", err)
	}
	return results, nil
}

// QueryToMap 通用查询函数：将结果转换为 []map[string]any
// 多用于查询
func QueryToMap(query string, args ...any) []map[string]any {
	rows, err := GlobalDB.Query(query, args...)
	if err != nil {
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)
	results, err := formatResultData(rows)
	if err != nil {
		logs.AppFatal("解析阐述出错！！！！")
	}

	return results
}

// ExecuteSQL 执行指定sql语句
// 多用于 删除 修改 插入
func ExecuteSQL(query string, args ...any) (sql.Result, error) {
	// 执行SQL语句
	result, err := GlobalDB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("SQL执行失败: %w", err)
	}

	return result, nil
}

// BuildInQuery
// 构建 IN语句的查 (?,?...) 字符串 --- old
func BuildInQuery[T any](slice []T) (string, []any) {
	if len(slice) == 0 {
		return "1=0", nil
	}

	placeholders := strings.Repeat("?,", len(slice)-1) + "?"
	args := make([]any, len(slice))
	for i, v := range slice {
		args[i] = v
	}

	return fmt.Sprintf("(%s)", placeholders), args
}

// BuildInQuerySmp
// 构建 IN语句的查 (?,?...) 字符串 --- new 简约版的
func BuildInQuerySmp(num int) string {
	argsTemplate := "(" + strings.Repeat("?,", num-1) + "?" + ")"
	return argsTemplate
}

// Begin 开启事务
func Begin() (*sql.Tx, error) {
	tx, err := GlobalDB.Begin()
	if err != nil {
		logs.AppFatal("事务开启失败:", err)
		return nil, err
	}
	return tx, nil
}

// Rollback 回滚事务
func Rollback(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil && err != sql.ErrTxDone {
		logs.AppFatal("事务回滚失败:", err)
		return err
	}
	return nil
}

// Commit 提交事务
func Commit(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		logs.AppFatal("事务提交失败:", err)
		return err
	}
	return nil
}

// TxFunc 定义事务执行函数类型
type TxFunc func(tx *sql.Tx) error

// ExecTx 封装事务执行逻辑
func ExecTx(txFunc TxFunc) error {
	tx, err := Begin()
	if err != nil {
		return fmt.Errorf("事务开启失败: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			// 捕获panic并回滚
			_ = Rollback(tx)
			panic(p) // 重新抛出panic
		}
		if err != nil {
			// 普通错误回滚
			_ = Rollback(tx)
		}
	}()

	// 执行业务逻辑
	if err = txFunc(tx); err != nil {
		return fmt.Errorf("业务执行失败: %w", err)
	}

	// 提交事务
	if err = Commit(tx); err != nil {
		return fmt.Errorf("事务提交失败: %w", err)
	}
	return nil
}
