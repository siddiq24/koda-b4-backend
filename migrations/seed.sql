INSERT INTO users(email, password, role) VALUES
('admin@fore.com', '$argon2id$v=19$m=65536,t=3,p=4$nVh9LrnucqK4MkZi1Owhmg$U8cdlJ8doFHl3S4RFoW2V69/zma8dF7dsW0ttnWLd70', 'admin'),
('user1@fore.com', '$argon2id$v=19$m=65536,t=3,p=4$pHWWe5mAbQKnuS1CilIIMA$pgsX/wwf7AZnmDOGSOaM0dEni1FLf9EutCYlGnmk2bQ', 'user'),
('user2@fore.com', '$argon2id$v=19$m=65536,t=3,p=4$LeFCsef0fswy/nzgkn138A$jJxY9lg/Zn5v/roCVb1XeIyOadVcYWAZp3G9orna4ZA', 'user'),
('user3@fore.com', '$argon2id$v=19$m=65536,t=3,p=4$TF81u3i9Y0yNCRlE5yg+2g$nPIs9LwKIhwzxjP6O+Z8V6dfA709QB2xRllfscZQhjU', 'user'),
('user4@fore.com', '$argon2id$v=19$m=65536,t=3,p=4$vTWJu7fyPpLgYxbtwzyypA$aJL1Oix2FRWoqiykR309QhKBA4jsyh3qbjG4in1l08g', 'user'),
('user5@fore.com', '$argon2id$v=19$m=65536,t=3,p=4$s9LBNLTFB5CycQ1FKeRNdA$PpdvpQy9gUEupnh0nub0XRCQDBCtQyP4Yt7GHaLf7e8', 'user');

-- Insert data untuk tabel categories
INSERT INTO categories (name) VALUES
('Coffee'),
('Non-Coffee'),
('Fruit Tea'),
('Tea'),
('Food'),
('Ice Blended'),
('Signature Coffee'),
('Origin Coffee');

-- Insert data untuk tabel users

-- Insert data untuk tabel sizes
INSERT INTO sizes (name, additional_price) VALUES
('S', 0),
('M', 5000),
('L', 10000),
('XL', 15000);

-- Insert data untuk tabel variants
INSERT INTO variants (name, additional_price) VALUES
('Less Sugar', 0),
('More Sugar', 3000),
('No Sugar', 0),
('Less Ice', 0),
('More Ice', 0),
('No Ice', 0),
('Hot', 3000),
('Iced', 0);

INSERT INTO products (title, description, base_price, stock, category_id) VALUES
('Ice Americano', 'An espresso shot mixed with a glass of water, delivering an ideal balance of character, aroma, and flavor.', 22000, 200, 1),
('Iced Bumi Latte', 'The creamy and subtly sweet sensation of caramel and butterscotch sauce blends with authentic Indonesian coffee', 24000, 200, 1),
('Iced Cappuccino', 'Blend of espresso and milk with a thick layer of foam on top without additional sugar.', 29000, 200, 1),
('Double Iced Shaken Latte', 'Classic blend of 2 shot espresso with milk and cream.', 33000, 200, 1),
('Iced Café Latte', 'Espresso blend and milk with a thin layer of foam on top without additional sugar.', 29000, 200, 1),

('Nutty Oat Latte', 'Espresso from Indonesian specialty coffee beans combined with gluten-free oat milk and a nutty sensation from hazelnuts.', 39000, 200, 1),
('Iced Caramel Praline Macchiato', 'Latte with Praline and caramel sauce with a sweet taste and aroma.', 33000, 200, 1),
('Iced Dark Chocolate', 'Made from 100% dark chocolate and milk', 34000, 200, 2),
('Iced Almond Choco', 'Rich chocolate drink with almond flavor and fresh milk', 39000, 200, 2),
('Iced Matcha Green Tea', 'Fore`s special matcha powder is soft and creamy combined with fresh milk', 34000, 200, 2),

