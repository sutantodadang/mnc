PGDMP  "                    }            mnc    12.22    17.2 ,    @           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            A           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            B           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            C           1262    16384    mnc    DATABASE     n   CREATE DATABASE mnc WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';
    DROP DATABASE mnc;
                     postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                     postgres    false            D           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                        postgres    false    7            E           0    0    SCHEMA public    ACL     Q   REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;
                        postgres    false    7            �           1247    16724    type_source    TYPE     W   CREATE TYPE public.type_source AS ENUM (
    'TOPUP',
    'PAYMENT',
    'TRANSFER'
);
    DROP TYPE public.type_source;
       public               postgres    false    7            �           1247    16719    type_transaction    TYPE     K   CREATE TYPE public.type_transaction AS ENUM (
    'CREDIT',
    'DEBIT'
);
 #   DROP TYPE public.type_transaction;
       public               postgres    false    7            �            1259    16744    balances    TABLE     �  CREATE TABLE public.balances (
    balance_id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    user_id uuid NOT NULL,
    balance_amount numeric DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.314132+00'::timestamp with time zone NOT NULL,
    updated_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.314132+00'::timestamp with time zone NOT NULL
);
    DROP TABLE public.balances;
       public         heap r       postgres    false    7    7    7            �            1259    16833    balances_histories    TABLE     �  CREATE TABLE public.balances_histories (
    balance_history_id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    balance_id uuid NOT NULL,
    transaction_id uuid,
    balance_amount_before numeric DEFAULT 0 NOT NULL,
    balance_amount_after numeric DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.354085+00'::timestamp with time zone NOT NULL,
    updated_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.354085+00'::timestamp with time zone NOT NULL
);
 &   DROP TABLE public.balances_histories;
       public         heap r       postgres    false    7    7    7            �            1259    16387    goose_db_version    TABLE     �   CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now() NOT NULL
);
 $   DROP TABLE public.goose_db_version;
       public         heap r       postgres    false    7            �            1259    16385    goose_db_version_id_seq    SEQUENCE     �   ALTER TABLE public.goose_db_version ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME public.goose_db_version_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public               postgres    false    7    204            �            1259    16778    payments    TABLE     �  CREATE TABLE public.payments (
    payment_id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    user_id uuid NOT NULL,
    remarks character varying NOT NULL,
    payment_amount numeric DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.331104+00'::timestamp with time zone NOT NULL,
    updated_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.331104+00'::timestamp with time zone NOT NULL
);
    DROP TABLE public.payments;
       public         heap r       postgres    false    7    7    7            �            1259    16761    topups    TABLE     �  CREATE TABLE public.topups (
    top_up_id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    user_id uuid NOT NULL,
    top_up_amount numeric DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.322765+00'::timestamp with time zone NOT NULL,
    updated_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.322765+00'::timestamp with time zone NOT NULL
);
    DROP TABLE public.topups;
       public         heap r       postgres    false    7    7    7            �            1259    16817    transactions    TABLE       CREATE TABLE public.transactions (
    transaction_id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    user_id uuid NOT NULL,
    status character varying NOT NULL,
    transaction_type public.type_transaction NOT NULL,
    source_id uuid,
    source_type public.type_source NOT NULL,
    created_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.346507+00'::timestamp with time zone NOT NULL,
    updated_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.346507+00'::timestamp with time zone NOT NULL
);
     DROP TABLE public.transactions;
       public         heap r       postgres    false    7    7    647    7    644            �            1259    16795 	   transfers    TABLE     �  CREATE TABLE public.transfers (
    transfer_id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    source_user_id uuid NOT NULL,
    target_user_id uuid NOT NULL,
    remarks character varying NOT NULL,
    transfer_amount numeric DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.338419+00'::timestamp with time zone NOT NULL,
    updated_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.338419+00'::timestamp with time zone NOT NULL
);
    DROP TABLE public.transfers;
       public         heap r       postgres    false    7    7    7            �            1259    16731    users    TABLE     �  CREATE TABLE public.users (
    user_id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying,
    phone_number character varying NOT NULL,
    address text NOT NULL,
    pin character varying NOT NULL,
    created_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.30091+00'::timestamp with time zone NOT NULL,
    updated_at timestamp with time zone DEFAULT '2025-02-15 09:38:08.30091+00'::timestamp with time zone NOT NULL
);
    DROP TABLE public.users;
       public         heap r       postgres    false    7    7    7            8          0    16744    balances 
   TABLE DATA           _   COPY public.balances (balance_id, user_id, balance_amount, created_at, updated_at) FROM stdin;
    public               postgres    false    206   >       =          0    16833    balances_histories 
   TABLE DATA           �   COPY public.balances_histories (balance_history_id, balance_id, transaction_id, balance_amount_before, balance_amount_after, created_at, updated_at) FROM stdin;
    public               postgres    false    211   �>       6          0    16387    goose_db_version 
   TABLE DATA           N   COPY public.goose_db_version (id, version_id, is_applied, tstamp) FROM stdin;
    public               postgres    false    204   �?       :          0    16778    payments 
   TABLE DATA           h   COPY public.payments (payment_id, user_id, remarks, payment_amount, created_at, updated_at) FROM stdin;
    public               postgres    false    208   @       9          0    16761    topups 
   TABLE DATA           [   COPY public.topups (top_up_id, user_id, top_up_amount, created_at, updated_at) FROM stdin;
    public               postgres    false    207   t@       <          0    16817    transactions 
   TABLE DATA           �   COPY public.transactions (transaction_id, user_id, status, transaction_type, source_id, source_type, created_at, updated_at) FROM stdin;
    public               postgres    false    210   �@       ;          0    16795 	   transfers 
   TABLE DATA           �   COPY public.transfers (transfer_id, source_user_id, target_user_id, remarks, transfer_amount, created_at, updated_at) FROM stdin;
    public               postgres    false    209   �A       7          0    16731    users 
   TABLE DATA           s   COPY public.users (user_id, first_name, last_name, phone_number, address, pin, created_at, updated_at) FROM stdin;
    public               postgres    false    205   �B       F           0    0    goose_db_version_id_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.goose_db_version_id_seq', 22, true);
          public               postgres    false    203            �           2606    16845 *   balances_histories balances_histories_pkey 
   CONSTRAINT     x   ALTER TABLE ONLY public.balances_histories
    ADD CONSTRAINT balances_histories_pkey PRIMARY KEY (balance_history_id);
 T   ALTER TABLE ONLY public.balances_histories DROP CONSTRAINT balances_histories_pkey;
       public                 postgres    false    211            �           2606    16755    balances balances_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.balances
    ADD CONSTRAINT balances_pkey PRIMARY KEY (balance_id);
 @   ALTER TABLE ONLY public.balances DROP CONSTRAINT balances_pkey;
       public                 postgres    false    206            �           2606    16392 &   goose_db_version goose_db_version_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);
 P   ALTER TABLE ONLY public.goose_db_version DROP CONSTRAINT goose_db_version_pkey;
       public                 postgres    false    204            �           2606    16789    payments payments_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (payment_id);
 @   ALTER TABLE ONLY public.payments DROP CONSTRAINT payments_pkey;
       public                 postgres    false    208            �           2606    16772    topups topups_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.topups
    ADD CONSTRAINT topups_pkey PRIMARY KEY (top_up_id);
 <   ALTER TABLE ONLY public.topups DROP CONSTRAINT topups_pkey;
       public                 postgres    false    207            �           2606    16827    transactions transactions_pkey 
   CONSTRAINT     h   ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id);
 H   ALTER TABLE ONLY public.transactions DROP CONSTRAINT transactions_pkey;
       public                 postgres    false    210            �           2606    16806    transfers transfers_pkey 
   CONSTRAINT     _   ALTER TABLE ONLY public.transfers
    ADD CONSTRAINT transfers_pkey PRIMARY KEY (transfer_id);
 B   ALTER TABLE ONLY public.transfers DROP CONSTRAINT transfers_pkey;
       public                 postgres    false    209            �           2606    16743    users users_phone_number_key 
   CONSTRAINT     _   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_number_key UNIQUE (phone_number);
 F   ALTER TABLE ONLY public.users DROP CONSTRAINT users_phone_number_key;
       public                 postgres    false    205            �           2606    16741    users users_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public                 postgres    false    205            �           2606    16846 )   balances_histories fk_balance_id_balances    FK CONSTRAINT     �   ALTER TABLE ONLY public.balances_histories
    ADD CONSTRAINT fk_balance_id_balances FOREIGN KEY (balance_id) REFERENCES public.balances(balance_id);
 S   ALTER TABLE ONLY public.balances_histories DROP CONSTRAINT fk_balance_id_balances;
       public               postgres    false    211    206    2980            �           2606    16807 %   transfers fk_source_user_id_transfers    FK CONSTRAINT     �   ALTER TABLE ONLY public.transfers
    ADD CONSTRAINT fk_source_user_id_transfers FOREIGN KEY (source_user_id) REFERENCES public.users(user_id);
 O   ALTER TABLE ONLY public.transfers DROP CONSTRAINT fk_source_user_id_transfers;
       public               postgres    false    2978    205    209            �           2606    16812 %   transfers fk_target_user_id_transfers    FK CONSTRAINT     �   ALTER TABLE ONLY public.transfers
    ADD CONSTRAINT fk_target_user_id_transfers FOREIGN KEY (target_user_id) REFERENCES public.users(user_id);
 O   ALTER TABLE ONLY public.transfers DROP CONSTRAINT fk_target_user_id_transfers;
       public               postgres    false    209    2978    205            �           2606    16851 1   balances_histories fk_transaction_id_transactions    FK CONSTRAINT     �   ALTER TABLE ONLY public.balances_histories
    ADD CONSTRAINT fk_transaction_id_transactions FOREIGN KEY (transaction_id) REFERENCES public.transactions(transaction_id);
 [   ALTER TABLE ONLY public.balances_histories DROP CONSTRAINT fk_transaction_id_transactions;
       public               postgres    false    2988    210    211            �           2606    16756    balances fk_user_id_balances    FK CONSTRAINT     �   ALTER TABLE ONLY public.balances
    ADD CONSTRAINT fk_user_id_balances FOREIGN KEY (user_id) REFERENCES public.users(user_id);
 F   ALTER TABLE ONLY public.balances DROP CONSTRAINT fk_user_id_balances;
       public               postgres    false    2978    206    205            �           2606    16790    payments fk_user_id_payments    FK CONSTRAINT     �   ALTER TABLE ONLY public.payments
    ADD CONSTRAINT fk_user_id_payments FOREIGN KEY (user_id) REFERENCES public.users(user_id);
 F   ALTER TABLE ONLY public.payments DROP CONSTRAINT fk_user_id_payments;
       public               postgres    false    208    2978    205            �           2606    16773    topups fk_user_id_topups    FK CONSTRAINT     |   ALTER TABLE ONLY public.topups
    ADD CONSTRAINT fk_user_id_topups FOREIGN KEY (user_id) REFERENCES public.users(user_id);
 B   ALTER TABLE ONLY public.topups DROP CONSTRAINT fk_user_id_topups;
       public               postgres    false    207    2978    205            �           2606    16828 $   transactions fk_user_id_transactions    FK CONSTRAINT     �   ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_user_id_transactions FOREIGN KEY (user_id) REFERENCES public.users(user_id);
 N   ALTER TABLE ONLY public.transactions DROP CONSTRAINT fk_user_id_transactions;
       public               postgres    false    210    205    2978            8   g   x�����0c\9:��X>S�aL�%���h�i'�ԣIP?��/�L�h d�G�����g'�@�g�M}����*�M�����?C*��U)�Z�5�t�?>?      =   �   x���Mn� �ur����=g������4+�
