CREATE TABLE IF NOT EXISTS "document_type"
(
    "id"   INTEGER PRIMARY KEY,
    "name" VARCHAR(40) NOT NULL,
    "code" VARCHAR(4)  NOT NULL
);

CREATE TABLE IF NOT EXISTS "citizen"
(
    "id"               SERIAL PRIMARY KEY,
    "document"         VARCHAR(20)  NOT NULL,
    "document_type_id" INTEGER      NOT NULL,
    "expedition_date"  DATE         NOT NULL,
    "name"             VARCHAR(100) NOT NULL,
    "lastname"         VARCHAR(100) NOT NULL,
    "address"          VARCHAR(100),
    "gender"           VARCHAR(10)  NOT NULL,
    "age"              INTEGER      NOT NULL,
    CONSTRAINT "fk_document_type"
        FOREIGN KEY ("document_type_id")
            REFERENCES "document_type" ("id")
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT "uc_document"
        UNIQUE ("document", "document_type_id")
);
CREATE INDEX IF NOT EXISTS "idx_citizen_document" ON "citizen" ("document");