Fcgiroxy
========

A simple FastCGI to FastCGI proxy. You know, the stuff between your Webserver and PHP.

Fastcgi server is in golang core library. Fastcgi client cames from [Caddy](https://github.com/mholt/caddy) (double forked from https://code.google.com/p/go-fastcgi-client/ and http://bitbucket.org/PinIdea/fcgi_client).

Usage
-----

There is an autodoc

    $ fcgiproxy -h
    Usage of fcgiproxy:
      -listen string
            What Apache or Nginx see (default "/var/run/fcgiproxy.sock")
      -log string
            Log path (default "/var/log/fcgiproxy.log")
      -root string
            Where the PHP files live (default "/var/www")
      -target string
            Where is php-fpm (default "/var/run/php5-fpm.sock")

You can test with something like that :

    fcgiproxy -root /path/to/web/root -listen /path/to/unix_socket/for/apache -target /path/to/unix_socket/from/php -log /tmp/fcgiproxy.log

Licence
-------

Apache 2 Licence, © 2016 Mathieu Lecarme
