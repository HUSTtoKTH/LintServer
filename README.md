# LintServer

Developed base on [Go-clean-template](https://evrone.com/go-clean-template?utm_source=github&utm_campaign=go-clean-template)

## Run Server

```sh
# Postgres
$ make compose-up
# Run app with migrations
$ make run
```

## Run Tests

API DOC http://localhost:8080/swagger/index.html

Admin user create new rule for project1

```sh
$ curl --location --request POST 'http://localhost:8080/v1/lint/upload' \
--header 'Token: superadmin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "organization_id": 1,
    "project_id": 1,
    "rule": "ruledetail"
}'
```

Admin user update rule for project1
```sh
$ curl --location --request POST 'http://localhost:8080/v1/lint/upload' \
--header 'Token: superadmin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "organization_id": 1,
    "project_id": 1,
    "rule": "ruleupdate"
}'
```

Non-admin user update rule for project1
```sh
$ curl --location --request POST 'http://localhost:8080/v1/lint/upload' \
--header 'Token: user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "organization_id": 1,
    "project_id": 1,
    "rule": "ruleupdate"
}'
```

Non-admin user with access to project1 get rule for project1
```sh
$ curl --location --request GET 'http://localhost:8080/v1/lint/rule/1' \
--header 'Token: user' \
--data-raw ''
```

Non-admin user without access to project2 get rule for project2
```sh
$ curl --location --request GET 'http://localhost:8080/v1/lint/rule/2' \
--header 'Token: user' \
--data-raw ''
```

