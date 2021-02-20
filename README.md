# Gin User Boilerplate

This repository contains a boilerplate and/or reference use of user management system like: registration, login, logout, profile update, profile delete.

Technology used in this repository:

- Backend language: Go
- Framework: [Gin](https://github.com/gin-gonic/gin) (mostly for route handling)
- Database: MySQL
- ORM: [GORM](https://github.com/go-gorm/gorm)
- Authentication: [Go-JWT](https://github.com/dgrijalva/jwt-go)
- Session: [Gin-Session](https://github.com/gin-contrib/sessions)

<br/>
<br/>

### Caution Case

To register a user, once need to send JSON as following format:

```json
{
  "first_name": "Haris",
  "last_name": "Dipto",
  "email": "haris.dipto@gmail.com",
  "password": "12345",
  "birth_day": "2020-04-23T00:00:00Z"
}
```

Special case is for `birth_day` property. Its required to send date in the above mentioned format.

<br/>
<br/>

### Future Improvements

There are lots of cases where improvement is needed. I mostly did it for my personal usage. So I think its very less likely that I'll make those improvemnts in future. 😅😅
