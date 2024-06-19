package simple

//go:generate enumeration -v -f -i season_json.go -type Season_Nc_Ji -suffix _Nc_Ji -s -marshaljson identifier
//go:generate enumeration -v -f -i season_json.go -type Season_Nc_Jn -suffix _Nc_Jn -s -marshaljson number
//go:generate enumeration -v -f -i season_json.go -type Season_Nc_Js -suffix _Nc_Js -s

type (
	Season_Nc_Ji uint
	Season_Nc_Jn uint
	Season_Nc_Js uint
)

const (
	Spring_Nc_Ji, Summer_Nc_Ji, Autumn_Nc_Ji, Winter_Nc_Ji Season_Nc_Ji = 1, 2, 3, 4
)

const (
	Spring_Nc_Jn, Summer_Nc_Jn, Autumn_Nc_Jn, Winter_Nc_Jn Season_Nc_Jn = 1, 2, 3, 4
)

const (
	_            Season_Nc_Js = iota
	Spring_Nc_Js              // json:"Sprg"
	Summer_Nc_Js              // json:"Sumr"
	Autumn_Nc_Js              // json:"Autm"
	Winter_Nc_Js              // json:"Wint"
)

//=================================================================================================
//go:generate enumeration -v -f -i season_json.go -type Season_Uc_Ji -suffix _Uc_Ji -uc -s -marshaljson identifier
//go:generate enumeration -v -f -i season_json.go -type Season_Uc_Jn -suffix _Uc_Jn -uc -s -marshaljson number
//go:generate enumeration -v -f -i season_json.go -type Season_Uc_Js -suffix _Uc_Js -uc -s

type (
	Season_Uc_Ji uint
	Season_Uc_Jn uint
	Season_Uc_Js uint
)

const (
	Spring_Uc_Ji, Summer_Uc_Ji, Autumn_Uc_Ji, Winter_Uc_Ji Season_Uc_Ji = 1, 2, 3, 4
)

const (
	_ Season_Uc_Jn = iota
	Spring_Uc_Jn
	Summer_Uc_Jn
	Autumn_Uc_Jn // 3
	Winter_Uc_Jn
)

const (
	_            Season_Uc_Js = iota
	Spring_Uc_Js              // json:"Sprg"
	Summer_Uc_Js              // json:"Sumr"
	Autumn_Uc_Js              // json:"Autm"
	Winter_Uc_Js              // json:"Wint"
)
