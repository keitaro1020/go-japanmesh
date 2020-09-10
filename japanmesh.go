package japanmesh

import (
	"fmt"
	"math"
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

type MeshCode string
type MeshCodes []MeshCode

type GeoCode struct {
	Latitude  float64
	Longitude float64
}

// ToCode 緯度経度から地域メッシュコードを取得する。
// 算出式 : https://www.stat.go.jp/data/mesh/pdf/gaiyo1.pdf
func ToCode(geoCode GeoCode, level Level) (MeshCode, error) {
	// （１）緯度よりｐ，ｑ，ｒ，ｓ，ｔを算出
	p := math.Floor((geoCode.Latitude * 60) / 40)
	a := math.Mod(geoCode.Latitude*60, 40)
	q := math.Floor(a / 5)
	b := math.Mod(a, 5)
	r := math.Floor((b * 60) / 30)
	c := math.Mod(b*60, 30)
	s := math.Floor(c / 15)
	d := math.Mod(c, 15)
	t := math.Floor(d / 7.5)
	// 以下、８分の１地域メッシュ算出のため拡張
	ee := math.Mod(d, 7.5)
	uu := math.Floor(ee / 3.75)

	// （２）経度よりｕ，ｖ，ｗ，ｘ，ｙを算出
	u := math.Floor(geoCode.Longitude - 100)
	f := geoCode.Longitude - 100 - u
	v := math.Floor((f * 60) / 7.5)
	g := math.Mod(f*60, 7.5)
	w := math.Floor((g * 60) / 45)
	h := math.Mod(g*60, 45)
	x := math.Floor(h / 22.5)
	i := math.Mod(h, 22.5)
	y := math.Floor(i / 11.25)
	// 以下、８分の１地域メッシュ算出のため拡張
	jj := math.Mod(i, 11.25)
	zz := math.Floor(jj / 5.625)

	// （３）ｓ，ｘよりｍを算出，ｔ，ｙよりｎを算出
	m := s*2 + (x + 1)
	n := t*2 + (y + 1)
	// 以下、８分の１地域メッシュ算出のため拡張
	oo := uu*2 + (zz + 1)

	// （４）ｐ，ｑ，ｒ，ｕ，ｖ，ｗ，ｍ、ｎ、ooより地域メッシュ・コードを算出
	code1 := Level1Code(fmt.Sprintf("%.f%.f", p, u))
	if _, ok := level1Codes[code1]; !ok {
		return "", ErrInvalidArea
	}
	code := MeshCode(fmt.Sprintf("%.f%.f%.f%.f%.f%.f%.f%.f%.f", p, u, q, v, r, w, m, n, oo))
	return getCodeByLevel(code, level), nil
}

// ToGeoJSON
func ToGeoJSON(code MeshCode, properties map[string]interface{}) (*geojson.Feature, error) {
	if !isValidCode(code) {
		return nil, ErrInvalidMeshCode
	}
	lv1X, err := strconv.ParseFloat(string(code[2:4]), 64)
	if err != nil {
		return nil, err
	}
	lv1Y, err := strconv.ParseFloat(string(code[0:2]), 64)
	if err != nil {
		return nil, err
	}

	var minX, maxX, minY, maxY float64
	digit := code.getDigit()
	if digit >= level1Mesh.Digit {
		minX =
			level1Mesh.Section.Lng.Min +
				(lv1X-level1Mesh.Section.X.Min)*level1Mesh.Distance.Lng
		maxX = minX + level1Mesh.Distance.Lng
		minY =
			level1Mesh.Section.Lat.Min +
				(lv1Y-level1Mesh.Section.Y.Min)*level1Mesh.Distance.Lat
		maxY = minY + level1Mesh.Distance.Lat
	}

	if digit >= level2Mesh.Digit {
		lv2X, err := strconv.ParseFloat(string(code[5:6]), 64)
		if err != nil {
			return nil, err
		}
		lv2Y, err := strconv.ParseFloat(string(code[4:5]), 64)
		if err != nil {
			return nil, err
		}
		minX += lv2X * level2Mesh.Distance.Lng
		maxX = minX + level2Mesh.Distance.Lng
		minY += lv2Y * level2Mesh.Distance.Lat
		maxY = minY + level2Mesh.Distance.Lat
	}

	if digit >= level3Mesh.Digit {
		lv3X, err := strconv.ParseFloat(string(code[7:8]), 64)
		if err != nil {
			return nil, err
		}
		lv3Y, err := strconv.ParseFloat(string(code[6:7]), 64)
		if err != nil {
			return nil, err
		}
		minX += lv3X * level3Mesh.Distance.Lng
		maxX = minX + level3Mesh.Distance.Lng
		minY += lv3Y * level3Mesh.Distance.Lat
		maxY = minY + level3Mesh.Distance.Lat
	}

	if digit >= levelHalfMesh.Digit {
		lv4Num := code[8:9]
		var lv4X, lv4Y float64
		switch lv4Num {
		case "1":
			lv4X = 0
			lv4Y = 0
		case "2":
			lv4X = 1
			lv4Y = 0
		case "3":
			lv4X = 0
			lv4Y = 1
		default:
			lv4X = 1
			lv4Y = 1
		}
		minX += lv4X * levelHalfMesh.Distance.Lng
		maxX = minX + levelHalfMesh.Distance.Lng
		minY += lv4Y * levelHalfMesh.Distance.Lat
		maxY = minY + levelHalfMesh.Distance.Lat
	}

	if digit >= levelQuarterMesh.Digit {
		lv5Num := code[9:10]
		var lv5X, lv5Y float64
		switch lv5Num {
		case "1":
			lv5X = 0
			lv5Y = 0
		case "2":
			lv5X = 1
			lv5Y = 0
		case "3":
			lv5X = 0
			lv5Y = 1
		default:
			lv5X = 1
			lv5Y = 1
		}

		minX += lv5X * levelQuarterMesh.Distance.Lng
		maxX = minX + levelQuarterMesh.Distance.Lng
		minY += lv5Y * levelQuarterMesh.Distance.Lat
		maxY = minY + levelQuarterMesh.Distance.Lat
	}

	if digit >= levelOneEighthMesh.Digit {
		lv6Num := code[10:11]
		var lv6X, lv6Y float64
		switch lv6Num {
		case "1":
			lv6X = 0
			lv6Y = 0
		case "2":
			lv6X = 1
			lv6Y = 0
		case "3":
			lv6X = 0
			lv6Y = 1
		default:
			lv6X = 1
			lv6Y = 1
		}
		minX += lv6X * levelOneEighthMesh.Distance.Lng
		maxX = minX + levelOneEighthMesh.Distance.Lng
		minY += lv6Y * levelOneEighthMesh.Distance.Lat
		maxY = minY + levelOneEighthMesh.Distance.Lat
	}
	return createGeoJSON(minX, maxX, minY, maxY, properties), nil
}

// GetLevel
func GetLevel(code MeshCode) (Level, error) {
	switch code.getDigit() {
	case level1Mesh.Digit:
		return Level1, nil
	case level2Mesh.Digit:
		return Level2, nil
	case level3Mesh.Digit:
		return Level3, nil
	case levelHalfMesh.Digit:
		return LevelHalf, nil
	case levelQuarterMesh.Digit:
		return LevelQuarter, nil
	case levelOneEighthMesh.Digit:
		return LevelOneEighth, nil
	}
	return "", ErrInvalidMeshCode
}

// GetCodes
func GetCodes(code MeshCode) (MeshCodes, error) {
	if !isValidCode(code) {
		return nil, ErrInvalidMeshCode
	}
	codes := make(MeshCodes, 0)
	level, err := GetLevel(code)
	if err != nil {
		return nil, err
	}
	switch level {
	case Level1:
		// 2次メッシュ
		for y2 := 0; y2 < level2Mesh.Division.Y; y2++ {
			for x2 := 0; x2 < level2Mesh.Division.X; x2++ {
				codes = append(codes, MeshCode(fmt.Sprintf("%s%d%d", code, y2, x2)))
			}
		}
	case Level2:
		// 3次メッシュ
		for y3 := 0; y3 < level3Mesh.Division.Y; y3++ {
			for x3 := 0; x3 < level3Mesh.Division.X; x3++ {
				codes = append(codes, MeshCode(fmt.Sprintf("%s%d%d", code, y3, x3)))
			}
		}
	case Level3, LevelHalf, LevelQuarter, LevelOneEighth:
		// 4次,5次,6次メッシュ
		divisionNum := 4 // 分割数(=マスの数)
		for i := 1; i <= divisionNum; i++ {
			codes = append(codes, MeshCode(fmt.Sprintf("%s%d", code, i)))
		}
	}

	return codes, nil
}

// SplitCodeByLevel
func SplitCodeByLevel(code MeshCode) []MeshCode {
	var codes []MeshCode
	digit := code.getDigit()
	if digit >= level1Mesh.Digit {
		codes = append(codes, getCodeByLevel(code, Level1))
	}
	if digit >= level2Mesh.Digit {
		codes = append(codes, getCodeByLevel(code, Level2))
	}
	if digit >= level3Mesh.Digit {
		codes = append(codes, getCodeByLevel(code, Level3))
	}
	if digit >= levelHalfMesh.Digit {
		codes = append(codes, getCodeByLevel(code, LevelHalf))
	}
	if digit >= levelQuarterMesh.Digit {
		codes = append(codes, getCodeByLevel(code, LevelQuarter))
	}
	if digit >= levelOneEighthMesh.Digit {
		codes = append(codes, getCodeByLevel(code, LevelOneEighth))
	}
	return codes
}

func isValidCode(code MeshCode) bool {
	switch code.getDigit() {
	case
		level1Mesh.Digit,
		level2Mesh.Digit,
		level3Mesh.Digit,
		levelHalfMesh.Digit,
		levelQuarterMesh.Digit,
		levelOneEighthMesh.Digit:
		return true
	}
	return false
}

func getCodeByLevel(code MeshCode, level Level) MeshCode {
	switch level {
	case Level1:
		return code[0:level1Mesh.Digit]
	case Level2:
		return code[0:level2Mesh.Digit]
	case Level3:
		return code[0:level3Mesh.Digit]
	case LevelHalf:
		return code[0:levelHalfMesh.Digit]
	case LevelQuarter:
		return code[0:levelQuarterMesh.Digit]
	case LevelOneEighth:
		return code[0:levelOneEighthMesh.Digit]
	default:
		return code
	}
}

func createGeoJSON(minX, maxX, minY, maxY float64, properties map[string]interface{}) *geojson.Feature {
	// 北東 -> 北西 -> 南西 -> 南東 -> 北東
	coordinates := [][]float64{
		{maxX, maxY},
		{minX, maxY},
		{minX, minY},
		{maxX, minY},
		{maxX, maxY},
	}
	feature := geojson.NewFeature(geojson.NewPolygonGeometry([][][]float64{coordinates}))
	feature.Properties = properties
	return feature
}

func (code MeshCode) getDigit() int {
	return len(code)
}
