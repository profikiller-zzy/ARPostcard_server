package cflag

import (
	"ARPostcard_server/biz/infra"
	"ARPostcard_server/biz/model"
	"fmt"
)

func MakeMigration() {
	var err error

	// 对模型自动迁移
	err = infra.MysqlDB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.User{},
			&model.Prefab{},
			&model.Video{},
			&model.Image{},
			&model.Menu{})
	if err != nil {
		fmt.Printf(err.Error()) // 这里后续使用日志
		return
	}
	fmt.Println("数据表迁移成功！")
}
