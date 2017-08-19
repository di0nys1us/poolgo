CREATE EXTENSION pgcrypto;

CREATE TABLE "poolmate"."session" (
  "id"          BIGSERIAL NOT NULL,
  "deleted"     BOOLEAN   NOT NULL DEFAULT FALSE,
  "version"     BIGINT    NOT NULL DEFAULT 0,
  "user_id"     BIGINT    NOT NULL,
  "date"        DATE      NOT NULL,
  "pool_length" INT       NOT NULL DEFAULT 50,
  "calories"    INT       NOT NULL DEFAULT 0,

  PRIMARY KEY ("id")
);

CREATE TABLE "poolmate"."session_set" (
  "id"               BIGSERIAL NOT NULL,
  "deleted"          BOOLEAN   NOT NULL DEFAULT FALSE,
  "version"          BIGINT    NOT NULL DEFAULT 0,
  "session_id"       BIGINT    NOT NULL,
  "set_number"       INT       NOT NULL,
  "swimming_time"    BIGINT    NOT NULL DEFAULT 0,
  "rest_time"        BIGINT    NOT NULL DEFAULT 0,
  "laps"             INT       NOT NULL DEFAULT 0,
  "average_strokes"  INT       NOT NULL DEFAULT 0,
  "speed"            INT       NOT NULL DEFAULT 0,
  "efficiency_index" INT       NOT NULL DEFAULT 0,

  PRIMARY KEY ("id"),
  FOREIGN KEY ("session_id") REFERENCES "poolmate"."session" ("id")
);

CREATE TABLE "poolmate"."table_event" (
  "id"         BIGSERIAL    NOT NULL,
  "timestamp"  TIMESTAMP    NOT NULL DEFAULT current_timestamp,
  "event_name" VARCHAR(255) NOT NULL,
  "table_name" VARCHAR(255) NOT NULL,
  "row_id"     VARCHAR(255) NOT NULL,
  "user_id"    BIGINT       NOT NULL,

  PRIMARY KEY ("id")
);