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
