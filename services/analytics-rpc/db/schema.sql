CREATE TABLE analytics (
    id UUID PRIMARY KEY,
    url_id UUID NOT NULL,
    user_id UUID NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL, 
    
);