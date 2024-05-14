package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/f1rdavsi/reporter/internal/repository"
	"github.com/f1rdavsi/reporter/models"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

type ReportService struct {
	repo repository.Report
}

func NewReportService(repo repository.Report) *ReportService {
	return &ReportService{
		repo: repo,
	}
}

func (s *ReportService) GetReport(
	FromID, ToID int, ToType string,
	Limit, Page int, Type string,
	From, To time.Time,
) ([]models.Transaction, error) {
	var rep models.Report

	rep.FromID = FromID
	rep.ToID = ToID
	rep.ToType = ToType
	rep.Limit = Limit
	rep.Page = Page
	rep.Type = Type
	rep.From = From
	rep.To = To

	return s.repo.GetReport(rep)
}

func (s *ReportService) GetCSVReport(
	userFrom, userTo models.User,
	accountFrom, accountTo models.Account,
	tr []models.Transaction,
) (*os.File, error) {
	timeSign := fmt.Sprintf("%d", time.Now().UnixNano())
	filePath := fmt.Sprintf("%s_%s", timeSign, "report")
	filePath = strings.Replace(filePath, " ", "", 111)
	path := fmt.Sprintf("./files/reports/%s.csv", filePath)

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"Имя пользователя (отправитель)",
		"Название аккаунта (отправитель)",
		"Имя пользователя (получатель)",
		"Тип операции",
		"Комментарий",
		"Сумма",
		"Дата совершения операции",
	}

	err = writer.Write(header)
	if err != nil {
		return nil, err
	}

	for _, transaction := range tr {
		row := []string{
			userFrom.Username,
			accountFrom.Name,
			userTo.Username,
			accountTo.Name,
			transaction.Comment,
			cast.ToString(transaction.Amount),
			transaction.CreatedAt.String(),
		}
		err = writer.Write(row)
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

func (s *ReportService) GetExcelReport(
	userFrom, userTo models.User,
	accountFrom, accountTo models.Account,
	tr []models.Transaction,
) (*excelize.File, error) {
	excelFile := excelize.NewFile()

	sheet, err := excelFile.NewSheet("Отчёт")
	if err != nil {
		return nil, err
	}

	err = excelFile.SetCellValue("Отчёт", "A1", "Имя пользователя (отправитель)")
	if err != nil {
		return nil, err
	}

	err = excelFile.SetCellValue("Отчёт", "B1", "Название аккаунта (отправитель)")
	if err != nil {
		return nil, err
	}

	err = excelFile.SetCellValue("Отчёт", "C1", "Имя пользователя (получатель)")
	if err != nil {
		return nil, err
	}

	err = excelFile.SetCellValue("Отчёт", "D1", "Тип операции")
	if err != nil {
		return nil, err
	}

	err = excelFile.SetCellValue("Отчёт", "E1", "Комментарий")
	if err != nil {
		return nil, err
	}

	err = excelFile.SetCellValue("Отчёт", "F1", "Сумма")
	if err != nil {
		return nil, err
	}

	err = excelFile.SetCellValue("Отчёт", "G1", "Дата совершения операции")
	if err != nil {
		return nil, err
	}

	for i, transaction := range tr {
		i += 2

		err = excelFile.SetCellValue("Отчёт", "A"+strconv.Itoa(i), userFrom.Username)
		if err != nil {
			return nil, err
		}

		err = excelFile.SetCellValue("Отчёт", "B"+strconv.Itoa(i), accountFrom.Name)
		if err != nil {
			return nil, err
		}

		err = excelFile.SetCellValue("Отчёт", "C"+strconv.Itoa(i), userTo.Username)
		if err != nil {
			return nil, err
		}

		err = excelFile.SetCellValue("Отчёт", "D"+strconv.Itoa(i), accountTo.Name)
		if err != nil {
			return nil, err
		}

		err = excelFile.SetCellValue("Отчёт", "E"+strconv.Itoa(i), transaction.Comment)
		if err != nil {
			return nil, err
		}

		err = excelFile.SetCellValue("Отчёт", "F"+strconv.Itoa(i), transaction.Amount)
		if err != nil {
			return nil, err
		}

		err = excelFile.SetCellValue("Отчёт", "G"+strconv.Itoa(i), transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	excelFile.SetActiveSheet(sheet)

	timeSign := fmt.Sprintf("%d", time.Now().UnixNano())
	filePath := fmt.Sprintf("%s_%s", timeSign, "report")
	filePath = strings.Replace(filePath, " ", "", 111)

	err = excelFile.SaveAs(fmt.Sprintf("./files/reports/%s.xlsx", filePath))
	if err != nil {
		return nil, err
	}

	return excelFile, nil
}
