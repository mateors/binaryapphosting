# How to add sftp user in a linux system

contabo::root:MR657GDREx

## How to create Group?
> groupadd <groupName> \
> groupadd sftponly

## How to create user?
> useradd <userName> \
> useradd automan \
> passwd automan

Enter the password
> mteorsA2021

## Add user into a group
> usermod -G <groupName> <userName>
> #usermod -g sftponly automan
> ```usermod -G sftponly automan```

## Check the user and group added or not
> cat /etc/group

## Make directory for chroot jail
> mkdir -p /home/mastererp/automan.biz/

## verify directory owner is root
> ls -ld /home/mastererp/automan.biz/

## Change owner of the directory
chown automan /home/mastererp/automan.biz/www/

## change the home directory for a user
> usermod -d /home/mastererp/automan.biz/ automan

## change the user default bash to false
> usermod -s /sbin/nologin automan \
> usermod -s /bin/false automan

## verify everything set properly
> cat /etc/passwd
  
## Reference
* https://www.tecmint.com/restrict-sftp-user-home-directories-using-chroot/
  
