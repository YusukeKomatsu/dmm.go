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

>r = *dmm.request

| API | this library | description | e.g. | how to set parameter |
|---|---|---|---|---|
| api_id | ApiId | API ID | "KcZ2ymn6VPufm4XjxFu6" | r := New("KcZ2ymn6VPufm4XjxFu6", "dummy-999") |
| affiliate_id | AffiliateId | affiliate iD | dummy-999 | r := New("foobarbazbuzz", "dummy-999") |
| operation | Operation | API method name | ItemList | Nothing (operation exists ItemList **ONLY**) |
| version | Version | API version | 2.00 | Nothing (version 2 **ONLY**) |
| timestamp | Timestamp | timestamp | 2006-01-02 15:04:05 | Nothing(timestamp is automatically set) |
| site | Site | site name (DMM.com or DMM.co.jp) | DMM.co.jp | r.SetSite("DMM.com") |
| service | Service | target service | mono | r.SetService("mono") |
| floor | Floor | target floor | dvd | r.SetFloor("dvd") |
| hits | Length | maximum request length | 100 | r.SetLength(100) |
| offset | Offset | request data offset | 0 | r.SetOffset(0) |
| sort | Sort | response data sort | rank | r.SetSort("rank") |
| keyword | Keyword | request keyword | social network | r.SetKeyword("social network") |
