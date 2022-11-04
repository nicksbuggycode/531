-- name: LogLift :one
INSERT INTO logs (
  lift,weight,reps
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListLog :many
SELECT * FROM logs
ORDER BY lift;

-- name: SetMax :one
INSERT INTO maxes (
  lift,OneRepMax,TrainingMax
) VALUES (
  $1, $2, $3
)
RETURNING *;