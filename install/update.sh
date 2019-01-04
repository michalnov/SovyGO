#! /bin/bash
while true
do
	date +"%T"
	cd go/src/github.com/michalnov/basicAPI
	git pull
	cd /var/www/html/projects/pexeso/develop
	git pull
	cd /var/www/html/projects/pexeso/release
	git pull
	echo ""
	cd
	sleep 2m
done
