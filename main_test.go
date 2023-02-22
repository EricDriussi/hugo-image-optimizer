package main_test

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	setCWDToProjectRoot()
	os.Exit(code)
}

func setCWDToProjectRoot() {
	_, filename, _, _ := runtime.Caller(0)
	project_root := path.Join(path.Dir(filename), "../../../../..")
	if err := os.Chdir(project_root); err != nil {
		panic(err)
	}
}
