
## utkit


utkit is a tool to help us analyse the unit test (increasing, weighted) coverage

## Installation
to install the latest version
```shell
go install github.com/robertsong9972/utkit
```
specifically, you can specify the version to install
```shell
go install github.com/robertsong9972/utkit@v0.0.1
```
## Features (as of 01/20/2022)

* support delta code UT analysis
* support weighted code UT analysis

## Run
```shell
utkit
```

## An example of output from spulib:

```text
?       github.com/robertsong9972/utkit [no test files]
?       github.com/robertsong9972/utkit/internal/config [no test files]
ok      github.com/robertsong9972/utkit/internal/core   0.642s  coverage: 30.7% of statements
?       github.com/robertsong9972/utkit/internal/model  [no test files]
ok      github.com/robertsong9972/utkit/internal/util   0.339s  coverage: 66.7% of statements
/Users/songjiang/go/bin/gocov
/Users/songjiang/go/bin/gocov-xml
-------------
Diff Coverage
Diff: origin/master...HEAD, staged and unstaged changes
-------------
No lines with coverage information in this diff.
-------------

The packages involved in the statistics are as follows:
package:github.com/robertsong9972/utkit/internal/core
package:github.com/robertsong9972/utkit/internal/util
The weighted average coverage: 32.3% of statements

```
