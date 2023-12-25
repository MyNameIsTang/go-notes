package demo

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func InitGzipFile() {
	filename, _ := filepath.Abs("demo/products.txt")
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create("/Users/daguang/main/go-notes/demo/products.gz")
	if err != nil {
		panic(err)
	}
	gzipWriter := gzip.NewWriter(newFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, file)
	// os.WriteFile()
	if err != nil {
		panic(err)
	}
}

func InitGzipped() {
	filename, _ := filepath.Abs("demo/products.gz")
	var r *bufio.Reader
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], filename,
			err)
		os.Exit(0)
	}
	defer fi.Close()
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
