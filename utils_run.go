package easysdk

import (
	"log"
	"os/exec"
)

func RunCMD(Command string, Home string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", Command)
	cmd.Dir = Home

	retData, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	// stderr, _ := cmd.StderrPipe()
	// if err := cmd.Start(); err != nil {
	// 	log.Println(err)
	// 	return ""
	// }

	// retData := ""

	// scanner := bufio.NewScanner(stderr)
	// for scanner.Scan() {
	// 	retData += scanner.Text()
	// }
	// defer stderr.Close()

	//if err := cmd.Process.Kill(); err != nil {
	//	log.Println("failed to kill process: ", err)
	//}

	return string(retData), nil
}
