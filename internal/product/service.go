package product

type Service interface {
	getAll() ([]ProductResponse, error)
	getByID(id int) (*Product, error)
	create(product CreateProductRequest) error
	update(id uint, updateProductRequest UpdateProductRequest) error
	delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) getAll() ([]ProductResponse, error) {
	products, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ProductResponse{
			Id:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productResponses, nil
}

func (s *service) getByID(id int) (*Product, error) {
	product, err := s.repository.getByID(id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *service) create(product CreateProductRequest) error {
	productModel := Product{
		Name:  product.Name,
		Price: product.Price,
	}
	_, err := s.repository.create(productModel)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) update(id uint, product UpdateProductRequest) error {
	prod, err := s.getByID(int(id))
	if err != nil {
		return err
	}
	
	return s.repository.update(prod, product)
}

func (s *service) delete(id int) error {
	err := s.repository.delete(id)
	if err != nil {
		return err
	}
	return nil
}