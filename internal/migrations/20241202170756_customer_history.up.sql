CREATE TABLE
    IF NOT EXISTS customer_history (
        id UUID PRIMARY KEY,
        table_number INTEGER NOT NULL,
        status TEXT NOT NULL DEFAULT 'occupied' CHECK (status IN ('paid', 'occupied')),
        time_in TIMESTAMP NOT NULL DEFAULT NOW (),
        time_out TIMESTAMP,
        total_price INTEGER
    );