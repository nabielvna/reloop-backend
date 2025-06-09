package main

import (
	"log"
	"os"
	"reloop-backend/internal/db"
	"reloop-backend/internal/models"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	// Load .env if exists
	godotenv.Load()

	// Connect to database
	dbAddr := os.Getenv("DB_ADDR")
	if dbAddr == "" {
		dbAddr = "./database/reloop.db"
	}

	database, err := db.New(dbAddr, 30, 30, "15m")
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	log.Println("🌱 Seeding database with sample data...")

	// Seed Categories
	seedCategories(database)

	// Seed Users
	seedUsers(database)

	// Seed Sellers
	seedSellers(database)

	// Seed Items
	seedItems(database)

	log.Println("🎉 Database seeding completed!")
	log.Println("📊 Sample data created:")
	log.Println("   👤 Users: admin@reloop.com, seller@reloop.com, user@reloop.com")
	log.Println("   🏪 Sellers: Tech Store, Fashion Hub")
	log.Println("   📱 Items: MacBook Pro, iPhone 15, Gaming Laptop")
	log.Println("   🏷️ Categories: Electronics, Books, Clothing, etc.")
	log.Println("")
	log.Println("🔑 Test credentials:")
	log.Println("   Email: seller@reloop.com")
	log.Println("   Password: 12345678")
}

func seedCategories(db *gorm.DB) {
	categories := []models.Category{
		{Name: "Electronics", IsActive: true},
		{Name: "Books", IsActive: true},
		{Name: "Clothing", IsActive: true},
		{Name: "Home & Garden", IsActive: true},
		{Name: "Sports & Outdoor", IsActive: true},
		{Name: "Automotive", IsActive: true},
		{Name: "Beauty & Personal Care", IsActive: true},
		{Name: "Toys & Games", IsActive: true},
	}

	for _, category := range categories {
		var existing models.Category
		if err := db.Where("name = ?", category.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("❌ Failed to create category %s: %v", category.Name, err)
			} else {
				log.Printf("✅ Created category: %s", category.Name)
			}
		} else {
			log.Printf("⏭️  Category already exists: %s", category.Name)
		}
	}
}

func seedUsers(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("12345678"), 12)

	users := []models.User{
		{Email: "admin@reloop.com", PasswordHash: string(password), UserName: "Administrator", Role: "admin"},
		{Email: "seller@reloop.com", PasswordHash: string(password), UserName: "TechSeller", Role: "user"},
		{Email: "seller2@reloop.com", PasswordHash: string(password), UserName: "FashionSeller", Role: "user"},
		{Email: "user@reloop.com", PasswordHash: string(password), UserName: "RegularUser", Role: "user"},
		{Email: "user2@reloop.com", PasswordHash: string(password), UserName: "TestUser", Role: "user"},
	}

	for _, user := range users {
		var existing models.User
		if err := db.Where("email = ?", user.Email).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&user).Error; err != nil {
				log.Printf("❌ Failed to create user %s: %v", user.Email, err)
			} else {
				log.Printf("✅ Created user: %s (%s)", user.Email, user.Role)
			}
		} else {
			log.Printf("⏭️  User already exists: %s", user.Email)
		}
	}
}

func seedSellers(db *gorm.DB) {
	// Get users to make them sellers
	var seller1, seller2 models.User
	db.Where("email = ?", "seller@reloop.com").First(&seller1)
	db.Where("email = ?", "seller2@reloop.com").First(&seller2)

	sellers := []models.Seller{
		{
			UserID:             seller1.ID,
			BusinessName:       "Tech Store Indonesia",
			WhatsappNumber:     "+6281234567890",
			WhatsappLink:       "https://wa.me/6281234567890",
			VerificationStatus: "verified",
			AccountStatus:      "active",
		},
		{
			UserID:             seller2.ID,
			BusinessName:       "Fashion Hub Jakarta",
			WhatsappNumber:     "+6281234567891",
			WhatsappLink:       "https://wa.me/6281234567891",
			VerificationStatus: "verified",
			AccountStatus:      "active",
		},
	}

	for _, seller := range sellers {
		var existing models.Seller
		if err := db.Where("user_id = ?", seller.UserID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&seller).Error; err != nil {
				log.Printf("❌ Failed to create seller %s: %v", seller.BusinessName, err)
			} else {
				log.Printf("✅ Created seller: %s", seller.BusinessName)
			}
		} else {
			log.Printf("⏭️  Seller already exists: %s", seller.BusinessName)
		}
	}
}

