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

func GetFileContent(filePath string) (*FileContent, error) { 
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
	//fmt.Println(dateTime, "=>", timee)

	//fileSize := filestat.Size()
	//fmt.Println(fileSize)
	
	return &FileContent{
		fileName:   filestat.Name(),
		FileSize:   filestat.Size(),
		ModeText:   filestat.Mode().Perm().String(),
		ModTime:    dateTime,
		Content:    string(data),
		ModeNumber: fmt.Sprintf("%04o", filestat.Mode().Perm()),
	}, nil
	
}
```

## Reference
* https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner/
* https://stackoverflow.com/questions/7151261/append-to-a-file-in-go
