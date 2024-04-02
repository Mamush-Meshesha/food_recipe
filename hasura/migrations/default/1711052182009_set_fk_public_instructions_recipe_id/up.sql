alter table "public"."instructions" drop constraint "instructions_recipe_id_fkey",
  add constraint "instructions_recipe_id_fkey"
  foreign key ("recipe_id")
  references "public"."recipe"
  ("id") on update restrict on delete cascade;
