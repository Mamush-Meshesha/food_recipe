CREATE TABLE "public"."ingredients" ("id" serial NOT NULL, "recipe_id" integer NOT NULL, "name" text NOT NULL, PRIMARY KEY ("id") , FOREIGN KEY ("recipe_id") REFERENCES "public"."recipe"("id") ON UPDATE restrict ON DELETE restrict);
