CREATE SEQUENCE IF NOT EXISTS product_type_id_seq;

CREATE TABLE IF NOT EXISTS product_category (
    id int4 DEFAULT nextval('product_type_id_seq'::regclass) NOT NULL,
    name varchar(255) NOT NULL,
    props json NULL,
    CONSTRAINT product_type_pkey PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS product_status (
                                              id serial4 NOT NULL,
                                              name varchar(255) NOT NULL,
    CONSTRAINT product_status_pkey PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS products (
                                        id serial4 NOT NULL,
                                        name varchar(255) NOT NULL,
    description text NULL,
    quantity int4 NULL,
    image text NULL,
    sale int4 DEFAULT 0 NULL,
    price int4 NOT NULL,
    status int4 NULL,
    category int4 NULL,
    property json NULL,
    CONSTRAINT products_pkey PRIMARY KEY (id),
    CONSTRAINT products_product_status_fk FOREIGN KEY (status) REFERENCES product_status(id),
    CONSTRAINT products_product_type_fk FOREIGN KEY (category) REFERENCES product_category(id)
    );

CREATE TABLE IF NOT EXISTS products_property (
                                                 id serial4 NOT NULL,
                                                 name varchar(225) NULL,
    CONSTRAINT products_property_pkey PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS users_type (
                                          id serial4 NOT NULL,
                                          type varchar(225) NULL,
    CONSTRAINT users_type_pkey PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS users (
                                     id serial4 NOT NULL,
                                     firstname varchar(255) NULL,
    lastname varchar(255) NULL,
    username varchar(255) NOT NULL,
    type int4 NOT NULL,
    email varchar(255) NULL,
    password varchar(255) NOT NULL,
    phone bigint NOT NULL,
    wishlist int4 NULL,
    cart int4 NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_users_type_fk FOREIGN KEY (type) REFERENCES users_type(id)
    );