('Iced Classic Milo', 'A classic malt chocolate drink with creamy sauce and creamy milk', 25000, 200, 2),
('Iced Coco Peach Fusion', 'Refreshing flavors of peach, lychee, orange citrus, and coconut water', 29000, 200, 3),
('Hibiscus Lychee Peach Yakult', 'The goodness of berries, hibiscus, and yakult which provide freshness all day long.', 29000, 200, 3),
('Sunny Citrus Jasmine', 'A special blend of tropical fruit flavors, Manuka honey and jasmine tea that is both refreshing and flavorful.', 31000, 200, 3),
('Iced English Breakfast', 'Legendary fragrant classic black tea', 29000, 200, 4),

('Iced Green Tea Jasmine', 'Green tea with a touch of jasmine flavor', 29000, 200, 4),
('Iced Pure Chamomile', 'Chamomile flower tea has a rich honey aroma and imparts a golden color', 29000, 200, 4),
('Iced Creme Caramel Tea', 'Red tea and a blend of spices with a sweet taste', 29000, 200, 4),
('Iced Green Tea Mint', 'Green tea with refreshing mint flavor', 29000, 200, 4),
('Klapertart Croissant', 'Twist of Eastern Indonesia`s dessert (coconut custard, almond, and raisin) filled in flaky croissant', 33000, 50, 5),

('Beef Mentai Sandwich', 'Sandwich with sliced beef, mentai sauce, red cheddar cheese and red paprika', 39000, 50, 5),
('Cakalang Quiche', 'Savoury pie with egg, cream, cheese and cakalang fish filling', 36000, 50, 5),
('Pain au Tiramisu', 'Pastry with tiramisu flavoured almond paste and cocoa crumble topping', 36000, 50, 5),
('Matcha Strawberry Cake', 'Sliced cake with combination flavour of matcha and strawberry', 27000, 50, 5),
('Mushroom Truffle Sandwich', 'Sandwich with parsley topping bread, mushroom, truffle sauce and cheddar cheese', 42000, 50, 5),

('Cheesy Tuna Sandwich', 'Sandwich with black sesame topping bread, tuna mayo, tartar sauce, mozzarella cheese, red cheddar cheese, roasted onion and roasted green paprika.', 39000, 50, 5),
('Chicken Teriyaki Sandwich', 'Sandwich with white sesame topping bread, chicken teriyaki, teriyaki mayonnaise, cheddar cheese and roasted onion', 39000, 50, 5),
('Blueberry Cheese Muffin', 'Vanilla flavoured muffin with blueberry cream cheese filling and crumble topping', 36000, 50, 5),
('Choco Melt Muffin', 'Chocolate flavoured muffin with melted chocolate fillings and choco chips topping', 36000, 50, 5),
('Smoked Beef & Cheese Croissant', 'Danish filled with smoked beef and cheddar cheese', 36000, 50, 5),

('Triple Cheese Danish', 'Danish with cream cheese filling', 36000, 50, 5),
('Almond Croissant', 'Croissants sprinkled with powdered sugar and chopped almonds', 36000, 50, 5),
('Banana Chocolate Cake', 'Sponge cake with a mix of banana and chocolate', 27000, 50, 5),
('Cempedak Cake', 'Sponge cake with chunks of cempedak', 27000, 50, 5),
('Butter Croissant', 'Pastry with taste and aroma of butter', 24000, 50, 5),

('Pain au Chocolat', 'French pastry with chocolate filling', 29000, 50, 5),
('Kouign-Amann', 'Pastry with a layer of sugar and sprinkle of cinnamon', 29000, 50, 5),
('Caramel Praline Coffee Ice Blended', 'Ice blended latte with Praline and caramel sauce.', 33000, 60, 6),
('Matcha Ice Blended', 'Fore Coffee`s signature Creamy Matcha blend, fresh milk, with ice, just right to cool your day!', 33000, 60, 6),
('Strawberry Ice Blended', 'Refreshing strawberry blend with ice and milk', 33000, 60, 6),

