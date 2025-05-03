CREATE TABLE Transactions (
    transaction_id VARCHAR(50) PRIMARY KEY,
    wallet_id INT REFERENCES Wallets(wallet_id) ON DELETE CASCADE,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,  
    status VARCHAR(50) CHECK (status IN ('verified', 'completed')),
    payment_method VARCHAR(50) CHECK (payment_method IN ('credit_card')),
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);
