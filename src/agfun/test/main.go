package main

import (
	"agfun/agfun-service/log"
	"fmt"
	"io/ioutil"
	"os"

	"archive/zip"
	"github.com/mholt/archiver"
)

func main() {
	//r := archiver.NewRar()
	//
	//err := r.Walk("./test/我的测试.rar", func(f archiver.File) error {
	//	rh, ok := f.Header.(*rardecode.FileHeader)
	//	if !ok {
	//		return fmt.Errorf("expected header")
	//	}
	//	fmt.Println("FileName:", rh.Name)
	//
	//	content, err := ioutil.ReadAll(f)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Println("FileContent:", string(content))
	//
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = r.Unarchive("./test/我的测试.rar", "./test/我的测试")
	//if err != nil {
	//	panic(err)
	//}

	fzip()
	//stdZip()
}

func fzip() {
	z := archiver.DefaultZip
	err := z.Walk("./test/我的测试.zip", func(f archiver.File) error {
		zfh, ok := f.Header.(zip.FileHeader)
		if !ok {
			return fmt.Errorf("expected header")
		}
		fmt.Println("FileName:", zfh.Name)

		content, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		fmt.Println("FileContent:", string(content))

		return nil
	})
	if err != nil {
		panic(err)
	}

	err = z.Unarchive("./test/我的测试.zip", "./test/我的测试")
	if err != nil {
		panic(err)
	}
}

func stdZip() {
	r, err := zip.OpenReader("./test/我的测试.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, e := f.Open()
		if e != nil {
			log.Fatal(e)
		}
		buf, e := ioutil.ReadAll(rc)
		if e != nil {
			log.Fatal(e)
		}
		e = ioutil.WriteFile("./test/"+f.Name, buf, os.ModePerm)
		if e != nil {
			log.Fatal(e)
		}
		rc.Close()
	}
}
