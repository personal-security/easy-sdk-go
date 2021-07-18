package easysdk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

// FileExists check is file exist
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func FileStringReplace(FilePath string, Original string, Replace string) error {
	input, err := ioutil.ReadFile(FilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	output := bytes.Replace(input, []byte(Original), []byte(Replace), -1)

	if err = ioutil.WriteFile(FilePath, output, 0666); err != nil {
		fmt.Println(err)
		return err
	}

	time.Sleep(5 * time.Second)
	return nil
}

func FileStringReplaceRegEx(FilePath string, RegEx string, Replace string) error {
	input, err := ioutil.ReadFile(FilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	re := regexp.MustCompile(RegEx)
	s := re.ReplaceAllString(string(input), Replace)

	if err = ioutil.WriteFile(FilePath, []byte(s), 0666); err != nil {
		fmt.Println(err)
		return err
	}

	time.Sleep(5 * time.Second)
	return nil
}

func FileWriteBytes(FilePath string, Data []byte) error {
	file, err := os.OpenFile(
		FilePath,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	// Write bytes to file
	bytesWritten, err := file.Write(Data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("File : %s , saved %d bytes.\n", FilePath, bytesWritten)

	time.Sleep(5 * time.Second)
	return nil
}

func FolderExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
