package demo

import (
	"fmt"
	"os"
	"os/exec"
)

func InitExec() {
	// env := os.Environ()
	// procAttr := &os.ProcAttr{
	// 	Env: env,
	// 	Files: []*os.File{
	// 		os.Stdin,
	// 		os.Stdout,
	// 		os.Stderr,
	// 	},
	// }
	// pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	// if err != nil {
	// 	fmt.Printf("Error %v starting process!", err) //
	// 	os.Exit(1)
	// }
	// fmt.Printf("The process id is %v", pid)

	// pid, err := os.StartProcess("/bin/ps", []string{"ps", "-e", "-opid,ppid,comm"}, procAttr)
	// if err != nil {
	// 	fmt.Printf("Error %v starting process!", err) //
	// 	os.Exit(1)
	// }
	// fmt.Printf("The process id is %v", pid)

	cmd := exec.Command("ls", "-l")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)
}
