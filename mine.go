package main

import (
	"fmt"
	"os/exec"
	"github.com/kopoli/go-terminal-size"
)

func err_check(err error) {
	if err != nil {
		panic(err);
	}
}

const PASSWORD string = "quandale"

func main() {

	var UserIO string;

	for {
		size, err := tsize.GetSize();
		fmt.Print("Library password: ")
		fmt.Scanln(&UserIO);
		for i:=0;i<=size.Height;i++ {fmt.Println("")}
		if UserIO == PASSWORD {
			fmt.Println("\033[32mCORRECT PASSWORD");
			break
		} else {
			fmt.Println("\033[31mFAILED TO AUTHENTICATE:", UserIO, "\033[0m");
			err_check(err);
		}
	}

	_, err := exec.Command("open", "./.sysconfig/Library.app").Output();
	err_check(err);
	fmt.Println("\033[01mOPENING LIBRARY\033[0m");
}