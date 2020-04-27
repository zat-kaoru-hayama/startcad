package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var optionCadObject = flag.String("v", "BricscadApp.AcadApplication", "Ole automation name")

func mains(args []string) error {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	_cad, err := oleutil.CreateObject(*optionCadObject)
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
