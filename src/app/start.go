package app

import (
	"github.com/owlsn/apis/src/common"
)

func StartOn() {
	if !common.IsProd() {
		return
	}

	// 开启定时任务
	StartSchedule()
}
