package main

import (
    "gorm.io/gen"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    // 替换为你自己的数据库连接信息
    dsn := "root:123456@tcp(127.0.0.1:3306)/hulio_user?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn))
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }

    g := gen.NewGenerator(gen.Config{
        OutPath: "dao/query",   // 生成的代码输出目录
        Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
    })

    g.UseDB(db)

    // 自动为表生成结构体
    g.GenerateModel("user")        // 指定表名
    g.Execute()
}
