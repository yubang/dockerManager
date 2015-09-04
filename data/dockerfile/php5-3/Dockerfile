#打包php5.3环境

FROM centos:6

MAINTAINER yubang cybzxc@163.com

RUN yum -y install httpd
RUN yum -y install php
RUN yum -y install php-mysql
RUN yum -y install php-xml php-gd
RUN yum -y install php-mbstring
RUN yum -y install epel-release
RUN yum -y install php-mcrypt

RUN mkdir -v /var/script/
ADD script.sh /var/script/script.sh
ADD start.sh /var/script/start.sh
RUN /bin/bash /var/script/script.sh

