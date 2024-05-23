package product

type Service interface {
	getAll() ([]GetAllProductResponse, error)
	getByID(id int) (GetAllProductResponse, error)
	create(product CreateProductRequest) error
	update(product UpdateProductRequest) error
	delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) getAll() ([]GetAllProductResponse, error) {
	products, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	var productResponses []GetAllProductResponse
	for _, product := range products {
		productResponses = append(productResponses, GetAllProductResponse{
			Id:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productResponses, nil
}

func (s *service) getByID(id int) (GetAllProductResponse, error) {
	product, err := s.repository.getByID(id)
	if err != nil {
		return GetAllProductResponse{}, err
	}
	return GetAllProductResponse{
		Id:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
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

func (s *service) update(product UpdateProductRequest) error {
	productModel := Product{
		Name:  product.Name,
		Price: product.Price,
	}
	_, err := s.repository.update(productModel)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) delete(id int) error {
	err := s.repository.delete(id)
	if err != nil {
		return err
	}
	return nil
}