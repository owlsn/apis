package common

import (
	"github.com/owlsn/apis/src/common/config"
)

// 是否是正式环境
func IsProd() bool {
	return config.Conf.Env == "prod"
}
