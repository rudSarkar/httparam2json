# httparam2json

#### Useages:

```bash
$ go run main.go "fname=firstname&lname=lastname" | jq
{
  "fname": "firstname",
  "lname": "lastname"
}
```
