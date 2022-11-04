// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const listLog = `-- name: ListLog :many
SELECT lift, weight, reps, "calculatedMax" FROM logs
ORDER BY lift
`

func (q *Queries) ListLog(ctx context.Context) ([]Log, error) {
	rows, err := q.db.QueryContext(ctx, listLog)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Log
	for rows.Next() {
		var i Log
		if err := rows.Scan(
			&i.Lift,
			&i.Weight,
			&i.Reps,
			&i.CalculatedMax,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const logLift = `-- name: LogLift :one
INSERT INTO logs (
  lift,weight,reps
) VALUES (
  $1, $2, $3
)
RETURNING lift, weight, reps, "calculatedMax"
`

type LogLiftParams struct {
	Lift   sql.NullString `json:"lift"`
	Weight sql.NullInt32  `json:"weight"`
	Reps   sql.NullInt32  `json:"reps"`
}

func (q *Queries) LogLift(ctx context.Context, arg LogLiftParams) (Log, error) {
	row := q.db.QueryRowContext(ctx, logLift, arg.Lift, arg.Weight, arg.Reps)
	var i Log
	err := row.Scan(
		&i.Lift,
		&i.Weight,
		&i.Reps,
		&i.CalculatedMax,
	)
	return i, err
}

const setMax = `-- name: SetMax :one
INSERT INTO maxes (
  lift,OneRepMax,TrainingMax
) VALUES (
  $1, $2, $3
)
RETURNING lift, onerepmax, trainingmax
`

type SetMaxParams struct {
	Lift        sql.NullString `json:"lift"`
	Onerepmax   sql.NullInt32  `json:"onerepmax"`
	Trainingmax sql.NullInt32  `json:"trainingmax"`
}

func (q *Queries) SetMax(ctx context.Context, arg SetMaxParams) (Max, error) {
	row := q.db.QueryRowContext(ctx, setMax, arg.Lift, arg.Onerepmax, arg.Trainingmax)
	var i Max
	err := row.Scan(&i.Lift, &i.Onerepmax, &i.Trainingmax)
	return i, err
}