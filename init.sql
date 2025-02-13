CREATE TABLE public.product_category (
                                         id int4 DEFAULT nextval('product_type_id_seq'::regclass) NOT NULL,
                                         name varchar(255) NOT NULL,
                                         props json NULL,
                                         CONSTRAINT product_type_pkey PRIMARY KEY (id)
);

CREATE TABLE public.product_status (
                                       id serial4 NOT NULL,
                                       "name" varchar(255) NOT NULL,
                                       CONSTRAINT product_status_pkey PRIMARY KEY (id)
);

CREATE TABLE public.products (
                                 id serial4 NOT NULL,
                                 "name" varchar(255) NOT NULL,
                                 description text NULL,
                                 quantity int4 NULL,
                                 image _text NULL,
                                 sale int4 DEFAULT 0 NULL,
                                 price int4 NOT NULL,
                                 status int4 NULL,
                                 category int4 NULL,
                                 property json NULL,
                                 CONSTRAINT products_pkey PRIMARY KEY (id),
                                 CONSTRAINT products_product_status_fk FOREIGN KEY (status) REFERENCES public.product_status(id),
                                 CONSTRAINT products_product_type_fk FOREIGN KEY (category) REFERENCES public.product_category(id)
);

CREATE TABLE public.products_prorerty (
                                          id serial4 NOT NULL,
                                          "name" varchar(225) NULL,
                                          CONSTRAINT products_prorerty_pkey PRIMARY KEY (id)
);

CREATE TABLE public.users (
                              id serial4 NOT NULL,
                              firstname varchar(255) NULL,
                              lastname varchar(255) NULL,
                              username varchar(255) NOT NULL,
                              "type" int4 NOT NULL,
                              email varchar(255) NULL,
                              "password" varchar(255) NOT NULL,
                              phone int8 NOT NULL,
                              wishlist int4 NULL,
                              cart int4 NULL,
                              CONSTRAINT users_pkey PRIMARY KEY (id),
                              CONSTRAINT users_users_type_fk FOREIGN KEY ("type") REFERENCES public.users_type(id)
);

CREATE TABLE public.users_type (
                                   id serial4 NOT NULL,
                                   "type" varchar(225) NULL,
                                   CONSTRAINT users_type_pkey PRIMARY KEY (id)
);
