# Installing phpMyAdmin on Nginx

## Step 1 – Installing the Nginx Web Server


    sudo apt update

    sudo apt install nginx

Type the address that you receive in your web browser and it will take you to Nginx’s default landing page:

http://localhost


## Step 2 — Installing MySQL

    sudo apt install mysql-server

### mysql password setup
    sudo mysql_secure_installation


## Step 3 – Installing PHP

    sudo apt install php-fpm php-mysql


Create the root web directory for domain as follows:

    sudo mkdir /var/www/html

.

    sudo chown -R $USER:$USER /var/www/html

. 

    sudo nano /etc/nginx/sites-available/html
.

    server {
    listen 80;
    server_name html;
    root /var/www/html;

    index index.php index.html index.htm index.php;

    location / {
        try_files $uri $uri/ =404;
    }

    location ~ \.php$ {
        include snippets/fastcgi-php.conf;
        fastcgi_pass unix:/var/run/php/php7.4-fpm.sock;
     }

    location ~ /\.ht {
        deny all;
    }

}

. 

    sudo ln -s /etc/nginx/sites-available/html /etc/nginx/sites-enabled/

.

    sudo ln -s /etc/nginx/sites-available/your_domain /etc/nginx/sites-enabled/


.

    sudo nginx -t

.

    sudo systemctl reload nginx


## Step 5 –Testing PHP with Nginx

    nano /var/www/your_domain/info.php
 
 
 
## Step 6 Installing phpMyAdmin

    sudo apt install phpmyadmin

.

    sudo ln -s /usr/share/phpmyadmin /var/www/html/phpmyadmin

