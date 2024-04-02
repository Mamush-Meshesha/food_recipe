alter table "public"."comments"
  add constraint "comments_food_id_fkey"
  foreign key ("food_id")
  references "public"."food"
  ("id") on update restrict on delete restrict;
