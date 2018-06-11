package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

var hash_var string
var file []string

func main() {

	for _, f := range file {
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

		//fp := os.OpenFile(file,os.O_RDONLY)
		fp, err := os.Open(f)

		if err != nil {
			fmt.Println("read error.")
			os.Exit(1)
		}
		defer fp.Close()

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
		fmt.Printf("%x  %s\n", sha.Sum(nil), f)

		//fmt.Printf("%#v %#v\n", f, hash_var)
	}

}

func init() {
    const HELP="Usage: sha-sum [-hash <sha>] <file...>"
	//flag.StringVar(file, "f", "", "a file")
	flag.StringVar(&hash_var, "hash", "md5", "hash sha512,385,256,224,1 or md5 al.")
    help := flag.Bool("help",false,HELP)
	flag.Parse()

    if *help || flag.NArg() < 1 {
        fmt.Println(HELP)
        os.Exit(1)
    }

	file = flag.Args()

	//fmt.Println(file)

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
	//fmt.Println(f, hash_var)

	for _, f := range file {

		fi, err := os.Stat(f)
		if os.IsNotExist(err) {
			fmt.Println(f, "文件不存在.")
			os.Exit(1)
		}
		if fi.IsDir() {
			fmt.Println(f, "不是一个文件")
			os.Exit(1)
		}
	}
}
