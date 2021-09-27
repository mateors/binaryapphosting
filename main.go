package main

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

type FileContent struct {
	Content  string
	FileInfo fs.FileInfo
}

func (fc *FileContent) LastModified() string {

	filestat := fc.FileInfo
	timee := filestat.ModTime()
	dateTime := timee.Format("2006-01-02 15:04:05")
	return dateTime
}

func (fc *FileContent) ModeN() string {

	filestat := fc.FileInfo
	modeNumeric := fmt.Sprintf("%04o", filestat.Mode().Perm())
	return modeNumeric
}

func (fc *FileContent) ModeT() string {

	filestat := fc.FileInfo
	return filestat.Mode().String()
}

func main() {

	//timeStart := time.Now()

	//File manipulation
	// text := "This is a new text file with one line of text"
	// err := FileCreateUpdate(fileName, text)

	fileName := "hello.txt"
	content, err := GetFileContent(fileName)
	fmt.Println(err)
	fmt.Println(content.Content)
	fmt.Println(content.FileInfo.Size())
	fmt.Println(content.ModeT())
	fmt.Println(content.ModeN())
	fmt.Println(content.LastModified())

	// ms := time.Since(timeStart).Milliseconds()
	// fmt.Println(ms, content.GetContent(),
	// 	content.ModTime,
	// 	content.ModeText,
	// 	content.ModeNumber,
	// 	err)
}

func GetFileContent(filePath string) (*FileContent, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//os.File
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return &FileContent{Content: string(data), FileInfo: fileInfo}, nil
}

//FileCreate function will create file if not exist,
//Otherwise it will append to the end of the file
func FileCreateUpdate(fileName, text string) error {

	//file, err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString("\n" + text); err != nil {
		return err
	}
	return nil
}

func chownRecursive(username, path string) error {

	if !checkIfRootLogin() {
		return errors.New("only root user can execute")
	}
	cmd := fmt.Sprintf("chown -R %s:%s %s", username, username, path)
	return ExecCommand(cmd)
}

func checkIfRootLogin() bool {
	cu, _ := user.Current()
	uid, _ := strconv.Atoi(cu.Uid)
	return uid == 0
}

func chown(path, username string) error {

	if !checkIfRootLogin() {
		return errors.New("only root user allowed")
	}

	user, err := user.Lookup(username)
	if err != nil {
		return err
	}
	userID, _ := strconv.Atoi(user.Uid)
	groupID, _ := strconv.Atoi(user.Gid)

	err = os.Chown(path, userID, groupID) //-1
	if err != nil {
		return err
	}
	return nil
}

func userAdd(username, txtpassword string) {

	cmd := exec.Command("openssl", "passwd", "-1", txtpassword)
	passwordBytes, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	// remove whitespace (possibly a trailing newline)
	password := strings.TrimSpace(string(passwordBytes))
	fmt.Println("pass:>", password, len(password))
	cmd = exec.Command("useradd", "-p", password, "-s", "/bin/false", "-d", "/var/www/vhosts/automan.biz/", "-R", "/var/www/vhosts/automan.biz/", username)
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", b)

}

func CreateRandom(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	fmt.Println(b)
	return string(b)
}

//ExecCommand send command and receive error
func ExecCommand(command string) error {

	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		//return []byte{}, err
		return err
	}
	//response := string(bs)
	return nil
}

//commandExecute
func commandExecute(command string) (string, error) {

	//var command string = "sudo apt update"
	args := strings.Split(command, " ")
	//fmt.Println(runtime.GOOS)
	cmd := exec.Command(args[0], args[1:]...)
	if runtime.GOOS != "linux" {
		return "", errors.New("sorry only linux os supported")
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	//fmt.Println("**", stdoutBuf.String())

	cmd.StdinPipe()

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	//outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	outStr, errStr := stdoutBuf.String(), stderrBuf.String()
	if errStr != "" {
		return outStr, errors.New(errStr)
	}

	//fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	return outStr, nil
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func Execute(output_buffer *bytes.Buffer, stack ...*exec.Cmd) (err error) {
	var error_buffer bytes.Buffer
	pipe_stack := make([]*io.PipeWriter, len(stack)-1)
	i := 0
	for ; i < len(stack)-1; i++ {
		stdin_pipe, stdout_pipe := io.Pipe()
		stack[i].Stdout = stdout_pipe
		stack[i].Stderr = &error_buffer
		stack[i+1].Stdin = stdin_pipe
		pipe_stack[i] = stdout_pipe
	}
	stack[i].Stdout = output_buffer
	stack[i].Stderr = &error_buffer

	if err := call(stack, pipe_stack); err != nil {
		log.Fatalln(string(error_buffer.Bytes()), err)
	}
	return err
}

func call(stack []*exec.Cmd, pipes []*io.PipeWriter) (err error) {
	if stack[0].Process == nil {
		if err = stack[0].Start(); err != nil {
			return err
		}
	}
	if len(stack) > 1 {
		if err = stack[1].Start(); err != nil {
			return err
		}
		defer func() {
			if err == nil {
				pipes[0].Close()
				err = call(stack[1:], pipes[1:])
			}
		}()
	}
	return stack[0].Wait()
}
