-- Table Executivo
CREATE TABLE IF NOT EXISTS smartpay.executivo (
  id_executivo SERIAL PRIMARY KEY,
  nome_executivo VARCHAR(45) NOT NULL,
  email VARCHAR(45) NOT NULL,
  area VARCHAR(45) NOT NULL
);

-- Table Area
CREATE TABLE IF NOT EXISTS smartpay.area (
  id_area SERIAL PRIMARY KEY,
  nome_area VARCHAR(45) NOT NULL
);

-- Table Funcionario
CREATE TABLE IF NOT EXISTS smartpay.funcionario (
  id_funcionario SERIAL PRIMARY KEY,
  nome_funcionario VARCHAR(45) NOT NULL,
  cargo VARCHAR(45) NOT NULL,
  senioridade VARCHAR(45) NOT NULL,
  salario FLOAT NOT NULL,
  vt FLOAT NOT NULL,
  vr FLOAT NOT NULL
);

-- Table Orcamento_trimestral
CREATE TABLE IF NOT EXISTS smartpay.orcamento_trimestral (
  id_orcamento_trimestral SERIAL PRIMARY KEY,
  data_inicio TIMESTAMP NOT NULL,
  data_fim TIMESTAMP NOT NULL,
  orcamento_trimestral FLOAT NOT NULL
);

-- Table Centro_de_Custos
CREATE TABLE IF NOT EXISTS smartpay.centro_de_custos (
  id_centro_de_custos SERIAL PRIMARY KEY,
  nome_centro VARCHAR(45) NOT NULL,
  tipo VARCHAR(45) NOT NULL,
  orcamento_trimestral INT NOT NULL,
  fk_executivo INT NOT NULL,
  fk_area INT NOT NULL,
  fk_funcionario INT NOT NULL,
  fk_orcamento_trimestral INT NOT NULL,
  CONSTRAINT fk_centro_de_custos_executivo
    FOREIGN KEY (fk_executivo)
    REFERENCES smartpay.executivo (id_executivo)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT fk_centro_de_custos_area
    FOREIGN KEY (fk_area)
    REFERENCES smartpay.area (id_area)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT fk_centro_de_custos_funcionario
    FOREIGN KEY (fk_funcionario)
    REFERENCES smartpay.funcionario (id_funcionario)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT fk_centro_de_custos_orcamento_trimestral
    FOREIGN KEY (fk_orcamento_trimestral)
    REFERENCES smartpay.orcamento_trimestral (id_orcamento_trimestral)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

-- Table Gastos_variaveis
CREATE TABLE IF NOT EXISTS smartpay.gastos_variaveis (
  id_gastos_variaveis SERIAL PRIMARY KEY,
  tipo_variavel VARCHAR(45) NOT NULL,
  valor FLOAT NOT NULL,
  categoria_despesa VARCHAR(45) NOT NULL,
  desc_transacao VARCHAR(45),
  metodo_pagto VARCHAR(45),
  obs VARCHAR(45),
  data TIMESTAMP NOT NULL,
  fk_funcionario INT NOT NULL,
  CONSTRAINT fk_gastos_variaveis_funcionario
    FOREIGN KEY (fk_funcionario)
    REFERENCES smartpay.funcionario (id_funcionario)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);
