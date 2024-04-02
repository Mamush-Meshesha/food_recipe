CREATE OR REPLACE FUNCTION public.recipe_like(recipe_row recipe, hasura_session json)
RETURNS boolean
LANGUAGE sql
STABLE
AS $$
SELECT EXISTS (
    SELECT 1
    FROM likes A
    WHERE A.user_id = CAST((hasura_session ->> 'x-hasura-user-id') as INTEGER) AND A.recipe_id = recipe_row.id
);
$$;

CREATE OR REPLACE FUNCTION public.recipe_boomark(recipe_row recipe, hasura_session json)
RETURNS boolean
LANGUAGE sql
STABLE
AS $$
SELECT EXISTS (
    SELECT 1
    FROM bookmarks A
    WHERE A.user_id = CAST((hasura_session ->> 'x-hasura-user-id') as INTEGER) AND A.recipe_id = recipe_row.id
);
$$;

CREATE OR REPLACE FUNCTION public.rate_average(recipe_row public.recipe) 
RETURNS double precision
LANGUAGE sql 
STABLE
AS $$
SELECT AVG(rating_value) FROM rating WHERE recipe_id = recipe_row.id;
$$;

CREATE OR REPLACE FUNCTION public.insert_rating(user_id integer, recipe_id integer, rating_value integer)
RETURNS void
LANGUAGE sql
AS $$
INSERT INTO rating (user_id, recipe_id, rating_value)
VALUES (user_id, recipe_id, rating_value)
ON CONFLICT (user_id, recipe_id) 
DO UPDATE SET rating_value = EXCLUDED.rating_value;
$$;
