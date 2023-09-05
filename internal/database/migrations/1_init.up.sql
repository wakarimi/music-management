CREATE TABLE "directories"
(
    "dir_id"       SERIAL PRIMARY KEY,
    "path"         TEXT        NOT NULL UNIQUE,
    "date_added"   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "last_scanned" TIMESTAMPTZ
);

CREATE TABLE "covers"
(
    "cover_id"   SERIAL PRIMARY KEY,
    "dir_id"     INTEGER     NOT NULL,
    "path"       TEXT        NOT NULL,
    "name"       TEXT        NOT NULL,
    "size"       BIGINT      NOT NULL,
    "format"     TEXT        NOT NULL,
    "date_added" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("dir_id") REFERENCES "directories" ("dir_id")
);

CREATE TABLE "tracks"
(
    "track_id"   SERIAL PRIMARY KEY,
    "dir_id"     INTEGER     NOT NULL,
    "cover_id"   INTEGER,
    "path"       TEXT        NOT NULL,
    "name"       TEXT        NOT NULL,
    "size"       BIGINT      NOT NULL,
    "format"     TEXT        NOT NULL,
    "date_added" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("dir_id") REFERENCES "directories" ("dir_id"),
    FOREIGN KEY ("cover_id") REFERENCES "covers" ("cover_id")
);