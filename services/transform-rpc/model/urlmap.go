package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UrlMap struct {
	ID         int64     `db: "id"`
	short_code string    `db: "short_code"`
	url        string    `db: "url"`
	created_at time.Time `db: "created_at" json:"created_at,omitempty"`
}

type UrlMapModel struct {
	conn sqlx.SqlConn
}

func NewUrlMapModel(conn sqlx.SqlConn) *UrlMapModel {
	return &UrlMapModel{
		conn: conn,
	}
}

func (m *UrlMapModel) FindByShortCode(ctx context.Context, shortCode string) (*UrlMap, error) {
	var resp UrlMap
	query := "SELECT id,short_code,url,created_at FROM url_map WHERE short_code = ?"
	err := m.conn.QueryRowCtx(ctx, &resp, query, shortCode)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &resp, nil
}
