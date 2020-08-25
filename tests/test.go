/*
@Time : 2020/7/17 11:47
@Author : xuyiqing
@File : test.go
*/

package main

import (
	"fmt"
	"github.com/unidoc/unioffice/document"
)

func main()  {
	doc, err := document.Open("test1.docx")
	if err != nil {
		fmt.Println(err)
	}
	img := doc.Images
	for _, img := range img {
		fmt.Println(img.Format())
		fmt.Println(img.Path())
	}
	for _, para := range doc.Paragraphs() {
		for _, t := range para.Runs() {
			fmt.Println(t.Text())
		}
	}
}
