/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var clear = flag.Bool("c", false, "清空回收站")
var dir = flag.String("f", "", "删除目录(放入回收站 /tmp/rm)")
var help = flag.Bool("h", false, "帮助")

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *clear {
		fmt.Printf("execute rm -rf /tmp/rm \n")
		cmd := exec.Command("rm", "-rf", "/tmp/rm")

		output, err := cmd.CombinedOutput()
		if nil != err {

			fmt.Println(string(output))
			os.Exit(0)
		}
		return
	}

	if *dir == "" {
		flag.Usage()
		os.Exit(0)
	}
	if *dir == "/" || *dir == "/*" || strings.HasSuffix(*dir, "/ ") || strings.HasSuffix(*dir, "/* ") {
		fmt.Println("do not mv /!")
		os.Exit(0)
	}



	cmd := exec.Command("mkdir", "-p", "/tmp/rm/")

	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
	fileName:=fmt.Sprintf("/tmp/rm/%s-%d", *dir, time.Now().Unix())
	fmt.Printf("execute mv  %s /tmp/rm/%s\n", *dir, fileName)
	cmd = exec.Command("mv", "-b", *dir,fileName )

	output, err = cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
}
