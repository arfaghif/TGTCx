--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3
-- Dumped by pg_dump version 13.4

-- Started on 2021-10-02 20:23:02

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
-- TOC entry 202 (class 1259 OID 66116)
-- Name: banner_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.banner_tags (
    banner_id integer NOT NULL,
    tag_id integer NOT NULL
);


ALTER TABLE public.banner_tags OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 66108)
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
-- TOC entry 203 (class 1259 OID 66121)
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    tag character varying(16) NOT NULL
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 66100)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    age smallint NOT NULL,
    id integer NOT NULL,
    region character varying NOT NULL,
    gender "char" NOT NULL,
    tier character varying NOT NULL,
    wallet_balance integer,
    product_ref character varying
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3006 (class 0 OID 66116)
-- Dependencies: 202
-- Data for Name: banner_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banner_tags (banner_id, tag_id) FROM stdin;
\.


--
-- TOC entry 3005 (class 0 OID 66108)
-- Dependencies: 201
-- Data for Name: banners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banners (id, name, description, image_path, start_date, end_date) FROM stdin;
\.


--
-- TOC entry 3007 (class 0 OID 66121)
-- Dependencies: 203
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tags (id, tag) FROM stdin;
\.


--
-- TOC entry 3004 (class 0 OID 66100)
-- Dependencies: 200
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (age, id, region, gender, tier, wallet_balance, product_ref) FROM stdin;
\.


--
-- TOC entry 2867 (class 2606 OID 66120)
-- Name: banner_tags banner_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banner_tags
    ADD CONSTRAINT banner_tags_pkey PRIMARY KEY (banner_id, tag_id);


--
-- TOC entry 2865 (class 2606 OID 66115)
-- Name: banners banners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banners
    ADD CONSTRAINT banners_pkey PRIMARY KEY (id);


--
-- TOC entry 2871 (class 2606 OID 66125)
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- TOC entry 2863 (class 2606 OID 66107)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2868 (class 1259 OID 66131)
-- Name: fki_banner_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_banner_id ON public.banner_tags USING btree (banner_id);


--
-- TOC entry 2869 (class 1259 OID 66137)
-- Name: fki_tag_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_tag_id ON public.banner_tags USING btree (tag_id);


--
-- TOC entry 2872 (class 2606 OID 66126)
-- Name: banner_tags banner_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banner_tags
    ADD CONSTRAINT banner_id FOREIGN KEY (banner_id) REFERENCES public.banners(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- TOC entry 2873 (class 2606 OID 66132)
-- Name: banner_tags tag_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banner_tags
    ADD CONSTRAINT tag_id FOREIGN KEY (tag_id) REFERENCES public.tags(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


-- Completed on 2021-10-02 20:23:03

--
-- PostgreSQL database dump complete
--

