## PowerDNS Utility Program

* [DNS Record format](https://registry.terraform.io/providers/pan-net/powerdns/latest/docs)

# pdnsutil --version
> pdnsutil 4.2.1

# pdnsutil list-algorithms
5 - RSASHA1
7 - RSASHA1-NSEC3-SHA1
8 - RSASHA256
10 - RSASHA512
13 - ECDSAP256SHA256
14 - ECDSAP384SHA384
15 - ED25519
16 - ED448

* https://manpages.ubuntu.com/manpages/groovy/man1/pdnsutil.1.html
> pdnsutil add-zone-key example.nl KSK inactive ecdsa256

> pdnsutil list-keys mostain.net

https://www.sidn.nl/en/dnssec/dnssec-on-the-powerdns-authoritative-server

pdnsutil add-zone-key mostain.net KSK inactive ecdsa256

pdnsutil add-zone-key mostain.net ZSK inactive ecdsa256


pdnsutil list-keys mostain.net

pdnsutil show-zone mostain.net

#set-nsec3 ZONE ['HASH-ALGORITHM FLAGS ITERATIONS SALT']
pdnsutil set-nsec3 mostain.net '1 0 1 ab'

pdnsutil rectify-zone mostain.net

pdnsutil set-publish-cdnskey mostain.net
pdnsutil set-publish-cds mostain.net

pdnsutil activate-zone-key mostain.net 1
pdnsutil activate-zone-key mostain.net 2

pdnsutil check-zone mostain.net

pdnsutil show-zone mostain.net

pdnsutil rectify-zone mostain.net

pdnsutil list-zone mostain.net

pdnsutil rectify-all-zones


Rectifying mostain.net: Adding NSEC3 hashed ordering information for 'mostain.net'
Rectifying zerowastepac.com: Adding empty non-terminals for non-DNSSEC zone
Rectifying zerowastepac.com: Adding empty non-terminals for non-DNSSEC zone

pdnsutil create-zone mostain.net


pdnsutil add-record example.net _autodiscover._tcp srv "100 443 mail.example.com"

#add-record ZONE NAME TYPE [TTL] CONTENT

pdnsutil add-record mostain.com ns1.mostain.com. A 74.208.181.199

clear-zone ZONE

## Resources
* https://www.sidn.nl/en/dnssec/dnssec-on-the-powerdns-authoritative-server
* https://doc.powerdns.com/md/manpages
* https://bind9.readthedocs.io/en/latest/dnssec-guide.html#advanced-discussions-key-generation
* https://registry.terraform.io/providers/pan-net/powerdns/latest/docs/resources/record
* https://doc.powerdns.com/md/manpages/pdnsutil.1/
* https://maddy.email/tutorials/setting-up/
* https://www.howson.pro/dnssec-and-power-dns/
