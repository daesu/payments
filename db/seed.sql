--
-- Data for Name: scheme_payment_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scheme_payment_type (scheme_payment_type_id, value, payment_scheme, resource, created_at, updated_at) FROM stdin;
1	ImmediatePayment	FPS	Payment	1543509958	1543509958
\.


--
-- Data for Name: currency; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.currency (currency_id, name, abbrv, created_at, updated_at) FROM stdin;
a04c291f-80bc-4cad-9fbd-90592e4dd0ea	British Pound	GBP	1543509959	1543509959
b54f7a63-1f91-4234-9a68-6be5d0c564f3	Euro	EUR	1554540743	1554540743
4ef07148-beb8-4855-92b6-8e5e6d0c3f4d	US Dollar	USD	1554540743	1554540743
\.



--
-- Data for Name: organisation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organisation (organisation_id, created_at, updated_at) FROM stdin;
c551e221-5fbc-4aa7-84fc-85a6a8a26720	1543509959	1543509959
\.


--
-- Data for Name: customer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customer (customer_id, customer_name, customer_address, created_at, updated_at) FROM stdin;
c5d71434-fdf5-4570-a26e-c483584658e2	Lews Therin	Two Rivers, Andor	1543509959	1543509959
6bad2463-f118-441d-9bae-22d26533db6f	Jon Snow	123 Wall st, London	1543509959	1543509959
\.


--
-- Data for Name: bank; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bank (bank_id, bank_code, bank_name, bank_address) FROM stdin;
10042	IRONB	Iron Bank	Braavos
\.


--
-- Data for Name: account; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.account (account_id, bank_id, account_number, account_number_code, account_type, account_name, created_at, updated_at) FROM stdin;
5701249e-f33a-45a3-8722-e6917ccff6f0	10042	12345667	BBAN	0	Lews T	1543509959	1543509959
6eae6bb8-f7fb-425a-8af8-64adb616b739	10042	87654323	CDAN	0	J Snow	1543509959	1543509959
\.


--
-- Data for Name: customer_account; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customer_account (customer_account_id, customer_id, account_id, created_at, updated_at) FROM stdin;
2e938509-e5d2-4d54-9354-59dcd26683f1	c5d71434-fdf5-4570-a26e-c483584658e2	5701249e-f33a-45a3-8722-e6917ccff6f0	1543509959	1543509959
c984bcd0-5970-458b-bd28-2bb06c75af85	6bad2463-f118-441d-9bae-22d26533db6f	6eae6bb8-f7fb-425a-8af8-64adb616b739	1543509959	1543509959
\.


--
-- Data for Name: payment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment (payment_id, version, organisation_id, type, currency_id, amount, beneficary_id, debtor_id, end_to_end_reference, numeric_reference, payment_purpose, scheme_payment_type_id, reference, processing_date, created_at, updated_at) FROM stdin;
61b0c143-f1f9-457d-a889-80570b820348	0	c551e221-5fbc-4aa7-84fc-85a6a8a26720	Payment	a04c291f-80bc-4cad-9fbd-90592e4dd0ea	102.12	2e938509-e5d2-4d54-9354-59dcd26683f1	c984bcd0-5970-458b-bd28-2bb06c75af85	football boots	1002001	Paying for goods/services	1	Payment for goods	2001-09-28	1543509959	1543509959
fd54832d-d872-428b-a10d-17ddf782b4df	0	c551e221-5fbc-4aa7-84fc-85a6a8a26720	Payment	a04c291f-80bc-4cad-9fbd-90592e4dd0ea	10.99	c984bcd0-5970-458b-bd28-2bb06c75af85	2e938509-e5d2-4d54-9354-59dcd26683f1	pizza	1002001	Paying for delivery	1	Payment for service	2001-09-28	1543509959	1543509959
\.

