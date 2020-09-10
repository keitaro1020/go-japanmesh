package japanmesh

import "errors"

type Level string
type Level1Code string

type Mesh struct {
	// コード桁数
	Digit int
	// メッシュ分割数
	Division Division
	// 緯度経度の間隔(単位:度)
	Distance Distance
	Section  Section
}

type Division struct {
	X int
	Y int
}

type Distance struct {
	Lat float64
	Lng float64
}

type Section struct {
	X   SectionXY
	Y   SectionXY
	Lat SectionLatLng
	Lng SectionLatLng
}

type SectionXY struct {
	Min float64
	Max float64
}
type SectionLatLng struct {
	Min float64
}

const (
	Level1         Level = "1"
	Level2         Level = "2"
	Level3         Level = "3"
	LevelHalf      Level = "1/2"
	LevelQuarter   Level = "1/4"
	LevelOneEighth Level = "1/8"
)

var (
	ErrInvalidArea     = errors.New("invalid area")
	ErrInvalidMeshCode = errors.New("invalid meshcode")
)

// 第1次地域区画
var level1Mesh = Mesh{
	Digit: 4,
	Division: Division{
		X: 32,
		Y: 39,
	},
	// 緯度: 40分, 経度: 1度
	Distance: Distance{
		Lat: float64(40) / float64(60),
		Lng: 1,
	},
	// 日本の国土にかかる第１次地域区画
	Section: Section{
		X: SectionXY{
			Min: 22,
			Max: 53,
		},
		Y: SectionXY{
			Min: 30,
			Max: 68,
		},
		Lat: SectionLatLng{
			Min: 20,
		},
		Lng: SectionLatLng{
			Min: 122,
		},
	},
}

// 第2次地域区画
var level2Mesh = Mesh{
	Digit: 6,
	Division: Division{
		X: 8,
		Y: 8,
	},
	// 緯度: 5分, 経度: 7分30秒
	Distance: Distance{
		Lat: float64(5) / float64(60),
		Lng: float64(7)/float64(60) + float64(30)/float64(60)/float64(60),
	},
}

// 基準地域メッシュ(第3次地域区画)
var level3Mesh = Mesh{
	Digit: 8,
	Division: Division{
		X: 10,
		Y: 10,
	},
	// 緯度: 30秒, 経度: 45秒
	Distance: Distance{
		Lat: float64(30) / float64(60) / float64(60),
		Lng: float64(45) / float64(60) / float64(60),
	},
}

// 2分の1地域メッシュ
var levelHalfMesh = Mesh{
	Digit: 9,
	Division: Division{
		X: 2,
		Y: 2,
	},
	// 緯度: 15秒, 経度: 22.5秒
	Distance: Distance{
		Lat: float64(15) / float64(60) / float64(60),
		Lng: float64(22.5) / float64(60) / float64(60),
	},
}

// 4分の1地域メッシュ
var levelQuarterMesh = Mesh{
	Digit: 10,
	Division: Division{
		X: 2,
		Y: 2,
	},
	// 緯度: 7.5秒, 経度: 11.25秒
	Distance: Distance{
		Lat: float64(7.5) / float64(60) / float64(60),
		Lng: float64(11.25) / float64(60) / float64(60),
	},
}

// 8分の1地域メッシュ
var levelOneEighthMesh = Mesh{
	Digit: 11,
	Division: Division{
		X: 2,
		Y: 2,
	},
	// 緯度: 3.75秒, 経度: 5.625秒
	Distance: Distance{
		Lat: float64(3.75) / float64(60) / float64(60),
		Lng: float64(5.625) / float64(60) / float64(60),
	},
}

