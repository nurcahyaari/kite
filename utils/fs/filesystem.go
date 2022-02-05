package fs

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/nurcahyaari/kite/utils/logger"
	"golang.org/x/mod/modfile"
)

// checkTemplateLocationIsExist is a function for validate is the template location exist
func checkTemplateLocationIsExist(templateLocation string) error {
	return IsFileExist(templateLocation)
}

func ConcatDirPath(basePath, newDir string) string {
	basePath = ValidatePath(basePath)
	return fmt.Sprintf("%s%s", basePath, newDir)
}

// GetAppNameBasedOnGoMod will set the appname based on the go mod init
func GetAppNameBasedOnGoMod(goModName string) string {
	appName := strings.Split(goModName, "/")
	return appName[len(appName)-1]
}

func IsFolderExist(path string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return printErr(FolderWasCreated)
	}

	return nil
}

func IsFileExist(filelocation string) error {
	fileinfo, err := os.Stat(filelocation)
	if os.IsNotExist(err) {
		return printErr(FileNotFoundErr)
	}

	if fileinfo.IsDir() {
		return printErr(FileIsNotFileButADir)
	}

	return nil
}

func CreateFolder(path string) error {
	IsFileExist(path)

	if err := os.MkdirAll(path, 0755); err != nil {
		return printErr(MkdirErr)
	}

	return nil
}

func CreateFolderIsNotExist(path string) error {
	err := IsFolderExist(path)
	if err != nil {
		return err
	}
	CreateFolder(path)

	return nil
}

func CreateFile(path, fileName, template string) error {
	path = ValidatePath(path)

	resultFile, err := os.Create(fmt.Sprintf("%s%s", path, fileName))
	if err != nil {
		return printErr(MkfileErr)
	}

	err = WriteStringToFile(template, resultFile)
	if err != nil {
		return printErr(filesystemErr(err.Error()))
	}

	return nil
}

func CreateFileIfNotExist(path, fileName, template string) error {
	IsFileExist(fmt.Sprintf("%s%s", path, fileName))

	err := CreateFile(path, fileName, template)
	if err != nil {
		return err
	}

	return nil
}

func ReplaceFile(path, filename, template string) error {
	err := DeleteFile(ConcatDirPath(path, filename))
	if err != nil {
		return err
	}

	err = CreateFile(path, filename, template)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFolder(path string) {
	os.RemoveAll(path)
}

func DeleteFile(filepath string) error {
	err := os.Remove(filepath)
	return err
}

func GoFormat(path, goModName string) error {
	cmd := exec.Command("go", "fmt", fmt.Sprintf("%s/...", goModName))
	cmd.Dir = path
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return printErr(GoFmtErr)
	}

	return nil
}

func GoModInit(projectPath, goModName string) error {
	os.Chdir(projectPath)
	cmd := exec.Command("go", "mod", "init", goModName)
	if err := cmd.Run(); err != nil {
		return printErr(GoModInitErr)
	}
	return nil
}

func GetGoModName(gomodPath string) string {
	goModBytes, err := ioutil.ReadFile(fmt.Sprintf("%sgo.mod", gomodPath))
	if err != nil {
		logger.Errorln(err.Error())
		return ""
	}

	modName := modfile.ModulePath(goModBytes)

	return modName
}

func ValidatePath(path string) string {
	pathSplit := strings.Split(path, "")

	if pathSplit[len(pathSplit)-1] != "/" {
		pathSplit = append(pathSplit, "/")
	}

	return strings.Join(pathSplit, "")
}

func WriteStringToFile(template string, file *os.File) error {
	_, err := file.WriteString(template)

	if err != nil {
		return err
	}

	return nil
}

type ImportedPackages struct {
	Alias    string
	FilePath string
}

type DependencyInterface struct {
	Method string
}

type DependencyStruct struct {
	ObjectName     string
	ObjectDataType string
}

type Dependency struct {
	Name string
	DependencyInterface
	DependencyStruct
}

func ReadImportedPackages(dependency string) []ImportedPackages {
	importedPackages := []ImportedPackages{}

	dependencyStringSplit := strings.Split(dependency, "\n")
	regexFindImportOneLine := regexp.MustCompile(`import`)
	regexFindMultipleImport := regexp.MustCompile(`import \(`)
	regexEndOfMultipleImport := regexp.MustCompile(`\)`)
	regexCleaning := regexp.MustCompile("\"|[\t]*|[\n]*")
	regexFrontSpaceCleaning := regexp.MustCompile("^ +")
	regexEndSpaceCleaning := regexp.MustCompile(`\s+$`)
	regexSpaceCleaning := regexp.MustCompile("  *")
	foundMultipleImport := false
	for _, p := range dependencyStringSplit {

		matchEndOfMultipleInterface := regexEndOfMultipleImport.MatchString(p)
		if matchEndOfMultipleInterface {
			foundMultipleImport = false
		}

		if foundMultipleImport {
			repl := regexCleaning.ReplaceAll([]byte(p), []byte(""))
			repl = regexFrontSpaceCleaning.ReplaceAll(repl, []byte(""))
			repl = regexEndSpaceCleaning.ReplaceAll(repl, []byte(""))
			repl = regexSpaceCleaning.ReplaceAll(repl, []byte(" "))

			importedPackageString := strings.Split(string(repl), " ")
			alias := ""
			filePath := ""
			if len(importedPackageString) > 1 {
				alias = importedPackageString[0]
				filePath = importedPackageString[1]
			} else {
				filePath = importedPackageString[0]
			}

			if string(repl) == "" {
				continue
			}
			importedPackages = append(importedPackages, ImportedPackages{
				Alias:    alias,
				FilePath: filePath,
			})
		}

		matchMultipleImport := regexFindMultipleImport.MatchString(p)
		if matchMultipleImport {
			foundMultipleImport = true
		}
	}

	if len(importedPackages) == 0 {
		for _, p := range dependencyStringSplit {
			matchImport := regexFindImportOneLine.MatchString(p)
			if matchImport {
				p = string(regexFindImportOneLine.ReplaceAll([]byte(p), []byte("")))
				repl := regexCleaning.ReplaceAll([]byte(p), []byte(""))
				repl = regexFrontSpaceCleaning.ReplaceAll(repl, []byte(""))
				repl = regexEndSpaceCleaning.ReplaceAll(repl, []byte(""))
				repl = regexSpaceCleaning.ReplaceAll(repl, []byte(" "))
				importedPackageString := strings.Split(string(repl), " ")
				alias := ""
				filePath := ""
				if len(importedPackageString) > 1 {
					alias = importedPackageString[0]
					filePath = importedPackageString[1]
				} else {
					filePath = importedPackageString[0]
				}

				if string(repl) == "" {
					continue
				}
				importedPackages = append(importedPackages, ImportedPackages{
					Alias:    alias,
					FilePath: filePath,
				})
			}
		}
	}

	return importedPackages
}

