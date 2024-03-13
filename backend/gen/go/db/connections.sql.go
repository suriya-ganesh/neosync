// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: connections.sql

package db_queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	pg_models "github.com/nucleuscloud/neosync/backend/sql/postgresql/models"
)

const areConnectionsInAccount = `-- name: AreConnectionsInAccount :one
SELECT count(c.id) from neosync_api.connections c
INNER JOIN neosync_api.accounts a ON a.id = c.account_id
WHERE a.id = $1 and c.id = ANY($2::uuid[])
`

type AreConnectionsInAccountParams struct {
	AccountId     pgtype.UUID
	ConnectionIds []pgtype.UUID
}

func (q *Queries) AreConnectionsInAccount(ctx context.Context, db DBTX, arg AreConnectionsInAccountParams) (int64, error) {
	row := db.QueryRow(ctx, areConnectionsInAccount, arg.AccountId, arg.ConnectionIds)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createConnection = `-- name: CreateConnection :one
INSERT INTO neosync_api.connections (
  name, account_id, connection_config, created_by_id, updated_by_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, created_at, updated_at, name, account_id, connection_config, created_by_id, updated_by_id
`

type CreateConnectionParams struct {
	Name             string
	AccountID        pgtype.UUID
	ConnectionConfig *pg_models.ConnectionConfig
	CreatedByID      pgtype.UUID
	UpdatedByID      pgtype.UUID
}

func (q *Queries) CreateConnection(ctx context.Context, db DBTX, arg CreateConnectionParams) (NeosyncApiConnection, error) {
	row := db.QueryRow(ctx, createConnection,
		arg.Name,
		arg.AccountID,
		arg.ConnectionConfig,
		arg.CreatedByID,
		arg.UpdatedByID,
	)
	var i NeosyncApiConnection
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.AccountID,
		&i.ConnectionConfig,
		&i.CreatedByID,
		&i.UpdatedByID,
	)
	return i, err
}

const getConnectionById = `-- name: GetConnectionById :one
SELECT id, created_at, updated_at, name, account_id, connection_config, created_by_id, updated_by_id from neosync_api.connections WHERE id = $1
`

func (q *Queries) GetConnectionById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiConnection, error) {
	row := db.QueryRow(ctx, getConnectionById, id)
	var i NeosyncApiConnection
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.AccountID,
		&i.ConnectionConfig,
		&i.CreatedByID,
		&i.UpdatedByID,
	)
	return i, err
}

const getConnectionByNameAndAccount = `-- name: GetConnectionByNameAndAccount :one
SELECT c.id, c.created_at, c.updated_at, c.name, c.account_id, c.connection_config, c.created_by_id, c.updated_by_id from neosync_api.connections c
INNER JOIN neosync_api.accounts a ON a.id = c.account_id
WHERE a.id = $1 AND c.name = $2
`

type GetConnectionByNameAndAccountParams struct {
	AccountId      pgtype.UUID
	ConnectionName string
}

func (q *Queries) GetConnectionByNameAndAccount(ctx context.Context, db DBTX, arg GetConnectionByNameAndAccountParams) (NeosyncApiConnection, error) {
	row := db.QueryRow(ctx, getConnectionByNameAndAccount, arg.AccountId, arg.ConnectionName)
	var i NeosyncApiConnection
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.AccountID,
		&i.ConnectionConfig,
		&i.CreatedByID,
		&i.UpdatedByID,
	)
	return i, err
}

const getConnectionsByAccount = `-- name: GetConnectionsByAccount :many
SELECT c.id, c.created_at, c.updated_at, c.name, c.account_id, c.connection_config, c.created_by_id, c.updated_by_id from neosync_api.connections c
INNER JOIN neosync_api.accounts a ON a.id = c.account_id
WHERE a.id = $1
ORDER BY c.created_at DESC
`

