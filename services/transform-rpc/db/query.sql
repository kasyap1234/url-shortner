-- name: CreateURLMap :exec
INSERT INTO url_map (short_code, url)
VALUES ($1, $2);

-- name: GetURLByShortCode :one
SELECT short_code, url, created_at
FROM url_map
WHERE short_code = $1
LIMIT 1;
