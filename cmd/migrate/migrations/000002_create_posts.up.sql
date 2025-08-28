CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    title text NOT NULL,
    Content TEXT NOT NULL,
    create_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);