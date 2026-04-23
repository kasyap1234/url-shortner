-- name: CreateAnalytics :exec
INSERT INTO analytics (url_id, user_id, timestamp, page_title, page_url)
VALUES ($1, $2, $3, $4, $5);

-- name: GetAnalytics : many
SELECT * FROM analytics WHERE url_id = $1;
