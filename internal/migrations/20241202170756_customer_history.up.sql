CREATE TABLE
    IF NOT EXISTS customer_history (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        table_number INTEGER NOT NULL,
        status TEXT NOT NULL,
        time_in TIMESTAMP NOT NULL DEFAULT NOW (),
        time_out TIMESTAMP,
        total_price INTEGER
    );

-- CHECK (status IN ('paid', 'occupied'))