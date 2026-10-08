# 👑 Binghatti Royal CRM & Luxury Real Estate Platform

<p align="center">
  <img src="https://img.shields.io/badge/Backend-Go_1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Backend" />
  <img src="https://img.shields.io/badge/Dashboard-SvelteKit_5-FF3E00?style=for-the-badge&logo=svelte&logoColor=white" alt="SvelteKit Dashboard" />
  <img src="https://img.shields.io/badge/Landing-Astro_4-BC52EE?style=for-the-badge&logo=astro&logoColor=white" alt="Astro Landing" />
  <img src="https://img.shields.io/badge/Database-PostgreSQL_16-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/Container-Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
  <img src="https://img.shields.io/badge/Styling-Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white" alt="Tailwind CSS" />
</p>

---

## 🌟 Executive Summary

**Binghatti Royal CRM & Luxury Real Estate Platform** is an end-to-end, ultra-luxury real estate sales and VIP lead tracking system designed specifically for high-net-worth real estate developments in Dubai (inspired by Jacob & Co Residences and Bugatti Residences).

The ecosystem connects an ultra-responsive **Astro 4 Landing Page** to a high-speed **Go REST API Backend** and an executive **SvelteKit Admin CRM Dashboard**, providing real-time lead ingestion, Kanban pipeline tracking, unit inventory management, and financial ROI analytics.

---

## 🚀 Key Features

### 1. 📊 VIP Lead Kanban Pipeline (/leads)

