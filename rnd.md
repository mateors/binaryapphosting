# Research

YunoHost is server operating system 
soser
hoster

serveroperatingsystem.com
serveroperating.com
serveros.net
serveros.org

masterhost24.com
mastrhost.com

mpanel.org

**mastr.org**
golanghost.org

*host.pro*
mstrhost

easyhostingpanel.com
hostingpanelservice.com
hosting-master.net

### hostemaster.com
hostmastr.com

**mastrhost.com**
Host.Pro
HOST.PRO

yourhostshop.com
masterpanel.co

servicepanel.org
easywebhoster.com

### codewithmosh.activehosted.com
* [A Nice UI/UX](https://codewithmosh.activehosted.com/admin)

## What is reverse proxy?
A reverse proxy is a server that sits in front of web servers and forwards client (e.g. web browser) requests to those web servers. Reverse proxies are typically implemented to help increase security, performance, and reliability. In order to better understand how a reverse proxy works and the benefits it can provide, let’s first define what a proxy server is.

## What is a proxy server?
> A forward proxy, often called a proxy, proxy server, or web proxy, is a server that sits in front of a group of client machines. When those computers make requests to sites and services on the Internet, the proxy server intercepts those requests and then communicates with web servers on behalf of those clients, like a middleman.

![Proxy Server](https://www.cloudflare.com/img/learning/cdn/glossary/reverse-proxy/forward-proxy-flow.svg)

### There are a few reasons one might want to use a forward proxy:
* To avoid state or institutional browsing restrictions
* To block access to certain content
* To protect our identity online

## How is a reverse proxy different?
A reverse proxy is a server that sits in front of one or more web servers, intercepting requests from clients. This is different from a forward proxy, where the proxy sits in front of the clients. With a reverse proxy, when clients send requests to the origin server of a website, those requests are intercepted at the network edge by the reverse proxy server. The reverse proxy server will then send requests to and receive responses from the origin server.

The difference between a forward and reverse proxy is subtle but important. A simplified way to sum it up would be to say that a forward proxy sits in front of a client and ensures that no origin server ever communicates directly with that specific client. On the other hand, a reverse proxy sits in front of an origin server and ensures that no client ever communicates directly with that origin server.

![Reverse_Proxy](https://www.cloudflare.com/img/learning/cdn/glossary/reverse-proxy/reverse-proxy-flow.svg)

## Some of the benefits of a reverse proxy:
* Load balancing
* Protection from attacks 
* Global Server Load Balancing (GSLB)- CDN
* Caching
* Freeing up valuable resources on the origin server (SSL encryption)

## How to implement reverse proxy by our own?
* https://hackernoon.com/writing-a-reverse-proxy-in-just-one-line-with-go-c1edfa78c84b
* https://developer20.com/writing-proxy-in-go/
* https://blog.joshsoftware.com/2021/05/25/simple-and-powerful-reverseproxy-in-go/


## Reference
* https://www.cloudflare.com/learning/cdn/glossary/reverse-proxy/


Cloudflare customer's hosting provider, acting as a reverse proxy for websites.
