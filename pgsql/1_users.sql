CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username varchar(250) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    api_key UUID DEFAULT gen_random_uuid(),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at timestamp without time zone default (now() at time zone 'utc'),
    updated_at timestamp without time zone default (now() at time zone 'utc')
);


INSERT INTO public.users (id,username,"password",api_key,is_active,created_at,updated_at) VALUES 
('e5d9a6b6-cc76-4832-9fd9-69b76fad7542','admin','$2a$10$ObBLp7uclHmb/l3tSMTjC.Kf/sW2YV9Vs9ReCHiW/AkYBie1T9Ejq','cd4898b3-167e-4f60-9832-242a41e2d0ba',true,'2019-03-28 18:16:37.710','2019-03-28 18:16:37.710')
;