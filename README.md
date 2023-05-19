# Go Benchmark : JSON v. Protobuf v. Gojay v. Easyjson
This small benchmark is done to test the differences in size and speed between 
json and protobuf.

### Get the code and try it yourself

To run the benchmark:
```
go test --bench=.
```

Try changing the content of the data structure in proto_test.go and see the results

### Libraries used:
JSON library :
``` 
encoding/json 
francoispqt/gojay
mailru/easyjson
```

Protobuf library : 
``` 
github.com/golang/protobuf/proto
```

Results :

|   cpu   |   AMD Ryzen 5 5600H with Radeon Graphics     |
|------|--------|
| goos | lilnux |
|  goarch    |    amd64    |

Marshal Small Struct (2 property)

| name                              | time, ns/op | memory, B/op | allocs, allocs/op |
|-----------------------------------|-------------|-------|--------|
| encoding/json                     | 781.6 | 176 | 2 |
| Gojay                             | 222.7 | 512 | 1 |
| easyjson                          | 211.7  | 240 | 2 |
| protobuf                          | 158.5  | 32 | 1 |


Marshal Medium Struct (6 property)

| name                              | time, ns/op | memory, B/op | allocs, allocs/op |
|-----------------------------------|------|--------|--------|
| encoding/json                     |   3992   | 1056   | 5      |
| Gojay                             |    948.6  | 512    | 1      |
| easyjson                          |  1139  | 928    |    5    |
| protobuf                          |  466.8  | 176    |    1    |

Marshal Large Struct (10 property)

| name                              | time, ns/op | memory, B/op | allocs, allocs/op |
|-----------------------------------|-------------|--------------|-------------------|
| encoding/json                     |      8645       |      2010        | 6                 |
| Gojay                             |       2322      |        2048      | 2                 |
| easyjson                          |       2316      |        1416      | 6                 |
| protobuf                          |        1263     |      640        | 1                 |

Unmarshal Small Struct (2 property)

| name                              | time, ns/op | memory, B/op | allocs, allocs/op |
|-----------------------------------|-------------|--------------|-------------------|
| encoding/json                     |      1034       |      184        | 4                 |
| Gojay                             |       221.5      |       48       | 1                 |
| easyjson                          |        212.7     |        32      | 2                 |
| protobuf                          |        188.9     |       32       | 2                 |

Unmarshal Medium Struct (6 property)

| name                              | time, ns/op | memory, B/op | allocs, allocs/op |
|-----------------------------------|-------------|--------------|-------------------|
| encoding/json                     |      2875       | 320          | 4                 |
| Gojay                             |      885.1       | 240          | 1                 |
| easyjson                          |      972.4       |    168          | 2                 |
| protobuf                          |       493.8      |      152        | 2                 |

Unmarshal Large Struct (10 property)

| name                              | time, ns/op | memory, B/op | allocs, allocs/op |
|-----------------------------------|-------------|--------------|-------------------|
| encoding/json                     |      6521       |        800      | 6                 |
| Gojay                             |       2712      |      704        | 1                 |
| easyjson                          |       1986      |       648       | 4                 |
| protobuf                          |       1359      |       584       | 4                 |
