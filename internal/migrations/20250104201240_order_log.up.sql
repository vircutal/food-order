--SQL migration content here
CREATE TABLE
    IF NOT EXISTS order_log (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        customer_history_id UUID NOT NULL,
        menu_item_id UUID NOT NULL,
        menu_item_price REAL NOT NULL,
        quantity INTEGER NOT NULL,
        order_description TEXT,
        ordered_time TIMESTAMP NOT NULL,
        FOREIGN KEY (customer_history_id) REFERENCES customer_history (id),
        FOREIGN KEY (menu_item_id) REFERENCES menu_item (id)
    );