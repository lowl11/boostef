package ef

import "context"

func Execute(ctx context.Context, script string) error {
	_, err := Connection().ExecContext(ctx, script)
	if err != nil {
		return err
	}

	return nil
}

func ExecuteResult(ctx context.Context, script string) ([]map[string]any, error) {
	rows, err := Connection().QueryxContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]any, 5)
	for rows.Next() {
		item := make(map[string]any)
		if err = rows.MapScan(item); err != nil {
			return nil, err
		}

		if len(item) == 0 {
			continue
		}

		result = append(result, item)
	}

	return result, nil
}
