--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3
-- Dumped by pg_dump version 13.4

-- Started on 2021-10-03 11:02:12

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 206 (class 1259 OID 66200)
-- Name: banner_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.banner_tags (
    banner_id integer NOT NULL,
    tag_id integer NOT NULL
);


ALTER TABLE public.banner_tags OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 66196)
-- Name: banner_tags_banner_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.banner_tags_banner_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.banner_tags_banner_id_seq OWNER TO postgres;

--
-- TOC entry 3049 (class 0 OID 0)
-- Dependencies: 204
-- Name: banner_tags_banner_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.banner_tags_banner_id_seq OWNED BY public.banner_tags.banner_id;


--
-- TOC entry 205 (class 1259 OID 66198)
-- Name: banner_tags_tag_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.banner_tags_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.banner_tags_tag_id_seq OWNER TO postgres;

--
-- TOC entry 3050 (class 0 OID 0)
-- Dependencies: 205
-- Name: banner_tags_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.banner_tags_tag_id_seq OWNED BY public.banner_tags.tag_id;


--
-- TOC entry 208 (class 1259 OID 66209)
-- Name: banners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.banners (
    id integer NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    image_path character varying NOT NULL,
    start_date timestamp with time zone NOT NULL,
    end_date timestamp with time zone NOT NULL
);


ALTER TABLE public.banners OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 66207)
-- Name: banners_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.banners_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.banners_id_seq OWNER TO postgres;

--
-- TOC entry 3051 (class 0 OID 0)
-- Dependencies: 207
-- Name: banners_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.banners_id_seq OWNED BY public.banners.id;


--
-- TOC entry 211 (class 1259 OID 66232)
-- Name: tag_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tag_users (
    user_id integer NOT NULL,
    tag_id integer NOT NULL
);


ALTER TABLE public.tag_users OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 66230)
-- Name: tag_users_tag_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tag_users_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_users_tag_id_seq OWNER TO postgres;

--
-- TOC entry 3052 (class 0 OID 0)
-- Dependencies: 210
-- Name: tag_users_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tag_users_tag_id_seq OWNED BY public.tag_users.tag_id;


--
-- TOC entry 209 (class 1259 OID 66228)
-- Name: tag_users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tag_users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_users_user_id_seq OWNER TO postgres;

--
-- TOC entry 3053 (class 0 OID 0)
-- Dependencies: 209
-- Name: tag_users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tag_users_user_id_seq OWNED BY public.tag_users.user_id;


--
-- TOC entry 203 (class 1259 OID 66187)
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    tag character varying NOT NULL
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 66185)
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO postgres;

--
-- TOC entry 3054 (class 0 OID 0)
-- Dependencies: 202
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- TOC entry 201 (class 1259 OID 66176)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    age date NOT NULL,
    region character varying NOT NULL,
    gender "char" NOT NULL,
    tier character varying NOT NULL,
    wallet_balance integer NOT NULL,
    product_pref character varying NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 66174)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3055 (class 0 OID 0)
-- Dependencies: 200
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 2883 (class 2604 OID 66203)
-- Name: banner_tags banner_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banner_tags ALTER COLUMN banner_id SET DEFAULT nextval('public.banner_tags_banner_id_seq'::regclass);


--
-- TOC entry 2884 (class 2604 OID 66204)
-- Name: banner_tags tag_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banner_tags ALTER COLUMN tag_id SET DEFAULT nextval('public.banner_tags_tag_id_seq'::regclass);


--
-- TOC entry 2885 (class 2604 OID 66212)
-- Name: banners id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banners ALTER COLUMN id SET DEFAULT nextval('public.banners_id_seq'::regclass);


--
-- TOC entry 2886 (class 2604 OID 66235)
-- Name: tag_users user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_users ALTER COLUMN user_id SET DEFAULT nextval('public.tag_users_user_id_seq'::regclass);


