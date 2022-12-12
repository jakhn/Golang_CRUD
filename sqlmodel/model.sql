

CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid(),
    first_name varchar(24) NOT NULL,
    last_name varchar(36) NOT NULL
);

INSERT INTO users (first_name, last_name) VALUES
('Samandar', 'Foziljonov');


UPDATE 
    users 
SET 
    first_name = 'Ravshanbek', 
    last_name = 'Olimov' 
WHERE 
    id = '32d98624-3aa2-4550-954f-c91798794dd7';

