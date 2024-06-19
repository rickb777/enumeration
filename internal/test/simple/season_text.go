package simple

//go:generate enumeration -v -f -i season_text.go -type Season_Nc_Ti -suffix _Nc_Ti -s -marshaltext identifier
//go:generate enumeration -v -f -i season_text.go -type Season_Nc_Tn -suffix _Nc_Tn -s -marshaltext number
//go:generate enumeration -v -f -i season_text.go -type Season_Nc_Ts -suffix _Nc_Ts -s

type (
	Season_Nc_Ti uint
	Season_Nc_Tn uint
	Season_Nc_Ts uint
)

const (
	Spring_Nc_Ti, Summer_Nc_Ti, Autumn_Nc_Ti, Winter_Nc_Ti Season_Nc_Ti = 1, 2, 3, 4
)

const (
	Spring_Nc_Tn, Summer_Nc_Tn, Autumn_Nc_Tn, Winter_Nc_Tn Season_Nc_Tn = 1, 2, 3, 4
)

const (
	_            Season_Nc_Ts = iota
	Spring_Nc_Tt              // text:"Sprg"
	Summer_Nc_Tt              // text:"Sumr"
	Autumn_Nc_Tt              // text:"Autm"
	Winter_Nc_Tt              // text:"Wint"
)

//=================================================================================================
//go:generate enumeration -v -f -i season_text.go -type Season_Uc_Ti -suffix _Uc_Ti -uc -s -marshaltext identifier
//go:generate enumeration -v -f -i season_text.go -type Season_Uc_Tn -suffix _Uc_Tn -uc -s -marshaltext number
//go:generate enumeration -v -f -i season_text.go -type Season_Uc_Ts -suffix _Uc_Ts -uc -s

type (
	Season_Uc_Ti uint
	Season_Uc_Tn uint
	Season_Uc_Ts uint
)

const (
	Spring_Uc_Ti, Summer_Uc_Ti, Autumn_Uc_Ti, Winter_Uc_Ti Season_Uc_Ti = 1, 2, 3, 4
)

const (
	Spring_Uc_Tn, Summer_Uc_Tn, Autumn_Uc_Tn, Winter_Uc_Tn Season_Uc_Tn = 1, 2, 3, 4
)

const (
	_            Season_Uc_Ts = iota
	Spring_Uc_Tt              // text:"Sprg"
	Summer_Uc_Tt              // text:"Sumr"
	Autumn_Uc_Tt              // text:"Autm"
	Winter_Uc_Tt              // text:"Wint"
)
