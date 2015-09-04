#!/bin/bash
echo 'ServerName 127.0.0.1:80' >> /etc/httpd/conf/httpd.conf
echo '<VirtualHost *:80>' >> /etc/httpd/conf/httpd.conf
echo 'DocumentRoot /data/code' >> /etc/httpd/conf/httpd.conf
echo 'ServerName 127.0.0.1' >> /etc/httpd/conf/httpd.conf
echo 'ErrorLog /data/log/error.log' >> /etc/httpd/conf/httpd.conf
echo 'CustomLog /data/log/access.log common' >> /etc/httpd/conf/httpd.conf
echo '</VirtualHost>' >> /etc/httpd/conf/httpd.conf
chkconfig --level 12345 httpd on
mkdir -vp /data/code
mkdir -vp /data/log
service httpd start