('Chocolate Ice Blended', 'Ice blended special chocolate with selected milk', 36000, 60, 6),
('Iced Kopi Dari Tani', 'Bold flavors and aromas of Indonesian coffee pair with the richness of authentic Indonesian palm sugar', 24000, 70, 7),
('Iced Butterscotch Sea Salt Latte', 'A blend of butterscotch and house blend with a soft cream sea salt topping on top', 33000, 70, 7),
('Iced Buttercream Latte', 'Buttercream Coffee topping on caramel latte', 31000, 70, 7),
('Iced Aren Latte', 'The natural taste of palm sugar blends perfectly with premium espresso. Fore Coffee`s best selling menu', 29000, 70, 7),

('Iced Pandan Latte', 'Latte with a unique taste and aroma from natural pandan extract. Special menu from Fore Coffee', 29000, 70, 7),
('Iced Aceh Gayo', 'Aromatic coffee with notes of chocolate, butterscotch and spices', 24000, 80, 8),
('Iced Toraja Sapan', 'Aromatic coffee with notes of citrus, spices and molasses', 24000, 80, 8),
('Iced Bali Kintamani', 'Aromatic coffee with fruity and citrus notes', 24000, 80, 8),
('Iced Malty Latte', 'The irresistible Cafe Malt Latte revolutionized with a bold and sweet sensation.', 27000, 80, 8);


