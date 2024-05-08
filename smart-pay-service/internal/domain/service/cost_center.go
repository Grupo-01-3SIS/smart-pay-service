package service

import (
	"context"
	"fmt"
	"os"
	"regexp"
	log_zap "smart-pay-service/config/log"
	"smart-pay-service/internal/domain/entity"
	"smart-pay-service/internal/domain/gateway"
	s3_client "smart-pay-service/internal/s3"

	"github.com/trimmer-io/go-csv"
	"go.uber.org/zap"
)

var _ CostCenter = (*CostCenterService)(nil)

type CostCenter interface {
	UnmarshalMonthlyCosts(context.Context) ([]*entity.MonthlyCosts, error)
	UnmarshalEmployeeMonthlyCosts(context.Context) ([]*entity.EmployeeMonthlyCosts, error)
	UnmarshalCostCenterInfo(context.Context) ([]*entity.CostCenterInfo, error)
	RunService(context.Context) error
}

type CostCenterService struct {
	log *zap.Logger
	gtw gateway.CostCenter
}

func NewCostCenterService(gtw gateway.CostCenter) *CostCenterService {
	return &CostCenterService{
		log: log_zap.NewLogger().Named("layer-service"),
		gtw: gtw,
	}
}

func (cc *CostCenterService) RunService(ctx context.Context) error {
	objCC, err := cc.UnmarshalCostCenterInfo(ctx)
	if err != nil {
		return err
	}

	cc.log.Info("indo para a camada do database inserir as informações do centro de custos")
	idCC, err := cc.gtw.InsertCoastCenter(objCC)
	if err != nil {
		return err
	}

	objEMC, err := cc.UnmarshalEmployeeMonthlyCosts(ctx)
	if err != nil {
		return err
	}

	cc.log.Info("indo para a camada do database inserir as informações dos gastos dos funcionarios")
	err = cc.gtw.InsertCoastEmployee(objEMC, idCC)
	if err != nil {
		return err
	}

	objMM, err := cc.UnmarshalMonthlyCosts(ctx)
	if err != nil {
		return err
	}

	cc.log.Info("indo para a camada do database inserir as informações dos gastos variaveis")
	err = cc.gtw.InsertCoastVariable(objMM, idCC)
	if err != nil {
		return err
	}

	return nil
}

// Metodo para fazer o Marshal do csv custos-mensais...csv
func (cc *CostCenterService) UnmarshalMonthlyCosts(context.Context) ([]*entity.MonthlyCosts, error) {
	// Diretório onde estão os arquivos
	directory := fmt.Sprintf("%s/", s3_client.Path)

	// Regex para corresponder ao nome do arquivo
	fileRegex := regexp.MustCompile(`^custos-mensais.*\.csv$`)

	// Abrir o diretório
	dir, err := os.ReadDir(directory)
	if err != nil {
		cc.log.Error(err.Error())
		return nil, err
	}

	// Iterar sobre os arquivos e ler o que corresponde ao regex
	for _, file := range dir {
		if fileRegex.MatchString(file.Name()) {
			// Construir o caminho completo do arquivo
			filePath := fmt.Sprintf("%s/%s", directory, file.Name())

			// Ler o conteúdo do arquivo
			b, err := os.ReadFile(filePath)
			if err != nil {
				cc.log.Error(err.Error())
				return nil, err
			}

			// Processar o conteúdo do arquivo (por exemplo, fazer o unmarshal do CSV)
			monthlyCosts := make([]*entity.MonthlyCosts, 0)
			if err := csv.Unmarshal(b, &monthlyCosts); err != nil {
				cc.log.Error(err.Error())
				return nil, err
			}

			return monthlyCosts, nil
		}
	}

	// Se nenhum arquivo correspondente for encontrado, retorna um erro
	err = fmt.Errorf("Nenhum arquivo correspondente encontrado no diretório %s", directory)
	cc.log.Error(err.Error())
	return nil, err
}

// Metodo para fazer o Marshal do csv custos-funcionarios...csv
func (cc *CostCenterService) UnmarshalEmployeeMonthlyCosts(context.Context) ([]*entity.EmployeeMonthlyCosts, error) {
	// Diretório onde estão os arquivos
	directory := fmt.Sprintf("%s/", s3_client.Path)

	// Regex para corresponder ao nome do arquivo
	fileRegex := regexp.MustCompile(`^custos-funcionarios.*\.csv$`)

	// Abrir o diretório
	dir, err := os.ReadDir(directory)
	if err != nil {
		cc.log.Error(err.Error())
		return nil, err
	}

	// Iterar sobre os arquivos e ler o que corresponde ao regex
	for _, file := range dir {
		if fileRegex.MatchString(file.Name()) {
			// Construir o caminho completo do arquivo
			filePath := fmt.Sprintf("%s/%s", directory, file.Name())

			// Ler o conteúdo do arquivo
			b, err := os.ReadFile(filePath)
			if err != nil {
				cc.log.Error(err.Error())
				return nil, err
			}

			// Processar o conteúdo do arquivo (por exemplo, fazer o unmarshal do CSV)
			employeeMonthlyCosts := make([]*entity.EmployeeMonthlyCosts, 0)
			if err := csv.Unmarshal(b, &employeeMonthlyCosts); err != nil {
				cc.log.Error(err.Error())
				return nil, err
			}

			return employeeMonthlyCosts, nil
		}
	}

	// Se nenhum arquivo correspondente for encontrado, retorna um erro
	err = fmt.Errorf("Nenhum arquivo correspondente encontrado no diretório %s", directory)
	cc.log.Error(err.Error())
	return nil, err
}

