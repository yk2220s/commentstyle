# commentstyle

`commentstyle` is the linter to used to keep consistent comment styles.

## install

```
go get -u github.com/yk2220s/commentstyle/cmd/commentstyle
```

### use as golangci-lint plugin

need to build `commentstyle` as plugin

```
go build -buildmode=plugin -o path/to/plugin github.com/yk2220s/commentstyle/plugin/commentstyle
```

add lines below to your `.golangci.yml`

```yaml
linters
  enable:
  	- commentstyle
linters-settings:
  custom:
    commentstyle:
      path: path/to/plugin/commentstyle.so
```

## rules

each rule can be disabled with a flag.

### only-ascii

default: true

any comment string should be an ascii character.

### prefer-line-style

default: true

line-style (//-style) should be preferred than block-style (/\*-style).

## thanks

- [gostaticanalysis/skeleton](https://github.com/gostaticanalysis/skeleton)
