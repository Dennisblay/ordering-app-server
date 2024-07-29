CREATE TABLE "user"
(
    "id"                     INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "first_name"             varchar        NOT NULL,
    "last_name"              varchar        NOT NULL,
    "email"                  varchar UNIQUE NOT NULL,
    "phone"                  varchar        NOT NULL,
    "address"                varchar        NOT NULL,
    "password_hash"          varchar        NOT NULL,
    "password_updated_at"     timestampz DEFAULT '0001-01-01 00:00:00Z',
    "reset_token"            varchar,
    "reset_token_expires_at" timestampz,
    "created_at"             timestampz NOT NULL DEFAULT NOW(),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_user_password_hash ON "user" (password_hash);
CREATE INDEX idx_user_email ON "user" (email);
CREATE INDEX idx_user_phone ON "user" (phone);

CREATE TABLE "role"
(
    "id"         INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "name"       varchar UNIQUE NOT NULL,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_role_name ON "role" (name);

CREATE TABLE "user_role"
(
    "user_id" int,
    "role_id" int,
    PRIMARY KEY ("user_id", "role_id"),
    FOREIGN KEY ("user_id") REFERENCES "user" ("id"),
    FOREIGN KEY ("role_id") REFERENCES "role" ("id")
);

CREATE INDEX idx_user_role_user_id ON "user_role" (user_id);
CREATE INDEX idx_user_role_role_id ON "user_role" (role_id);

CREATE TABLE "session"
(
    "id"         INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "user_id"    int,
    "token"      varchar UNIQUE NOT NULL,
    "expires_at" timestampz,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now()),
    FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);

CREATE INDEX idx_session_user_id ON "session" (user_id);
CREATE INDEX idx_session_token ON "session" (token);

CREATE TABLE product
(
    "id"          INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "name"        varchar,
    "description" text,
    "category"    varchar,
    "price"       decimal(10, 2),
    "stock"       int,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_product_name ON product (name);
CREATE INDEX idx_product_category ON product (category);

CREATE TABLE "order"
(
    "id"            INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "user_id"       int,
    "total_price"   decimal(10, 2),
    "status"        varchar,
    "delivery_date" date,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_order_user_id ON "order" (user_id);
CREATE INDEX idx_order_status ON "order" (status);

CREATE TABLE order_item
(
    "id"         INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "order_id"   int,
    "product_id" int,
    "quantity"   int,
    "price"      decimal(10, 2),
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_order_item_order_id ON order_item (order_id);
CREATE INDEX idx_order_item_product_id ON order_item (product_id);

CREATE TABLE review
(
    "id"         INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "user_id"    int,
    "product_id" int,
    "rating"     int,
    "comment"    text,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_review_user_id ON review (user_id);
CREATE INDEX idx_review_product_id ON review (product_id);

CREATE TABLE booking
(
    "id"          INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "user_id"     int,
    "event_date"  date,
    "event_type"  varchar,
    "guest_count" int,
    "location"    varchar,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_booking_user_id ON booking (user_id);
CREATE INDEX idx_booking_event_date ON booking (event_date);

CREATE TABLE payment
(
    "id"             INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "order_id"       int,
    "amount"         decimal(10, 2),
    "payment_method" varchar,
    "payment_date"   timestampz,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_payment_order_id ON payment (order_id);
CREATE INDEX idx_payment_payment_method ON payment (payment_method);

CREATE TABLE notification
(
    "id"         INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "user_id"    int,
    "message"    text,
    "sent_at"    timestampz,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_notification_user_id ON notification (user_id);
CREATE INDEX idx_notification_sent_at ON notification (sent_at);

CREATE TABLE "social_media"
(
    "id"         INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "user_id"    int,
    "platform"   varchar,
    "link"       varchar,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_social_media_user_id ON "social_media" (user_id);
CREATE INDEX idx_social_media_platform ON "social_media" (platform);

CREATE TABLE testimonial
(
    "id"         INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "user_id"    int,
    "message"    text,
    "created_at"             timestampz not null DEFAULT (now()),
    "updated_at"             timestampz not null  DEFAULT (now())
);

CREATE INDEX idx_testimonial_user_id ON testimonial (user_id);

ALTER TABLE "order"
    ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE order_item
    ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE order_item
    ADD FOREIGN KEY ("product_id") REFERENCES product ("id");

ALTER TABLE review
    ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE review
    ADD FOREIGN KEY ("product_id") REFERENCES product ("id");

ALTER TABLE booking
    ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE payment
    ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE notification
    ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "social_media"
    ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE testimonial
    ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
