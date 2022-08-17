package test

//=================================================================================================
//go:generate enumeration -v -f -i season_text.go -type Season_Nc_Ti -suffix _Nc_Ti -marshaltext identifier

type Season_Nc_Ti uint

const (
	Spring_Nc_Ti, Summer_Nc_Ti, Autumn_Nc_Ti, Winter_Nc_Ti Season_Nc_Ti = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Nc_Tn -suffix _Nc_Tn -marshaltext number

type Season_Nc_Tn uint

const (
	Spring_Nc_Tn, Summer_Nc_Tn, Autumn_Nc_Tn, Winter_Nc_Tn Season_Nc_Tn = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Nc_To -suffix _Nc_To -marshaltext ordinal

type Season_Nc_To uint

const (
	Spring_Nc_To, Summer_Nc_To, Autumn_Nc_To, Winter_Nc_To Season_Nc_To = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Nc_Tt -suffix _Nc_Tt

type Season_Nc_Tt uint

const (
	_            Season_Nc_Tt = iota
	Spring_Nc_Tt              // text:"Sprg"
	Summer_Nc_Tt              // text:"Sumr"
	Autumn_Nc_Tt              // text:"Autm"
	Winter_Nc_Tt              // text:"Wint"
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Nc_Ta -suffix _Nc_Ta

type Season_Nc_Ta uint

const (
	_            Season_Nc_Ta = iota
	Spring_Nc_Ta              // all:"Sprg"
	Summer_Nc_Ta              // all:"Sumr"
	Autumn_Nc_Ta              // all:"Autm"
	Winter_Nc_Ta              // all:"Wint"
)

//=================================================================================================
//go:generate enumeration -v -f -i season_text.go -type Season_Ic_Ti -suffix _Ic_Ti -ic -marshaltext identifier

type Season_Ic_Ti uint

const (
	Spring_Ic_Ti, Summer_Ic_Ti, Autumn_Ic_Ti, Winter_Ic_Ti Season_Ic_Ti = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Ic_Tn -suffix _Ic_Tn -ic -marshaltext number

type Season_Ic_Tn uint

const (
	Spring_Ic_Tn, Summer_Ic_Tn, Autumn_Ic_Tn, Winter_Ic_Tn Season_Ic_Tn = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Ic_To -suffix _Ic_To -ic -marshaltext ordinal

type Season_Ic_To uint

const (
	Spring_Ic_To, Summer_Ic_To, Autumn_Ic_To, Winter_Ic_To Season_Ic_To = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Ic_Tt -suffix _Ic_Tt -ic

type Season_Ic_Tt uint

const (
	_            Season_Ic_Tt = iota
	Spring_Ic_Tt              // text:"Sprg"
	Summer_Ic_Tt              // text:"Sumr"
	Autumn_Ic_Tt              // text:"Autm"
	Winter_Ic_Tt              // text:"Wint"
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Ic_Ta -suffix _Ic_Ta -ic

type Season_Ic_Ta uint

const (
	_            Season_Ic_Ta = iota
	Spring_Ic_Ta              // all:"Sprg"
	Summer_Ic_Ta              // all:"Sumr"
	Autumn_Ic_Ta              // all:"Autm"
	Winter_Ic_Ta              // all:"Wint"
)

//=================================================================================================
//go:generate enumeration -v -f -i season_text.go -type Season_Uc_Ti -suffix _Uc_Ti -uc -marshaltext identifier

type Season_Uc_Ti uint

const (
	Spring_Uc_Ti, Summer_Uc_Ti, Autumn_Uc_Ti, Winter_Uc_Ti Season_Uc_Ti = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Uc_Tn -suffix _Uc_Tn -uc -marshaltext number

type Season_Uc_Tn uint

const (
	Spring_Uc_Tn, Summer_Uc_Tn, Autumn_Uc_Tn, Winter_Uc_Tn Season_Uc_Tn = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_text.go -type Season_Uc_To -suffix _Uc_To -uc -marshaltext ordinal

type Season_Uc_To uint

const (
	Spring_Uc_To, Summer_Uc_To, Autumn_Uc_To, Winter_Uc_To Season_Uc_To = 1, 2, 3, 4
)
