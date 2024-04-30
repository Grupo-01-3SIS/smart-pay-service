package entity

const filePatternCCI = `^centro-de-custos.*\.csv$`

type CostCenterInfo struct {
	AreaName       string     `csv:"nome_Area"`
	ExecutiveName  string     `csv:"nome_Executivo"`
	Email          string     `csv:"email"`
	TypeCostCenter string     `csv:"tipo_Cc"`
	Value          int        `csv:"orcamento_tri"`
	StartDate      DateFormat `csv:"dat_inicio_orcamento"`
	EndDate        DateFormat `csv:"dat_fim_orcamento"`
}
