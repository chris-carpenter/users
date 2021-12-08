CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE emails (
  id uuid DEFAULT uuid_generate_v4(),
  email_type VARCHAR(15),
  address VARCHAR(50),
  PRIMARY KEY (id)
);

CREATE TABLE phone_numbers (
    id uuid DEFAULT uuid_generate_v4(),
    phone_number_type VARCHAR(15),
    number VARCHAR(15),
    PRIMARY KEY (id)
);

ALTER TABLE users DROP COLUMN user_id;

ALTER TABLE users ADD COLUMN id UUID DEFAULT uuid_generate_v4();
ALTER TABLE users ADD PRIMARY KEY (id);

ALTER TABLE users ADD COLUMN group_name VARCHAR(50);

ALTER TABLE users ADD COLUMN emails UUID;
ALTER TABLE users ADD FOREIGN KEY (emails) REFERENCES emails (id);

ALTER TABLE users ADD COLUMN phone_numbers UUID;
ALTER TABLE users ADD FOREIGN KEY (phone_numbers) REFERENCES phone_numbers (id);
