#!/bin/sh
echo "1 - Install GO enviroment? y/n"
read option
if[ $option == 'y' ] || [ $option == 'Y' ]
then
	sh setup_GO_env.sh
fi
echo "2 - Create / Re create database? y/n"
read option
if[ $option == 'y' ] || [ $option == 'Y' ]
then
	sh setup_DB.sh
else
	echo "2 - Update database? y/n"
	read option
	if[ $option == 'y' ] || [ $option == 'Y' ]
	then
		sh setup_GO_env.sh
	fi
fi
