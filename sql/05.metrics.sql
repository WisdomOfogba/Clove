-- 05.metrics.sql
-- Tables for metrics and analytics

CREATE TABLE vendor_metrics (
    metrics_id TEXT PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    total_ratings INTEGER NOT NULL DEFAULT 0,
    average_rating DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_comments INTEGER NOT NULL DEFAULT 0,
    positive_comments INTEGER NOT NULL DEFAULT 0,
    negative_comments INTEGER NOT NULL DEFAULT 0,
    positive_comment_ratio DOUBLE PRECISION NOT NULL DEFAULT 0,
    response_rate DOUBLE PRECISION NOT NULL DEFAULT 0,
    average_response_time DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_orders INTEGER NOT NULL DEFAULT 0,
    order_fulfillment_rate DOUBLE PRECISION NOT NULL DEFAULT 0,
    last_updated TIMESTAMP NOT NULL
);

CREATE INDEX idx_vendor_metrics_customer_id ON vendor_metrics(customer_id);

CREATE TABLE platform_metrics (
    metrics_id TEXT PRIMARY KEY,
    total_vendors INTEGER NOT NULL DEFAULT 0,
    approved_vendors INTEGER NOT NULL DEFAULT 0,
    restricted_vendors INTEGER NOT NULL DEFAULT 0,
    flagged_vendors INTEGER NOT NULL DEFAULT 0,
    pending_manual_review INTEGER NOT NULL DEFAULT 0,
    avg_trust_score DOUBLE PRECISION NOT NULL DEFAULT 0,
    fraud_attempts INTEGER NOT NULL DEFAULT 0,
    total_orders INTEGER NOT NULL DEFAULT 0,
    total_revenue BIGINT NOT NULL DEFAULT 0,
    last_updated TIMESTAMP NOT NULL
);