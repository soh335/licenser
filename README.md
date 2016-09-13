# licenser

my gathering and concatenating license file tool

## INSTALL

```
$ go get github.com/soh335/licenser
```

## USAGE

```
$ licenser --carthage --dir /path/to/lib1 --dir /path/to/lib2 > LICENSE.html
```

### CONFIG

Default config file name is `.licenser.json`. And such format.

```json
{
  "lib1": "LIB1",
  "lib2": "LIB2"
}
```

It is used for section name of library. If not exist, it is directory name.

## TODO

* support cocoapods

## LICENSE

MIT
