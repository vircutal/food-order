CREATE TABLE
    IF NOT EXISTS customer (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        table_number INTEGER NOT NULL,
        status TEXT NOT NULL,
        time_in TIMESTAMP NOT NULL DEFAULT NOW (),
        time_out TIMESTAMP,
        payment_time TIMESTAMP,
        total_price REAL
    );

-- CHECK (status IN ('paid', 'occupied'))