- **6 Sales Pipeline Stages**: New Leads, Contacted, Qualified, Viewing Scheduled, Closed Won, Closed Lost.
- **Ultra-Visible Glowing Count Badges**: High-contrast, color-coded pill indicators placed inline with column headers.
- **HTML5 Drag & Drop**: Fluid lead stage transitions with dynamic golden drag-over glow (#D4AF37).
- **Real-time Status Overrides & Persistence**: Multi-layered state persistence via Go REST API and local state synchronization.
- **Instant Search & Filter Bar**: Instant client lookup by full name, email, WhatsApp number, budget range, and investment purpose.
- **CSV Exporter**: Single-click CSV data export for sales advisors.
- **VIP Lead Inspector Modal**: Detailed lead history view with instant WhatsApp direct messaging link.

### 2. 🏢 Unit Inventory Management (/inventory)

- **Real-Time Catalog Summary Bar**: Instant tracking of Available, Reserved, and Sold luxury penthouse units.
- **Unit Layout & Project Filter**: Instant filtering by layout type (Sky Penthouse, Mansion, Villa) and project name.
- **Floor Plan PDF Links**: Quick access floor plan blueprints for clients.
- **Add & Edit Luxury Unit Modal**: Full CRUD functionality to manage live real estate units.

### 3. 📈 Executive Overview & Analytics (/ & /analytics)

- **KPI Summary Cards**: Total active VIP leads, gross pipeline volume, conversion rate, and revenue figures.
- **Interactive Lead Funnel & ROI Charts**: Visual breakdowns of sales stage distribution and campaign conversion metrics.
- **Platform Shortcuts & Status Indicators**: Live PostgreSQL sync status indicator.

### 4. ⚙️ CRM Settings (/settings)

- **Advisor Profile Management**: Profile avatar, email, and advisor role settings.
- **API Key Security**: Key rotation and secure backend authentication.
- **Currency & Regional Preferences**: Currency switcher (USD, AED, EUR, GBP) and timezone settings.
- **Notification Controls**: Email alerts, instant WhatsApp notifications, and lead assignment triggers.

### 5. 💎 Ultra-Luxury Astro Landing Page (:4321)

- **Interactive VIP Registration Form**: Connects directly to backend API :8085 to push new leads straight to the sales team dashboard.
- **Responsive Architecture**: Fully mobile & tablet optimized with 2-column payment plan cards.
- **23K Gold Glassmorphism Theme**: Curated dark luxury aesthetics (#0A0A0A, #D4AF37, #121218).

---

## 🏗️ Architecture & Technology Stack

`mermaid
graph TD;
    A[Astro VIP Landing Page - :4321] -->|POST /api/v1/leads| B(Go REST Backend API - :8085);
    C[SvelteKit Admin CRM Dashboard - :5173] -->|GET / PUT / POST| B;
    B -->|SQL Queries| D[(PostgreSQL 16 Database)];
`

| Layer                | Technology    | Key Libraries & Frameworks                  |
| :------------------- | :------------ | :------------------------------------------ |
| **Frontend Landing** | Astro 4       | Tailwind CSS, Google Fonts (Syne, Inter)    |
| **Admin Dashboard**  | SvelteKit     | Svelte 5, Tailwind CSS, Lucide Vector Icons |
| **Backend REST API** | Go 1.22+      | Chi Router, CORS middleware, PG Driver      |
| **Database**         | PostgreSQL 16 | Auto-Migrations, Custom Enum Types          |
| **Containerization** | Docker        | Docker Compose, Nginx Alpine Reverse Proxy  |

---

## 📂 Repository Structure

`.
├── backend/                  # Go REST API Server
│   ├── cmd/server/           # Application Entry Point
│   ├── internal/             # Handlers, Repositories, Models & Database
│   ├── go.mod
│   └── go.sum
├── dashboard/                # SvelteKit CRM Dashboard
│   ├── src/
│   │   ├── lib/              # Svelte Stores & API Helpers
│   │   ├── routes/           # /leads, /inventory, /analytics, /settings
│   │   └── app.css           # Luxury Gold Custom Dropdowns & Theme
│   ├── package.json
│   └── svelte.config.js
├── landing/                  # Astro 4 VIP Landing Page
│   ├── src/
│   │   ├── components/       # VIPForm.astro, Footer.astro, PaymentPlan.astro
│   │   └── pages/index.astro
│   └── package.json
├── docker-compose.yml        # Development Docker Orchestration
├── docker-compose.prod.yml   # Production Docker Orchestration
└── README.md                 # Project Documentation`

---

## ⚡ Quick Start Guide (Local Development)

### Prerequisites

Make sure you have the following installed on your machine:

- **Node.js**: v18.0.0 or higher
- **Go**: v1.22.0 or higher
- **PostgreSQL**: v16.0 or higher (or Docker)

---

### Step 1: Clone the Repository

`ash
git clone https://github.com/your-username/binghatti-royal-crm.git
cd binghatti-royal-crm
`

---

### Step 2: Start the Go Backend Server (:8085)

`ash
cd backend
go mod tidy
go run cmd/server/main.go
`
_The Go REST API will start listening on http://localhost:8085._

---

### Step 3: Start the SvelteKit CRM Dashboard (:5173)

Open a new terminal window:
`ash
cd dashboard
npm install
npm run dev
`
_The CRM Dashboard will be accessible at http://localhost:5173._

---

### Step 4: Start the Astro Landing Page (:4321)

Open a new terminal window:
`ash
cd landing
npm install
npm run dev
`
_The VIP Landing Page will be accessible at http://localhost:4321._

---

## 🐳 Docker Deployment (One-Command Setup)

You can run the entire infrastructure (PostgreSQL, Go Backend, Svelte Dashboard, Astro Landing) using Docker Compose:

`ash
docker-compose up --build -d
`

To stop all services:
`ash
docker-compose down
`

---

## 🔌 API Reference Endpoints

| HTTP Method                                                          | Endpoint                  | Description                               |
| :------------------------------------------------------------------- | :------------------------ | :---------------------------------------- |
| GET                                                                  | /api/v1/health            | Service health status check               |
| GET                                                                  | /api/v1/leads             | Fetch all VIP leads                       |
| POST                                                                 | /api/v1/leads             | Create a new lead (Landing form / Manual) |
| PUT                                                                  | /api/v1/leads/{id}/status | Update lead sales stage (                 |
| ew, contacted, qualified, iewing_scheduled, closed_won, closed_lost) |
| GET                                                                  | /api/v1/units             | Fetch unit inventory list                 |
| POST                                                                 | /api/v1/units             | Create new real estate unit               |

---

## 🎨 Theme & Design System Guidelines

- **Primary Gold Accent**: #D4AF37
- **Dark Off-Black Background**: #0A0A0C
- **Off-Black Container**: #121218
- **Border Overlay**:
  gba(255, 255, 255, 0.1) /
  gba(212, 175, 55, 0.3)
- **Typography**: Syne (Headlines) & Inter (Body & Data Tables)

---

## 📄 License

This project is licensed under the **MIT License**.

---

<p align="center">
  Crafted by Anwarul Karim for Luxury Real Estate Automation
</p>
