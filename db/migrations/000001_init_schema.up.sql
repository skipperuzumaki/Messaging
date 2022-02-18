CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "messageGroup" (
  "id" bigserial PRIMARY KEY,
  "identifier" varchar UNIQUE NOT NULL,
  "from_user" bigint NOT NULL,
  "to_user" bigint NOT NULL
);

CREATE TABLE "message" (
  "id" bigserial PRIMARY KEY,
  "group" varchar NOT NULL,
  "message" varchar NOT NULL,
  "sent_at" timestamptz NOT NULL DEFAULT (now()),
  "read" boolean NOT NULL DEFAULT (FALSE),
  "sent_from" bigint NOT NULL,
  "sent_to" bigint NOT NULL
);

ALTER TABLE "messageGroup" ADD FOREIGN KEY ("from_user") REFERENCES "users" ("id");

ALTER TABLE "messageGroup" ADD FOREIGN KEY ("to_user") REFERENCES "users" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("group") REFERENCES "messageGroup" ("identifier");

ALTER TABLE "message" ADD FOREIGN KEY ("sent_from") REFERENCES "users" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("sent_to") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "messageGroup" ("from_user");

CREATE INDEX ON "messageGroup" ("to_user");

CREATE INDEX ON "messageGroup" ("identifier");

CREATE INDEX ON "message" ("group");

CREATE INDEX ON "message" ("sent_at");

CREATE INDEX ON "message" ("read");
