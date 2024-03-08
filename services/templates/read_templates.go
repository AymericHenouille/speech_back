package templates

import (
	"database/sql"
	"math"
	"speechs/api/api_controller"
	"speechs/databases"
)

func normalizeLimitOffset(limit, offset, total int) (int, int) {
	offset = int(math.Min(float64(offset), math.Max(float64(total-1), 0)))
	limit = int(math.Min(float64(limit), math.Max(float64(total-offset), 0)))
	return limit, offset
}

func readTemplates(db *sql.DB, limit, offset int) ([]Template, error) {
	rows, err := databases.QueryFile(db, "scripts/templates/read_templates.sql", limit, offset)
	if err != nil {
		return []Template{}, err
	}

	templates := make([]Template, limit)
	for index := 0; rows.Next(); index++ {
		template := Template{}
		err = rows.Scan(&template.Id, &template.Name, &template.Description, &template.Content)
		if err != nil {
			return templates, err
		}
		templates[index] = template
	}

	return templates, nil
}

func ReadTemplates(db *sql.DB, limit, offset int) (api_controller.ReadResponse[Template], error) {
	count, err := ReadTemplateCount(db)
	if err != nil {
		return api_controller.ReadResponse[Template]{}, err
	}

	limit, offset = normalizeLimitOffset(limit, offset, count)
	templates, err := readTemplates(db, limit, offset)
	if err != nil {
		return api_controller.ReadResponse[Template]{}, err
	}

	return api_controller.ReadResponse[Template]{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Datas:  templates,
	}, nil
}

func ReadTemplateByIds(db *sql.DB, ids []int) ([]Template, error) {
	rows, err := databases.QueryFile(db, "scripts/templates/read_templates_by_id", ids)
	if err != nil {
		return []Template{}, err
	}
	var templates = make([]Template, len(ids))
	for index := 0; rows.Next(); index++ {
		template := Template{}
		err = rows.Scan(&template.Id, &template.Name, &template.Description, &template.Content)
		if err != nil {
			return []Template{}, err
		}
		templates[index] = template
	}
	return templates, nil
}

func ReadTemplateCount(db *sql.DB) (int, error) {
	var count int
	rows, err := databases.QueryFile(db, "scripts/templates/read_templates_count.sql")
	if err != nil {
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&count)
	return count, err
}
