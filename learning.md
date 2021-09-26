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


## Reference
* https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner/
