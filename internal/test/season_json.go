package test

//=================================================================================================
//go:generate enumeration -v -f -i season_json.go -type Season_Nc_Ji -suffix _Nc_Ji -marshaljson identifier

type Season_Nc_Ji uint

const (
	Spring_Nc_Ji, Summer_Nc_Ji, Autumn_Nc_Ji, Winter_Nc_Ji Season_Nc_Ji = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_json.go -type Season_Nc_Jn -suffix _Nc_Jn -marshaljson number

type Season_Nc_Jn uint

const (
	Spring_Nc_Jn, Summer_Nc_Jn, Autumn_Nc_Jn, Winter_Nc_Jn Season_Nc_Jn = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_json.go -type Season_Nc_Jj -suffix _Nc_Jj

type Season_Nc_Jj uint

const (
	_            Season_Nc_Jj = iota
	Spring_Nc_Jj              // json:"Sprg"
	Summer_Nc_Jj              // json:"Sumr"
	Autumn_Nc_Jj              // json:"Autm"
	Winter_Nc_Jj              // json:"Wint"
)

//=================================================================================================
//go:generate enumeration -v -f -i season_json.go -type Season_Ic_Ji -suffix _Ic_Ji -ic -marshaljson identifier

type Season_Ic_Ji uint

const (
	Spring_Ic_Ji, Summer_Ic_Ji, Autumn_Ic_Ji, Winter_Ic_Ji Season_Ic_Ji = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_json.go -type Season_Ic_Jn -suffix _Ic_Jn -ic -marshaljson number

type Season_Ic_Jn uint

const (
	Spring_Ic_Jn, Summer_Ic_Jn, Autumn_Ic_Jn, Winter_Ic_Jn Season_Ic_Jn = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_json.go -type Season_Ic_Jj -suffix _Ic_Jj -ic

type Season_Ic_Jj uint

const (
	_            Season_Ic_Jj = iota
	Spring_Ic_Jj              // json:"Sprg"
	Summer_Ic_Jj              // json:"Sumr"
	Autumn_Ic_Jj              // json:"Autm"
	Winter_Ic_Jj              // json:"Wint"
)

//=================================================================================================
//go:generate enumeration -v -f -i season_json.go -type Season_Uc_Ji -suffix _Uc_Ji -uc -marshaljson identifier

type Season_Uc_Ji uint

const (
	Spring_Uc_Ji, Summer_Uc_Ji, Autumn_Uc_Ji, Winter_Uc_Ji Season_Uc_Ji = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_json.go -type Season_Uc_Jn -suffix _Uc_Jn -uc -marshaljson number -z

type Season_Uc_Jn uint

const (
	_ Season_Uc_Jn = iota
	Spring_Uc_Jn
	Summer_Uc_Jn
	Autumn_Uc_Jn // 3
	Winter_Uc_Jn
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_json.go -type Season_Uc_Jj -suffix _Uc_Jj -uc

type Season_Uc_Jj uint

const (
	_            Season_Uc_Jj = iota
	Spring_Uc_Jj              // json:"Sprg"
	Summer_Uc_Jj              // json:"Sumr"
	Autumn_Uc_Jj              // json:"Autm"
	Winter_Uc_Jj              // json:"Wint"
)
