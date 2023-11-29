CREATE TABLE "news" (
    "id" uuid PRIMARY KEY,
    "title" varchar(200) NOT NULL,
    "text" text NOT NULL,
    "status" smallint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "accepted_by" uuid NULL,
    "created_by" uuid NULL
);

CREATE TABLE "media" (
    "id" uuid PRIMARY KEY,
    "name" varchar(200) NOT NULL,
    "path" path NOT NULL,
    "type" smallint NOT NULL
);

CREATE TABLE "news_media" (
    "new_id" uuid REFERENCES news(id) ON DELETE CASCADE,
    "media_id" uuid REFERENCES media(id) ON DELETE CASCADE,
    PRIMARY KEY (new_id, media_id)
);

CREATE TABLE "tags" (
    "id" uuid PRIMARY KEY,
    "text" varchar(200) NOT NULL,
    "user_id" uuid NOT NULL
);

CREATE TABLE "news_tags" (
    "new_id" uuid REFERENCES news(id) ON DELETE CASCADE,
    "tag_id" uuid REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (new_id, tag_id)
);

CREATE INDEX ON tags("text");
CREATE INDEX ON news("text");
CREATE INDEX ON news("title");