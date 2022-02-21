# Softeather VPN
> ssh root@139.59.9.243

> ssh -i /home/rezaul/github root@139.59.9.243

## Server Installation
> apt-get update -y

> apt-get install build-essential gnupg2 gcc make -y

> wget http://www.softether-download.com/files/softether/v4.38-9760-rtm-2021.08.17-tree/Linux/SoftEther_VPN_Server/64bit_-_Intel_x64_or_AMD64/softether-vpnserver-v4.38-9760-rtm-2021.08.17-linux-x64-64bit.tar.gz


> tar -xvzf softether-vpnserver-v4.38-9760-rtm-2021.08.17-linux-x64-64bit.tar.gz

> cd vpnserver \
> make

> cd ..\
> mv vpnserver /usr/local/


> cd /usr/local/vpnserver/

> chmod 600 *

> chmod 700 vpnserver

> chmod 700 vpncmd

## create a systemd file
> nano /etc/init.d/vpnserver

```
#!/bin/sh
# chkconfig: 2345 99 01
# description: SoftEther VPN Server
DAEMON=/usr/local/vpnserver/vpnserver
LOCK=/var/lock/subsys/vpnserver
test -x $DAEMON || exit 0
case "$1" in
start)
$DAEMON start
touch $LOCK
;;
stop)
$DAEMON stop
rm $LOCK
;;
restart)
$DAEMON stop
sleep 3
$DAEMON start
;;
*)
echo "Usage: $0 {start|stop|restart}"
exit 1
esac
exit 0
```

> mkdir /var/lock/subsys

> chmod 755 /etc/init.d/vpnserver


> /etc/init.d/vpnserver start

> update-rc.d vpnserver defaults

## Configure vpnserver

> cd /usr/local/vpnserver \
> ./vpncmd


## UserCreate mostain

> UserPasswordSet mostain
> mateors2022

## Install softeather vpn client
> apt-get install build-essential gnupg2 gcc make -y

> wget http://www.softether-download.com/files/softether/v4.38-9760-rtm-2021.08.17-tree/Linux/SoftEther_VPN_Client/64bit_-_Intel_x64_or_AMD64/softether-vpnclient-v4.38-9760-rtm-2021.08.17-linux-x64-64bit.tar.gz

> tar -xvzf softether-vpnclient-v4.38-9760-rtm-2021.08.17-linux-x64-64bit.tar.gz

> mkdir vpnscript

> cd vpnscript
> wget https://raw.githubusercontent.com/mfaizanse/intellexlab-files/main/softether-vpn-client/remove-client.sh \
> wget https://raw.githubusercontent.com/mfaizanse/intellexlab-files/main/softether-vpn-client/setup-client.sh \
> wget https://raw.githubusercontent.com/mfaizanse/intellexlab-files/main/softether-vpn-client/vpn-connect.sh \
> wget https://raw.githubusercontent.com/mfaizanse/intellexlab-files/main/softether-vpn-client/vpn-disconnect.sh \
> wget https://raw.githubusercontent.com/mfaizanse/intellexlab-files/main/softether-vpn-client/vpn_config


```bash
sudo apt -f install
```

> go to on your vpnclient folder 

### 01
```bash
make
```
### 02

```bash
sudo vpnclient start
```
### 03
```bash
vpncmd
```
> choose 3 & check the weather Softether VPN status
```bash
Check
```
> Please Check All are `pass` or not if pass then press `ctrl+c`

> again type `vpncmd` &chosse 2 then press enter then come to show `Connected to VPN Client "localhost" 
VPN Client>`

> type here `remoteenable` press enter

> type - `niccreate` press Enter.

> now set the virtual adapter name : `nic`

> type `accountcreate` press Enter & type name of vpn connection setting : `mostain`

> then type the host name & port number : `139.59.9.243:443`

> now type Virtual hub name : `myhub`

> then type connecting user name : `mostain`

> again virtual adapter name : `nic`

> now completed successfully... Check ur information 
```bash
Destination VPN Server Host Name and Port Number: 139.59.9.243:443
Destination Virtual Hub Name: myhub
Connecting User Name: mostain
Used Virtual Network Adapter Name: VPN
Password : mateors2022
confirm password : same..
Specify standard or radius : standard

```

### 04
> check accountlist by type `accountlist`
> show offline your `Status`

### 05 
> Create Password on your account by type `accountpassword mostain`
> Password : `mateors2022` Retype this

### 06
> now specify standard or radius : `standard`

### 07
> type now `accountstartupset`

### 08
> type `accountconnect mostain`

### 09
> Check account list `accountlist`
> wait a sec & show connected....

### 10
> now uncheck the IPV4 config file `net.ipv4.ip_forward=1` & save it...

```bash
sudo nano /etc/sysctl.conf
```
### 11
> check ipv4 status
 
```bash
sudo sysctl -p
```

### 12
> now check to `ifconfig` for `vpn_nic` available or not...

### 13
> check vpn_nic ipaddress 

```bash
ip addr show vpn_nic
```
### 14
> enable dhclient ip 

```bash
sudo dhclient vpn_nic
```
> again check ip come out or not
```bash
ip addr show vpn_nic
```

### 15
> check ip neighbour
```bash
ip neigh
```
### 16
> check ip route
```bash
ip route
```

### 17
> now added ip route
```bash 
sudo ip route add 139.59.9.243/32 via 192.168.1.1 dev ens33
```
> retype 

```bash
sudo ip route del default 192.168.1.1 dev ens33
```
> check again ip ok or not
```bash
ip route
```
### 18

> login super user
```bash sudo su
```
> now type 
```bash
echo "nameserver 8.8.8.8">>"/etc/resolv.conf"
```
> type `exit`

> now `ping 8.8.8.8 -c4`
> now check neighbour ip 
```bash
ip neigh
```
> ip reachable or not if reachable then on otherwise type-
```bash
sudo dhclient vpn_nic
```
### 19

> now check your ip 
```bash
wget -qO- icanhazip.com
```
> now your vpn ip is shown & also check on webbrowser `whatismyip.com`

### Important Section - create script for vpn start & stop service


## Resource
* [vpn-server-on-ubuntu-20-04](https://cloudinfrastructureservices.co.uk/how-to-install-softether-vpn-server-on-ubuntu-20-04)
* [softether systemd](https://blog.yasithab.com/ubuntu/install-softether-vpn-on-ubuntu-1604)
* [Video tutorial Creating scripts to start and stop VPN Client](https://www.youtube.com/watch?v=i2zN1IFKNYU)
* [softether-vpn-client-on-linux](https://anuradha-15.medium.com/installation-guide-of-softether-vpn-client-on-linux-54a405a0ae2c)
* [softether-vpn-client start from Part-II](https://www.cactusvpn.com/tutorials/how-to-set-up-softether-vpn-client-on-linux/)

