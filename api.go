package dmm

import (
    "encoding/xml"
    "fmt"
    "net/http"
    "net/url"
    "io"
    "regexp"
    "strconv"
    "time"

    "golang.org/x/text/encoding/japanese"
    "golang.org/x/text/transform"
)

const (
    SITE_ALLAGES   = "DMM.com"
    SITE_ADULT     = "DMM.co.jp"
    OPERATION_LIST = "ItemList"
    DEFAULT_LENGTH = 20
    DEFAULT_OFFSET = 1
)

var (
    API_URL      = "http://affiliate-api.dmm.com/"
    API_VERSION  = "2.0.0"
)

type Request struct {
    ApiId        string
    AffiliateId  string
    Operation    string
    Version      string
    Timestamp    string
    Site         string
    Service      string
    Floor        string
    Length       int64
    Offset       int64
    Sort         string
    Keyword      string
}

type Response struct {
    Result  Result  `xml:"result"`
}

type Result struct {
    ResultCount   int64  `xml:"result_count"`
    TotalCount    int64  `xml:"total_count"`
    FirstPosition int64  `xml:"first_position"`
    Items         ItemList `xml:"items"`
}

type ItemList struct {
    Item []Item `xml:"item"`
}

type Item struct {              
    ServiceName        string `xml:"service_name"`
    FloorName          string `xml:"floor_name"`
    CategoryName       string `xml:"category_name"`
    ContentId          string `xml:"content_id"`
    ProductId          string `xml:"product_id"`
    Title              string `xml:"title"`
    Url                string `xml:"URL"`
    UrlMoble           string `xml:"URLsp"`
    AffiliateUrl       string `xml:"affiliateURL"`
    AffiliateUrlMobile string `xml:"affiliateURLsp"`
    Date               string `xml:"date"`
    JANCode            string `xml:"jancode"`
    ProductCode        string `xml:"maker_product"`
    ISBN               string `xml:"isbn"`
    Stock              string `xml:"stock"`
    ImageUrl           ImageUrlList `xml:"imageURL"`
    SampleImageUrl     SampleImageUrlList `xml:"sampleImageURL"`
    PriceInformation   PriceInformation `xml:"prices"`
    ItemInformation    ItemInformation `xml:"iteminfo"`
    BandaiInformation  BandaiInformation `xml:"bandaiinfo"`
    CdInformation      CdInformation `xml:"cdinfo"`
}

type ImageUrlList struct {
    List  string `xml:"list"`
    Small string `xml:"small"`
    Large string `xml:"large"`
}

type SampleImageUrlList struct {
    Sample_s SmallSampleList  `xml:"sample_s"`
}

type SmallSampleList struct {
    Image []string `xml:"image"`
}

type PriceInformation struct {
    Price     string `xml:"price"`
    PriceAll  string `xml:"price_all"`
    RetailPrice string `xml:"list_price"`
    Distributions DistributionList `xml:"deliveries"`
}

type DistributionList struct {
    Distribution []Distribution `xml:"delivery"`
}

type Distribution struct {
    Type  string `xml:"type"`
    Price string `xml:"price"`
}

type ItemInformation struct {
    Maker     ItemComponent   `xml:"maker"`
    Label     ItemComponent   `xml:"label"`
    Series    ItemComponent   `xml:"series"`
    Keywords  []ItemComponent `xml:"keyword"`
    Genres    []ItemComponent `xml:"genre"`
    Actors    []ItemComponent `xml:"actor"`
    Artists   []ItemComponent `xml:"artist"`
    Authors   []ItemComponent `xml:"author"`
    Directors []ItemComponent `xml:"director"`
    Fighters  []ItemComponent `xml:"fighter"`
    Colors    []ItemComponent `xml:"color"`
    Sizes     []ItemComponent `xml:"size"`
}

type ItemComponent struct {
    Id   string `xml:"id"`
    Name string `xml:"name"`
}

type BandaiInformation struct {
    TitleCode string `xml:"titlecode"`
}

type CdInformation struct {
    Kind string `xml:"kind"`
}

