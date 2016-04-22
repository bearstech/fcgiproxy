Fcgiroxy
========

A simple FastCGI to FastCGI proxy. You know, the stuff between your Webserver and PHP.

Fastcgi server is in golang core library. Fastcgi client cames from [Caddy](https://github.com/mholt/caddy) (double forked from https://code.google.com/p/go-fastcgi-client/ and http://bitbucket.org/PinIdea/fcgi_client).

Usage
-----

    fcgiproxy /path/to/web/root /path/to/unix_socket/for/apache /path/to/unix_socket/from/php

Licence
-------

Apache 2 Licence, © 2016 Mathieu Lecarme
