# httparam2json

#### Install:

```bash
go install https://github.com/rudSarkar/httparam2json@latest
```

#### Usage:

```bash
$ httparam2json "fname=firstname&lname=lastname" | jq
{
  "fname": "firstname",
  "lname": "lastname"
}
```
