package main

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type FileContent struct {
	fileName   string
	FileSize   int64
	ModeText   string
	ModeNumber string
	ModTime    string
	Content    string
}

func (fc *FileContent) GetContent() string {
	return fc.Content
}

func (fc *FileContent) LastModified() string {
	return fc.ModTime
}

func (fc *FileContent) Name() string {
	return fc.fileName
}

func (fc *FileContent) Mode() string {
	return fc.ModeNumber
}

func (fc *FileContent) Size() int64 {
	return fc.FileSize
}

func main() {

	// Reference
	//(https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html)

	//step-1
	// cmd := exec.Command("ls", "-lah")
	// if runtime.GOOS == "windows" {
	// 	cmd = exec.Command("tasklist")
	// }
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	//step-2
	// cmd := exec.Command("ls", "-lah")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	//step - 2 (modified)
	// var b bytes.Buffer
	// cmd := exec.Command("ls", "-lah")
	// cmd.Stdout = &b
	// cmd.Stderr = &b
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// bs := b.Bytes()
	// fmt.Println(string(bs))

	//step-3
	// cmd := exec.Command("ls", "-lah")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// fmt.Printf("combined out:\n%s\n", string(out))

	//data, err := ioutil.ReadAll(os.Stdin)
	// data, err := io.ReadAll(os.Stdin)
	// fmt.Println(data, err)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	//Step-4: Capture stdout and stderr separately
	// bs, err := exec.Command("lscpu").Output()
	// fmt.Println(string(bs), err)

	//step-5
	// cmd := exec.Command("lscpu")
	// var stdout, stderr bytes.Buffer
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	// fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	//Step-6
	//Capture output but also show progress
	// cmd := exec.Command("sudo", "apt", "update")

	// var stdout, stderr []byte
	// var errStdout, errStderr error
	// stdoutIn, _ := cmd.StdoutPipe()
	// stderrIn, _ := cmd.StderrPipe()
	// err := cmd.Start()
	// if err != nil {
	// 	log.Fatalf("cmd.Start() failed with '%s'\n", err)
	// }

	// // cmd.Wait() should be called only after we finish reading
	// // from stdoutIn and stderrIn.
	// // wg ensures that we finish
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	// 	wg.Done()
	// }()

	// stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)

	// wg.Wait()

	// err = cmd.Wait()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// if errStdout != nil || errStderr != nil {
	// 	log.Fatal("failed to capture stdout or stderr\n")
	// }
	// outStr, errStr := string(stdout), string(stderr)
	// fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	//Step-7
	// var command string = "sudo apt update"
	// args := strings.Split(command, " ")
	// //fmt.Println(args[0], args[1:]...)
	// //cmd := exec.Command("sudo", "apt", "update")
	// cmd := exec.Command(args[0], args[1:]...)
	// if runtime.GOOS == "windows" {
	// 	cmd = exec.Command("tasklist")
	// }

	// var stdoutBuf, stderrBuf bytes.Buffer
	// cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	// cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	// fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	// fmt.Println(os.Environ(), len(os.Environ()))
	// for i, v := range os.Environ() {
	// 	fmt.Println(i, v)
	// }

	//fmt.Println(os.Getenv("USERNAME"))

	// path, err := exec.LookPath("useradd")
	// if err != nil {
	// 	fmt.Printf("Error: %s\n", err.Error())
	// } else {
	// 	fmt.Printf("'ls' executable is in '%s'\n", path)
	// }

	//useradd automan
	// res, err := commandExecute("sudo apt update")
	// fmt.Println(">>", res, err)

	//trial
	// cmd := exec.Command("cat")

	//***
	// res, err := commandExecute("sudo useradd automan")
	// fmt.Println(">>", res, err)

	// cmd := exec.Command("openssl", "passwd", "-1")
	// //cmd := exec.Command("sudo", "apt", "update")
	// var stdin, stdout, stderr bytes.Buffer
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// cmd.Stdin = &stdin

	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// fmt.Println("IN:", stdin.String())
	// fmt.Println("OUT:", stdout.String())
	// fmt.Println("ERR:", stderr.String())

	//userAdd("automan", "test123")
	// bs, err := execCommand("groupadd sftponly")
	// fmt.Println(bs, err)

	// bs, err := execCommand("usermod -G sftponly automan")
	// fmt.Println(bs, err)

	// bs, err := execCommand("mkdir -p /var/www/vhosts/automan.biz/www2")
	// fmt.Println(bs, err)

	//user info
	// cu, err := user.Current()
	// fmt.Println(cu.Name, cu.Username, cu.Uid, cu.Gid, cu.HomeDir, err)

	//os ==> user function
	// uinfo, err := user.Lookup("automan")
	// fmt.Println(uinfo.Username, uinfo.Gid, uinfo.Uid, uinfo.HomeDir, err)

	// ug, err := user.LookupGroupId(uinfo.Gid)
	// fmt.Println(ug.Name, ug.Name, err)

	// ginfo, err := user.LookupGroup("sftponly")
	// fmt.Println(ginfo.Name, ginfo.Gid, err)

	// kevin:x:1005:1006::/home/kevin:/usr/bin/zsh
	// line := "kevin:x:1005:1006::/home/kevin:/usr/bin/zsh"
	// parts := strings.SplitN(string(line), ":", 8)
	// fmt.Println(parts, len(parts))

	// username := "automan"
	// nameC := make([]byte, len(username)+1)
	// copy(nameC, username)
	// fmt.Println(nameC)

	// fmt.Println(syscall.Getuid())
	// val, isFound := syscall.Getenv("PATH")
	// fmt.Println(val, isFound)

	// err = os.Mkdir("test", os.FileMode(0775))
	// fmt.Println(err)
	// uid, _ := strconv.Atoi(uinfo.Uid)
	// gid, _ := strconv.Atoi(uinfo.Gid)
	// err = os.Chown("test", uid, gid)
	// fmt.Println(err)

	// if err := syscall.Setuid(1001); err != nil {
	// 	fmt.Printf("%v\n", err)
	// }

	// cmd := exec.Command("mkdir", "rassel")
	// cmd.SysProcAttr = &syscall.SysProcAttr{}
	// cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 0, Gid: 0}
	// bs, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("Result:", string(bs))

	//execute bash command
	// bs, err := execCommand("bash bashscript.sh")
	// ipaddr := strings.TrimSpace(string(bs))
	// fmt.Println(ipaddr, len(ipaddr), err)

	//cat /etc/passwd | cut -d: -f 1

	// bs, err := execCommand("bash -c cat /etc/passwd") //cut -d : -f 1
	// fmt.Println(string(bs), err)

	//golang command pipeline example
	timeStart := time.Now()

	// var b bytes.Buffer
	// if err := Execute(&b,
	// 	exec.Command("cat", "/etc/passwd"),
	// 	exec.Command("cut", "-d", ":", "-f", "1"),
	// 	exec.Command("grep", "rassel"),
	// ); err != nil {
	// 	log.Fatalln(err)
	// }
	// //io.Copy(os.Stdout, &b)
	// fmt.Println(b.String())

	// bs, err := execCommand("bash bashscript.sh")
	// ipaddr := strings.TrimSpace(string(bs))
	// fmt.Println(ipaddr, len(ipaddr), err)

	//err := os.Chown("/var/www/vhosts/automan.biz", 1001, 0)
	//err := chown("/var/www/vhosts/automan.biz/www/", "automan")

	// path := "/var/www/vhosts/automan.biz/www/"
	// //err := chownRecursive("mostain", path)

	// fs, err := os.Lstat(path)
	// fmt.Println(fs.Mode())
	// fmt.Printf("\n%04o\n", fs.Mode().Perm())

	//File manipulation
	fileName := "hello.txt"
	// text := "This is a new text file with one line of text"
	// err := FileCreateUpdate(fileName, text)

	content, err := GetFileContent(fileName)

	ms := time.Since(timeStart).Milliseconds()
	fmt.Println(ms, content.GetContent(),
		content.ModTime,
		content.ModeText,
		content.ModeNumber,
		err)
}

