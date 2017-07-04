CREATE TABLE "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "version_idx" ON "schema_migration" (version);
CREATE TABLE "movies" (
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"description" text NOT NULL,
"url" TEXT NOT NULL,
"ratings" integer NOT NULL
);
CREATE TABLE "reviews" (
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
"id" TEXT PRIMARY KEY,
"content" text NOT NULL,
"title" TEXT NOT NULL,
"movie_id" char(36) NOT NULL
);
CREATE INDEX "reviews_movie_id_idx" ON "reviews" (movie_id);
