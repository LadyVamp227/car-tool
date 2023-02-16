CREATE TABLE IF NOT EXISTS cars(
    id         serial PRIMARY KEY,
    brand      VARCHAR(50)  NOT NULL,
    model      VARCHAR(50)  NOT NULL,
    price      VARCHAR(50)  NOT NULL,
    type       VARCHAR(50)  NOT NULL,
    usage_type VARCHAR(50)  NOT NULL,
    year       VARCHAR(50)  NOT NULL,
    mileage    VARCHAR(50)  NOT NULL,
    engine     VARCHAR(50)  NOT NULL,
    features   VARCHAR(50)  NOT NULL,
    logo_url   VARCHAR(255) NOT NULL,
    created_on TIMESTAMP    NOT NULL
    );