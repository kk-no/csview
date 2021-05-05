# csview

csview is a CLI tool that allows you to view csv in a browser.  
This is for studying go:embed.

## run
```shell
$ make run
or
$ make build
$ ./csview -f sample.csv 
```

query sample
```
where Country = Japan
```

## TODO
- Support Update
- Support other search keyword (e.g. "AND", "OR" and "BETWEEN")