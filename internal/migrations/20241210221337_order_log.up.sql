--SQL migration content here
CREATE TABLE
    IF NOT EXISTS order_log (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        customer_id UUID NOT NULL,
        food_id UUID NOT NULL,
        food_price REAL NOT NULL,
        quantity INTEGER NOT NULL,
        order_description TEXT,
        ordered_time TIMESTAMP NOT NULL,
        FOREIGN KEY (customer_id) REFERENCES customer_history (id),
        FOREIGN KEY (food_id) REFERENCES menu (id)
    );