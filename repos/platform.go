package repos

import (
	"go-rbac/model"
	"log"
)

func Save(model *model.Platform) {
	_, err := db.Exec(PLATFORM_INSERT_SQL, &model.PlatformCode, &model.PlatformName, &model.SignKey, &model.SignType)
	if err != nil {
		log.Println(err.Error())
		log.Fatal("插入错误")
	}

}
