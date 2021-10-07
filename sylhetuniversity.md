## sylhetuniversity.com Virtual & DNS Hosting
sudo mkdir -p /var/www/sylhetuniversity.com/public_html


sudo chown -R $USER:$USER /var/www/sylhetuniversity.com/public_html

sudo chown -R nginx:nginx /var/www/sylhetuniversity.com/public_html


r = 4
w = 2
x = 1


sudo chmod -R 755 /var/www


nano /var/www/sylhetuniversity.com/public_html/index.html


#sudo cp /etc/nginx/sites-available/default /etc/nginx/sites-available/sylhetuniversity.com.conf

nano /etc/nginx/sites-available/sylhetuniversity.com.conf

```
server {

	listen 80;
	server_name sylhetuniversity.com www.sylhetuniversity.com;

	root /var/www/sylhetuniversity.com/public_html;

	# Add index.php to the list if you are using PHP
	index index.html;

	location / {
		try_files $uri $uri/ =404;
	}
}
```


sudo ln -s /etc/nginx/sites-available/sylhetuniversity.com.conf /etc/nginx/sites-enabled/


> systemctl restart nginx
> service nginx retstart


pdnsutil add-record ZONE NAME TYPE [TTL] CONTENT

ZONE = sylhetuniversity.com.
NAME = www.sylhetuniversity.com.
TYPE = A
TTL = 3600
CONTENT = "74.208.181.199"

pdnsutil create-zone sylhetuniversity.com.

pdnsutil add-record sylhetuniversity.com. www. A 3600 "74.208.181.199"

#pdnsutil add-record sylhetuniversity.com. www.sylhetuniversity.com. A "74.208.181.199"


pdnsutil check-zone sylhetuniversity.com

pdnsutil add-record sylhetuniversity.com. www. A 3600 "74.208.181.199"


19 |         3 | sylhetuniversity.com                          | SOA  | a.misconfigured.powerdns.server hostmaster.sylhetuniversity.com 1 10800 3600 604800 3600 | 3600 |    0 |        0 | NULL                             |    1 |


20 |         3 | www.sylhetuniversity.com.sylhetuniversity.com | A    | 74.208.181.199                                                                           | 3600 |    0 |        0 | NULL                             |    1 |



pdnsutil rectify-zone sylhetuniversity.com.

pdnsutil add-record ZONE NAME TYPE [TTL] CONTENT
pdnsutil add-record sylhetuniversity.com. ns1 NS 74.208.181.199

pdnsutil delete-rrset sylhetuniversity.com. ns1 NS

pdnsutil create-zone sylhetuniversity.com ns1.sylhetuniversity.com


#pdnsutil delete-zone sylhetuniversity.com
####################################################

pdnsutil create-zone sylhetuniversity.com ns1.sylhetuniversity.com

pdnsutil add-record sylhetuniversity.com ns1 A 74.208.181.199

pdnsutil add-record sylhetuniversity.com "" A 74.208.181.199

pdnsutil add-record sylhetuniversity.com www A 74.208.181.199

pdnsutil list-zone sylhetuniversity.com

groupadd sftponly
useradd sylhetuniv

passwd sylhetuniv
*****

usermod -g sftponly sylhetuniv && usermod -G sftponly sylhetuniv

mkdir -p /var/www/sylhetuniversity.com

ls -ld /var/www/sylhetuniversity.com

chown sylhetuniv /var/www/sylhetuniversity.com/public_html/

chown -R sylhetuniv /var/www/sylhetuniversity.com/public_html/

usermod -d /var/www/sylhetuniversity.com/ sylhetuniv

usermod -s /bin/false sylhetuniv

nano /etc/ssh/sshd_config

service sshd restart

## pdnsutil 
* [Tutorial of pdnstl](https://gist.github.com/ahupowerdns/38932960414f02c786f7)
