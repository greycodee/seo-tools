package google

import (
	"errors"
	"os"
)

type IndexService struct {

}

func getGoogleEnv() error {
	googleEnv := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if googleEnv == "" {
		return errors.New("请设置 GOOGLE_APPLICATION_CREDENTIALS 环境变量")
	}
	return nil
}
