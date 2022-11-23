package test

//go:generate enumeration -v -f -i season_sql.go -type Season_Nc_Si -suffix _Nc_Si -store identifier

type Season_Nc_Si uint

const (
	Spring_Nc_Si, Summer_Nc_Si, Autumn_Nc_Si, Winter_Nc_Si Season_Nc_Si = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_sql.go -type Season_Nc_Sn -suffix _Nc_Sn -store number

type Season_Nc_Sn uint

const (
	Spring_Nc_Sn, Summer_Nc_Sn, Autumn_Nc_Sn, Winter_Nc_Sn Season_Nc_Sn = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_sql.go -type Season_Nc_Ss -suffix _Nc_Ss

type Season_Nc_Ss uint

const (
	_            Season_Nc_Ss = iota
	Spring_Nc_Ss              // sql:"Sprg"
	Summer_Nc_Ss              // sql:"Sumr"
	Autumn_Nc_Ss              // sql:"Autm"
	Winter_Nc_Ss              // sql:"Wint"
)

//=================================================================================================
//go:generate enumeration -v -f -i season_sql.go -type Season_Ic_Si -suffix _Ic_Si -ic -store identifier

type Season_Ic_Si uint

const (
	Spring_Ic_Si, Summer_Ic_Si, Autumn_Ic_Si, Winter_Ic_Si Season_Ic_Si = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_sql.go -type Season_Ic_Sn -suffix _Ic_Sn -ic -store number

type Season_Ic_Sn uint

const (
	Spring_Ic_Sn, Summer_Ic_Sn, Autumn_Ic_Sn, Winter_Ic_Sn Season_Ic_Sn = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_sql.go -type Season_Ic_Ss -suffix _Ic_Ss -ic

type Season_Ic_Ss uint

const (
	_            Season_Ic_Ss = iota
	Spring_Ic_Ss              // sql:"Sprg"
	Summer_Ic_Ss              // sql:"Sumr"
	Autumn_Ic_Ss              // sql:"Autm"
	Winter_Ic_Ss              // sql:"Wint"
)

//=================================================================================================
//go:generate enumeration -v -f -i season_sql.go -type Season_Uc_Si -suffix _Uc_Si -uc -store identifier

type Season_Uc_Si uint

const (
	Spring_Uc_Si, Summer_Uc_Si, Autumn_Uc_Si, Winter_Uc_Si Season_Uc_Si = 1, 2, 3, 4
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season_sql.go -type Season_Uc_Sn -suffix _Uc_Sn -uc -store number

type Season_Uc_Sn uint

const (
	Spring_Uc_Sn, Summer_Uc_Sn, Autumn_Uc_Sn, Winter_Uc_Sn Season_Uc_Sn = 1, 2, 3, 4
)
