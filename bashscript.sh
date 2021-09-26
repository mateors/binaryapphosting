#! /usr/bin/bash

#echo $(ifconfig | grep broadcast | awk '{print $2}')

echo "$(cat /etc/passwd | cut -d : -f 1)"