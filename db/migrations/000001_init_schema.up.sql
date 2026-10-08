-- Create extension for UUID generation if not exists
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create Enum Types
CREATE TYPE user_role AS ENUM ('super_admin', 'sales_agent');
CREATE TYPE unit_status AS ENUM ('available', 'reserved', 'sold');
CREATE TYPE lead_status AS ENUM ('new', 'contacted', 'qualified', 'viewing_scheduled', 'closed_won', 'closed_lost');

-- 1. Users Table (Admin / Sales Agents)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role user_role NOT NULL DEFAULT 'sales_agent',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 2. Projects Table (e.g. Binghatti Mercedes-Benz Places, Bugatti Residences)
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(100) UNIQUE NOT NULL,
    title VARCHAR(150) NOT NULL,
    tagline TEXT,
    location_name VARCHAR(150) NOT NULL,
    starting_price_aed NUMERIC(15, 2) NOT NULL,
    handover_date VARCHAR(50),
    description TEXT,
    brochure_url TEXT,
    hero_video_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 3. Units Table (Bedrooms, Penthouse, Floor Plans)
CREATE TABLE units (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    unit_type VARCHAR(50) NOT NULL,
    size_sqft NUMERIC(10, 2) NOT NULL,
    price_aed NUMERIC(15, 2) NOT NULL,
    floor_plan_image_url TEXT NOT NULL,
    floor_plan_pdf_url TEXT NOT NULL,
    status unit_status NOT NULL DEFAULT 'available',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 4. Leads Table (VIP Registrations)
CREATE TABLE leads (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id) ON DELETE SET NULL,
    full_name VARCHAR(100) NOT NULL,
    whatsapp_number VARCHAR(30) NOT NULL,
    email VARCHAR(150) NOT NULL,
    budget_range VARCHAR(50),
    investment_purpose VARCHAR(50),
    lead_source VARCHAR(50) NOT NULL DEFAULT 'Website VIP Form',
    status lead_status NOT NULL DEFAULT 'new',
    assigned_agent_id UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 5. Performance Indexes
CREATE INDEX idx_leads_status ON leads(status);
CREATE INDEX idx_leads_created_at ON leads(created_at DESC);
CREATE INDEX idx_projects_slug ON projects(slug);
CREATE INDEX idx_units_project_id ON units(project_id);
CREATE INDEX idx_leads_project_id ON leads(project_id);