// 第１次地域区画の全メッシュコード
// https://www.e-stat.go.jp/pdf/gis/primary_mesh_jouhou.pdf
var level1Codes = map[Level1Code]interface{}{
	"3036": struct{}{},
	"3622": struct{}{},
	"3623": struct{}{},
	"3624": struct{}{},
	"3631": struct{}{},
	"3641": struct{}{},
	"3653": struct{}{},
	"3724": struct{}{},
	"3725": struct{}{},
	"3741": struct{}{},
	"3823": struct{}{},
	"3824": struct{}{},
	"3831": struct{}{},
	"3841": struct{}{},
	"3926": struct{}{},
	"3927": struct{}{},
	"3928": struct{}{},
	"3942": struct{}{},
	"4027": struct{}{},
	"4028": struct{}{},
	"4040": struct{}{},
	"4042": struct{}{},
	"4128": struct{}{},
	"4129": struct{}{},
	"4142": struct{}{},
	"4229": struct{}{},
	"4230": struct{}{},
	"4328": struct{}{},
	"4329": struct{}{},
	"4429": struct{}{},
	"4440": struct{}{},
	"4529": struct{}{},
	"4530": struct{}{},
	"4531": struct{}{},
	"4540": struct{}{},
	"4629": struct{}{},
	"4630": struct{}{},
	"4631": struct{}{},
	"4728": struct{}{},
	"4729": struct{}{},
	"4730": struct{}{},
	"4731": struct{}{},
	"4739": struct{}{},
	"4740": struct{}{},
	"4828": struct{}{},
	"4829": struct{}{},
	"4830": struct{}{},
	"4831": struct{}{},
	"4839": struct{}{},
	"4928": struct{}{},
	"4929": struct{}{},
	"4930": struct{}{},
	"4931": struct{}{},
	"4932": struct{}{},
	"4933": struct{}{},
	"4934": struct{}{},
	"4939": struct{}{},
	"5029": struct{}{},
	"5030": struct{}{},
	"5031": struct{}{},
	"5032": struct{}{},
	"5033": struct{}{},
	"5034": struct{}{},
	"5035": struct{}{},
	"5036": struct{}{},
	"5038": struct{}{},
	"5039": struct{}{},
	"5129": struct{}{},
	"5130": struct{}{},
	"5131": struct{}{},
	"5132": struct{}{},
	"5133": struct{}{},
	"5134": struct{}{},
	"5135": struct{}{},
	"5136": struct{}{},
	"5137": struct{}{},
	"5138": struct{}{},
	"5139": struct{}{},
	"5229": struct{}{},
	"5231": struct{}{},
	"5232": struct{}{},
	"5233": struct{}{},
	"5234": struct{}{},
	"5235": struct{}{},
	"5236": struct{}{},
	"5237": struct{}{},
	"5238": struct{}{},
	"5239": struct{}{},
	"5240": struct{}{},
	"5332": struct{}{},
	"5333": struct{}{},
	"5334": struct{}{},
	"5335": struct{}{},
	"5336": struct{}{},
	"5337": struct{}{},
	"5338": struct{}{},
	"5339": struct{}{},
	"5340": struct{}{},
	"5432": struct{}{},
	"5433": struct{}{},
	"5435": struct{}{},
	"5436": struct{}{},
	"5437": struct{}{},
	"5438": struct{}{},
	"5439": struct{}{},
	"5440": struct{}{},
	"5531": struct{}{},
	"5536": struct{}{},
	"5537": struct{}{},
	"5538": struct{}{},
	"5539": struct{}{},
	"5540": struct{}{},
	"5541": struct{}{},
	"5636": struct{}{},
	"5637": struct{}{},
	"5638": struct{}{},
	"5639": struct{}{},
	"5640": struct{}{},
	"5641": struct{}{},
	"5738": struct{}{},
	"5739": struct{}{},
	"5740": struct{}{},
	"5741": struct{}{},
	"5839": struct{}{},
	"5840": struct{}{},
	"5841": struct{}{},
	"5939": struct{}{},
	"5940": struct{}{},
	"5941": struct{}{},
	"5942": struct{}{},
	"6039": struct{}{},
	"6040": struct{}{},
	"6041": struct{}{},
	"6139": struct{}{},
	"6140": struct{}{},
	"6141": struct{}{},
	"6239": struct{}{},
	"6240": struct{}{},
	"6241": struct{}{},
	"6243": struct{}{},
	"6339": struct{}{},
	"6340": struct{}{},
	"6341": struct{}{},
	"6342": struct{}{},
	"6343": struct{}{},
	"6439": struct{}{},
	"6440": struct{}{},
	"6441": struct{}{},
	"6442": struct{}{},
	"6443": struct{}{},
	"6444": struct{}{},
	"6445": struct{}{},
	"6540": struct{}{},
	"6541": struct{}{},
	"6542": struct{}{},
	"6543": struct{}{},
	"6544": struct{}{},
	"6545": struct{}{},
	"6546": struct{}{},
	"6641": struct{}{},
	"6642": struct{}{},
	"6643": struct{}{},
	"6644": struct{}{},
	"6645": struct{}{},
	"6646": struct{}{},
	"6647": struct{}{},
	"6740": struct{}{},
	"6741": struct{}{},
	"6742": struct{}{},
	"6747": struct{}{},
	"6748": struct{}{},
	"6840": struct{}{},
	"6841": struct{}{},
	"6842": struct{}{},
	"6847": struct{}{},
	"6848": struct{}{},
}
