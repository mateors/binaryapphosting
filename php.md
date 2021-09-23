# PHP configuration tips & Tricks
Can I add a second php.ini file that supercedes the first?

> PHP loads additional .ini files from a certain directory. On a Debian 8 box, that directory is: /etc/php5/apache2/conf.d.

> Files in that directory are loaded in alphabetical order, so naming a file 00-overrides.ini (for example) will cause it to be loaded first. Settings set in that file will override settings set in the default php.ini.

Nginx uses PHP FPM and php. ini is usually located in /etc/php/7.4/fpm/php

By default, in Ubuntu 14.04 every php-fpm pool should be configured in a file inside the directory /etc/php5/fpm/pool.d. Every file with the extensions .conf in this directory is automatically loaded in the php-fpm global configuration.

So for our new site let’s create a new file /etc/php5/fpm/pool.d/site1.conf. You can do this with your favorite editor like this:

sudo vim /etc/php5/fpm/pool.d/site1.conf

* https://superuser.com/questions/911420/can-i-add-a-second-php-ini-file-that-supercedes-the-first
* https://devanswers.co/ubuntu-php-php-ini-configuration-file/
* https://www.digitalocean.com/community/tutorials/how-to-host-multiple-websites-securely-with-nginx-and-php-fpm-on-ubuntu-14-04
