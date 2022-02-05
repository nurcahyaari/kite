package pkg

import (
	"fmt"
	"os/exec"

	"github.com/nurcahyaari/kite/utils/logger"
)

type AppPackages struct {
	Packages []string
}

func (s *AppPackages) AddPackage(pkg string) {
	s.Packages = append(s.Packages, pkg)
}

func (s AppPackages) GetPackages() []string {
	return s.Packages
}

func (s AppPackages) CheckPackageAvailable(pkg string) bool {
	for _, p := range s.Packages {
		if p == pkg {
			return true
		}
	}

	return false
}

func (s AppPackages) InstallPackage() error {
	logger.Infoln("installing packages")
	for _, p := range s.Packages {
		logger.Info(fmt.Sprintf("installing %s...", p))
		cmd := exec.Command("go", "get", p)
		if err := cmd.Run(); err != nil {
			return err
		}
		logger.InfoSuccessln("success")
	}
	return nil
}
