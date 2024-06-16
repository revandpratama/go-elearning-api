# E-learning API

Simple API using repository pattern and GIN http library.


## Features
- User registration and login
- Authentication and Authorization with JWT library
- KRS and Score creation, update, and deletion

## Installation

### Packages needed

- GIN

```bash
go get -u github.com/gin-gonic/gin
```

- Gorm

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

- Viper

```bash
go get github.com/spf13/viper
```

- JWT

```bash
go get -u github.com/golang-jwt/jwt/v5
```

## Usage

Run the app
```bash
go run main.go
```

## Routes

### Auth

| Method | Action         |      URL      |
| ------ | -------------- | :-----------: |
| POST    | Login          |  /api/login   |
| POST   | Register User  | /api/register |
| POST   | Delete Session |  /api/logout  |
### User

| Method | Action         |      URL      |
| ------ | -------------- | :-----------: |
| GET    | Get KRS by userID          |  /api/users/{userID}/krs   |
| GET    | Get Score by userID          |  /api/users/{userID}/score  |

### KRS

| Method | Action         |      URL      |
| ------ | -------------- | :-----------: |
| GET    | Get all KRS         |  /api/krs   |
| GET   | Get KRS by ID  | /api/krs/{krsID} |
| POST   | Create KRS  | /api/krs |
| PUT   | Update KRS |  /api/krs/{krsID}  |
| DELETE | Delete KRS |  /api/krs/{krsID}  |
### Score

| Method | Action         |      URL      |
| ------ | -------------- | :-----------: |
| GET    | Get all Score         |  /api/score   |
| GET   | Get Score by ID  | /api/score/{scoreID} |
| POST   | Create Score  | /api/score |
| PUT   | Update Score |  /api/score/{scoreID}  |
| DELETE | Delete Score |  /api/score/{scoreID}  |





## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.


## License

[MIT](https://choosealicense.com/licenses/mit/)
