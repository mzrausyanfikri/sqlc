// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package querytest

import (
	"context"
)

const joinWhereClause = `-- name: JoinWhereClause :many
SELECT foo.barid
FROM foo
JOIN bar ON bar.id = barid
WHERE owner = ?
`

func (q *Queries) JoinWhereClause(ctx context.Context, owner string) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, joinWhereClause, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var barid int64
		if err := rows.Scan(&barid); err != nil {
			return nil, err
		}
		items = append(items, barid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
