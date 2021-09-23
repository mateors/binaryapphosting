# Nginx Tips & Tricks

### Nginx X-Accel Explained
* https://blog.horejsek.com/nginx-x-accel-explained/
* https://www.nginx.com/resources/wiki/start/topics/examples/x-accel/

### Nginx -- static file serving confusion with root & alias
> There is a very important difference between the root and the alias directives. This difference exists in the way the path specified in the root or the alias is processed.

> In case of the root directive, full path is appended to the root including the location part, whereas in case of the alias directive, only the portion of the path NOT including the location part is appended to the alias.


To illustrate:

Let's say we have the config

```
location /static/ {
    root /var/www/app/static/;
    autoindex off;
}
```
In this case the final path that Nginx will derive will be

``` /var/www/app/static/static ```

This is going to return 404 since there is no static/ within static/

This is because the location part is appended to the path specified in the root. Hence, with root, the correct way is

```
location /static/ {
    root /var/www/app/;
    autoindex off;
}
```
On the other hand, with alias, the location part gets dropped. So for the config

```
location /static/ {
    alias /var/www/app/static/;
    autoindex off;           ↑
}                            |
                             pay attention to this trailing slash
```
the final path will correctly be formed as

```
/var/www/app/static
```

In a way this makes sense. The alias just let's you define a new path to represent an existing "real" path. The location part is that new path, and so it gets replaced with the real path. Think of it as a symlink.

Root, on the other hand is not a new path, it contains some information that has to be collated with some other info to make the final path. And so, the location part is used, not dropped.

## A picture is worth a thousand words
```for root:```

![root](https://i.stack.imgur.com/vCgqh.png)

```for alias:```

![alias](https://i.stack.imgur.com/Oahx0.png)

* https://stackoverflow.com/questions/10631933/nginx-static-file-serving-confusion-with-root-alias

### Serving-static-content
* https://docs.nginx.com/nginx/admin-guide/web-server/serving-static-content/
* https://docs.nginx.com/nginx/admin-guide/web-server/web-server/
*

## Nginx Virtual Host

* https://www.keycdn.com/support/nginx-virtual-host
* https://www.digitalocean.com/community/tutorials/how-to-set-up-nginx-server-blocks-virtual-hosts-on-ubuntu-16-04
* https://webautomation.io/pde/github-search-results-extractor/373/
* https://serverspace.io/support/help/nginx-virtual-hosts-on-ubuntu-20-04/