// Metodo para fazer o Marshal do csv centro-de-custos...csv
func (cc *CostCenterService) UnmarshalCostCenterInfo(context.Context) ([]*entity.CostCenterInfo, error) {
	// Diretório onde estão os arquivos
	directory := fmt.Sprintf("%s/", s3_client.Path)

	// Regex para corresponder ao nome do arquivo
	fileRegex := regexp.MustCompile(`^centro-de-custos.*\.csv$`)

	// Abrir o diretório
	dir, err := os.ReadDir(directory)
	if err != nil {
		cc.log.Error(err.Error())
		return nil, err
	}

	// Iterar sobre os arquivos e ler o que corresponde ao regex
	for _, file := range dir {
		if fileRegex.MatchString(file.Name()) {
			// Construir o caminho completo do arquivo
			filePath := fmt.Sprintf("%s/%s", directory, file.Name())

			// Ler o conteúdo do arquivo
			b, err := os.ReadFile(filePath)
			if err != nil {
				cc.log.Error(err.Error())
				return nil, err
			}

			// Processar o conteúdo do arquivo (por exemplo, fazer o unmarshal do CSV)
			costCenterInfo := make([]*entity.CostCenterInfo, 0)
			if err := csv.Unmarshal(b, &costCenterInfo); err != nil {
				cc.log.Error(err.Error())
				return nil, err
			}

			return costCenterInfo, nil
		}
	}

	// Se nenhum arquivo correspondente for encontrado, retorna um erro
	err = fmt.Errorf("Nenhum arquivo correspondente encontrado no diretório %s", directory)
	cc.log.Error(err.Error())
	return nil, err
}

// Metodo generico que faz o marshal dos 3 Csv
func (cc *CostCenterService) UnmarshalCSVData(ctx context.Context, filenamePattern string) (interface{}, error) {
	// Diretório onde estão os arquivos
	directory := fmt.Sprintf("%s/", s3_client.Path)

	// Regex para corresponder ao nome do arquivo
	fileRegex := regexp.MustCompile(filenamePattern)

	// Abrir o diretório
	dir, err := os.ReadDir(directory)
	if err != nil {
		cc.log.Error(err.Error())
		return nil, err
	}

	// Iterar sobre os arquivos e ler o que corresponde ao regex
	for _, file := range dir {
		if fileRegex.MatchString(file.Name()) {
			// Construir o caminho completo do arquivo
			filePath := fmt.Sprintf("%s/%s", directory, file.Name())

			// Ler o conteúdo do arquivo
			b, err := os.ReadFile(filePath)
			if err != nil {
				cc.log.Error(err.Error())
				return nil, err
			}

			// Processar o conteúdo do arquivo (por exemplo, fazer o unmarshal do CSV)
			var data interface{}
			switch filenamePattern {
			case `^custos-mensais.*\.csv$`:
				monthlyCosts := make([]*entity.MonthlyCosts, 0)
				if err := csv.Unmarshal(b, &monthlyCosts); err != nil {
					cc.log.Error(err.Error())
					return nil, err
				}
				data = monthlyCosts
			case `^custos-funcionarios.*\.csv$`:
				employeeMonthlyCosts := make([]*entity.EmployeeMonthlyCosts, 0)
				if err := csv.Unmarshal(b, &employeeMonthlyCosts); err != nil {
					cc.log.Error(err.Error())
					return nil, err
				}
				data = employeeMonthlyCosts
			case `^centro-de-custos.*\.csv$`:
				costCenterInfo := make([]*entity.CostCenterInfo, 0)
				if err := csv.Unmarshal(b, &costCenterInfo); err != nil {
					cc.log.Error(err.Error())
					return nil, err
				}
				data = costCenterInfo
			default:
				err = fmt.Errorf("Padrão de nome de arquivo não reconhecido: %s", filenamePattern)
				cc.log.Error(err.Error())
				return nil, err
			}

			return data, nil
		}
	}

	// Se nenhum arquivo correspondente for encontrado, retorna um erro
	err = fmt.Errorf("Nenhum arquivo correspondente encontrado no diretório %s", directory)
	cc.log.Error(err.Error())
	return nil, err
}
