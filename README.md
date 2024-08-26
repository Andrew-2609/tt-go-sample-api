# 🧢 TT GoLang Sample API 🧢

# Overall Idea and Explanations

This project was built to give a general idea of the concepts I've learnt with GoLang. Although the overall idea is kinda silly (especially an SQS queue to "require" employees from the HR team), what matters most is the project architecture and scalability.

I've implemented unit and integration tests that cover most of the API, except the module I implemented last night (2024-08-26), because it was almost 11PM and my wife was almost sleeping and wouldn't allow me to work any further 😅.

# 🥷 API

This project is an API for registering and listing employees. Yes, it's that simple, and it only has four routes (including the `GET /health` that may be used for a Kubernetes health check, for example). But the way was paved to extend the application's routes easily and rapidly, and the fact that there's a whole testing structure would facilitate and foster the development of unit and integration tests for whatever new route that may be implemented.

## Routes

⚠️ This section exists because I didn't get the time to implement a proper OpenAPI specification for the project.

* `GET /health` - returns the application's health status
* `GET /api/v1/employees` - returns a paginated list of employees. I only implemented the basic logic of pagination, so there's no filtering of any kind
* `POST /api/v1/employees` - registers a new employee
* `POST /api/v1/employees/hr` - requires a new employee from the "HR" team. I should've used the word "solicitates" or "requests" instead, but I was sleepy by the time I implemented this

## Web Framework

I used [fiber](https://github.com/gofiber/fiber) as this project's web framework, but I'm more used with [gin](https://github.com/gin-gonic/gin). The implementation doesn't change that much, and in fact I suppose I can learn any web framework for GoLang pretty easily, since the standard `http` library of GoLang is already fantastic and is the base of it all.

# 🛢️ Database

This project utilizes PostgreSQL as its relational database, making use of the libraries [migrate](https://github.com/golang-migrate/migrate) and [sqlc](https://sqlc.dev/) to create and run migrations, and to generate SQL code based on SQL queries with sqlc keywords, respectively.

# 🛠️ How to run the project

## Environment

After cloning this repository, you'll have to **create a `.env` file on the root directory** to put the local environment variables. You can get the values from `.env.example`. They should work as they are, but know that if you wanna change them, you'll have to adapt them wherever needed!

## Commands

After this, you can simply run the following commands to have the following outcomes:

* `make run` - will run the `docker-compose.yml` file and start the application on a `scratch` container. Then, you'll be able to send requests to the API

* `make coverage` - will run the tests of the application and generate a coverage output file. It should be 61.3% by the time I'm typing this

---

# 💡 Ideas that I didn't have the time to implement

## Dependency Injection

I wanted to use [wire](https://github.com/google/wire) to inject the dependencies between repositories -> use cases -> web handles, but I didn't have the time to implemente it, alghouth I know how to do it.

So you'll notice that I simply called the methods to create the dependencies for the web handlers directly on `server/route.go`.

## Cloud Environment Variables

I wanted to use [AWS Secrets Manager](https://aws.amazon.com/pt/secrets-manager) on [localstack](https://github.com/localstack/localstack) to enable the environment variables loading simulating a Cloud environment. I've used localstack to simulate SQS queues, SNS topics, S3 buckets, DynamoDB tables, etc., but never to load secrets from a Secrets Manager, so it would take me some precious time to both learn it and implemented on the last weekend.

## SonarQube

I wanted to implement a conteinerized SonarQube for this API, but I had a hard time configuring this bloody thing locally many times, so I didn't have the patience nor the time to try to implement it during this weekend.

Just know that it was a nice goal to achieve, and that in a real world project there would be at least a `sonar-project.properties` file to enable SonarQube to capture coverage data during a pipeline, for example.