INSERT INTO products_images( product_id, image ) VALUES
(1,'https://static.fore.coffee/product/Americano%20Iced.jpg'),
(2,'https://static.fore.coffee/product/Bumi%20Latte%20w%20Badge.jpg'),
(3, 'https://static.fore.coffee/product/Capucino%20Iced%20(1).jpg'),
(4, 'https://static.fore.coffee/product/Double%20Iced%20Shaken%20Latte%20(1).jpg'),
(5, 'https://static.fore.coffee/product/Americano%20Iced.jpg'),
(6, 'https://static.fore.coffee/product/Nutty%20Oat%20Latte%20Iced.jpg'),
(7, 'https://static.fore.coffee/product/Nutty%20Oat%20Latte%20Iced.jpg'),
(8, 'https://static.fore.coffee/product/darkchocolate-01.jpg'),
(9, 'https://static.fore.coffee/product/almondchocoiced173.jpg'),
(10, 'https://static.fore.coffee/product/almondchocoiced173.jpg'),
(11, 'https://static.fore.coffee/product/classicmiloiced173.jpg'),
(12, 'https://static.fore.coffee/product/Coco%20Peach%20Fusion%20(3).jpg'),
(13, 'https://static.fore.coffee/product/hibiscuslychee173.jpg'),
(14, 'https://static.fore.coffee/product/sunnycitrus173.jpg'),
(15, 'https://static.fore.coffee/product/englishbreakfasticed173.jpg'),
(15, 'https://static.fore.coffee/product/englishbreakfast173.jpg'),
(16, 'https://static.fore.coffee/product/greenteajasmineiced173.jpg'),
(16, 'https://static.fore.coffee/product/greenteajasmine173.jpg'),
(17, 'https://static.fore.coffee/product/purechamomileiced173.jpg'),
(18, 'https://static.fore.coffee/product/cremecaramelteaiced173.jpg'),
(19, 'https://static.fore.coffee/product/greenteaminticed173.jpg'),
(20, 'https://static.fore.coffee/product/klapetart%20Thumbnail.png'),
(21, 'https://static.fore.coffee/product/Creamy%20Beef%20Mentai%20Sandwich.png'),
(22, 'https://static.fore.coffee/product/Cakalang%20Quiche.png'),
(23, 'https://static.fore.coffee/product/Pain%20au%20Tiramisu.png'),
(24, 'https://static.fore.coffee/product/Strawberry%20Matcha%20Cake%20Zoomed%20(1).png'),
(25, 'https://static.fore.coffee/product/Mushroom%20Truffle%20Sandwich.png'),
(26, 'https://static.fore.coffee/product/Cheesy%20Tuna%20Sandwich.png'),
(27, 'https://static.fore.coffee/product/Chicken%20Teriyaki%20Sandwich.png'),
(28, 'https://static.fore.coffee/product/Blueberry%20Cheese%20Muffin.png'),
(29, 'https://static.fore.coffee/product/Choco%20Melt%20Muffin.png'),
(30, 'https://static.fore.coffee/product/Smoked%20Beef%20_%20Cheese%20Croissant.jpg'),
(31, 'https://static.fore.coffee/product/Triple%20Cheese%20Danish.jpg'),
(32, 'https://static.fore.coffee/product/Almond%20Croissant%20-1.jpg'),
(33, 'https://static.fore.coffee/product/thumbbbb.jpg'),
(34, 'https://static.fore.coffee/product/cempedak-80.jpg'),
(35, 'https://static.fore.coffee/product/Butter%20Croissant%20_-80.jpg'),
(36, 'https://static.fore.coffee/product/painauchocolat2403.jpg'),
(37, 'https://static.fore.coffee/product/Kouign%20amann-.jpg'),
(38, 'https://static.fore.coffee/product/caramelpralinecoffee173.jpg'),
(39, 'https://static.fore.coffee/product/matchablended173.jpg'),
(40, 'https://static.fore.coffee/product/strawberryblend173.jpg'),
(41, 'https://static.fore.coffee/product/chocolateblend173.jpg'),
(42, 'https://static.fore.coffee/product/Kopi%20dari%20Tani%20w%20Badge.jpg'),
(43, 'https://static.fore.coffee/product/Butterscoth%20Iced.jpg'),
(44, 'https://static.fore.coffee/product/Buttercream%20Latte%20(1).jpg'),
(45, 'https://static.fore.coffee/product/Aren%20Latte%20Ice.jpg'),
(45, 'https://static.fore.coffee/product/Aren%20Latte%20Hot.jpg'),
(46, 'https://static.fore.coffee/product/Pandan%20Latte%20Iced.jpg'),
(46, 'https://static.fore.coffee/product/Pandan%20Latte%20Hot.jpg'),
(47, 'https://static.fore.coffee/product/ICED%20COD.jpg'),
(47, 'https://static.fore.coffee/product/COD_satuan-01.jpg'),
(48, 'https://static.fore.coffee/product/ICED%20COD.jpg'),
(48, 'https://static.fore.coffee/product/COD_satuan-02.jpg'),
(49, 'https://static.fore.coffee/product/ICED%20COD.jpg'),
(49, 'https://static.fore.coffee/product/COD_satuan-03.jpg'),
(50, 'https://static.fore.coffee/product/Malty%20Latte.jpg')
;

