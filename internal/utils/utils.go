package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/mod/modfile"
)

func CapitalizeFirstLetter(s string) string {
	r := []rune(s)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}

func ConcatDirPath(dir ...string) string {
	newPath := ""
	for i, d := range dir {
		path := d
		if i+1 < len(dir) {
			path = AddSlashOnPath(path)
		}
		newPath = fmt.Sprintf("%s%s", newPath, path)
	}
	return newPath
}

func RemoveOnlySlashPath(dirPaths []string) []string {
	newDirPaths := []string{}
	for _, p := range dirPaths {
		if p != "/" {
			newDirPaths = append(newDirPaths, p)
		}
	}
	return newDirPaths
}

func AddSlashOnPath(path string) string {
	pathSplit := strings.Split(path, "")
	if len(pathSplit) > 0 {
		if pathSplit[len(pathSplit)-1] != "/" {
			pathSplit = append(pathSplit, "/")
		}

		return strings.Join(pathSplit, "")
	}
	return path
}

func RemoveSlashFirstAndLast(path string) string {
	pathSplit := strings.Split(path, "")
	if pathSplit[0] == "/" {
		pathSplit[0] = ""
	}
	if pathSplit[len(pathSplit)-1] == "/" {
		pathSplit[len(pathSplit)-1] = ""
	}
	return strings.Join(pathSplit, "")
}

func ReadFile(filePath string) (string, error) {
	var fileValue string
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileValue += fmt.Sprintf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return fileValue, nil
}

func GoGenerateRun(projectPath string) error {
	os.Chdir(projectPath)
	cmd := exec.Command("go", "generate", ".")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func GetGoModName(gomodPath string) string {
	path := ConcatDirPath(gomodPath, "go.mod")
	goModBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	modName := modfile.ModulePath(goModBytes)
	return modName
}

func IsFolderHasGoMod(path string) bool {
	s := GetGoModName(path)
	return s != ""
}

func GetProjectPath(path string, pointingPath string) string {
	pathSplit := strings.Split(path, pointingPath)
	return pathSplit[0]
}

func GetLastDirPath(path string) string {
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	paths := strings.Split(path, "/")
	return strings.ToLower(paths[len(paths)-1])
}

func GoFormat(path, goModName string) error {
	cmd := exec.Command("go", "fmt", fmt.Sprintf("%s/...", goModName))
	cmd.Dir = path
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Gitinit(projectPath string) error {
	os.Chdir(projectPath)
	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func GoModInit(projectPath, goModName string) error {
	os.Chdir(projectPath)
	cmd := exec.Command("go", "mod", "init", goModName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func GetImportPathBasedOnProjectPath(projectPath, gomodName string) string {
	s := strings.Split(projectPath, gomodName)
	s = RemoveOnlySlashPath(s)
	if len(s) == 0 {
		return ""
	}
	return ConcatDirPath(gomodName, RemoveSlashFirstAndLast(s[len(s)-1]))
}

func GetGoFilesInPath(path string) ([]string, error) {
	paths := []string{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".go" {
			paths = append(paths, ConcatDirPath(path, file.Name()))
		}
	}

	return paths, nil
}