func GetFileContent(filePath string) (*FileContent, error) {

	//var fc FileContent{}

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

	//io.Copy(file,)

	filestat, err := file.Stat()
	if err != nil {
		//fmt.Println("Could not able to get the file stat")
		return nil, err
	}

	timee := filestat.ModTime()
	dateTime := timee.Format("2006-01-02 15:04:05")
	//fmt.Println(dateTime, "=>", timee)

	//fileSize := filestat.Size()
	//fmt.Println(string(data), fileSize)
	// offset := fileSize - 1
	// lastLineSize := 0
	// fmt.Println(fileSize)

	// for {
	// 	b := make([]byte, 1)
	// 	//fmt.Println(b, offset)
	// 	n, err := file.ReadAt(b, offset)
	// 	if err != nil {
	// 		//fmt.Println("Error reading file ", err)
	// 		fmt.Println("---", err.Error())
	// 		break
	// 	}
	// 	char := string(b[0])
	// 	fmt.Println(b, offset, "char:", char)
	// 	if char == "\n" {
	// 		fmt.Println(char, "break")
	// 		break
	// 	}
	// 	offset--
	// 	lastLineSize += n
	// }

	// lastLine := make([]byte, lastLineSize)
	// _, err = file.ReadAt(lastLine, offset+1)
	// if err != nil {
	// 	//fmt.Println("Could not able to read last line with offset", offset, "and lastline size", lastLineSize)
	// 	return err
	// }
	return &FileContent{
		fileName:   filestat.Name(),
		FileSize:   filestat.Size(),
		ModeText:   filestat.Mode().Perm().String(),
		ModTime:    dateTime,
		Content:    string(data),
		ModeNumber: fmt.Sprintf("%04o", filestat.Mode().Perm()),
	}, nil
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

	if _, err = file.WriteString(text); err != nil {
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
