# Linux Users, Groups file permission etc analysis

stat /etc/passwd

stat -c '%a' /etc/passwd

 stat -c '%A %a %n' /etc/passwd


at first you have to create linux system user
using system user you have to add/make virtual host folder

All directory permission is 0755
stat /var/www/vhosts/

stat /var/www/vhosts/mateors.com/
0710
psacln

cat /etc/group

psaserv:x:1002:psaadm,psaftp,www-data,nginx

stat logs
0700


$ compgen -g

root
daemon
bin
sys
adm
tty
disk
lp
mail
news
uucp
man
proxy
kmem
dialout
fax
voice
cdrom
floppy
tape
sudo
audio
dip
www-data *
backup
operator
list
irc
src
gnats
shadow
utmp
video
sasl
plugdev
staff
games
users
nogroup
systemd-journal
systemd-timesync
systemd-network
systemd-resolve
systemd-bus-proxy
input
crontab
syslog
netdev
lxd
messagebus
uuidd
mlocate
ssh
ntp
mysql
psaadm
swkey-data
popuser
psaserv
psacln
sw-cp-server
lock-manager
ssl-cert
postfix
postdrop
dovecot
dovenull
roundcube_sysgroup
horde_sysgroup
nginx
bind
psaftp
psasb
docker
drweb
sambashare
debian-spamd
tomcat8
magicspam
tortix
plesksendmail

Display group memberships for each Linux user:
groups {USERNAME}
groups mateors
mateors : psacln

Linux List all members of a group using members command

members {GROUPNAME}

grep -i --color '^root' /etc/group

 grep -i --color '^psaserv' /etc/group
psaserv:x:1002:psaadm,psaftp,www-data,nginx

https://devconnected.com/how-to-list-users-and-groups-on-linux/

cat /etc/passwd
hostparkbd:x:10031:1003::/var/www/vhosts/hostparkbd.net:/bin/false


cat /etc/shadow


https://serverfault.com/questions/355292/show-all-users-and-their-groups-vice-versa

List Usernames using the /etc/passwd fil
 cat /etc/passwd | cut -d: -f1

List users and their groups
for user in $(awk -F: '{print $1}' /etc/passwd); do groups $user; done


List groups and their users:
cat /etc/group | awk -F: '{print $1, $3, $4}' | while read group gid members; do
    members=$members,$(awk -F: "\$4 == $gid {print \",\" \$1}" /etc/passwd);
    echo "$group: $members" | sed 's/,,*/ /g';
done

1Am2_1lo

https://www.cyberciti.biz/faq/unix-linux-check-if-port-is-in-use-command/

sudo lsof -i:22 ## see a specific port such as 22 ##

systemctl --type=service --state=active
systemctl --type=service --state=running


service --status-all

  mailman
 [ + ]  mdadm
 [ - ]  mdadm-waitidle
 [ - ]  mountall-bootclean.sh
 [ - ]  mountall.sh
 [ - ]  mountdevsubfs.sh
 [ - ]  mountkernfs.sh
 [ - ]  mountnfs-bootclean.sh
 [ - ]  mountnfs.sh
 [ + ]  mysql
 [ + ]  networking
 [ + ]  nginx

 [ + ]  php7.0-fpm
 [ + ]  plesk-php54-fpm
 [ + ]  plesk-php55-fpm
 [ + ]  plesk-php56-fpm
 [ - ]  plesk-php70-fpm
 [ - ]  plesk-php71-fpm
 [ + ]  plesk-php72-fpm


https://www.cyberciti.biz/faq/check-running-services-in-rhel-redhat-fedora-centoslinux/

List service and their open ports
netstat -tulpn

To view processes associated with a particular service (cgroup)
systemd-cgtop


https://superuser.com/questions/911420/can-i-add-a-second-php-ini-file-that-supercedes-the-first

Can I add a second php.ini file that supercedes the first?

https://stackoverflow.com/questions/39669204/how-to-overwrite-some-php-defaults-configuration

6

Yeah should be possible. First, create a PHP file somewhere in your document root called info.php and put <?php phpinfo() ?> in it. Then call it from the browser and look for Scan this dir for additional .ini files. On Ubuntu it is probably something like /etc/php5/apache2/conf.d.

In that directory, you can put your own custom ini file (call it something like 99.overrides.php.ini) so it gets parsed last.

In that file, put whatever additional config you want, restart Apache or the web server and the changes will take effect.

Nginx -- static file serving confusion with root & alias

https://stackoverflow.com/questions/10631933/nginx-static-file-serving-confusion-with-root-alias

There is a very important difference between the root and the alias directives. This difference exists in the way the path specified in the root or the alias is processed.

In case of the root directive, full path is appended to the root including the location part, whereas in case of the alias directive, only the portion of the path NOT including the location part is appended to the alias.

## how-do-i-save-terminal-output-to-a-file?
* https://askubuntu.com/questions/420981/how-do-i-save-terminal-output-to-a-file

> $SHELL --version

### tee is useful if you want to be able to capture command output while also viewing it live. 
> command |& tee ~/outputfile.txt \
> ls -la |& tee ~/outputfile.txt \
