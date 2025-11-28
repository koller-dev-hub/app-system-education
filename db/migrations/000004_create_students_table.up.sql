CREATE TABLE IF NOT EXISTS students (
    id UUID PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    enrollment_code VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(50),
    date_of_birth TIMESTAMP,
    cpf VARCHAR(20) UNIQUE,
    rg VARCHAR(20),
    
    address_zip_code VARCHAR(20),
    address_city VARCHAR(100),
    address_state VARCHAR(50),
    address_street VARCHAR(255),
    address_country VARCHAR(50),
    
    school_id UUID REFERENCES schools(id) ON DELETE SET NULL,
    school_grade VARCHAR(50),
    school_class VARCHAR(50),
    school_shift VARCHAR(20),
    enrollment_date TIMESTAMP,
    
    guardian_name VARCHAR(255),
    guardian_phone VARCHAR(50),
    guardian_email VARCHAR(255),
    guardian_cpf VARCHAR(20),
    
    is_active BOOLEAN DEFAULT TRUE,
    observations TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
