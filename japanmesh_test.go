package japanmesh

import (
	"encoding/json"
	"reflect"
	"testing"

	geojson "github.com/paulmach/go.geojson"
)

func TestToCode(t *testing.T) {
	type args struct {
		geoCode GeoCode
		level   Level
	}
	tests := []struct {
		name    string
		args    args
		want    MeshCode
		wantErr bool
	}{
		{name: "level1", args: args{geoCode: GeoCode{Latitude: 35.70078, Longitude: 139.71475}, level: Level1}, want: "5339", wantErr: false},
		{name: "level2", args: args{geoCode: GeoCode{Latitude: 35.70078, Longitude: 139.71475}, level: Level2}, want: "533945", wantErr: false},
		{name: "level3", args: args{geoCode: GeoCode{Latitude: 35.70078, Longitude: 139.71475}, level: Level3}, want: "53394547", wantErr: false},
		{name: "level1-2", args: args{geoCode: GeoCode{Latitude: 35.70078, Longitude: 139.71475}, level: LevelHalf}, want: "533945471", wantErr: false},
		{name: "level1-4", args: args{geoCode: GeoCode{Latitude: 35.70078, Longitude: 139.71475}, level: LevelQuarter}, want: "5339454711", wantErr: false},
		{name: "level1-8", args: args{geoCode: GeoCode{Latitude: 35.70078, Longitude: 139.71475}, level: LevelOneEighth}, want: "53394547112", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCode(tt.args.geoCode, tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToGeoJSON(t *testing.T) {
	jsonLv1 := geojson.NewFeature(geojson.NewPolygonGeometry([][][]float64{{
		{139, 36.666666666666664},
		{138, 36.666666666666664},
		{138, 36},
		{139, 36},
		{139, 36.666666666666664},
	}}))
	jsonLv2 := geojson.NewFeature(geojson.NewPolygonGeometry([][][]float64{{
		{139.75, 35.75},
		{139.625, 35.75},
		{139.625, 35.666666666666664},
		{139.75, 35.666666666666664},
		{139.75, 35.75},
	}}))
	jsonLv3 := geojson.NewFeature(geojson.NewPolygonGeometry([][][]float64{{
		{139.725, 35.70833333333333},
		{139.7125, 35.70833333333333},
		{139.7125, 35.699999999999996},
		{139.725, 35.699999999999996},
		{139.725, 35.70833333333333},
	}}))
	jsonLvHalf := geojson.NewFeature(geojson.NewPolygonGeometry([][][]float64{{
		{139.71875, 35.704166666666666},
		{139.7125, 35.704166666666666},
		{139.7125, 35.699999999999996},
		{139.71875, 35.699999999999996},
		{139.71875, 35.704166666666666},
	}}))
	jsonLvQuarter := geojson.NewFeature(geojson.NewPolygonGeometry([][][]float64{{
		{139.71562500000002, 35.70208333333333},
		{139.7125, 35.70208333333333},
		{139.7125, 35.699999999999996},
		{139.71562500000002, 35.699999999999996},
		{139.71562500000002, 35.70208333333333},
	}}))
	jsonLvOneEight := geojson.NewFeature(geojson.NewPolygonGeometry([][][]float64{{
		{139.71562500000002, 35.70104166666666},
		{139.7140625, 35.70104166666666},
		{139.7140625, 35.699999999999996},
		{139.71562500000002, 35.699999999999996},
		{139.71562500000002, 35.70104166666666},
	}}))

	type args struct {
		code       MeshCode
		properties map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *geojson.Feature
		wantErr bool
	}{
		{name: "level1", args: args{code: "5438", properties: nil}, want: jsonLv1, wantErr: false},
		{name: "level2", args: args{code: "533945", properties: nil}, want: jsonLv2, wantErr: false},
		{name: "level3", args: args{code: "53394547", properties: nil}, want: jsonLv3, wantErr: false},
		{name: "level1-2", args: args{code: "533945471", properties: nil}, want: jsonLvHalf, wantErr: false},
		{name: "level1-4", args: args{code: "5339454711", properties: nil}, want: jsonLvQuarter, wantErr: false},
		{name: "level1-8", args: args{code: "53394547112", properties: nil}, want: jsonLvOneEight, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToGeoJSON(tt.args.code, tt.args.properties)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToGeoJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotjsn, _ := json.Marshal(got)
			wantjsn, _ := json.Marshal(tt.want)
			if !reflect.DeepEqual(gotjsn, wantjsn) {
				t.Errorf("ToGeoJSON() got = %v, want %v", string(gotjsn), string(wantjsn))
			}
		})
	}
}

func TestGetLevel(t *testing.T) {
	type args struct {
		code MeshCode
	}
	tests := []struct {
		name    string
		args    args
		want    Level
		wantErr bool
	}{
		{name: "level1", args: args{code: "5339"}, want: Level1, wantErr: false},
		{name: "level2", args: args{code: "533945"}, want: Level2, wantErr: false},
		{name: "level3", args: args{code: "53394547"}, want: Level3, wantErr: false},
		{name: "level1-2", args: args{code: "533945471"}, want: LevelHalf, wantErr: false},
		{name: "level1-4", args: args{code: "5339454711"}, want: LevelQuarter, wantErr: false},
		{name: "level1-8", args: args{code: "53394547112"}, want: LevelOneEighth, wantErr: false},
		{name: "level1-8", args: args{code: "1"}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLevel(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLevel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCodes(t *testing.T) {
	lv2Codes := []MeshCode{
		"533900", "533901", "533902", "533903", "533904", "533905", "533906", "533907",
		"533910", "533911", "533912", "533913", "533914", "533915", "533916", "533917",
		"533920", "533921", "533922", "533923", "533924", "533925", "533926", "533927",
		"533930", "533931", "533932", "533933", "533934", "533935", "533936", "533937",
		"533940", "533941", "533942", "533943", "533944", "533945", "533946", "533947",
		"533950", "533951", "533952", "533953", "533954", "533955", "533956", "533957",
		"533960", "533961", "533962", "533963", "533964", "533965", "533966", "533967",
		"533970", "533971", "533972", "533973", "533974", "533975", "533976", "533977",
	}
	lv3Codes := []MeshCode{
		"53394500", "53394501", "53394502", "53394503", "53394504", "53394505", "53394506", "53394507", "53394508", "53394509",
		"53394510", "53394511", "53394512", "53394513", "53394514", "53394515", "53394516", "53394517", "53394518", "53394519",
		"53394520", "53394521", "53394522", "53394523", "53394524", "53394525", "53394526", "53394527", "53394528", "53394529",
		"53394530", "53394531", "53394532", "53394533", "53394534", "53394535", "53394536", "53394537", "53394538", "53394539",
		"53394540", "53394541", "53394542", "53394543", "53394544", "53394545", "53394546", "53394547", "53394548", "53394549",
		"53394550", "53394551", "53394552", "53394553", "53394554", "53394555", "53394556", "53394557", "53394558", "53394559",
		"53394560", "53394561", "53394562", "53394563", "53394564", "53394565", "53394566", "53394567", "53394568", "53394569",
		"53394570", "53394571", "53394572", "53394573", "53394574", "53394575", "53394576", "53394577", "53394578", "53394579",
		"53394580", "53394581", "53394582", "53394583", "53394584", "53394585", "53394586", "53394587", "53394588", "53394589",
		"53394590", "53394591", "53394592", "53394593", "53394594", "53394595", "53394596", "53394597", "53394598", "53394599",
	}
	lvHalfCodes := []MeshCode{
		"533945471", "533945472", "533945473", "533945474",
	}
	lvQuarterCodes := []MeshCode{
		"5339454711", "5339454712", "5339454713", "5339454714",
	}
	lvOneEightCodes := []MeshCode{
		"53394547111", "53394547112", "53394547113", "53394547114",
	}

	type args struct {
		code MeshCode
	}
	tests := []struct {
		name    string
		args    args
		want    MeshCodes
		wantErr bool
	}{
		{name: "level1->level2list", args: args{code: "5339"}, want: lv2Codes, wantErr: false},
		{name: "level2->level3list", args: args{code: "533945"}, want: lv3Codes, wantErr: false},
		{name: "level3->level1-2list", args: args{code: "53394547"}, want: lvHalfCodes, wantErr: false},
		{name: "level1-2->level1-4list", args: args{code: "533945471"}, want: lvQuarterCodes, wantErr: false},
		{name: "level1-4->level1-8list", args: args{code: "5339454711"}, want: lvOneEightCodes, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCodes(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitCodeByLevel(t *testing.T) {
	type args struct {
		code MeshCode
	}
	tests := []struct {
		name string
		args args
		want []MeshCode
	}{
		{name: "level1", args: args{code: "5339"}, want: []MeshCode{"5339"}},
		{name: "level2", args: args{code: "533945"}, want: []MeshCode{"5339", "533945"}},
		{name: "level3", args: args{code: "53394547"}, want: []MeshCode{"5339", "533945", "53394547"}},
		{name: "level1-2", args: args{code: "533945471"}, want: []MeshCode{"5339", "533945", "53394547", "533945471"}},
		{name: "level1-4", args: args{code: "5339454711"}, want: []MeshCode{"5339", "533945", "53394547", "533945471", "5339454711"}},
		{name: "level1-8", args: args{code: "53394547112"}, want: []MeshCode{"5339", "533945", "53394547", "533945471", "5339454711", "53394547112"}},
		{name: "level1-8", args: args{code: "1"}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitCodeByLevel(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitCodeByLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
