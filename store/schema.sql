-- Drop existing foreign key constraint (if exists)
ALTER TABLE IF EXISTS car
DROP CONSTRAINT IF EXISTS fk_engine_id;

-- Truncate car table to clear existing data
TRUNCATE TABLE car;

-- Truncate engine table to clear existing data
TRUNCATE TABLE engine;    

-- Create engine table
CREATE TABLE IF NOT EXISTS engine (
    id INT AUTO_INCREMENT PRIMARY KEY,
    displacement INT NOT NULL,
    no_of_cylinders INT NOT NULL,
    car_range INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS car (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    year VARCHAR(4) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    fuel_type VARCHAR(50) NOT NULL,
    engine_id INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Add foreign key constraint on engine_id in car table
ALTER TABLE car
ADD CONSTRAINT fk_engine_id
FOREIGN KEY (engine_id)
REFERENCES engine(id)
ON DELETE CASCADE;

-- Insert dummy data into the engine table
INSERT INTO engine (id, displacement, no_of_cylinders, car_range)
VALUES
    (1, 2000, 4, 600),
    (2, 1600, 4, 550),
    (3, 3000, 6, 700),
    (4, 1800, 4, 500);

-- Insert dummy data into the car table
INSERT INTO car (id, name, year, brand, fuel_type, engine_id, price)
VALUES
    (10, 'Honda Civic', '2023', 'Honda', 'Gasoline', 1, 25000.00),
    (11, 'Toyota Corolla', '2022', 'Toyota', 'Gasoline', 2, 22000.00),
    (12, 'Ford Mustang', '2024', 'Ford', 'Gasoline', 3, 40000.00),
    (13, 'BMW 3 Series', '2023', 'BMW', 'Gasoline', 4, 35000.00);