# Softeather VPN
> ssh root@139.59.9.243

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

> chmod 755 *

> nano vpn_config

```
CLIENT_DIR="/home/mostain/vpnclient"
NIC_NAME="NIC"
ACCOUNT_NAME="mostain"
VPN_HOST_IPv4="139.59.9.243"
LOCAL_GATEWAY="192.168.0.1"
```

## Resource
* [vpn-server-on-ubuntu-20-04](https://cloudinfrastructureservices.co.uk/how-to-install-softether-vpn-server-on-ubuntu-20-04)
* [softether systemd](https://blog.yasithab.com/ubuntu/install-softether-vpn-on-ubuntu-1604)
* [Video tutorial Creating scripts to start and stop VPN Client](https://www.youtube.com/watch?v=i2zN1IFKNYU)
* [softether-vpn-client-on-linux](https://anuradha-15.medium.com/installation-guide-of-softether-vpn-client-on-linux-54a405a0ae2c)
