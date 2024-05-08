package entity

type Executivo struct {
	Id_Executivo   int    `db:"Id_Executivo"`
	Nome_Executivo string `db:"nome_Executivo"`
	Email          string `db:"email"`
	Area           string `db:"area"`
}

type Area struct {
	IdArea    int    `db:"idArea"`
	Nome_Area string `db:"nome_Area"`
}

type Funcionario struct {
	IdFuncionarios    int     `db:"idFuncionarios"`
	Nome_Funcionarios string  `db:"nome_funcionarios"`
	Cargo             string  `db:"cargo"`
	Senioridade       string  `db:"senioridade"`
	Salario           float64 `db:"salario"`
}

type OrcamentoTrimestral struct {
	Idorcamento_Trimestral int     `db:"idorcamento_trimestral"`
	Data_Inicio            string  `db:"data_inicio"`
	Data_Fim               string  `db:"data_fim"`
	Orcamento_Trimestral   float64 `db:"orcamento_trimestral"`
}

type CentroDeCustos struct {
	IdCentro_de_Custos      int    `db:"idCentro_de_Custos"`
	Nome_Centro             string `db:"nome_Centro"`
	Tipo                    string `db:"tipo"`
	Orcamento_Trimestral    int    `db:"orcamento_Trimestral"`
	Fk_Executivo            int    `db:"fk_executivo"`
	Fk_Area                 int    `db:"fk_Area"`
	Fk_Funcionario          int    `db:"fk_funcionario"`
	Fk_Orcamento_Trimestral int    `db:"fk_orcamento_trimestral"`
}

type GastosVariaveis struct {
	IdGastos_variaveis int `db:"idGastos_variaveis"`
	//Tipo_Variavel      string `db:"tipo_variavel"`
	Valor             string `db:"valor"`
	Categoria_Despesa string `db:"categoria_despesa"`
	Desc_Transacao    string `db:"desc_transacao"`
	Metodo_Pagto      string `db:"metodo_pagto"`
	Obs               string `db:"obs"`
	Data              string `db:"data"`
	Fk_Funcionarios   int    `db:"fk_funcionarios"`
}
