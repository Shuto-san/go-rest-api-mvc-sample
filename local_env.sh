#!/bin/sh

echo "export mysql_username=root" >> ~/.bash_profile
echo "export mysql_password=gomvcpass" >> ~/.bash_profile
echo "export mysql_host=127.0.0.1" >> ~/.bash_profile
echo "export mysql_dbname=gomvcdb" >> ~/.bash_profile

source ~/.bash_profile
