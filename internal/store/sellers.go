package store

type Seller struct {
	UserID             uint   `gorm:"primaryKey"`
	BusinessName       string `gorm:"not null;size:255"`
	WhatsappNumber     string `gorm:"not null;size:25"`
	WhatsappLink       string `gorm:"size:255"`
	ProfilePictureURL  string `gorm:"size:255"`
	VerificationStatus string `gorm:"size:50"`
	AccountStatus      string `gorm:"type:varchar(25);not null;default:pending_verification"`

	User  *User  `gorm:"foreignKey:UserID"`
	Items []Item `gorm:"foreignKey:SellerID"`
}