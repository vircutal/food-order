--SQL migration content here
CREATE TABLE
    IF NOT EXISTS menu_item (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        menu_id UUID NOT NULL,
        menu_item_name TEXT NOT NULL,
        menu_item_price REAL NOT NULL,
        menu_item_description TEXT,
        menu_item_image_url TEXT,
        FOREIGN KEY (menu_id) REFERENCES menu (id)
    )