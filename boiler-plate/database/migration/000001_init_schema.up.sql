CREATE TABLE public.sign_ups (
    user_name varchar(255) NOT NULL,
    user_email varchar(255) NOT NULL,
    users_id int8 NOT NULL,
    password varchar(255) NOT NULL,
    CONSTRAINT sign_up_email_key UNIQUE (user_email)
);