package main

import (
	"crypto/sha512"
	"crypto/sha256"
    "crypto/sha1"
	"crypto/md5"
    "hash"
	"flag"
	"fmt"
	"io"
	"os"
)

var file, hash_var string

func main() {
	//fp := os.OpenFile(file,os.O_RDONLY)
	fp, err := os.Open(file)

	if err != nil {
		fmt.Println("read error.")
		os.Exit(1)
	}
	defer fp.Close()

    var sha hash.Hash
    switch hash_var {
    case "sha512":
	    sha = sha512.New()
    case "sha384":
	    sha = sha512.New384()
    case "sha256":
        sha = sha256.New()
    case "sha224":
        sha = sha256.New224()
    case "sha1":
	    sha = sha1.New()
    case "md5":
        sha = md5.New()
    default:
        sha = md5.New()
    }

    //var sha = select_sha()

	buf := make([]byte, 4096)

	//for n, err :=fp.Read(buf);n != 0 && err != io.EOF;n, err = fp.Read(buf){
	for {
		n, err := fp.Read(buf)

		if err != nil && err != io.EOF {
			fmt.Println("error :", err)
			os.Exit(2)
		}

		if err == io.EOF {
			//fmt.Println("len(buf) and n", len(buf), n)
			break
		} else {
			//fmt.Println("读成功的时候err值:",err)
			sha.Write(buf[:n])
		}
	}
	fmt.Printf("%x  %s\n", sha.Sum(nil), file)

	//fmt.Printf("%#v %#v\n", file, hash_var)
}

func init() {
	flag.StringVar(&file, "f", "", "a file")
	flag.StringVar(&hash_var, "hash", "md5", "hash sha512,385,256,224,1 or md5 al.")
	flag.Parse()

	fi, err := os.Stat(file)
	if os.IsNotExist(err) {
		fmt.Println(file, "文件不存在.")
		os.Exit(1)
	}
	if fi.IsDir() {
		fmt.Println(file, "不是一个文件")
		os.Exit(1)
	}

	switch hash_var {
	case "sha512":
	case "sha384":
	case "sha256":
	case "sha224":
	case "sha1":
	case "md5":
	default:
		fmt.Println("hash 必须是sha512, sha384, sha256, sha224,sha1和md5中的一种")
		os.Exit(1)
	}
	fmt.Println(file, hash_var)
}

