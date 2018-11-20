package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mholt/archiver"
	"github.com/nwaples/rardecode"
)

func main() {
	r := archiver.NewRar()

	err := r.Walk("./test/我的测试.rar", func(f archiver.File) error {
		rh, ok := f.Header.(*rardecode.FileHeader)
		if !ok {
			return fmt.Errorf("expected header")
		}
		fmt.Println("FileName:", rh.Name)

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

	err = r.Unarchive("./test/我的测试.rar", "./test/我的测试")
	if err != nil {
		panic(err)
	}
}