// creates new client
func New(api_id, affiliate_id string) (*Request) {
    t := time.Now()
    ts := t.Format("2006-01-02 15:04:05")

    return &Request{
        ApiId:       api_id,
        AffiliateId: affiliate_id,
        Operation:   OPERATION_LIST,
        Version:     API_VERSION,
        Timestamp:   ts,
        Site:        SITE_ALLAGES,
        Length:      DEFAULT_LENGTH,
        Offset:      DEFAULT_OFFSET,
    }
}

// set site parameter
// site value: DMM.com or DMM.co.jp
func (req *Request) SetSite(site string) *Request {
    req.Site = site
    return req
}

// set service parameter
// see.) https://affiliate.dmm.com/api/reference/com/all/
func (req *Request) SetService(service string) *Request {
    req.Service = service
    return req
}

// set floor parameter
// see.) https://affiliate.dmm.com/api/reference/com/all/
func (req *Request) SetFloor(floor string) *Request {
    req.Floor = floor
    return req
}

// set hits parameter
func (req *Request) SetLength(length int64) *Request {
    req.Length = length
    return req
}

// set offset parameter
func (req *Request) SetOffset(offset int64) *Request {
    req.Offset = offset
    return req
}

// set sort parameter
func (req *Request) SetSort(sort string) *Request {
    req.Sort = sort
    return req
}

// set keyword parameter
func (req *Request) SetKeyword(keyword string) *Request {
    req.Keyword = keyword
    return req
}

// does API request
func (req *Request) Execute() (*Response, error) {
    reqUrl, err := requestUrl(req)
    if err != nil {
        return nil, err
    }

    resp, err := http.Get(reqUrl)
    if err != nil {
        return nil, fmt.Errorf("Error at API request:%#v", err)
    }
    defer resp.Body.Close()

    decoder := xml.NewDecoder(resp.Body)
    decoder.CharsetReader = conversion

    var result Response
    err = decoder.Decode(&result)
    if err != nil {
        return nil, fmt.Errorf("Error at decoding:%#v", err)
    }
    return &result, nil
}

// converts XML charaset from EUC-JP to UTF-8
func conversion(charset string, input io.Reader) (io.Reader, error) {
    if charset == "euc-jp" {
        dc := transform.NewReader(input, japanese.EUCJP.NewDecoder())
        return dc, nil
    }
    if charset == "utf-8" {
        return input, nil
    }
    return nil, fmt.Errorf("unsupported charset: %q", charset)
}

// creates web api request url
func requestUrl(req *Request) (string, error) {
    if req.ApiId == "" {
        return "", fmt.Errorf("set invalid parameter. ApiId")
    }
    if !validateAffiliateId(req.AffiliateId) {
        return "", fmt.Errorf("set invalid parameter. AffiliateId")
    }

    if req.Site != SITE_ALLAGES && req.Site != SITE_ADULT {
        return "", fmt.Errorf("set invalid parameter. Site:%s", req.Site)
    }
    if req.Length < 1 || req.Length > 200 {
        return "", fmt.Errorf("out of range. (0 < Length < 201) Length:%s", req.Length)
    }
    if req.Offset < 1 {
        return "", fmt.Errorf("out of range. (Offset > 0) Offset:%s", req.Offset)
    }

    queries := url.Values{}
    queries.Set("api_id", req.ApiId)
    queries.Set("affiliate_id", req.AffiliateId)
    queries.Set("operation", req.Operation)
    queries.Set("version", req.Version)
    queries.Set("timestamp", req.Timestamp)
    queries.Set("site", req.Site)
    queries.Set("hits", strconv.FormatInt(req.Length, 10))
    queries.Set("offset", strconv.FormatInt(req.Offset, 10))

    if (req.Service != "") {
        queries.Set("service", req.Service)
    }
    if (req.Floor != "") {
        queries.Set("floor", req.Floor)
    }
    if (req.Sort != "") {
        queries.Set("sort", req.Sort)
    }
    if (req.Keyword != "") {
        queries.Set("keyword", req.Keyword)
    }
    return API_URL + "?" + queries.Encode(), nil
}

// validates affiliate_id
// example value: dummy-999
// affiliate number range: 990 ~ 999
func validateAffiliateId(affiliate_id string) bool {
    return regexp.MustCompile(`^.+-99[0-9]$`).Match([]byte(affiliate_id))
}