func (q *Queries) GetConnectionsByAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiConnection, error) {
	rows, err := db.Query(ctx, getConnectionsByAccount, accountid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NeosyncApiConnection
	for rows.Next() {
		var i NeosyncApiConnection
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.AccountID,
			&i.ConnectionConfig,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getConnectionsByIds = `-- name: GetConnectionsByIds :many
SELECT id, created_at, updated_at, name, account_id, connection_config, created_by_id, updated_by_id from neosync_api.connections WHERE id = ANY($1::uuid[])
`

func (q *Queries) GetConnectionsByIds(ctx context.Context, db DBTX, dollar_1 []pgtype.UUID) ([]NeosyncApiConnection, error) {
	rows, err := db.Query(ctx, getConnectionsByIds, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NeosyncApiConnection
	for rows.Next() {
		var i NeosyncApiConnection
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.AccountID,
			&i.ConnectionConfig,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isConnectionInAccount = `-- name: IsConnectionInAccount :one
SELECT count(c.id) from neosync_api.connections c
INNER JOIN neosync_api.accounts a ON a.id = c.account_id
WHERE a.id = $1 and c.id = $2
`

type IsConnectionInAccountParams struct {
	AccountId    pgtype.UUID
	ConnectionId pgtype.UUID
}

func (q *Queries) IsConnectionInAccount(ctx context.Context, db DBTX, arg IsConnectionInAccountParams) (int64, error) {
	row := db.QueryRow(ctx, isConnectionInAccount, arg.AccountId, arg.ConnectionId)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const isConnectionNameAvailable = `-- name: IsConnectionNameAvailable :one
SELECT count(c.id) from neosync_api.connections c
INNER JOIN neosync_api.accounts a ON a.id = c.account_id
WHERE a.id = $1 and c.name = $2
`

type IsConnectionNameAvailableParams struct {
	AccountId      pgtype.UUID
	ConnectionName string
}

func (q *Queries) IsConnectionNameAvailable(ctx context.Context, db DBTX, arg IsConnectionNameAvailableParams) (int64, error) {
	row := db.QueryRow(ctx, isConnectionNameAvailable, arg.AccountId, arg.ConnectionName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const removeConnectionById = `-- name: RemoveConnectionById :exec
DELETE FROM neosync_api.connections WHERE id = $1
`

func (q *Queries) RemoveConnectionById(ctx context.Context, db DBTX, id pgtype.UUID) error {
	_, err := db.Exec(ctx, removeConnectionById, id)
	return err
}

const removeConnectionByNameAndAccount = `-- name: RemoveConnectionByNameAndAccount :exec
DELETE FROM neosync_api.connections WHERE name = $1 and account_id = $2
`

type RemoveConnectionByNameAndAccountParams struct {
	Name      string
	AccountID pgtype.UUID
}

func (q *Queries) RemoveConnectionByNameAndAccount(ctx context.Context, db DBTX, arg RemoveConnectionByNameAndAccountParams) error {
	_, err := db.Exec(ctx, removeConnectionByNameAndAccount, arg.Name, arg.AccountID)
	return err
}

const updateConnection = `-- name: UpdateConnection :one
UPDATE neosync_api.connections
SET name = $1, connection_config = $2,
updated_by_id = $3
WHERE id = $4
RETURNING id, created_at, updated_at, name, account_id, connection_config, created_by_id, updated_by_id
`

type UpdateConnectionParams struct {
	Name             string
	ConnectionConfig *pg_models.ConnectionConfig
	UpdatedByID      pgtype.UUID
	ID               pgtype.UUID
}

func (q *Queries) UpdateConnection(ctx context.Context, db DBTX, arg UpdateConnectionParams) (NeosyncApiConnection, error) {
	row := db.QueryRow(ctx, updateConnection,
		arg.Name,
		arg.ConnectionConfig,
		arg.UpdatedByID,
		arg.ID,
	)
	var i NeosyncApiConnection
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.AccountID,
		&i.ConnectionConfig,
		&i.CreatedByID,
		&i.UpdatedByID,
	)
	return i, err
}