--
-- TOC entry 2887 (class 2604 OID 66236)
-- Name: tag_users tag_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_users ALTER COLUMN tag_id SET DEFAULT nextval('public.tag_users_tag_id_seq'::regclass);


--
-- TOC entry 2882 (class 2604 OID 66190)
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- TOC entry 2881 (class 2604 OID 66179)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3038 (class 0 OID 66200)
-- Dependencies: 206
-- Data for Name: banner_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banner_tags (banner_id, tag_id) FROM stdin;
\.


--
-- TOC entry 3040 (class 0 OID 66209)
-- Dependencies: 208
-- Data for Name: banners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banners (id, name, description, image_path, start_date, end_date) FROM stdin;
\.


--
-- TOC entry 3043 (class 0 OID 66232)
-- Dependencies: 211
-- Data for Name: tag_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tag_users (user_id, tag_id) FROM stdin;
\.


--
-- TOC entry 3035 (class 0 OID 66187)
-- Dependencies: 203
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tags (id, tag) FROM stdin;
\.


--
-- TOC entry 3033 (class 0 OID 66176)
-- Dependencies: 201
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, age, region, gender, tier, wallet_balance, product_pref) FROM stdin;
\.


--
-- TOC entry 3056 (class 0 OID 0)
-- Dependencies: 204
-- Name: banner_tags_banner_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.banner_tags_banner_id_seq', 1, false);


--
-- TOC entry 3057 (class 0 OID 0)
-- Dependencies: 205
-- Name: banner_tags_tag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.banner_tags_tag_id_seq', 1, false);


--
-- TOC entry 3058 (class 0 OID 0)
-- Dependencies: 207
-- Name: banners_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.banners_id_seq', 1, false);


--
-- TOC entry 3059 (class 0 OID 0)
-- Dependencies: 210
-- Name: tag_users_tag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tag_users_tag_id_seq', 1, false);


--
-- TOC entry 3060 (class 0 OID 0)
-- Dependencies: 209
-- Name: tag_users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tag_users_user_id_seq', 1, false);


--
-- TOC entry 3061 (class 0 OID 0)
-- Dependencies: 202
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tags_id_seq', 1, false);


--
-- TOC entry 3062 (class 0 OID 0)
-- Dependencies: 200
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- TOC entry 2893 (class 2606 OID 66206)
-- Name: banner_tags banner_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banner_tags
    ADD CONSTRAINT banner_tags_pkey PRIMARY KEY (banner_id, tag_id);


--
-- TOC entry 2895 (class 2606 OID 66217)
-- Name: banners banners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banners
    ADD CONSTRAINT banners_pkey PRIMARY KEY (id);


--
-- TOC entry 2897 (class 2606 OID 66238)
-- Name: tag_users tag_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_users
    ADD CONSTRAINT tag_users_pkey PRIMARY KEY (user_id, tag_id);


--
-- TOC entry 2891 (class 2606 OID 66195)
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- TOC entry 2889 (class 2606 OID 66184)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2898 (class 2606 OID 66218)
-- Name: banner_tags banner_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banner_tags
    ADD CONSTRAINT banner_id FOREIGN KEY (banner_id) REFERENCES public.banners(id) NOT VALID;


--
-- TOC entry 2899 (class 2606 OID 66223)
-- Name: banner_tags tag_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banner_tags
    ADD CONSTRAINT tag_id FOREIGN KEY (tag_id) REFERENCES public.tags(id) NOT VALID;


--
-- TOC entry 2901 (class 2606 OID 66244)
-- Name: tag_users tag_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_users
    ADD CONSTRAINT tag_id FOREIGN KEY (tag_id) REFERENCES public.tags(id) NOT VALID;


--
-- TOC entry 2900 (class 2606 OID 66239)
-- Name: tag_users user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_users
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES public.users(id) NOT VALID;


-- Completed on 2021-10-03 11:02:13

--
-- PostgreSQL database dump complete
--

