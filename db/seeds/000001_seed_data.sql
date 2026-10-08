-- Seed Script for Initial Binghatti Data & Admin Accounts

-- 1. Insert Default Super Admin Account (Email: sales@binghatti.com / Password: password123)
INSERT INTO users (id, full_name, email, password_hash, role)
VALUES (
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11',
    'Senior Sales Advisor',
    'sales@binghatti.com',
    '$2a$10$7R0Z4QnL4W4cTjG7F5hW9.3D8H2U0B1V3N4C5M6K7L8P9Q0R1S2T3U4V5W', -- Bcrypt hash for password123
    'super_admin'
) ON CONFLICT (email) DO NOTHING;

-- 2. Insert Featured Luxury Projects
INSERT INTO projects (id, slug, title, tagline, location_name, starting_price_aed, handover_date, description, brochure_url, hero_video_url)
VALUES 
(
    'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22',
    'mercedes-benz-places',
    'Mercedes-Benz Places by Binghatti',
    'Iconic Architecture in Downtown Dubai',
    'Downtown Dubai Canal Frontage',
    8800000.00,
    'Q4 2026',
    'A 65-story architectural masterpiece integrating smart home technology and aerodynamic Mercedes-Benz design cues.',
    'https://www.binghatti.com/brochures/mercedes-benz-places.pdf',
    'https://www.binghatti.com/videos/mercedes-benz-places.mp4'
),
(
    'c2eebc99-9c0b-4ef8-bb6d-6bb9bd380a33',
    'bugatti-residences',
    'Bugatti Residences by Binghatti',
    'Hyper-Living in Business Bay',
    'Business Bay Prime Waterfront',
    19000000.00,
    'Q4 2027',
    'Inspired by the iconic Riviera and Bugatti hypercars, featuring private car lifts into sky mansions.',
    'https://www.binghatti.com/brochures/bugatti-residences.pdf',
    'https://www.binghatti.com/videos/bugatti-residences.mp4'
),
(
    'd3eebc99-9c0b-4ef8-bb6d-6bb9bd380a44',
    'jacob-and-co-residences',
    'Burj Binghatti Jacob & Co Residences',
    'The Tallest Residential Tower in the World',
    'Business Bay Financial Center',
    8000000.00,
    'Q2 2028',
    'Crown-shaped residential skyscraper adorned with haute-horlogerie jewels and 360-degree skyline views.',
    'https://www.binghatti.com/brochures/jacob-co.pdf',
    'https://www.binghatti.com/videos/jacob-co.mp4'
) ON CONFLICT (slug) DO NOTHING;

-- 3. Insert Initial Units & Floor Plans
INSERT INTO units (id, project_id, unit_type, size_sqft, price_aed, floor_plan_image_url, floor_plan_pdf_url, status)
VALUES
(
    'e4eebc99-9c0b-4ef8-bb6d-6bb9bd380a55',
    'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22',
    '1 Bedroom Suite',
    1050.00,
    1450000.00,
    'https://images.unsplash.com/photo-1560448204-e02f11c3d0e2?auto=format&fit=crop&w=800&q=80',
    'https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf',
    'available'
),
(
    'f5eebc99-9c0b-4ef8-bb6d-6bb9bd380a66',
    'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22',
    '2 Bedroom Royal',
    1650.00,
    2850000.00,
    'https://images.unsplash.com/photo-1600607687939-ce8a6c25118c?auto=format&fit=crop&w=800&q=80',
    'https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf',
    'reserved'
),
(
    'a6eebc99-9c0b-4ef8-bb6d-6bb9bd380a77',
    'c2eebc99-9c0b-4ef8-bb6d-6bb9bd380a33',
    'Sky Penthouse Mansion',
    4200.00,
    19500000.00,
    'https://images.unsplash.com/photo-1600566753376-12c8ab7fb75b?auto=format&fit=crop&w=800&q=80',
    'https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf',
    'available'
) ON CONFLICT DO NOTHING;

-- 4. Insert Initial Sample VIP Leads
INSERT INTO leads (id, project_id, full_name, whatsapp_number, email, budget_range, investment_purpose, lead_source, status)
VALUES
(
    'b7eebc99-9c0b-4ef8-bb6d-6bb9bd380a88',
    'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22',
    'Alexander Vance',
    '+971 50 888 9911',
    'alexander.vance@investor.com',
    '$1M+',
    'Investment',
    'Website VIP Form',
    'new'
),
(
    'c8eebc99-9c0b-4ef8-bb6d-6bb9bd380a99',
    'c2eebc99-9c0b-4ef8-bb6d-6bb9bd380a33',
    'Fatima Al-Maktoum',
    '+971 52 333 4455',
    'fatima@dubai-capital.ae',
    '$500k-$1M',
    'Self-use',
    'Website VIP Form',
    'contacted'
) ON CONFLICT DO NOTHING;
