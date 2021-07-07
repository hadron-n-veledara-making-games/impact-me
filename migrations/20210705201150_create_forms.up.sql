CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TYPE sex_enum AS ENUM ('F', 'M');
CREATE TYPE sex_preference_enum AS ENUM ('F', 'M', 'U');
CREATE TYPE server_enum AS ENUM ('EU', 'ASIA', 'AMERICA', 'SAR');


CREATE TABLE forms (
    id bigserial not null primary key,
    
     user_id  integer REFERENCES users,

     age  integer not null,
     consider_age  boolean not null,

     sex  sex_enum not null,
     sex_preference  sex_preference_enum not null,  
    
     server  server_enum not null,

     description  text not null;

     photo_url  text default null,

     created_at timestamptz not null default now(),
     updated_at timestamptz not null default now(),
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON forms
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE INDEX forms_user_id
  ON forms (user_id);