a���'�,�Aҥf��n�V�(PrL @S�h�)���$�P�C0!�g"��nA�{�>A,T�?�Ωr!,��ʾ�zw�E�(�r��N��#I�@ݡ&�p�Ys�n�W����:��H¸�nR�a�a[�|+F��^��I���IM�7l\�e�'��	k��q1��k7��y|�IL�Oxhu���#���I/����:S� 7�P�o��$�?��c��_I�L      6   g   x�uι�0����X�]��jq�?ι�|���~&(o�f^���n�\A�0~D_�u`�`�*�`�sY��4G:�E%:�p���s�h�1�%����J���n?2�      :   a   x���1
�0F�=��D��-F��j�A�^@��>�4�R3yU����B�,6��`�jb����5ǵ��<%�%ĥ�0&�]J���ߍSc��u#�      9   U   x����� �3�»�'�j��Gн����G�</���6Q�Շ8 .���N����GXʆ֓uXL��(�����-� $      <     x����N�0E��+�QV3~{��Kv$��d��O)��D��q�=SL����C��ہ2�	 TC�c E�C�'��s�8���kC�2�7�;���64
TY���A<�p�p��Y�� �ݖDsց���J��KS��N�/��I�q␍�[�p��?�<]ޞ�q���)��%)�OR�<�N��1X�.�@U�.�Fd���Q�I�p�A�Z1P��3�8%I��Pd���Z��]0[�5��d�$��f����rw1�8�m���B�      ;   �   x���;�0�����
���Y�8�t�ظ��Љ�x�����S��jI`u�I�[ &�( @�`�IH�rMY�7��<��>e���Ъ#�6e���G�FH�s�1߾u�搚,�Zb�>Cݟ�=n���ꕡ�i�g�K�t��|��6��SQ��;�����?��R(����}0Ƽ �'�      7   �   x����N�@���,��;�;)H�MQL���R��@!}z�M�4=��ܜ��'K���vm�`3�+[h�l�.ǒ! p�Pa�b�zl�o��!5~ac�]bq�Xy����i�>�U��q��q�ng2- ���Wt�)�|�.T��֔p��U�y&��� Hv�oi �@�L����q|����T���~Y�;(
���.)7��~�i���n˒�#�\Fض���D�=�Z�.�0��^R�9��8X>�a?�Ti�     