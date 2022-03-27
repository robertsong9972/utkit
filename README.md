
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

## An example of output:

```text
?       github.com/robertsong9972/utkit [no test files]
?       github.com/robertsong9972/utkit/internal/config [no test files]
ok      github.com/robertsong9972/utkit/internal/core   0.355s  coverage: 32.2% of statements
?       github.com/robertsong9972/utkit/internal/model  [no test files]
ok      github.com/robertsong9972/utkit/internal/util   0.101s  coverage: 50.0% of statements
-------------
Diff Coverage
Diff: origin/master...HEAD, staged and unstaged changes
-------------
internal/core/module_utils.go (0.0%): Missing lines 16-17,29,35,46,51-53,55-57,59-61,63-67,69-71,73
-------------
Total:   23 lines
Missing: 23 lines
Coverage: 0%
-------------

----------------------------
Packages involved in config:
    github.com/robertsong9972/utkit/internal/core
    github.com/robertsong9972/utkit/internal/util
verage coverage: 32.8%
----------------------------
The weighted average coverage: 32.8% of statements


```
