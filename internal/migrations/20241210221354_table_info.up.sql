--SQL migration content here
CREATE TABLE
    IF NOT EXISTS table_info (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        table_number INTEGER NOT NULL,
        status TEXT NOT NULL DEFAULT 'available' CHECK (status IN ('available', 'occupied', 'reserved'))
    );

INSERT INTO
    table_info (table_number, status)
VALUES
    (1, 'available'),
    (2, 'available'),
    (3, 'available'),
    (4, 'available'),
    (5, 'available');