CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS videos (
    video_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    uploader_id UUID,
    category VARCHAR(50),
    region VARCHAR(10),
    thumbnail_url TEXT,
    is_active BOOLEAN DEFAULT true
);

INSERT INTO videos (title, category, region, thumbnail_url) VALUES
('Taylor Swift - Cruel Summer', 'music', 'US', 'https://i.ytimg.com/vi/1.jpg'),
('MrBeast Epic Challenge', 'entertainment', 'US', 'https://i.ytimg.com/vi/2.jpg')
ON CONFLICT DO NOTHING;