--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.13
-- Dumped by pg_dump version 10.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: candidates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.candidates (
    id bigint NOT NULL,
    name character varying(2044) NOT NULL,
    age integer NOT NULL,
    vote_count integer NOT NULL,
    code integer NOT NULL,
    election_id bigint NOT NULL
);


ALTER TABLE public.candidates OWNER TO postgres;

--
-- Name: elections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.elections (
    id bigint NOT NULL,
    name character varying(2044) NOT NULL,
    election_time timestamp without time zone NOT NULL,
    district character varying(2044) NOT NULL
);


ALTER TABLE public.elections OWNER TO postgres;

--
-- Name: voters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.voters (
    id bigint NOT NULL,
    name character varying(2044) NOT NULL,
    district character varying(2044) NOT NULL,
    age bigint NOT NULL
);


ALTER TABLE public.voters OWNER TO postgres;

--
-- Data for Name: candidates; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: elections; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.elections VALUES (1, 'Pemilihan Lurah', '2019-04-14 00:00:00', 'Winterfell');


--
-- Data for Name: voters; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Name: candidates candidates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.candidates
    ADD CONSTRAINT candidates_pkey PRIMARY KEY (id);


--
-- Name: elections elections_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.elections
    ADD CONSTRAINT elections_pkey PRIMARY KEY (id);


--
-- Name: voters voters_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.voters
    ADD CONSTRAINT voters_pkey PRIMARY KEY (id);


--
-- Name: index_election_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX index_election_id ON public.candidates USING btree (election_id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

