-- Drop foreign keys
ALTER TABLE testimonial DROP CONSTRAINT testimonial_user_id_fkey;
ALTER TABLE "social_media" DROP CONSTRAINT social_media_user_id_fkey;
ALTER TABLE notification DROP CONSTRAINT notification_user_id_fkey;
ALTER TABLE payment DROP CONSTRAINT payment_order_id_fkey;
ALTER TABLE booking DROP CONSTRAINT booking_user_id_fkey;
ALTER TABLE review DROP CONSTRAINT review_user_id_fkey;
ALTER TABLE review DROP CONSTRAINT review_product_id_fkey;
ALTER TABLE order_item DROP CONSTRAINT order_item_product_id_fkey;
ALTER TABLE order_item DROP CONSTRAINT order_item_order_id_fkey;
ALTER TABLE "order" DROP CONSTRAINT "order_user_id_fkey";
ALTER TABLE "session" DROP CONSTRAINT session_user_id_fkey;
ALTER TABLE "user_role" DROP CONSTRAINT user_role_user_id_fkey;
ALTER TABLE "user_role" DROP CONSTRAINT user_role_role_id_fkey;

-- Drop indexes
DROP INDEX IF EXISTS idx_user_email;
DROP INDEX IF EXISTS idx_user_phone;
DROP INDEX IF EXISTS idx_role_name;
DROP INDEX IF EXISTS idx_user_role_user_id;
DROP INDEX IF EXISTS idx_user_role_role_id;
DROP INDEX IF EXISTS idx_session_user_id;
DROP INDEX IF EXISTS idx_session_token;
DROP INDEX IF EXISTS idx_product_name;
DROP INDEX IF EXISTS idx_product_category;
DROP INDEX IF EXISTS idx_order_user_id;
DROP INDEX IF EXISTS idx_order_status;
DROP INDEX IF EXISTS idx_order_item_order_id;
DROP INDEX IF EXISTS idx_order_item_product_id;
DROP INDEX IF EXISTS idx_review_user_id;
DROP INDEX IF EXISTS idx_review_product_id;
DROP INDEX IF EXISTS idx_booking_user_id;
DROP INDEX IF EXISTS idx_booking_event_date;
DROP INDEX IF EXISTS idx_payment_order_id;
DROP INDEX IF EXISTS idx_payment_payment_method;
DROP INDEX IF EXISTS idx_notification_user_id;
DROP INDEX IF EXISTS idx_notification_sent_at;
DROP INDEX IF EXISTS idx_social_media_user_id;
DROP INDEX IF EXISTS idx_social_media_platform;
DROP INDEX IF EXISTS idx_testimonial_user_id;

-- Drop tables
DROP TABLE IF EXISTS testimonial;
DROP TABLE IF EXISTS "social_media";
DROP TABLE IF EXISTS notification;
DROP TABLE IF EXISTS payment;
DROP TABLE IF EXISTS booking;
DROP TABLE IF EXISTS review;
DROP TABLE IF EXISTS order_item;
DROP TABLE IF EXISTS "order";
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS "session";
DROP TABLE IF EXISTS "user_role";
DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "user";
