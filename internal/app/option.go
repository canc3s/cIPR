package app

import (
	"flag"
	"fmt"
	"github.com/canc3s/cIPR/internal/tools"
	"github.com/canc3s/cIPR/pkg/qqwry"
	"log"
)

const banner = `
 ▄████████  ▄█     ▄███████▄    ▄████████ 
███    ███ ███    ███    ███   ███    ███ 
███    █▀  ███▌   ███    ███   ███    ███ 
███        ███▌   ███    ███  ▄███▄▄▄▄██▀ 
███        ███▌ ▀█████████▀  ▀▀███▀▀▀▀▀   
███    █▄  ███    ███        ▀███████████ 
███    ███ ███    ███          ███    ███ 
████████▀  █▀    ▄████▀        ███    ███ 
                               ███    ███
											v`

// Version is the current version of C
const Version = `0.0.1`

type Options struct {
	BlackList			[]string
	DatPath				string                 // Target is a single URL/Domain to scan usng a template
	InputFile			string                 // Targets specifies the targets to scan using templates.
	Silent				bool
}

func ParseOptions() *Options {
	options := &Options{}
	showBanner()
	flag.Parse()

	options.InputFile = flag.Arg(0)

	InitConf(options)
	options.validateOptions()

	return options
}

func (options *Options) validateOptions() {
	if options.DatPath != "" && !tools.FileExists(options.DatPath) {
		log.Println("文件 ", options.DatPath, " 文件不存在，尝试从网络获取最新纯真 IP 库")
		qqwry.Download(options.DatPath)
	}
	if options.InputFile == "" {
		log.Fatalf("用法:\t./cIPR domain.txt\n\n")
	}
	if options.InputFile != "" && !tools.FileExists(options.InputFile) {
		log.Fatalf("文件 %s 不存在!\n", options.InputFile)
	}
}

// showBanner is used to show the banner to the user
func showBanner() {
	fmt.Printf("%s%s\n", banner,Version)
	fmt.Printf("\thttps://github.com/canc3s/cIPR\n")

	//gologger.Labelf("请谨慎使用,您应对自己的行为负责\n")
	//gologger.Labelf("开发人员不承担任何责任，也不对任何滥用或损坏负责.\n")
}