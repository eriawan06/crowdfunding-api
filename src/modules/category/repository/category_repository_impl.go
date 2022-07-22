package repository

import (
	"crowdfunding-api/src/modules/category/model/dto"
	"crowdfunding-api/src/modules/category/model/entity"
	e "crowdfunding-api/src/utils/errors"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository *CategoryRepositoryImpl) Create(category entity.Category) error {
	result := repository.DB.Create(&category)

	var mySqlErr *mysql.MySQLError
	if errors.As(result.Error, &mySqlErr) && mySqlErr.Number == 1062 {
		result.Error = e.ErrDuplicateKey
	}
	return result.Error
}

func (repository *CategoryRepositoryImpl) Update(category entity.Category, categoryId uint) error {
	result := repository.DB.
		Model(&entity.Category{}).
		Where("id=?", categoryId).
		Updates(map[string]interface{}{"name": category.Name, "is_active": category.IsActive})
	return result.Error
}

func (repository *CategoryRepositoryImpl) Delete(categoryId uint) error {
	result := repository.DB.Delete(&entity.Category{}, categoryId)
	return result.Error
}

func (repository *CategoryRepositoryImpl) FindAll(filter dto.FilterCategory) ([]entity.Category, error) {
	var categories []entity.Category

	where, whereVals := BuildFilter(filter)

	var buildWhereQuery string
	if where != nil {
		buildWhereQuery = strings.Join(where, " AND ")
	}

	result := repository.DB.Where(buildWhereQuery, whereVals...).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func BuildFilter(filter dto.FilterCategory) (where []string, whereVal []interface{}) {

	if filter.Name != "" {
		where = append(where, "name LIKE ?")
		whereVal = append(whereVal, "%"+filter.Name+"%")
	}

	if filter.IsActive != "" {
		where = append(where, "is_active = ?")

		boolValue, _ := strconv.ParseBool(filter.IsActive)
		whereVal = append(whereVal, &boolValue)
	}

	return
}

func (repository *CategoryRepositoryImpl) FindOne(categoryId uint) (entity.Category, error) {
	var category entity.Category

	result := repository.DB.Where("id = ?", categoryId).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result.Error = e.ErrDataNotFound
		}
		return category, result.Error
	}

	return category, nil
}
