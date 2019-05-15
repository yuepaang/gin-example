package logging

import (
	"fmt"
	"time"

	"github.com/ypeng7/data-microservices/pkg/setting"
)

func GetLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func GetLogFileName() string {
	return fmt.Sprintf(
		"%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogSaveExt,
	)
}
