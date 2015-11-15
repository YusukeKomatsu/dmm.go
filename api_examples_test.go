package dmm

import "fmt"

func ExampleNew() {
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
}

func ExampleNew_OneLinerExecute() {
    rst, err := New("foobarbazbuzz", "dummy-999").SetSite(SITE_ADULT).SetLength(1).Execute()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rst)
    }
}