# Golang Performace Tracker


## File Read Challange
* https://marcellanz.com/post/file-read-challenge/
* [Source code](https://github.com/marcellanz/file-read-challenge-go/blob/master/rev4/readfile4.go)

## Append to a file in golang
* https://stackoverflow.com/questions/7151261/append-to-a-file-in-go

```golang
f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
if err != nil {
    panic(err)
}

defer f.Close()

if _, err = f.WriteString(text); err != nil {
    panic(err)
}

```

## Effective Go
* https://golang.org/doc/effective_go

> go doc -all

```golang

type FileContent struct{

	FileName string
	Size int
	ModeText string
	ModeNumber string
	ModTime string
	Content string

}

func GetFileContent(filePath string) (string, error) { 
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	//os.File
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	filestat, err := file.Stat()
	if err != nil {
		return "", err
	}

	timee := filestat.ModTime()
	dateTime := timee.Format("2006-01-02 15:04:05")
	fmt.Println(dateTime, "=>", timee)

	fileSize := filestat.Size()
	fmt.Println(fileSize)
	return string(data),nil
	
}
```

## Reference
* https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner/
* https://stackoverflow.com/questions/7151261/append-to-a-file-in-go
