PROJS := app1 app2 app3
OUTS := ./bin/app1 ./bin/app2 ./bin/app3

all: build-deploy

.PHONY: ${PROJS}
build-deploy: bin ${PROJS}

bin:
	mkdir bin

${PROJS}:
	go build -o ./bin/ $@
	nohup ./bin/$@ &

kill:
	killall app1 app2 app3

nginx-setup: nginx-conf
	mkdir /etc/nginx/sites-available
	mkdir /etc/nginx/sites-enabled

nginx-conf:
	cp app.conf /etc/nginx/sites-available
	-ln -s /etc/nginx/sites-available/app.conf /etc/nginx/sites-enabled/
	systemctl restart nginx

	@echo Please enter following config on nginx http block:
	@echo include /etc/nginx/sites-enabled/*;