func seedItems(db *gorm.DB) {
	// Get categories and sellers
	var electronicsCategory, clothingCategory models.Category
	var techSeller, fashionSeller models.Seller

	db.Where("name = ?", "Electronics").First(&electronicsCategory)
	db.Where("name = ?", "Clothing").First(&clothingCategory)
	db.Where("business_name = ?", "Tech Store Indonesia").First(&techSeller)
	db.Where("business_name = ?", "Fashion Hub Jakarta").First(&fashionSeller)

	items := []models.Item{
		// Tech Store Items
		{
			SellerID:    techSeller.UserID,
			CategoryID:  electronicsCategory.ID,
			Name:        "MacBook Pro 14-inch M3",
			Description: "Latest MacBook Pro with M3 chip, 16GB unified memory, 512GB SSD storage. Perfect for developers and creative professionals.",
			Price:       32000000.00,
			Quantity:    5,
			Status:      "approved",
		},
		{
			SellerID:    techSeller.UserID,
			CategoryID:  electronicsCategory.ID,
			Name:        "iPhone 15 Pro Max",
			Description: "iPhone 15 Pro Max 256GB, Titanium Natural. Latest Apple smartphone with Advanced Pro camera system.",
			Price:       22000000.00,
			Quantity:    8,
			Status:      "approved",
		},
		{
			SellerID:    techSeller.UserID,
			CategoryID:  electronicsCategory.ID,
			Name:        "Gaming Laptop ASUS ROG Strix",
			Description: "ASUS ROG Strix G16, Intel Core i7-13650HX, GeForce RTX 4060, 16GB DDR5, 1TB SSD. Ultimate gaming performance.",
			Price:       18500000.00,
			Quantity:    3,
			Status:      "pending",
		},
		{
			SellerID:    techSeller.UserID,
			CategoryID:  electronicsCategory.ID,
			Name:        "iPad Air 5th Generation",
			Description: "iPad Air with M1 chip, 64GB Wi-Fi model. Perfect for productivity and creativity on the go.",
			Price:       8500000.00,
			Quantity:    10,
			Status:      "approved",
		},
		{
			SellerID:    techSeller.UserID,
			CategoryID:  electronicsCategory.ID,
			Name:        "AirPods Pro 2nd Generation",
			Description: "AirPods Pro with MagSafe Charging Case. Advanced Active Noise Cancellation and Spatial Audio.",
			Price:       3200000.00,
			Quantity:    15,
			Status:      "approved",
		},

		// Fashion Hub Items
		{
			SellerID:    fashionSeller.UserID,
			CategoryID:  clothingCategory.ID,
			Name:        "Premium Cotton T-Shirt",
			Description: "100% premium cotton t-shirt, available in multiple colors. Comfortable daily wear for men and women.",
			Price:       150000.00,
			Quantity:    50,
			Status:      "approved",
		},
		{
			SellerID:    fashionSeller.UserID,
			CategoryID:  clothingCategory.ID,
			Name:        "Denim Jacket Vintage Style",
			Description: "Classic vintage denim jacket, premium quality cotton denim. Perfect for casual and semi-formal occasions.",
			Price:       450000.00,
			Quantity:    20,
			Status:      "approved",
		},
		{
			SellerID:    fashionSeller.UserID,
			CategoryID:  clothingCategory.ID,
			Name:        "Sneakers Premium Sport",
			Description: "Premium sport sneakers with comfortable cushioning. Suitable for running, gym, and casual wear.",
			Price:       850000.00,
			Quantity:    25,
			Status:      "pending",
		},
	}

	for _, item := range items {
		var existing models.Item
		if err := db.Where("name = ? AND seller_id = ?", item.Name, item.SellerID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&item).Error; err != nil {
				log.Printf(" Failed to create item %s: %v", item.Name, err)
			} else {
				log.Printf("reated item: %s (Rp %.0f)", item.Name, item.Price)
			}
		} else {
			log.Printf("⏭️  Item already exists: %s", item.Name)
		}
	}
}
