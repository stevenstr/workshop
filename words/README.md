# Search 

This is a simple task which test your ability to write effective parallel code.

The main goal is implement function `SearchWord`, this function should fetch documents from site, parse it and find count of occurrences for `word`.
Keep it in mind, word contains english alphabet characters only.

Algorithm should use as much processor cores as possible, must not have race conditions and produce stable result.

You can test proper work of your code using command `go test -v ./...`.

After performing task please create merge request to master branch, and wait until the end of Continues Integration pipeline.

As metric we will use result of benchmark rounded to seconds, in case of result the same result for two users we chose the earliest solution. 
