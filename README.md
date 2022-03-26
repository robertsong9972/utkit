
## utkit


utkit is a tool to help us analyse the unit test (increasing, weighted) coverage

## Installation
to install the latest version
```shell
go get git.garena.com/shopee/listing/Common/utkit
```
specifically, you can specify the version to install
```shell
go get git.garena.com/shopee/listing/Common/utkit@v0.0.1
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
ok      git.garena.com/shopee/listing/spu/spulib/bloom  (cached)        coverage: 91.3% of statements
ok      git.garena.com/shopee/listing/spu/spulib/cron   (cached)        coverage: 70.0% of statements
?       git.garena.com/shopee/listing/spu/spulib/event  [no test files]
ok      git.garena.com/shopee/listing/spu/spulib/hash   (cached)        coverage: 80.0% of statements
?       git.garena.com/shopee/listing/spu/spulib/kafka  [no test files]
?       git.garena.com/shopee/listing/spu/spulib/kafka/monitor  [no test files]
?       git.garena.com/shopee/listing/spu/spulib/notify [no test files]
?       git.garena.com/shopee/listing/spu/spulib/orm    [no test files]
ok      git.garena.com/shopee/listing/spu/spulib/redis  (cached)        coverage: 78.7% of statements
?       git.garena.com/shopee/listing/spu/spulib/spudb/cache    [no test files]
?       git.garena.com/shopee/listing/spu/spulib/spudb/db_proto [no test files]
?       git.garena.com/shopee/listing/spu/spulib/spudb/db_proto/gen     [no test files]
?       git.garena.com/shopee/listing/spu/spulib/spudb/db_proto/gen/mock        [no test files]
ok      git.garena.com/shopee/listing/spu/spulib/spudb/query    (cached)        coverage: 24.4% of statements
?       git.garena.com/shopee/listing/spu/spulib/spudb/query/mock       [no test files]
?       git.garena.com/shopee/listing/spu/spulib/trace  [no test files]
ok      git.garena.com/shopee/listing/spu/spulib/utils  (cached)        coverage: 86.9% of statements
/Users/jiang.song/go/bin/gocov
/Users/jiang.song/go/bin/gocov-xml
-------------
Diff Coverage
Diff: origin/master...HEAD, staged and unstaged changes
-------------
No lines with coverage information in this diff.
-------------

The packages involved in the statistics are as follows:
package:git.garena.com/shopee/listing/spu/spulib/bloom
package:git.garena.com/shopee/listing/spu/spulib/cron
package:git.garena.com/shopee/listing/spu/spulib/hash
package:git.garena.com/shopee/listing/spu/spulib/redis
package:git.garena.com/shopee/listing/spu/spulib/spudb/query
package:git.garena.com/shopee/listing/spu/spulib/utils
The weighted average coverage is: 51.5% of statements

```
For more detail, go to confluence and get started,thanks!
- [UTKIT](https://confluence.shopee.io/pages/viewpage.action?pageId=917279478)

