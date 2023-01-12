package main

import (
	"go_learn_web/pkg/mysql"
	"gorm.io/gen"
	"path"
	"runtime"
	"strings"
)

/**
参考文档 https://blog.csdn.net/qq_43792385/article/details/121602049
*/

func main() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	root = strings.TrimRight(root, "/pkg/gorm")

	g := gen.NewGenerator(gen.Config{
		OutPath: root + "/dao/query",
	})

	db := mysql.GetGormDB()
	g.UseDB(db)
	g.ApplyBasic(g.GenerateModelAs("t_user", "User"))
	g.Execute()
}
