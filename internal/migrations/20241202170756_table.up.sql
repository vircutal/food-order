--SQL migration content here
CREATE TABLE
    IF NOT EXISTS table_info (
        id UUID PRIMARY KEY,
        table_number INTEGER NOT NULL,
        status TEXT CHECK (status IN ('paid', 'occupied', 'reserved')),
        time_in TIMESTAMP DEFAULT NOW () NOT NULL,
        time_out TIMESTAMP
    );