-- Seed data untuk products_variants
INSERT INTO products_variants (product_id, variant_id) VALUES
-- Minuman dingin biasanya memiliki variant sugar dan ice
(1, 1), (1, 2), (1, 3), (1, 4), (1, 5), (1, 6), (1, 8),  -- Ice Americano
(2, 1), (2, 2), (2, 3), (2, 4), (2, 5), (2, 6), (2, 8),  -- Iced Bumi Latte
(3, 1), (3, 2), (3, 3), (3, 4), (3, 5), (3, 6), (3, 8),  -- Iced Cappuccino
(4, 1), (4, 2), (4, 3), (4, 4), (4, 5), (4, 6), (4, 8),  -- Double Iced Shaken Latte
(5, 1), (5, 2), (5, 3), (5, 4), (5, 5), (5, 6), (5, 8),  -- Iced Café Latte
-- Minuman kopi lainnya
(6, 1), (6, 2), (6, 3), (6, 4), (6, 5), (6, 6), (6, 7), (6, 8),  -- Nutty Oat Latte
(7, 1), (7, 2), (7, 3), (7, 4), (7, 5), (7, 6), (7, 8),  -- Iced Caramel Praline Macchiato
(8, 1), (8, 2), (8, 3), (8, 4), (8, 5), (8, 6), (8, 8),  -- Iced Dark Chocolate
(9, 1), (9, 2), (9, 3), (9, 4), (9, 5), (9, 6), (9, 8),  -- Iced Almond Choco
-- Minuman tea based
(10, 1), (10, 2), (10, 3), (10, 4), (10, 5), (10, 6), (10, 8),  -- Iced Matcha Green Tea
(11, 1), (11, 2), (11, 3), (11, 4), (11, 5), (11, 6), (11, 8),  -- Iced Classic Milo
(12, 1), (12, 2), (12, 3), (12, 4), (12, 5), (12, 6), (12, 8),  -- Iced Coco Peach Fusion
(13, 1), (13, 2), (13, 3), (13, 4), (13, 5), (13, 6), (13, 8),  -- Hibiscus Lychee Peach Yakult
(14, 1), (14, 2), (14, 3), (14, 4), (14, 5), (14, 6), (14, 8),  -- Sunny Citrus Jasmine
(15, 1), (15, 2), (15, 3), (15, 4), (15, 5), (15, 6), (15, 8),  -- Iced English Breakfast
(16, 1), (16, 2), (16, 3), (16, 4), (16, 5), (16, 6), (16, 8),  -- Iced Green Tea Jasmine
(17, 1), (17, 2), (17, 3), (17, 4), (17, 5), (17, 6), (17, 8),  -- Iced Pure Chamomile
(18, 1), (18, 2), (18, 3), (18, 4), (18, 5), (18, 6), (18, 8),  -- Iced Creme Caramel Tea
(19, 1), (19, 2), (19, 3), (19, 4), (19, 5), (19, 6), (19, 8),  -- Iced Green Tea Mint
-- Makanan/pastry (hanya tersedia dalam keadaan normal)
(20, 1),  -- Klapertart Croissant (default Less Sugar)
(21, 1),  -- Beef Mentai Sandwich
(22, 1),  -- Cakalang Quiche
(23, 1),  -- Pain au Tiramisu
(24, 1),  -- Matcha Strawberry Cake
(25, 1),  -- Mushroom Truffle Sandwich
(26, 1),  -- Cheesy Tuna Sandwich
(27, 1),  -- Chicken Teriyaki Sandwich
(28, 1),  -- Blueberry Cheese Muffin
(29, 1),  -- Choco Melt Muffin
(30, 1),  -- Smoked Beef & Cheese Croissant
(31, 1),  -- Triple Cheese Danish
(32, 1),  -- Almond Croissant
(33, 1),  -- Banana Chocolate Cake
(34, 1),  -- Cempedak Cake
(35, 1),  -- Butter Croissant
(36, 1),  -- Pain au Chocolat
(37, 1),  -- Kouign-Amann
-- Ice Blended drinks
(38, 1), (38, 2), (38, 3), (38, 4), (38, 5), (38, 6),  -- Caramel Praline Coffee Ice Blended
(39, 1), (39, 2), (39, 3), (39, 4), (39, 5), (39, 6),  -- Matcha Ice Blended
(40, 1), (40, 2), (40, 3), (40, 4), (40, 5), (40, 6),  -- Strawberry Ice Blended
(41, 1), (41, 2), (41, 3), (41, 4), (41, 5), (41, 6),  -- Chocolate Ice Blended
-- Minuman kopi spesial
(42, 1), (42, 2), (42, 3), (42, 4), (42, 5), (42, 6), (42, 7), (42, 8),  -- Iced Kopi Dari Tani
(43, 1), (43, 2), (43, 3), (43, 4), (43, 5), (43, 6), (43, 7), (43, 8),  -- Iced Butterscotch Sea Salt Latte
(44, 1), (44, 2), (44, 3), (44, 4), (44, 5), (44, 6), (44, 7), (44, 8),  -- Iced Buttercream Latte
(45, 1), (45, 2), (45, 3), (45, 4), (45, 5), (45, 6), (45, 7), (45, 8),  -- Iced Aren Latte
(46, 1), (46, 2), (46, 3), (46, 4), (46, 5), (46, 6), (46, 7), (46, 8),  -- Iced Pandan Latte
(47, 1), (47, 2), (47, 3), (47, 4), (47, 5), (47, 6), (47, 7), (47, 8),  -- Iced Aceh Gayo
(48, 1), (48, 2), (48, 3), (48, 4), (48, 5), (48, 6), (48, 7), (48, 8),  -- Iced Toraja Sapan
(49, 1), (49, 2), (49, 3), (49, 4), (49, 5), (49, 6), (49, 7), (49, 8),  -- Iced Bali Kintamani
(50, 1), (50, 2), (50, 3), (50, 4), (50, 5), (50, 6), (50, 7), (50, 8);  -- Iced Malty Latte


