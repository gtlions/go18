package tenxunloc

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"

	"gopkg.in/square/go-jose.v2/json"
)

type Location struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

type AddressComponents struct {
	Province     string `json:"province,omitempty"`
	City         string `json:"city,omitempty"`
	District     string `json:"district,omitempty"`
	Street       string `json:"street,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

type AdInfo struct {
	Adcode string `json:"adcode,omitempty"`
}

type ResultLoc struct {
	Title             string            `json:"title,omitempty"`
	Location          Location          `json:"location,omitempty"`
	AddressComponents AddressComponents `json:"address_components,omitempty"`
	AdInfo            AdInfo            `json:"ad_info,omitempty"`
	Reliability       int               `json:"reliability,omitempty"`
	Level             int               `json:"level,omitempty"`
}
type RespLoc struct {
	Status  int       `json:"status,omitempty"`
	Message string    `json:"message,omitempty"`
	Result  ResultLoc `json:"result,omitempty"`
}

type Element struct {
	Distance int `json:"distance,omitempty"`
	Duration int `json:"duration,omitempty"`
}
type Row struct {
	Elements []Element `json:"elements,omitempty"`
}

type ResultGeoDistance struct {
	Rows []Row `json:"rows,omitempty"`
}

type RespGeoDistance struct {
	Status  int               `json:"status,omitempty"`
	Message string            `json:"message,omitempty"`
	Result  ResultGeoDistance `json:"result,omitempty"`
}

// Address2Geo 腾讯位置服务-地址解析（地址转坐标）
//
// address 地址
//
// accessKey 访问Key
//
// secretKey 访问密钥
func Address2Geo(address, accessKey, secretKey string) (rsp RespLoc, err error) {
	params := url.Values{}
	Url, err := url.Parse("https://apis.map.qq.com/ws/geocoder/v1")
	if err != nil {
		return RespLoc{}, err
	}
	key := accessKey
	sk := secretKey
	params.Add("key", key)
	params.Add("address", address)
	paramIdx := make([]string, 0)
	for k := range params {
		paramIdx = append(paramIdx, k)
	}
	sort.Strings(paramIdx)
	sortedQueryString := ""
	for _, v := range paramIdx {
		sortedQueryString += "&" + v + "=" + params[v][0]
	}
	sortedQueryString = sortedQueryString[1:]
	stringToSign := fmt.Sprintf("%s?%s%s", Url.Path, sortedQueryString, sk)
	h := md5.New()
	h.Write([]byte(stringToSign))
	sign := hex.EncodeToString(h.Sum(nil))
	params.Set("sig", sign)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	respLoc := RespLoc{}
	if err := json.Unmarshal(body, &respLoc); err != nil {
		return RespLoc{}, err
	}
	return respLoc, nil
}

// Address2Geo 腾讯位置服务-地址解析（地址转坐标）
//
// mode 模式： driving-驾车 walking-步行 bicycling-自行车
//
// from 起点经纬度，逗号分隔
//
// to 终点经纬度，逗号分隔
//
// accessKey 访问Key
//
// secretKey 访问密钥
func GeoDistance(mode, from, to, accessKey, secretKey string) (rsp RespGeoDistance, err error) {
	params := url.Values{}
	Url, err := url.Parse("https://apis.map.qq.com/ws/distance/v1/matrix")
	if err != nil {
		return RespGeoDistance{}, err
	}
	key := accessKey
	sk := secretKey
	params.Add("key", key)
	params.Add("mode", mode)
	params.Add("from", from)
	params.Add("to", to)
	paramIdx := make([]string, 0)
	for k := range params {
		paramIdx = append(paramIdx, k)
	}
	sort.Strings(paramIdx)
	sortedQueryString := ""
	for _, v := range paramIdx {
		sortedQueryString += "&" + v + "=" + params[v][0]
	}
	sortedQueryString = sortedQueryString[1:]
	stringToSign := fmt.Sprintf("%s?%s%s", Url.Path, sortedQueryString, sk)
	h := md5.New()
	h.Write([]byte(stringToSign))
	sign := hex.EncodeToString(h.Sum(nil))
	params.Set("sig", sign)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	respGeoDistance := RespGeoDistance{}
	if err := json.Unmarshal(body, &respGeoDistance); err != nil {
		return RespGeoDistance{}, err
	}
	return respGeoDistance, nil
}
