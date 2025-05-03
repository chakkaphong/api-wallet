# 🚀 Project api-wallet (Go + Clean Architecture)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)



## 📃 Description

This Go project is designed for testing and evaluating. The project serves as a hands-on exam to assess problem-solving skills and the ability to write clean.

## 🔦 Key Focus Areas

- Clean Architecture
- Core language features: structs, slices, maps, and interfaces
- REST API


## ⚠️ Requirements

- go >= 1.23.0
- migrate >= 4.18.3


## 📦 Installation

First, install [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### MacOS

```bash
$ brew install golang-migrate
```

### Windows

Using [scoop](https://scoop.sh/)

```bash
$ scoop install migrate
```


## 🎉 Running the API
### Development
To start the application in development mode please follow this step. Make sure you have installed `migrate`


Assume your have pre-installed [docker](https://www.docker.com/)
```bash
docker-compose -f docker-compose.dev.yml up -d
```

#### Set up environment
In configs folder you can create new file `.env`. you can use example config from this file `.env.example` or change it for your purposes. make sure your have running `docker-compose` from above or you can use your own database.
for the example

```
POSTGRES_HOST= 
POSTGRES_PORT= 
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
```

#### Start the application

```bash
make run
```

api server listening on `http://localhost:9001` 🎊🎊🎊