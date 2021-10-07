# Coder.cf


## DNS Configuration for coder.cf

pdnsutil create-zone coder.cf ns1.coder.cf

pdnsutil list-zone coder.cf


pdnsutil add-record coder.cf ns1 A 74.208.181.199

pdnsutil add-record coder.cf "" A "74.208.181.199"

pdnsutil add-record coder.cf "www" A "74.208.181.199"

pdnsutil check-zone coder.cf

## Step-2
> nginx virtual hosting configuration

sudo mkdir -p /var/www/coder.cf/public_html

nano /var/www/coder.cf/public_html/index.html

nano /etc/nginx/sites-available/coder.cf.conf

```
server {

	listen 80;
	server_name coder.cf www.coder.cf;

	root /var/www/coder.cf/public_html;

	# Add index.php to the list if you are using PHP
	index index.html;

	location / {
		try_files $uri $uri/ =404;
	}
}
```
