# Clove: The Reputation Layer for Multi-Merchant Platforms

**Clove** is a real-time vendor verification and trust-scoring engine designed for e-commerce and delivery ecosystems like Jumia, Chowdeck, and Glovo. We bridge the gap between merchant onboarding and consumer safety.

## 🚀 The Problem

In emerging e-commerce markets, "what you ordered vs. what you got" is a systemic issue. Standard KYC is static—it only checks if a business is real at the start. It doesn't account for declining service quality or fraudulent behavior over time.

## 💡 The Solution

Clove provides a **Dynamic Trust Score** for vendors by combining:

* **Automated KYC:** Verifying business legitimacy using AI-powered document parsing.
* **Progressive Monitoring:** Analyzing customer review sentiment using LLMs to detect quality drops.
* **Safety Interventions:** A "Warning Nudge" system that alerts customers before they transact with high-risk vendors.
* **Search Influence:** An API-ready scoring system that allows platforms to demote "bad actors" in search results.

## 🛠️ Tech Stack

* **Backend:** Golang (for high-performance concurrency)
* **Frontend:** React (for a snappy, responsive dashboard)
* **Database:** PostgreSQL (Relational data for complex merchant-review mappings)
* **Intelligence:** Google Gemini Pro API (Sentiment analysis & Multimodal image verification)

## 🏗️ Core Architecture

1. **The Engine:** A Go-based service that aggregates merchant data.
2. **The Scorer:** A logic layer that weights KYC data against real-time customer feedback.
3. **The API:** Seamless integration for 3rd party platforms to fetch a `vendor_trust_score`.

## 👥 The Team

* **[Chibx](https://github.com/chibx/)** – Lead Backend / Systems Architect
* **[Wisdom Ofogba](https://github.com/WisdomOfogba/)** – Frontend Engineer / UI/UX
