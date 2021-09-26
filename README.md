# OS Command Execution and Gets the output

https://stackoverflow.com/questions/24095661/os-exec-sudo-command-in-go/24095983#24095983

https://stackoverflow.com/questions/66854550/golang-linux-add-an-user-with-help-of-exec

https://stackoverflow.com/questions/66854550/golang-linux-add-an-user-with-help-of-exec

> export PATH=$PATH:/usr/local/go/bin

-d /var/www/vhosts/automan.biz/

userdel -r automan

contabo::root:MR657GDREx

sudo netstat -tulpn | grep LISTEN

sudo apt install openssh-server

nc -v -z 127.0.0.1 22

ls -ld /var/www/vhosts/automan.biz/

> usermod -G \
> *usermod -g sftponly automan* \
> usermod -G sftponly automan

## How to permanently set $PATH on Linux/Unix
### You need to add it to your ~/.profile or ~/.bashrc file. 
> nano .profile
```
 export PATH = "$PATH:/usr/local/go/bin"

```
chown -R automan /var/www/vhosts/automan.biz/
### After that execute the following
> It will automatically update your path for the remainder of the session. \
> source ~/.profile

## Reference
* https://stackoverflow.com/questions/19649772/what-is-github-compare-pull-request
* https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
* https://stackoverflow.com/questions/20895552/how-can-i-read-from-standard-input-in-the-console
* https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux-unix