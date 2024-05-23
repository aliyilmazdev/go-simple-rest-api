package product

import "gorm.io/gorm"

type Repository interface {
	getAll() ([]Product, error)
	getByID(id int) (Product, error)
	create(product Product) (Product, error)
	update(product Product) (Product, error)
	delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) getAll() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) getByID(id int) (Product, error) {
	var product Product
	if err := r.db.First(&product, id).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repository) create(product Product) (Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repository) update(product Product) (Product, error) {

	if err := r.db.Save(&product).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repository) delete(id int) error {
	if err := r.db.Delete(&Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

