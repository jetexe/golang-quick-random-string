# go quick random string

This package was made from answer on [StackOwerflow][Stack]. Thanks [@icza][icza] for help

## Banchmark

Command | count | ns/op | B/op | allocs/op
------- | ----- | ----- | ---- | ---------
`RandStringBytesMaskImprSrcUnsafe(4)` |10000000|46.4 ns/op|4 B/op |1 allocs/op
`RandStringBytesMaskImprSrcUnsafe(32)`|10000000|221 ns/op |32 B/op|1 allocs/op

[Stack]:https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
[icza]:https://github.com/icza
