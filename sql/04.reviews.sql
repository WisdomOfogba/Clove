-- 04.reviews.sql
-- Table for reviews

CREATE TABLE reviews (
    id BIGINT PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    vendor_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    meal_id BIGINT NOT NULL REFERENCES meals(id) ON DELETE CASCADE,
    rating BIGINT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    sentiment TEXT,
    sentiment_score DOUBLE PRECISION,
    created_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_reviews_customer_id ON reviews(customer_id);
CREATE INDEX idx_reviews_vendor_id ON reviews(vendor_id);
CREATE INDEX idx_reviews_meal_id ON reviews(meal_id);