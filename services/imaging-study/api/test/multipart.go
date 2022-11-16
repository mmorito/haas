package main

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
)

const (
	exitOK int = iota
	exitErr
	targetPath = "/Users/mmorito/work/multipart"
)

func main() {
	var filenameList []string
	err := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		filenameList = append(filenameList, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	if len(filenameList) == 0 {
		os.Exit(exitOK)
		return
	}

	body := new(bytes.Buffer)
	mpWriter := multipart.NewWriter(body)

	boundary := "this_is_boundary"
	if err := mpWriter.SetBoundary(boundary); err != nil {
		panic(err)
	}

	for _, v := range filenameList {
		bytes, err := ioutil.ReadFile(v)
		if err != nil {
			panic(err)
		}

		part := make(textproto.MIMEHeader)
		part.Set("Content-Type", "application/dicom")
		writer, err := mpWriter.CreatePart(part)
		if err != nil {
			panic(err)
		}

		writer.Write(bytes)
	}

	mpWriter.Close()

	err = os.WriteFile(targetPath+"/study.multipart", body.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}
