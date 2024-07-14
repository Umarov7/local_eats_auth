INSERT INTO users (id, username, email, password_hash, full_name, user_type, address, phone_number, bio, specialties, years_of_experience, is_verified)
VALUES
('123e4567-e89b-12d3-a456-426614174001', 'john_doe', 'john.doe@example.com', 'hashed_password_1', 'John Doe', 'chef', '123 Main St, Anytown, USA', '+1234567890', 'I am a chef specializing in Italian cuisine.', '{Italian, Pasta, Pizza}', 10, true),
('123e4567-e89b-12d3-a456-426614174002', 'jane_smith', 'jane.smith@example.com', 'hashed_password_2', 'Jane Smith', 'chef', '456 Elm St, Anytown, USA', '+1987654321', 'Passionate about baking and desserts.', '{Baking, Desserts, French}', 8, true),
('123e4567-e89b-12d3-a456-426614174003', 'chef_mike', 'mike@example.com', 'hashed_password_3', 'Mike Johnson', 'chef', '789 Oak St, Anytown, USA', '+1122334455', 'Specializing in seafood and Mediterranean dishes.', '{Seafood, Mediterranean}', 12, true),
('123e4567-e89b-12d3-a456-426614174004', 'sushi_master', 'sushi@example.com', 'hashed_password_4', 'Sakura Tanaka', 'chef', '101 Pine St, Anytown, USA', '+1357924680', 'Experienced in traditional Japanese sushi.', '{Japanese, Sushi}', 15, true),
('123e4567-e89b-12d3-a456-426614174005', 'pasta_lover', 'pasta@example.com', 'hashed_password_5', 'Carlo Rossi', 'chef', '222 Cedar St, Anytown, USA', '+1443322110', 'Enthusiast for all things pasta.', '{Italian, Pasta}', 5, true),
('123e4567-e89b-12d3-a456-426614174006', 'baker_extraordinaire', 'baker@example.com', 'hashed_password_6', 'Emily Baker', 'chef', '333 Maple St, Anytown, USA', '+1662777889', 'Dedicated to perfecting the art of baking.', '{Baking, Cakes}', 7, true),
('123e4567-e89b-12d3-a456-426614174007', 'chef_gordon', 'gordon@example.com', 'hashed_password_7', 'Gordon Ramsey', 'chef', '444 Walnut St, Anytown, USA', '+1555099444', 'Renowned chef known for culinary expertise.', '{International, Fine Dining}', 20, true),
('123e4567-e89b-12d3-a456-426614174008', 'bbq_master', 'bbq@example.com', 'hashed_password_8', 'Frank Miller', 'chef', '555 Pineapple St, Anytown, USA', '+1777000111', 'Passionate about smoking meats.', '{BBQ, Grilling}', 9, true),
('123e4567-e89b-12d3-a456-426614174009', 'vegan_chef', 'vegan@example.com', 'hashed_password_9', 'Lily Green', 'chef', '666 Orange St, Anytown, USA', '+1888999222', 'Specializing in creative vegan cuisine.', '{Vegan, Plant-based}', 6, true),
('123e4567-e89b-12d3-a456-426614174010', 'wine_connoisseur', 'wine@example.com', 'hashed_password_10', 'Sophia Wright', 'chef', '777 Peach St, Anytown, USA', '+1999888777', 'Passionate about exploring wine varieties.', '{Wine, Tasting}', 3, true);

INSERT INTO kitchens (id, owner_id, name, description, cuisine_type, address, phone_number, rating, total_orders)
VALUES
('223e4567-e89b-12d3-a456-426614174001', '123e4567-e89b-12d3-a456-426614174001', 'Taste of Italy', 'Authentic Italian restaurant.', 'Italian', '123 Main St, Anytown, USA', '+1234567890', 4.5, 200),
('223e4567-e89b-12d3-a456-426614174002', '123e4567-e89b-12d3-a456-426614174002', 'Sweet Delights Bakery', 'Specializes in cakes and pastries.', 'Bakery', '456 Elm St, Anytown, USA', '+1987654321', 4.8, 350),
('223e4567-e89b-12d3-a456-426614174003', '123e4567-e89b-12d3-a456-426614174003', 'Mediterranean Flavors', 'Serves Mediterranean cuisine.', 'Mediterranean', '789 Oak St, Anytown, USA', '+1122334455', 4.2, 150),
('223e4567-e89b-12d3-a456-426614174004', '123e4567-e89b-12d3-a456-426614174004', 'Sushi House', 'Offers a wide range of sushi.', 'Japanese', '101 Pine St, Anytown, USA', '+1357924680', 4.7, 280),
('223e4567-e89b-12d3-a456-426614174005', '123e4567-e89b-12d3-a456-426614174005', 'Pasta Paradise', 'Dedicated to Italian pasta dishes.', 'Italian', '222 Cedar St, Anytown, USA', '+1443322110', 4.3, 180),
('223e4567-e89b-12d3-a456-426614174006', '123e4567-e89b-12d3-a456-426614174006', "Emily's Bakery", 'Known for exquisite cakes.', 'Bakery', '333 Maple St, Anytown, USA', '+1662777889', 4.9, 400),
('223e4567-e89b-12d3-a456-426614174007', '123e4567-e89b-12d3-a456-426614174007', "Gordon's Fine Dining", 'Fine dining experience.', 'International', '444 Walnut St, Anytown, USA', '+1555099444', 4.6, 320),
('223e4567-e89b-12d3-a456-426614174008', '123e4567-e89b-12d3-a456-426614174008', "Frank's BBQ Joint", 'Famous for smoked BBQ.', 'BBQ', '555 Pineapple St, Anytown, USA', '+1777000111', 4.4, 250),
('223e4567-e89b-12d3-a456-426614174009', '123e4567-e89b-12d3-a456-426614174009', 'Green Leaf Cafe', 'Creative vegan dishes.', 'Vegan', '666 Orange St, Anytown, USA', '+1888999222', 4.1, 120),
('223e4567-e89b-12d3-a456-426614174010', '123e4567-e89b-12d3-a456-426614174010', 'Wine & Dine', 'Offers a variety of wines.', 'Wine Bar', '777 Peach St, Anytown, USA', '+1999888777', 4.0, 100);
