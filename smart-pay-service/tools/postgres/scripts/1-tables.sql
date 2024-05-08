-- Table Orcamento_trimestral
CREATE TABLE IF NOT EXISTS smartpay.orcamento_trimestral (
  idorcamento_trimestral SERIAL PRIMARY KEY,
  data_inicio TIMESTAMP NOT NULL,
  data_fim TIMESTAMP NOT NULL,
  orcamento_trimestral DOUBLE PRECISION NOT NULL
);

-- Table Centro_de_Custos
CREATE TABLE IF NOT EXISTS smartpay.centro_de_custos (
  idcentro_de_custos SERIAL PRIMARY KEY,
  nome_centro VARCHAR(45) NOT NULL,
  tipo VARCHAR(45) NOT NULL,
  orcamento_trimestral INT NOT NULL,
  fk_orcamento_trimestral INT NOT NULL,
  CONSTRAINT fk_centro_de_custos_orcamento_trimestral1
    FOREIGN KEY (fk_orcamento_trimestral)
    REFERENCES smartpay.orcamento_trimestral (idorcamento_trimestral)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

-- Table Executivo
CREATE TABLE IF NOT EXISTS smartpay.executivo (
  id_executivo SERIAL PRIMARY KEY,
  nome_executivo VARCHAR(45) NOT NULL,
  email VARCHAR(45) NOT NULL,
  area VARCHAR(45) NOT NULL,
  fk_centro_de_custos INT NOT NULL,
  CONSTRAINT fk_executivo_centro_de_custos1
    FOREIGN KEY (fk_centro_de_custos)
    REFERENCES smartpay.centro_de_custos (idcentro_de_custos)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

-- Table Funcionario
CREATE TABLE IF NOT EXISTS smartpay.funcionario (
  idfuncionarios SERIAL PRIMARY KEY,
  nome_funcionarios VARCHAR(45) NOT NULL,
  cargo VARCHAR(45) NOT NULL,
  senioridade VARCHAR(45) NOT NULL,
  salario DOUBLE PRECISION NOT NULL,
  fk_centro_de_custos INT NOT NULL,
  CONSTRAINT fk_funcionario_centro_de_custos1
    FOREIGN KEY (fk_centro_de_custos)
    REFERENCES smartpay.centro_de_custos (idcentro_de_custos)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

-- Table Area
CREATE TABLE IF NOT EXISTS smartpay.area (
  idarea SERIAL PRIMARY KEY,
  nome_area VARCHAR(45) NOT NULL,
  fk_centro_de_custos INT NOT NULL,
  CONSTRAINT fk_area_centro_de_custos1
    FOREIGN KEY (fk_centro_de_custos)
    REFERENCES smartpay.centro_de_custos (idcentro_de_custos)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

-- Table Gastos_variaveis
CREATE TABLE IF NOT EXISTS smartpay.gastos_variaveis (
  idgastos_variaveis SERIAL PRIMARY KEY,
  -- tipo_variavel VARCHAR(255),
  valor DOUBLE PRECISION NOT NULL,
  categoria_despesa VARCHAR(45) NOT NULL,
  desc_transacao VARCHAR(255),
  metodo_pagto VARCHAR(45),
  obs VARCHAR(255),
  data DATE NOT NULL,
  responsavel VARCHAR(100) NOT NULL,
  fk_centro_de_custos INT NOT NULL,
  CONSTRAINT fk_gastos_variaveis_centro_de_custos1
    FOREIGN KEY (fk_centro_de_custos)
    REFERENCES smartpay.centro_de_custos (idcentro_de_custos)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

