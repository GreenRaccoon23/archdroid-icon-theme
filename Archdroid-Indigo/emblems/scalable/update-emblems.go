package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	buffer     bytes.Buffer
	verbose    bool
	numUpdated int
)

func init() {
	flag.BoolVar(&verbose, "v", false, "list every file updated")
	flag.Parse()
}

func main() {
	dir := "./"
	exts := []string{".svg", ".png", ".gif"}
	for _, e := range exts {
		emblemize(dir, e)
		iconize(dir, e)
	}
	if verbose {
		fmt.Println()
	}
	fmt.Printf("Updated %d emblems\n", numUpdated)
}

/*========================================
             Local Functions
========================================*/
func emblemize(dir, ext string) {
	pics := LsNonSyms(dir, ext)
	for _, p := range pics {
		FileAddPrefix(p.Name(), "emblem-")
	}
}

func iconize(dir, ext string) {
	pics := LsNonSyms(dir, ext)
	for _, p := range pics {
		genIcon(p)
		if verbose {
			fmt.Println(p.Name())
		}
		numUpdated += 1
	}
}

func genIcon(emblem os.FileInfo) {
	baseName := FileRemExt(emblem.Name())
	iconName := Concat(baseName, ".icon")
	if InFile(iconName, "DisplayName[af]") {
		return
	}
	iconFile := FileOverwrite(iconName)

	fmtdName := fmtName(baseName)
	dn := "DisplayName"
	data := fmt.Sprintf(
		"\n%v \n\n%s=%v \n%s[en_CA]=%v \n%s[en_GB]=%v",
		"[Icon Data]",
		dn,
		fmtdName,
		dn,
		fmtdName,
		dn,
		fmtdName,
	)

	StringToFile(data, iconFile)
}

func fmtName(fileName string) (fmtd string) {
	dePre := strings.TrimPrefix(fileName, "emblem-")
	reSpace := strings.Replace(dePre, "-", " ", -1)
	toTitle := strings.Title(reSpace)
	fmtd = toTitle
	return
}

/*========================================
              Global Functions
========================================*/
func LogErr(err error) {
	if err == nil {
		return
	}
	fmt.Println(err)
}

func Str(slice []string) (concatenated string) {
	for _, s := range slice {
		buffer.WriteString(s)
	}
	concatenated = buffer.String()
	buffer.Reset()
	return
}

func Concat(args ...string) string {
	return Str(args)
}

func IsFile(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true
	}
	return false
}

func IsFileLink(fileName string) bool {
	fi, err := os.Lstat(fileName)
	LogErr(err)
	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true
	}
	return false
}

func InFile(fileName, pattern string) bool {
	if IsFile(fileName) == false {
		return false
	}
	content := FileToString(fileName)
	if strings.Contains(content, pattern) == false {
		return false
	}
	return true
}

func LsNonSyms(dir, pattern string) (matches []os.FileInfo) {
	all, _ := ioutil.ReadDir(dir)
	for _, f := range all {
		name := f.Name()
		isLink := IsFileLink(name)
		hasPattern := strings.Contains(name, pattern)

		if isLink {
			continue
		}
		if hasPattern == false {
			continue
		}

		matches = append(matches, f)
	}
	return
}

func FileAddPrefix(fileName, prefix string) {
	isCorrectName := strings.HasPrefix(fileName, prefix)
	if isCorrectName {
		return
	}
	newName := Concat(prefix, fileName)
	os.Rename(fileName, newName)
}

func FileRemExt(fileName string) (base string) {
	ext := filepath.Ext(fileName)
	base = strings.Replace(fileName, ext, "", -1)
	return
}

func FileRemove(fileName string) {
	if IsFile(fileName) {
		return
	}
	os.Remove(fileName)
}

func FileCreate(fileName string) *os.File {
	file, err := os.Create(fileName)
	LogErr(err)
	return file
}

func FileOverwrite(fileName string) *os.File {
	FileRemove(fileName)
	file := FileCreate(fileName)
	return file
}

func FileToString(fileName string) (fileString string) {
	file, err := ioutil.ReadFile(fileName)
	LogErr(err)
	fileString = string(file)
	return
}

func StringToFile(s string, file *os.File) {
	b := []byte(s)
	_, err := file.Write(b)
	LogErr(err)
	return
}
