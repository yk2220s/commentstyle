# commentstyle

`commentstyle` is the linter to used to keep consistent comment styles.

## install

```
go get -u github.com/yk2220s/commentstyle/cmd/commentstyle
```

## rules

each rule can be disabled with a flag.

### only-ascii

default: true

any comment string should be an ascii character.

### prefer-line-style

default: true

line-style (//-style) should be preferred than block-style (/*-style).
