package pkg

import (
	"github.com/pro911/authApis/internal/pkg/conf"
	"github.com/pro911/authApis/internal/pkg/drives"
)

func InitProjects() {
	conf.MustInitConf()
	drives.MustInitMySQL()
	drives.MustInitMongo()
	drives.MustInitRedis()
}
