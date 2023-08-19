# httparam2json

#### Install:

```bash
go install https://github.com/rudSarkar/httparam2json@latest
```

#### Useages:

```bash
$ go run main.go "fname=firstname&lname=lastname" | jq
{
  "fname": "firstname",
  "lname": "lastname"
}
```
