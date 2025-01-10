--SQL migration content here
CREATE TABLE
    IF NOT EXISTS restaurant (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        restaurant_name TEXT NOT NULL,
        branch TEXT NOT NULL
    )