package repository

import (
	"backend/pkg/logutil"
	"backend/pkg/tmplutil"
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type BaseSQLRepository struct {
	db         *sqlx.DB
	tmplEngine *tmplutil.SQLTemplateEngine
}

func NewBaseSQLRepository(db *sqlx.DB, queryPath string) *BaseSQLRepository {
	return &BaseSQLRepository{
		db:         db,
		tmplEngine: tmplutil.NewSQLTemplateEngine(queryPath),
	}
}

// 📌 `Queryx` を使用した汎用的な結果取得 (単一/複数結果に対応)
func (r *BaseSQLRepository) ExecuteQuery(section string, params map[string]interface{}, scanFunc func(*sqlx.Rows) error) error {
	query, _, err := r.tmplEngine.RenderQuery(section, params)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	r.logQuery(query, params)

	query, args, err := sqlx.Named(query, params)
	if err != nil {
		return fmt.Errorf("failed to bind parameters: %w", err)
	}

	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	return scanFunc(rows)
}

// 📌 `NamedExec` を使用した非クエリ実行 (INSERT, UPDATE, DELETE)
func (r *BaseSQLRepository) ExecuteNonQuery(section string, params map[string]interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query, _, err := r.tmplEngine.RenderQuery(section, params)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare query in section '%s': %w", section, err)
	}

	r.logQuery(query, params)

	result, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return 0, fmt.Errorf("query execution failed in section '%s': %w", section, err)
	}

	return result.RowsAffected()
}

// 📌 クエリログの出力
func (r *BaseSQLRepository) logQuery(query string, params map[string]interface{}) {
	logutil.Info("Executing SQL Query: %s", query)
	logutil.Info("With Parameters: %+v", params)
}
