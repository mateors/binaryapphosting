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
  
> SSHD Config file ``` /etc/ssh/sshd_config```
> we need to set the following changes to the file
  
```
# override default of no subsystems
#Subsystem      sftp    /usr/libexec/openssh/sftp-server
Subsystem sftp internal-sftp

# For sftp user only
 Match Group sftponly
        PasswordAuthentication yes
        ChrootDirectory %h
        ForceCommand internal-sftp
        X11Forwarding no
        AllowTcpForwarding no
```

## Restart sshd service
> service sshd restart
  
## Reference
* https://www.tecmint.com/restrict-sftp-user-home-directories-using-chroot
* https://serverfault.com/questions/807074/enable-password-login-for-sftp-while-keeping-authentication-by-ssh-keys
* https://www.redhat.com/sysadmin/suid-sgid-sticky-bit
* https://www.youtube.com/watch?v=DDCLmAodbPc
* https://www.cyberciti.biz/faq/how-to-hide-files-and-directories-in-linux/
  
