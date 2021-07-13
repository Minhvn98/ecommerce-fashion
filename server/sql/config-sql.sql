use ecommerce_fashion;
SET autocommit = OFF;
ALTER TABLE products ADD FULLTEXT (name);