# RabbitMQ Tester

This is a simple tool for testing RabbitMQ functionality. It allows you to send a message to a specified queue and then read it back, which is useful for verifying the functionality of your RabbitMQ setup.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go programming language
- RabbitMQ server

### Installation

1. Clone the repository:
```
git clone https://github.com/MParvin/rabbitmq-tester.git
```
2. Navigate to the project directory:
```
cd rabbitmq-tester
```
3. Build the program:
```
go build -o rabbitmq-tester
```
4. Install the program:
```
go install
```

## Configuration

1. Copy `config_example.yml` to `config.yml` and edit it.
2. Specify the RabbitMQ host in the format `amqp://USERNAME:PASSWORD@HOST:PORT`.
3. If you want to write on a specific queue, you can change the Queue in `config.yml`.
4. This program will write a specified message to the queue and read it back. You can change this message in `config.yml`.

## Usage

To send a message to the queue:
```
rabbitmq-tester send -c config.yml

```

To read a message from the queue:

```
rabbitmq-tester receive -c config.yml
```

To run send and receive commands in one command:
```
rabbitmq-tester -c config.yml
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
