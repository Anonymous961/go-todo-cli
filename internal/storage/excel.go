package storage

import (
	"fmt"
	"strconv"
	"time"

	"github.com/anonymous961/todo-cli/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelStorage struct {
	filePath string
}

func NewExcelStorage(filePath string) *ExcelStorage {
	return &ExcelStorage{filePath: filePath}
}

func (s *ExcelStorage) ensureFileExists() error {
	f, err := excelize.OpenFile(s.filePath)
	if err != nil {
		f := excelize.NewFile()
		f.NewSheet("Todos")
		f.DeleteSheet("Sheet1")
		fmt.Printf("New sheet created")

		headers := []string{"ID", "Task", "Complete", "Category", "Due Date", "Priority", "Created At"}

		for i, header := range headers {
			cell, _ := excelize.CoordinatesToCellName(i+1, i)
			f.SetCellValue("Todos", cell, header)
		}
		return f.SaveAs(s.filePath)
	}

	defer f.Close()

	if Index, _ := f.GetSheetIndex("Todos"); Index == -1 {
		f.NewSheet("Todos")
		return f.Save()
	}
	return nil
}

func (s *ExcelStorage) Add(todo *models.Todo) error {
	if err := s.ensureFileExists(); err != nil {
		return err
	}
	f, err := excelize.OpenFile(s.filePath)

	if err != nil {
		return err
	}
	defer f.Close()

	rows, err := f.GetRows("Todos")
	if err != nil {
		return err
	}
	// fmt.Printf("The number of rows are %v", rows)

	rowNum := len(rows) + 1

	f.SetCellValue("Todos", "A"+strconv.Itoa(rowNum), todo.ID)
	f.SetCellValue("Todos", "B"+strconv.Itoa(rowNum), todo.Task)
	f.SetCellValue("Todos", "C"+strconv.Itoa(rowNum), todo.Complete)
	f.SetCellValue("Todos", "D"+strconv.Itoa(rowNum), todo.Category)
	f.SetCellValue("Todos", "E"+strconv.Itoa(rowNum), todo.DueDate.Format(time.RFC3339))
	f.SetCellValue("Todos", "F"+strconv.Itoa(rowNum), todo.Priority)
	f.SetCellValue("Todos", "G"+strconv.Itoa(rowNum), todo.CreatedAt.Format(time.RFC3339))

	return f.Save()
}

func (s *ExcelStorage) List() ([]*models.Todo, error) {
	if err := s.ensureFileExists(); err != nil {
		return nil, err
	}

	f, err := excelize.OpenFile(s.filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Todos")

	if err != nil {
		return nil, err
	}

	var todos []*models.Todo

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 7 {
			continue
		}

		dueDate, _ := time.Parse(time.RFC3339, row[4])
		createdAt, _ := time.Parse(time.RFC3339, row[6])
		complete, _ := strconv.ParseBool(row[2])
		priority, _ := strconv.Atoi(row[5])

		todos = append(todos, &models.Todo{
			ID:        row[0],
			Task:      row[1],
			Complete:  complete,
			Category:  row[3],
			DueDate:   dueDate,
			Priority:  priority,
			CreatedAt: createdAt,
		})
	}

	return todos, nil
}
