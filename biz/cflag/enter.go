package cflag

import "flag"

type Options struct {
	Mysql bool
}

// Parse 解析命令参数，并对不同的命令行参数的值来执行不同的操作
func Parse() {
	mysqlFlag := flag.Bool("mysql", false, "auto migrate database")
	flag.Parse()
	var option = Options{
		Mysql: *mysqlFlag,
	}
	Execute(option)
}

func Execute(options Options) {
	if options.Mysql {
		MakeMigration()
		return
	}
}
