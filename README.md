<<<<<<< HEAD
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


## Linux Pipe in golang
* https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go

## Reference
* https://stackoverflow.com/questions/19649772/what-is-github-compare-pull-request
* https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
* https://stackoverflow.com/questions/20895552/how-can-i-read-from-standard-input-in-the-console
* https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux-unix
* https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go
* https://stackoverflow.com/questions/41932480/file-permission-mask-in-go
=======
# Services we are going to implement

## PowerDNS

## Dynamic DNS
[How to build Dynamic DNS](http://mkaczanowski.com/golang-build-dynamic-dns-service-go/)

## Maddy Emailserver

## PostFix, DoveCot, Mysql, Spamassassin

* https://www.cyberpunk.rs/mail-server-setup-postfix
* https://www.linode.com/docs/guides/email-with-postfix-dovecot-and-mysql/
* https://ixnfo.com/en/mail-server-postfix-dovecot-mysql.html
* https://www.howtoforge.com/greylisting_postfix_postgrey
* [DigitalOcean Email Setup Guide](https://www.digitalocean.com/community/tutorials/how-to-configure-a-mail-server-using-postfix-dovecot-mysql-and-spamassassin)

## Postgrey [Greylisting](https://github.com/schweikert/postgrey)
> Greylisting in short means that when someone wants to deliver a mail to your mail server it will simply reply -Please come back later-?. That is something all RFC-compliant mail servers do and when they do come back the mail is accepted. Most spammers and spam software are not compliant and not patient enough to try again. You will be surprised to see how effective this is. 
> https://www.vpnreactor.com/greylisting-org-acquistion/

## Email Client
* [Server Maddy](https://brianlovin.com/hn/27557542)
* [Client MailPile](https://www.mailpile.is)

## SFTP chroot

## Reverse Proxy
> Expose a local server behind a NAT or firewall to the internet.
* [FRP](https://github.com/mateors/frp)

## VPN Service
[Pangolin](https://github.com/xitongsys/pangolin)

## Parentral Control DNS
[PcDNS](https://github.com/meggarr/pcdns)

## Ad Blocker 
[Blocky](https://github.com/0xERR0R/blocky)

## WebRTC
 > Distributed real-time communication system, the goal is to chat anydevice, anytime, anywhere
 
* [PION](https://github.com/pion/webrtc)
* [ION](https://github.com/pion/ion)
* [Learn WebRTC](https://webrtcforthecurious.com)
>>>>>>> origin/main
