CREATE TABLE wallet (
  id serial PRIMARY KEY NOT NULL,
  wallet_number text UNIQUE NOT NULL,
  is_active text NOT NULL,
  first_name text,
  last_name text,
  email text NOT NULL,
  secretkey text UNIQUE NOT NULL,
  bvn text UNIQUE NOT NULL,
  dob DATE,
  currency text NOT NULL,
  created_time TIMESTAMP default current_timestamp,
  modified_time TIMESTAMP default current_timestamp
);

CREATE TABLE transactions (
  id serial  PRIMARY KEY,
  transaction_ref text UNIQUE NOT NULL,
  transaction_type text NOT NULL,
  transaction_timestamp TIMESTAMP NOT NULL,
  amount numeric(20,2) NOT NULL,
  secretkey text UNIQUE NOT NULL,
  transaction_status text NOT NULL,
  transaction_description text NOT NULL,
  balance numeric(20,2) NOT NULL,
  created_time TIMESTAMP default current_timestamp,
  modified_time TIMESTAMP default current_timestamp,
  wallet_id int references wallet(id)
);