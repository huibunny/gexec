-- 2022-08-04

CREATE TABLE public.t_app (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	tenant_id uuid NULL,
	"name" varchar(256) NULL,
	cover varchar(256) NULL,
	icon varchar(256) NULL,
	brief varchar(256) NULL,
	CONSTRAINT t_app_pk PRIMARY KEY (id)
);


CREATE TABLE public.t_app_type_rel (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	tenant_id uuid NULL,
	openid varchar(256) NULL,
	unionid varchar(256) NULL,
	CONSTRAINT t_app_type_rel_pk PRIMARY KEY (id)
);
