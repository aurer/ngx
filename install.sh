#!/bin/bash

if [ -d /etc/nginx ]; then
	cp ./vhost-example.conf /etc/nginx/
	echo "Copied vhost-example.conf into /etc/nginx/"

	for dir in {$HOME/bin,/usr/local/bin,/usr/local/sbin}; do
		if [ -d $dir ]; then
			cp ./ngx $dir/
			echo "Copied ngx script into $dir"
			break
		fi
	done
	unset dir
fi
