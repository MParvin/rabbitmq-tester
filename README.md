# rabbitmq-tester
Test rabbitmq for functionality

copy config_example.yml to config.yml and edit it
````
RabbitHost: "amqp://USERNAME:PASSWORD@HOST:PORT"
````
for example
````
RabbitHost: "amqp://guest:guest@10.10.10.100:5672"
````

if you want to write on specific queue, you can change Queue in config.yml
````
Queue: test
````

This program will write specified message to queue and read it back, you can change this message in config.yml

build
````
go build -o rabbitmq-tester
````

install
````
go install
````

run send command to send message to queue
````
rabbitmq-tester send -c config.yml
````

and receive command to read message from queue
````
rabbitmq-tester receive -c config.yml
````

to run send and receive commands in one command
````
rabbitmq-tester -c config.yml
````