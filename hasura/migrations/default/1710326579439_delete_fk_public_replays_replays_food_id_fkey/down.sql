alter table "public"."replays"
  add constraint "replays_food_id_fkey"
  foreign key ("food_id")
  references "public"."food"
  ("id") on update restrict on delete restrict;
