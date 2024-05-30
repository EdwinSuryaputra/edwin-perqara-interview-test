package repository

import (
	"context"
	"fmt"
)

func (r *Repository) GetDrinks(ctx context.Context) (output []GetDrinkOutput, err error) {
	rows, err := r.Db.QueryContext(ctx, `SELECT id, public_id, name, stock
			FROM storage_drinks WHERE deleted_by IS NULL and deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for rows.Next() {
		var row GetDrinkOutput
		err := rows.Scan(&row.Id, &row.DrinkPublicId, &row.Name, &row.Stock)
		if err != nil {
			return nil, err
		}
		output = append(output, row)
	}

	return
}

func (r *Repository) GetDrinkByPublicId(ctx context.Context, publicId string) (output GetDrinkOutput, err error) {
	err = r.Db.QueryRowContext(ctx, `SELECT id, public_id, name, stock FROM storage_drinks WHERE public_id = $1 
	AND deleted_at IS NULL AND deleted_by IS NULL`, publicId).
		Scan(&output.Id, &output.DrinkPublicId, &output.Name, &output.Stock)
	if err != nil {
		return
	}

	return
}

func (r *Repository) InsertDrink(ctx context.Context, input InsertDrinkInput) error {
	statement, err := r.Db.PrepareContext(ctx, `INSERT INTO storage_drinks (public_id, name, stock, created_at, created_by, updated_at, updated_by) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}
	defer func() {
		err = statement.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = statement.ExecContext(ctx, input.DrinkPublicId, input.Name, input.Stock,
		input.CreatedAt, input.CreatedBy, input.CreatedAt, input.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateDrink(ctx context.Context, input UpdateDrinkInput) error {
	statement, err := r.Db.PrepareContext(ctx, `UPDATE storage_drinks 
		SET name = $1, stock = $2, updated_at = $3, updated_by = $4
		WHERE id = $5 AND deleted_at IS NULL AND deleted_by IS NULL`)
	if err != nil {
		return err
	}
	defer func() {
		err = statement.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = statement.ExecContext(ctx, input.Name, input.Stock, input.UpdatedAt, input.UpdatedBy, input.DrinkId)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteDrink(ctx context.Context, input DeleteDrinkInput) error {
	statement, err := r.Db.PrepareContext(ctx, `UPDATE storage_drinks 
		SET deleted_at = $1, deleted_by = $2
		WHERE id = $3 AND deleted_at IS NULL AND deleted_by IS NULL`)
	if err != nil {
		return err
	}
	defer func() {
		err = statement.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = statement.ExecContext(ctx, input.DeletedAt, input.DeletedBy, input.DrinkId)
	if err != nil {
		return err
	}

	return nil
}
