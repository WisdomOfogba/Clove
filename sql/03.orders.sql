-- 03.orders.sql
-- Tables for orders and order items

CREATE TABLE orders (
    id TEXT PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    vendor_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status TEXT NOT NULL DEFAULT 'pending',
    total_amount DECIMAL(15,2) NOT NULL,  -- in kobo
    delivery_fee DECIMAL(15,2) NOT NULL,  -- in kobo
    delivery_address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_orders_customer_id ON orders(customer_id);
CREATE INDEX idx_orders_vendor_id ON orders(vendor_id);
CREATE INDEX idx_orders_status ON orders(status);

CREATE TABLE order_items (
    id BIGINT PRIMARY KEY,
    order_id TEXT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    meal_id BIGINT NOT NULL REFERENCES meals(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL,
    price DECIMAL(15,2) NOT NULL,  -- per item in kobo
    created_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_meal_id ON order_items(meal_id);