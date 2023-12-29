all: build run

build:
	go build .

run:
	nohup ./app -port 5001 &
	nohup ./app -port 5002 &
	nohup ./app -port 5003 &
	nohup ./app -port 5004 &
	nohup ./app -port 5005 &
	sleep 1

kill:
	-killall app

show:
	curl localhost:5001/i
	curl localhost:5002/i
	curl localhost:5003/i
	curl localhost:5004/i
	curl localhost:5005/i

setups: haproxy-setup nginx-setup

tests: kill run haproxy-test nginx-test

haproxy-setup:
	cp haproxy.cfg /etc/haproxy/
	systemctl restart haproxy

nginx-setup:
	cp app.conf /etc/nginx/sites-available
	-ln -s /etc/nginx/sites-available/app.conf /etc/nginx/sites-enabled/
	systemctl restart nginx

haproxy-test:
	ab -c 350 -n 20000 localhost:6000/c

nginx-test:
	ab -c 350 -n 20000 localhost:5000/c
