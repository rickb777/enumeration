package simple

//go:generate enumeration -v -f -i season_sql.go -type Season_Nc_Si -suffix _Nc_Si -s -store identifier
//go:generate enumeration -v -f -i season_sql.go -type Season_Nc_Sn -suffix _Nc_Sn -s -store number
//go:generate enumeration -v -f -i season_sql.go -type Season_Nc_Ss -suffix _Nc_Ss -s

type (
	Season_Nc_Si uint
	Season_Nc_Sn uint
	Season_Nc_Ss uint
)

const (
	Spring_Nc_Si, Summer_Nc_Si, Autumn_Nc_Si, Winter_Nc_Si Season_Nc_Si = 1, 2, 3, 4
)

const (
	Spring_Nc_Sn, Summer_Nc_Sn, Autumn_Nc_Sn, Winter_Nc_Sn Season_Nc_Sn = 1, 2, 3, 4
)

const (
	_            Season_Nc_Ss = iota
	Spring_Nc_Ss              // sql:"Sprg"
	Summer_Nc_Ss              // sql:"Sumr"
	Autumn_Nc_Ss              // sql:"Autm"
	Winter_Nc_Ss              // sql:"Wint"
)

//=================================================================================================
//go:generate enumeration -v -f -i season_sql.go -type Season_Uc_Si -suffix _Uc_Si -uc -s -store identifier
//go:generate enumeration -v -f -i season_sql.go -type Season_Uc_Sn -suffix _Uc_Sn -uc -s -store number

type (
	Season_Uc_Si uint
	Season_Uc_Sn uint
)

const (
	Spring_Uc_Si, Summer_Uc_Si, Autumn_Uc_Si, Winter_Uc_Si Season_Uc_Si = 1, 2, 3, 4
)

const (
	Spring_Uc_Sn, Summer_Uc_Sn, Autumn_Uc_Sn, Winter_Uc_Sn Season_Uc_Sn = 1, 2, 3, 4
)
