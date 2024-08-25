CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS employees (
    "id" BIGSERIAL NOT NULL,
    "publicId" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "name" CHARACTER VARYING(70) NOT NULL,
    "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
    "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
    CONSTRAINT "PK_EMPLOYEE_ID" PRIMARY KEY ("id"),
    CONSTRAINT "UQ_EMPLOYEE_PUB_ID" UNIQUE ("publicId")
);

-- Got this information from LinkedIn
INSERT INTO employees ("name") VALUES
    ('Igor Marinelli'),
    ('Gabriel Lima'),
    ('Leonardo Vieira'),
    ('Karine Klann'),
    ('João Baroni'),
    ('Willian Freitas'),
    ('Pedro Henrique'),
    ('Paloma Lima'),
    ('Lucas Bivar'),
    ('Guilherme Sembeneli'),
    ('Filipe Toyoshima'),
    ('Lucas Simão');