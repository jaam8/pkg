package postgres

import "github.com/jackc/pgx/v5"

func ScanRowsInStruct[T any](rows pgx.Rows) ([]T, error) {
	return pgx.CollectRows(rows, pgx.RowToStructByName[T])
}

func ScanRowInStruct[T any](rows pgx.Rows) (T, error) {
	return pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[T])
}

func CollectRows[T any](rows pgx.Rows, fn pgx.RowToFunc[T]) ([]T, error) {
	return pgx.CollectRows(rows, fn)
}

func RowTo[T any](row pgx.CollectableRow) (T, error) {
	return pgx.RowTo[T](row)
}
