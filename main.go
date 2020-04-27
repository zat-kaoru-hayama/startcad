package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var optionCadName = flag.String("v", "", "BricsCAD class name")

func getCadApplicationName() (string, error) {
	k, err := registry.OpenKey(registry.CLASSES_ROOT,
		`BricscadApp.AcadApplication\CurVer`,
		registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	val, _, err := k.GetStringValue("")
	if err != nil {
		return "", err
	}
	return val, nil
}

func mains(args []string) error {
	cadname := *optionCadName
	if cadname == "" {
		var err error
		cadname, err = getCadApplicationName()
		if err != nil {
			return fmt.Errorf("getCadApplicationName: %w", err)
		}
	}
	fmt.Println(cadname)
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	_cad, err := oleutil.CreateObject(cadname)
	if err != nil {
		return err
	}
	cad, err := _cad.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer cad.Release()

	cad.PutProperty("Visible", true)

	_doc, err := cad.GetProperty("ActiveDocument")
	if err != nil {
		return err
	}
	doc := _doc.ToIDispatch()
	defer doc.Release()

	for _, text := range args {
		doc.CallMethod("SendCommand", text+"\r")
	}
	return nil
}

func main() {
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
