CREATE TABLE public.posts
(
    id character varying(36),
    title character varying(20) NOT NULL,
    content character varying(500),
    created_at numeric(14),
    updated_at numeric(14),
    PRIMARY KEY (id)
);


