# pass
password generator written in Go.

```shell script
$ get  clone git@github.com:pyihe/pass.git
$ cd pass 
$ go install
```

```shell script
$ pass gen mysql #generate password for mysql and save
$ pass get #list all password 
$ pass get mysql  #get mysql password
$ pass set mysql mysqlpass #reset mysql password
$ pass del mysql  #delete mysql password
```

# Extend(TODO)

set the password to mysql or redis or mongodb, etc. Not only generate and save to file.