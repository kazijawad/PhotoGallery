package models

import "github.com/jinzhu/gorm"

// Services represents all our indivudal services and
// our DB. It is mostly a container resource.
type Services struct {
	db      *gorm.DB
	User    UserService
	Gallery GalleryService
	Image   ImageService
}

// NewServices is used to create a new services container
// and open a connection to our DB.
func NewServices(dialect, connectionInfo string) (*Services, error) {
	db, err := gorm.Open(dialect, connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &Services{
		db:      db,
		User:    NewUserService(db),
		Gallery: NewGalleryService(db),
		Image:   NewImageService(),
	}, nil
}

// Close closes the database connection
func (s *Services) Close() error {
	return s.db.Close()
}

// AutoMigrate will attempt to automatically migrate all tables
func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}, &Gallery{}).Error
}

// DestructiveReset drops all tables and rebuilds them
func (s *Services) DestructiveReset() error {
	err := s.db.DropTableIfExists(&User{}, &Gallery{}).Error
	if err != nil {
		return err
	}
	return s.AutoMigrate()
}
