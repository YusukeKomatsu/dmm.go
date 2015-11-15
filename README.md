# dmm.go
DMM Web API Client

see: [DMM Affiliate](https://affiliate.dmm.com/)

## Installation

Standard `go get`:

```
$ go get github.com/YusukeKomatsu/dmm.go
```

## Usage

For usage and examples see the [Godoc](https://godoc.org/github.com/YusukeKomatsu/dmm.go).

## Example

```
api := New("foobarbazbuzz", "dummy-990")
api.SetSite(SITE_ALLAGES)
api.SetService("mono")
api.SetFloor("dvd")
api.SetSort("date")
api.SetLength(1)
result, err := api.Execute()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(result)
}
```

OR

```
rst, err := New("foobarbazbuzz", "dummy-999").SetSite(SITE_ADULT).SetLength(1).Execute()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(rst)
}
```

## Request parameter

| API | this library | description |
|---|---|---|
| api_id | ApiId | API ID |
| affiliate_id | AffiliateId | affiliate iD |
| operation | Operation | API method name |
| version | Version | API version |
| timestamp | Timestamp | timestamp |
| site | Site | site name (DMM.com or DMM.co.jp) |
| service | Service | target service |
| floor | Floor | target floor |
| hits | Length | maximum request length |
| offset | Offset | request data offset |
| sort | Sort | response data sort |
| keyword | Keyword | request keyword |
