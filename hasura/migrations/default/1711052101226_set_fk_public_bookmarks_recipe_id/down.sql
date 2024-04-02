alter table "public"."bookmarks" drop constraint "bookmarks_recipe_id_fkey",
  add constraint "bookmarks_post_id_fkey"
  foreign key ("recipe_id")
  references "public"."recipe"
  ("id") on update restrict on delete restrict;