func ReadInterfaceWithMethod(dependency string) []DependencyInterface {
	dependencyInterfaces := []DependencyInterface{}
	dependencyStringSplit := strings.Split(dependency, "\n")
	regexFindInterface := regexp.MustCompile(`type (.*) interface {`)
	regexFindEndOfInterface := regexp.MustCompile(`}`)
	regexCleaning := regexp.MustCompile("[\t]|[\n]")
	regexFrontSpaceCleaning := regexp.MustCompile("^ +")
	regexEndSpaceCleaning := regexp.MustCompile(`\s+$`)
	regexSpaceCleaning := regexp.MustCompile("  *")
	foundInterface := false
	for _, p := range dependencyStringSplit {
		matchEndOfInterface := regexFindEndOfInterface.MatchString(p)

		if matchEndOfInterface {
			foundInterface = false
		}

		if foundInterface {
			repl := regexCleaning.ReplaceAll([]byte(p), []byte(""))
			repl = regexFrontSpaceCleaning.ReplaceAll(repl, []byte(""))
			repl = regexEndSpaceCleaning.ReplaceAll(repl, []byte(""))
			repl = regexSpaceCleaning.ReplaceAll(repl, []byte(" "))
			dependencyInterfaces = append(dependencyInterfaces, DependencyInterface{
				Method: string(repl),
			})
		}

		matchInterface := regexFindInterface.MatchString(p)
		if matchInterface {
			foundInterface = true
		}
	}
	return dependencyInterfaces
}

func ReadStructWithObject(dependency string) []DependencyStruct {
	dependencyStruct := []DependencyStruct{}
	dependencyStringSplit := strings.Split(dependency, "\n")
	regexFindStruct := regexp.MustCompile(`type (.*) struct \{`)
	regexFindEndOfStruct := regexp.MustCompile(`\}`)
	regexCleaning := regexp.MustCompile("[\t]|[\n]")
	regexFrontSpaceCleaning := regexp.MustCompile("^ +")
	regexEndSpaceCleaning := regexp.MustCompile(`\s+$`)
	regexSpaceCleaning := regexp.MustCompile("  *")
	foundStruct := false
	for _, p := range dependencyStringSplit {
		matchEndOfInterface := regexFindEndOfStruct.MatchString(p)

		if matchEndOfInterface {
			foundStruct = false
		}

		if foundStruct {
			repl := regexCleaning.ReplaceAll([]byte(p), []byte(""))
			repl = regexFrontSpaceCleaning.ReplaceAll(repl, []byte(""))
			repl = regexEndSpaceCleaning.ReplaceAll(repl, []byte(""))
			repl = regexSpaceCleaning.ReplaceAll(repl, []byte(" "))
			obj := strings.Split(string(repl), " ")
			dependencyStruct = append(dependencyStruct, DependencyStruct{
				ObjectName:     obj[0],
				ObjectDataType: obj[1],
			})
		}

		matchInterface := regexFindStruct.MatchString(p)
		if matchInterface {
			foundStruct = true
		}
	}
	return dependencyStruct
}

func ReadMethodImpl(dependency string) []string {
	methodImpl := []string{}
	dependencyStringSplit := strings.Split(dependency, "\n")
	regexFindInterface := regexp.MustCompile(`func \(.*\) (.*) \{`)
	regexCleaning := regexp.MustCompile("[\n]*")
	foundInterface := false

	funcImpl := ""
	for _, p := range dependencyStringSplit {
		matchInterface := regexFindInterface.MatchString(p)
		if matchInterface {
			foundInterface = true
		}

		if foundInterface {
			repl := regexCleaning.ReplaceAll([]byte(p), []byte(""))
			if string(repl) != "" {
				funcImpl += fmt.Sprintf("%s\n", string(repl))
			}
		}
	}

	funcSplit := strings.Split(funcImpl, "func (")

	for _, f := range funcSplit {
		if f == "" {
			continue
		}
		methodImpl = append(methodImpl, fmt.Sprintf("func (%s", f))
	}

	return methodImpl
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
