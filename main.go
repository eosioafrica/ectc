package main
/*
import "github.com/eosioafrica/ecte/ecte"

var env ecte.Environment

func main() {


}*/

import (
	"fmt"
	"os/exec"
	"strings"
	"bytes"
	"log"
)

func main() {

	fmt.Printf("Hello, I am from 'hello_go.go' file.");
	fmt.Printf("\nNow, I am going to execute 'hello_sh.sh' file\n");
	cmd := exec.Command("bash", "/home/khosi/go/src/github.com/eosioafrica/ecte/ecte/.ecte/assets/get_dependency_installer.sh");
	cmd.Stdin = strings.NewReader("");
	var out bytes.Buffer;
	cmd.Stdout = &out;
	err := cmd.Run();
	if err != nil {
		log.Fatal(err);
	}
	fmt.Printf("Output \n",out.String());
}
