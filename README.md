# go-japanmesh

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

## About

JIS 規格で定められている地域メッシュを扱うユーティリティのGoでの実装です。  
地域メッシュコード、緯度経度の相互変換がおこなえます。

地域メッシュの区分は下記の通りです。  

レベル | 区画の種類 | コード桁数 | 一辺の長さ
:-|:-|:-|:-
1 | 第１次地域区画 | 4桁 | 約80km
2 | 第２次地域区画 | 6桁 | 約10km
3 | 基準地域メッシュ(第３次地域区画) | 8桁 | 約1km
1/2 | ２分の１地域メッシュ | 9桁 | 約500m
1/4 | ４分の１地域メッシュ | 10桁 | 約250m
1/8 | ８分の１地域メッシュ | 11桁 | 約125m

## Installation
```cassandraql
$ go get -u github.com/keitaro1020/go-japanmesh
```
## Usage
```go
import "github.com/keitaro1020/go-japanmesh"
```

### japanmesh.ToCode(geoCode japanmesh.GeoCode, level Level)

指定した緯度経度(WGS84)から、地域メッシュコードを取得します。  

```go
	code, _ := japanmesh.ToCode(japanmesh.GeoCode{
		Latitude:  35.70078,
		Longitude: 139.71475,
	}, japanmesh.Level3)
	fmt.Println(code)
	// => "53394547"
```

### japanmesh.ToGeoJSON(code[, properties])

指定した地域メッシュコードから、ポリゴンデータ(GeoJSON)を取得します。  

```go
	jsn, _ := japanmesh.ToGeoJSON("53394547", nil)
	jsnStr, _ := json.Marshal(jsn)
	fmt.Println(string(jsnStr))
	// => {
	//  "type": "Feature",
	//  "geometry": {
	//    "type": "Polygon",
	//    "coordinates": [
	//      [
	//        [139.725, 35.70833333333333],
	//        [139.7125, 35.70833333333333],
	//        [139.7125, 35.699999999999996],
	//        [139.725, 35.699999999999996],
	//        [139.725, 35.70833333333333]
	//      ]
	//    ]
	//  },
	//  "properties": null
	//}
```

### japanmesh.GetLevel(code)

指定した地域メッシュコードのレベルを取得します。  

```go
	level, _ := japanmesh.GetLevel("53394547")
	fmt.Println(level)
	// => 3
```

### japanmesh.GetCodes(code)
指定した地域メッシュコードの直下のレベルの地域メッシュコードを取得します。  

```go
	codes, _ := japanmesh.GetCodes("53394547")
	fmt.Println(codes)
	// => [533945471 533945472 533945473 533945474]
```

## Author

[keitaro shishido](https://github.com/keitaro1020)

## License

This project is licensed under the terms of the [MIT license](https://github.com/keitaro1020/go-japanmesh/blob/master/LICENSE).

## Acknowledgments

- inspired by [japanmesh](https://github.com/qazsato/japanmesh)