-- Insert data untuk tabel products_sizes (menghubungkan semua products dengan semua sizes)
INSERT INTO products_sizes (product_id, size_id)
SELECT p.id, s.id 
FROM products p, sizes s;

-- Insert data untuk tabel promos (promo berlaku untuk semua products)
INSERT INTO promos (title, description, discount, start, "end") VALUES
('Grand Opening', 'Special discount for all products during grand opening', 15.0, '2024-01-01 00:00:00', '2025-12-31 23:59:59'),
('Weekend Special', 'Weekend special discount for all products', 10.0, '2024-01-01 00:00:00', '2025-12-31 23:59:59'),
('Coffee Lover', 'Special discount for all coffee products', 20.0, '2024-01-01 00:00:00', '2025-12-31 23:59:59'),
('Summer Sale', 'Summer special discount for all products', 25.0, '2024-06-01 00:00:00', '2025-12-31 23:59:59'),
('Flash Sale', 'Flash sale for all products - limited time!', 30.0, '2024-03-01 00:00:00', '2025-12-07 23:59:59');

-- Insert data untuk tabel products_promos (menghubungkan products dengan promos)
INSERT INTO products_promos (product_id, promo_id) VALUES
(1, 1), (2, 1), (3, 1),  -- Beberapa produk dapat promo Grand Opening
(4, 2), (5, 2),          -- Beberapa produk dapat promo Weekend Special
(6, 3), (7, 3), (8, 3),  -- Beberapa produk dapat promo Coffee Lover
(9, 4), (10, 4);         -- Beberapa produk dapat promo Summer Sale


-- Karena promo berlaku untuk semua products, tidak perlu mapping spesifik

-- Insert data untuk tabel deliveries
INSERT INTO deliveries (name) VALUES
('Pickup in Store'),
('Home Delivery'),
('Express Delivery'),
('Grab Delivery'),
('Gojek Delivery');

-- Insert data untuk tabel status
INSERT INTO status (name) VALUES
('Pending'),
('Paid'),
('Preparing'),
('Ready for Pickup'),
('On Delivery'),
('Completed'),
('Cancelled');

-- Insert data untuk tabel payment_methods
INSERT INTO payment_methods( name, image, no_va) VALUES
('MANDIRI', 'https://cdn3.iconfinder.com/data/icons/banks-in-indonesia-logo-badge/100/Mandiri-512.png', ''),
('BRI', 'https://cdn3.iconfinder.com/data/icons/banks-in-indonesia-logo-badge/100/BRI-512.png', ''),
('BTN', 'https://cdn3.iconfinder.com/data/icons/banks-in-indonesia-logo-badge/100/Bank_BTN-512.png', ''),
('BCA', 'https://cdn3.iconfinder.com/data/icons/banks-in-indonesia-logo-badge/100/BCA-512.png', ''),
('OVO', 'https://i.pinimg.com/736x/61/c9/8a/61c98a1dffc2e04424d592564cef941f.jpg', ''),
('DANA', 'https://i.pinimg.com/1200x/cb/aa/03/cbaa0388892e0a154353c2a1cb8b3fee.jpg', ''),
('QRIS', 'https://i.pinimg.com/1200x/43/41/38/434138932e81512fe14236d29f09c7c3.jpg', '')
;

