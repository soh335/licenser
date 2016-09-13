[![wercker status](https://app.wercker.com/status/820a24506696f80642adcaff8a4de4f5/s/master "wercker status")](https://app.wercker.com/project/byKey/820a24506696f80642adcaff8a4de4f5)

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

It is used for section name of library. If not exist, directory name is used.

## TODO

* support cocoapods

## LICENSE

MIT
