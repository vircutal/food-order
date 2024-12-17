--SQL migration content here
CREATE TABLE
    IF NOT EXISTS table_info (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        table_number INTEGER NOT NULL,
        status TEXT NOT NULL
    );

--CHECK (status IN ('available', 'occupied', 'reserved'))
-- INSERT INTO
--     table_info (table_number)
-- VALUES
--     (1),
--     (2),
--     (3),
--     (4),
--     (5);