-- Insert data untuk tabel tags
INSERT INTO tags (name) VALUES
('Best Seller'),
('New'),
('Seasonal'),
('Limited Edition'),
('Vegan'),
('Gluten Free'),
('Sweet'),
('Refreshing'),
('Creamy'),
('Strong');

-- Insert data untuk tabel products_tags (menghubungkan products dengan tags)
INSERT INTO products_tags (product_id, tag_id) VALUES
(1, 1), (1, 10),   -- Ice Americano: Best Seller, Strong
(2, 1), (2, 7),    -- Iced Bumi Latte: Best Seller, Sweet
(3, 8), (3, 9),    -- Iced Cappuccino: Refreshing, Creamy
(6, 6), (6, 5),    -- Nutty Oat Latte: Gluten Free, Vegan
(10, 2), (10, 9),  -- Iced Matcha Green Tea: New, Creamy
(15, 8), (15, 3),  -- Sunny Citrus Jasmine: Refreshing, Seasonal
(20, 4), (20, 2),  -- Klapertart Croissant: Limited Edition, New
(25, 1), (25, 7),  -- Mushroom Truffle Sandwich: Best Seller, Sweet
(30, 1), (30, 9),  -- Caramel Praline Coffee Ice Blended: Best Seller, Creamy
(35, 1), (35, 7);  -- Iced Aren Latte: Best Seller, Sweet

-- Insert data untuk tabel profiles

-- Insert data untuk tabel carts (contoh keranjang belanja)
INSERT INTO carts (user_id, product_id, size_id, varian_id, qty, subtotal, product_name) VALUES
(2, 1, 2, 8, 2, 54000, 'Ice Americano - M - Iced'),
(2, 3, 3, 8, 1, 39000, 'Iced Cappuccino - L - Iced'),
(3, 5, 2, 3, 1, 34000, 'Iced Café Latte - M - No Sugar'),
(3, 10, 3, 8, 2, 88000, 'Iced Matcha Green Tea - L - Iced'),
(4, 20, 1, NULL, 1, 33000, 'Klapertart Croissant');

-- Insert data untuk tabel orders_products (contoh detail order)
INSERT INTO orders_products (invoice, product_id, size_id, varian_id, qty, subTotal, name) VALUES
('INV-001-2024', 1, 2, 8, 2, 54000, 'Ice Americano - M - Iced'),
('INV-001-2024', 3, 3, 8, 1, 39000, 'Iced Cappuccino - L - Iced'),
('INV-002-2024', 5, 2, 3, 1, 34000, 'Iced Café Latte - M - No Sugar'),
('INV-003-2024', 10, 3, 8, 2, 88000, 'Iced Matcha Green Tea - L - Iced'),
('INV-004-2024', 20, 1, NULL, 1, 33000, 'Klapertart Croissant');
 TABLE orders_products;

