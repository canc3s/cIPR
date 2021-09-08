package app

import (
	"flag"
	"fmt"
	"github.com/canc3s/cIPR/internal/tools"
	"github.com/canc3s/cIPR/pkg/qqwry"
	"log"
	"os"
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
const Version = `0.0.2`

type Options struct {
	BlackList			[]string
	DatPath				string                 // Target is a single URL/Domain to scan usng a template
	InputFile			string                 // Targets specifies the targets to scan using templates.
	GoroutineNum		int
	Silent				bool
}

func ParseOptions() *Options {
	options := &Options{}
	flag.IntVar(&options.GoroutineNum, "t", 10, "并发数量")
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
		log.Fatalf("用法:\n./cIPR domain.txt\n./cIPR -t 20 domain.txt\n默认并发10\n\n")
		os.Exit(0)
	}
	if options.InputFile != "" && !tools.FileExists(options.InputFile) {
		log.Fatalf("文件 %s 不存在!\n", options.InputFile)
		os.Exit(0)
	}
}

// showBanner is used to show the banner to the user
func showBanner() {
	fmt.Printf("%s%s\n", banner,Version)
	fmt.Printf("\thttps://github.com/canc3s/cIPR\n")

	//gologger.Labelf("请谨慎使用,您应对自己的行为负责\n")
	//gologger.Labelf("开发人员不承担任何责任，也不对任何滥用或损坏负责.\n")
}