--SQL migration content here
CREATE TABLE
    IF NOT EXISTS menu (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        restaurant_id UUID NOT NULL,
        menu_name TEXT NOT NULL,
        FOREIGN KEY (restaurant_id) REFERENCES restaurant (id) ON DELETE CASCADE
    )