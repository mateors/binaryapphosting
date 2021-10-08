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

    sudo mkdir /var/www/domain.com

.

    sudo chown -R $USER:$USER /var/www/domain.com

. 

    sudo nano /etc/nginx/sites-available/domain.com
.

    server {
    listen 80;
    server_name domain.com;
    root /var/www/domain.com;

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

    sudo ln -s /etc/nginx/sites-available/domain.com.conf /etc/nginx/sites-enabled/



.

    sudo nginx -t

.

    sudo systemctl reload nginx


## Step 5 –Testing PHP with Nginx

    nano /var/www/domain.com/info.php

    ```
        <?php phpinfo(); ?>
    ```
 
 
 
## Step 6 Installing phpMyAdmin

    sudo apt install phpmyadmin

.

    sudo ln -s /usr/share/phpmyadmin /var/www/domain.com/phpmyadmin



## Step 7 Changing phpMyAdmin’s Default Location
    cd /var/www/domain.com/
.

    sudo mv phpmyadmin newpath

.

    https://domain.com/newpath

## Step 8 — Disabling Root Login

    sudo nano /etc/phpmyadmin/conf.d/pma_secure.php

Past code 
        
    <?php

    # PhpMyAdmin Settings
    # This should be set to a random string of at least 32 chars
    $cfg['blowfish_secret'] = 'CHANGE_THIS_TO_A_STRING_OF_32_RANDOM_CHARACTERS';

    $i=0;
    $i++;

    $cfg['Servers'][$i]['auth_type'] = 'cookie';
    $cfg['Servers'][$i]['AllowNoPassword'] = false;
    $cfg['Servers'][$i]['AllowRoot'] = false;

    ?>


## Step 9 — Creating an Authentication Gateway

    openssl passwd

    sudo nano /etc/nginx/pma_pass

Past

    user:9YHV.p60.Cg6I

Open Nginx configuration file

    sudo nano /etc/nginx/sites-available/domain.com
.

    server {
            . . .

            location / {
                    try_files $uri $uri/ =404;
            }

            location ^~ /hiddenlink/ {
                    auth_basic "Admin Login";
                    auth_basic_user_file /etc/nginx/pma_pass;

                    location ~ \.php$ {
                            include snippets/fastcgi-php.conf;
                            fastcgi_pass unix:/var/run/php/php7.4-fpm.sock;
                    }
            }
        . . .
    }


```
sudo nginx -t

```
```
sudo systemctl reload nginx
```

```
https://domain.com/newpath
```

## Step 10 — Setting Up Access via Encrypted Tunnels

### Setting Up IP-Based Access Control on Nginx

```
allow hostname_or_IP;
deny all;
```

```
sudo nano /etc/nginx/sites-available/domain.com
```

```
server {
        . . .

        location ^~ /hiddenlink/ {
                satisfy all; #requires both conditions

                allow 203.0.113.0; #allow your IP
                allow 127.0.0.1; #allow localhost via SSH tunnels
                deny all; #deny all other sources

                auth_basic "Admin Login";
                auth_basic_user_file /etc/nginx/pma_pass;

                location ~ \.php$ {
                        include snippets/fastcgi-php.conf;
                        fastcgi_pass unix:/var/run/php/php7.4-fpm.sock;
                }
        }

        . . .
```

```
sudo nginx -t
```

```
sudo systemctl reload nginx
```
```
ssh user@server_domain_or_IP -L 8000:localhost:80 -L 8443:localhost:443 -N
```


- user: the Ubuntu user profile to connect to on the server where phpMyAdmin is running

- -L 8000:localhost:80 redirects HTTP traffic on port 8000

- -L 8443:localhost:443 redirects HTTPS traffic on port 8443

- -N: prevents the execution of remote commands 


```
http://localhost:8000/newpath

https://localhost:8443/newpath
```