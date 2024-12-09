--SQL migration content here
CREATE TABLE
    IF NOT EXISTS menu (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        food_name TEXT NOT NULL,
        food_price DECIMAL NOT NULL,
        food_description TEXT,
        food_image_url TEXT
    )