-- Insert data untuk tabel orders (contoh order dengan promo)
INSERT INTO orders (user_id, email, fullname, phone, address, payment_method_id, delivery_id, total_order, invoice, status_id, promo_id) VALUES
(2, 'user1@example.com', 'John Doe', '081298765432', 'Jl. Thamrin No. 45, Jakarta', 1, 2, 79050, 'INV-001-2024', 6, 1),  -- Dengan promo Grand Opening 15%
(3, 'user2@example.com', 'Sarah Connor', '085712345678', 'Jl. Gatot Subroto No. 67, Jakarta', 3, 1, 30600, 'INV-002-2024', 6, 2),  -- Dengan promo Weekend Special 10%
(4, 'john.doe@email.com', 'Michael Brown', '087812345679', 'Jl. MH Thamrin No. 89, Jakarta', 5, 3, 66000, 'INV-003-2024', 3, 3),  -- Dengan promo Coffee Lover 20%
(2, 'user1@example.com', 'John Doe', '081298765432', 'Jl. Thamrin No. 45, Jakarta', 2, 1, 24750, 'INV-004-2024', 2, 4);  -- Dengan promo Summer Sale 25%

-- Insert data contoh untuk users_tokens (jika diperlukan untuk testing)


-- Insert data contoh untuk forgot_pass_token (jika diperlukan untuk testing)

UPDATE products
SET is_favorite = TRUE
WHERE id < 10;

SELECT * 
FROM products p
LEFT JOIN products_variants pv ON pv.product_id = p.id
LEFT JOIN variants v ON v.id = pv.variant_id
WHERE p.id=2;

SELECT *
FROM products_variants
WHERE product_id = 2;


SELECT 
			o.invoice,
			o.fullname as cust_name,
			o.phone as cust_phone,
			o.email as cust_email,
			o.address as cust_address,
			pm.name as payment_method,
			d.name as delivery_method,
			s.name as status,
			o.total_order as total,
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'product_name', op.name,
					'quantity', op.qty,
					'subtotal', op.subtotal,
					'size', COALESCE(sz.name, ''),
					'variant', COALESCE(v.name, ''),
					'image', i.image
				)
			) as items,
			o.created_at,
			o.updated_at
		FROM orders o
		LEFT JOIN payment_methods pm ON o.payment_method_id = pm.id
		LEFT JOIN deliveries d ON o.delivery_id = d.id
		LEFT JOIN status s ON o.status_id = s.id
		LEFT JOIN orders_products op ON o.invoice = op.invoice
		LEFT JOIN sizes sz ON op.size_id = sz.id
		LEFT JOIN variants v ON op.varian_id = v.id
		LEFT JOIN products_images i ON i.product_id = op.product_id
		WHERE o.invoice = 'VIA-1763757135-9' AND o.user_id = 9
		GROUP BY 
			o.invoice, 
			o.fullname, 
			o.phone, 
			o.email, 
			o.address, 
			pm.name, 
			d.name, 
			s.name, 
			o.total_order,
			o.created_at,
			o.updated_at;

SELECT COALESCE(json_agg(json_build_object('id', s.id, 'name', s.name)), '[]') 
			FROM products_sizes ps LEFT JOIN sizes s ON ps.size_id = s.id WHERE ps.product_id = 5;

UPDATE products
SET deleted_at = NULL;

SELECT id, title, description FROM promos WHERE "end" >= CURRENT_DATE;

SELECT p.id, p.title, p.description, p.base_price, COALESCE(pr.discount, 0) AS discount, c.name AS category_name, COALESCE(ARRAY_AGG(DISTINCT i.image) FILTER (WHERE i.image IS NOT NULL), '{}') AS images, COALESCE(json_agg(DISTINCT jsonb_build_object('id', sz.id, 'name', sz.name)) FILTER (WHERE sz.id IS NOT NULL), '[]') AS sizes FROM products p LEFT JOIN categories c ON c.id = p.category_id LEFT JOIN products_images i ON i.product_id = p.id LEFT JOIN products_sizes ps ON ps.product_id = p.id LEFT JOIN sizes sz ON sz.id = ps.size_id LEFT JOIN products_promos pp ON pp.product_id = p.id LEFT JOIN promos pr ON pr.id = pp.promo_id WHERE p.deleted_at IS NULL GROUP BY p.id, c.name, pr.discount ORDER BY p.id DESC LIMIT 10 OFFSET 0