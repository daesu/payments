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
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: payment_scheme_type_enum; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.payment_scheme_type_enum AS ENUM (
    'FPS',
    'Bacs',
    'SEPAINSTANT'
);


--
-- Name: resource_enum; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.resource_enum AS ENUM (
    'Payment'
);


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: account; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.account (
    account_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    bank_id integer,
    account_number bigint NOT NULL,
    account_number_code character varying(20),
    account_type bigint DEFAULT 0 NOT NULL,
    account_name character varying(50),
    created_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    updated_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL
);


--
-- Name: bank; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.bank (
    bank_id integer NOT NULL,
    bank_code character varying(20),
    bank_name character varying(100),
    bank_address character varying(255),
    created_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    updated_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL
);


--
-- Name: currency; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.currency (
    currency_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name character varying(50) NOT NULL,
    abbrv character varying(20) NOT NULL,
    created_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    updated_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL
);


--
-- Name: customer; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.customer (
    customer_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    customer_name character varying(100),
    customer_address character varying(255),
    created_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    updated_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL
);


--
-- Name: customer_account; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.customer_account (
    customer_account_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    customer_id uuid,
    account_id uuid,
    created_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    updated_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL
);


--
-- Name: organisation; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.organisation (
    organisation_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    created_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    updated_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL
);


--
-- Name: payment; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.payment (
    payment_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    organisation_id uuid NOT NULL,
    type public.resource_enum DEFAULT 'Payment'::public.resource_enum NOT NULL,
    currency_id uuid NOT NULL,
    amount numeric NOT NULL,
    beneficary_id uuid NOT NULL,
    debtor_id uuid NOT NULL,
    end_to_end_reference character varying(255),
    numeric_reference bigint DEFAULT 0 NOT NULL,
    payment_purpose character varying(255),
    scheme_payment_type_id integer NOT NULL,
    reference character varying(255),
    processing_date date,
    created_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    updated_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: scheme_payment_type; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.scheme_payment_type (
    scheme_payment_type_id integer NOT NULL,
    value character varying(100) NOT NULL,
    payment_scheme public.payment_scheme_type_enum NOT NULL,
    resource public.resource_enum NOT NULL,
    created_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    updated_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL
);


--
-- Name: scheme_payment_type_scheme_payment_type_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.scheme_payment_type_scheme_payment_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: scheme_payment_type_scheme_payment_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.scheme_payment_type_scheme_payment_type_id_seq OWNED BY public.scheme_payment_type.scheme_payment_type_id;


--
-- Name: scheme_payment_type scheme_payment_type_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scheme_payment_type ALTER COLUMN scheme_payment_type_id SET DEFAULT nextval('public.scheme_payment_type_scheme_payment_type_id_seq'::regclass);


--
-- Name: account account_bank_id_account_number_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.account
    ADD CONSTRAINT account_bank_id_account_number_key UNIQUE (bank_id, account_number);


--
-- Name: account account_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.account
    ADD CONSTRAINT account_pkey PRIMARY KEY (account_id);


--
-- Name: bank bank_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bank
    ADD CONSTRAINT bank_pkey PRIMARY KEY (bank_id);


--
-- Name: currency currency_name_abbrv_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.currency
    ADD CONSTRAINT currency_name_abbrv_key UNIQUE (name, abbrv);


--
-- Name: currency currency_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.currency
    ADD CONSTRAINT currency_pkey PRIMARY KEY (currency_id);


--
-- Name: customer_account customer_account_customer_id_account_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.customer_account
    ADD CONSTRAINT customer_account_customer_id_account_id_key UNIQUE (customer_id, account_id);


--
-- Name: customer_account customer_account_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.customer_account
    ADD CONSTRAINT customer_account_pkey PRIMARY KEY (customer_account_id);


--
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (customer_id);


--
-- Name: organisation organisation_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.organisation
    ADD CONSTRAINT organisation_pkey PRIMARY KEY (organisation_id);


--
-- Name: payment payment_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_pkey PRIMARY KEY (payment_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: scheme_payment_type scheme_payment_type_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scheme_payment_type
    ADD CONSTRAINT scheme_payment_type_pkey PRIMARY KEY (scheme_payment_type_id);


--
-- Name: scheme_payment_type scheme_payment_type_value_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scheme_payment_type
    ADD CONSTRAINT scheme_payment_type_value_key UNIQUE (value);


--
-- Name: account account_bank_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.account
    ADD CONSTRAINT account_bank_id_fkey FOREIGN KEY (bank_id) REFERENCES public.bank(bank_id);


--
-- Name: customer_account customer_account_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.customer_account
    ADD CONSTRAINT customer_account_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account(account_id);


--
-- Name: customer_account customer_account_customer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.customer_account
    ADD CONSTRAINT customer_account_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.customer(customer_id);


--
-- Name: payment payment_beneficary_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_beneficary_id_fkey FOREIGN KEY (beneficary_id) REFERENCES public.customer_account(customer_account_id) ON DELETE CASCADE;


--
-- Name: payment payment_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES public.currency(currency_id) ON DELETE CASCADE;


--
-- Name: payment payment_debtor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_debtor_id_fkey FOREIGN KEY (debtor_id) REFERENCES public.customer_account(customer_account_id) ON DELETE CASCADE;


--
-- Name: payment payment_organisation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_organisation_id_fkey FOREIGN KEY (organisation_id) REFERENCES public.organisation(organisation_id);


--
-- Name: payment payment_scheme_payment_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_scheme_payment_type_id_fkey FOREIGN KEY (scheme_payment_type_id) REFERENCES public.scheme_payment_type(scheme_payment_type_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20190405064821');
