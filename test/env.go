package test

import (
	"fmt"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/provider/app"
	"github.com/gohade/hade/framework/provider/env"
	"os"
	"path/filepath"
)

const (
	BasePath = ""
)

// GetRootDirectory 获取当前项目根目录
func GetRootDirectory() (string, error) {
	executable, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(executable)
	for {
		if _, err := os.Stat(filepath.Join(dir, ".go-root")); err == nil {
			return dir, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("unable to find project root")
}

func GetBasePath() string {
	root, err := GetRootDirectory()
	if err != nil {
		return ""
	}
	return root
}

func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewHadeContainer()
	// 绑定App服务提供者
	var basePath = BasePath
	if basePath == "" {
		basePath = GetBasePath()
	}
	container.Bind(&app.HadeAppProvider{BaseFolder: basePath})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.HadeTestingEnvProvider{})
	return container
}
