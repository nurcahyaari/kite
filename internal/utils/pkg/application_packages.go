package pkg

import (
	"os/exec"
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
	for _, p := range s.Packages {
		cmd := exec.Command("go", "get", p)
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
