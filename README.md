Calculator still in progress

# Calculator
The application is evaluating math expressions as a strings

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Installing

### Step - 1
First you need to copy the project on your local machine:

>`git clone https://github.com/nvasilev98/calculator.git`

 ### Step - 2
 In order to get all dependencies run:

>`go mod tidy`

## Running The Tests
Execute unit tests

>`scripts/unit`

Execute integration tests

>`scripts/integration`

## Running The Application
 Execute the following command in the main directory:

 >`go run cmd/consoleapp/main.go ["Expression"]`

 ## Built With
 - Golang
 - Ginkgo & Gomega - framework for testing
 - GoMock - framework for mocking
 




