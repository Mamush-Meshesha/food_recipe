alter table "public"."catagories"
  add constraint "catagories_recipe_id_fkey"
  foreign key (recipe_id)
  references "public"."recipe"
  (id) on update restrict on delete cascade;
alter table "public"."catagories" alter column "recipe_id" drop not null;
alter table "public"."catagories" add column "recipe_id" int4;
