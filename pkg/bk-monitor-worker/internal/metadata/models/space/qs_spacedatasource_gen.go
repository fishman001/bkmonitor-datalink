// Code generated by go-queryset. DO NOT EDIT.
package space

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set SpaceDataSourceQuerySet

// SpaceDataSourceQuerySet is an queryset type for SpaceDataSource
type SpaceDataSourceQuerySet struct {
	db *gorm.DB
}

// NewSpaceDataSourceQuerySet constructs new SpaceDataSourceQuerySet
func NewSpaceDataSourceQuerySet(db *gorm.DB) SpaceDataSourceQuerySet {
	return SpaceDataSourceQuerySet{
		db: db.Model(&SpaceDataSource{}),
	}
}

func (qs SpaceDataSourceQuerySet) w(db *gorm.DB) SpaceDataSourceQuerySet {
	return NewSpaceDataSourceQuerySet(db)
}

func (qs SpaceDataSourceQuerySet) Select(fields ...SpaceDataSourceDBSchemaField) SpaceDataSourceQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *SpaceDataSource) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *SpaceDataSource) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) All(ret *[]SpaceDataSource) error {
	return qs.db.Find(ret).Error
}

// BkDataIdEq is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) BkDataIdEq(bkDataId uint) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("bk_data_id = ?", bkDataId))
}

// BkDataIdGt is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) BkDataIdGt(bkDataId uint) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("bk_data_id > ?", bkDataId))
}

// BkDataIdGte is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) BkDataIdGte(bkDataId uint) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("bk_data_id >= ?", bkDataId))
}

// BkDataIdIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) BkDataIdIn(bkDataId ...uint) SpaceDataSourceQuerySet {
	if len(bkDataId) == 0 {
		qs.db.AddError(errors.New("must at least pass one bkDataId in BkDataIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bk_data_id IN (?)", bkDataId))
}

// BkDataIdLt is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) BkDataIdLt(bkDataId uint) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("bk_data_id < ?", bkDataId))
}

// BkDataIdLte is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) BkDataIdLte(bkDataId uint) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("bk_data_id <= ?", bkDataId))
}

// BkDataIdNe is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) BkDataIdNe(bkDataId uint) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("bk_data_id != ?", bkDataId))
}

// BkDataIdNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) BkDataIdNotIn(bkDataId ...uint) SpaceDataSourceQuerySet {
	if len(bkDataId) == 0 {
		qs.db.AddError(errors.New("must at least pass one bkDataId in BkDataIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bk_data_id NOT IN (?)", bkDataId))
}

// Count is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Delete is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) Delete() error {
	return qs.db.Delete(SpaceDataSource{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(SpaceDataSource{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(SpaceDataSource{})
	return db.RowsAffected, db.Error
}

// FromAuthorizationEq is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) FromAuthorizationEq(fromAuthorization bool) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("from_authorization = ?", fromAuthorization))
}

// FromAuthorizationIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) FromAuthorizationIn(fromAuthorization ...bool) SpaceDataSourceQuerySet {
	if len(fromAuthorization) == 0 {
		qs.db.AddError(errors.New("must at least pass one fromAuthorization in FromAuthorizationIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("from_authorization IN (?)", fromAuthorization))
}

// FromAuthorizationNe is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) FromAuthorizationNe(fromAuthorization bool) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("from_authorization != ?", fromAuthorization))
}

// FromAuthorizationNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) FromAuthorizationNotIn(fromAuthorization ...bool) SpaceDataSourceQuerySet {
	if len(fromAuthorization) == 0 {
		qs.db.AddError(errors.New("must at least pass one fromAuthorization in FromAuthorizationNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("from_authorization NOT IN (?)", fromAuthorization))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) GetUpdater() SpaceDataSourceUpdater {
	return NewSpaceDataSourceUpdater(qs.db)
}

// IdEq is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) IdEq(id int) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("id = ?", id))
}

// IdGt is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) IdGt(id int) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("id > ?", id))
}

// IdGte is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) IdGte(id int) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("id >= ?", id))
}

// IdIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) IdIn(id ...int) SpaceDataSourceQuerySet {
	if len(id) == 0 {
		qs.db.AddError(errors.New("must at least pass one id in IdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", id))
}

// IdLt is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) IdLt(id int) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("id < ?", id))
}

// IdLte is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) IdLte(id int) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("id <= ?", id))
}

// IdNe is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) IdNe(id int) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("id != ?", id))
}

// IdNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) IdNotIn(id ...int) SpaceDataSourceQuerySet {
	if len(id) == 0 {
		qs.db.AddError(errors.New("must at least pass one id in IdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", id))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) Limit(limit int) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) Offset(offset int) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs SpaceDataSourceQuerySet) One(ret *SpaceDataSource) error {
	return qs.db.First(ret).Error
}

// OrderAscByBkDataId is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderAscByBkDataId() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("bk_data_id ASC"))
}

// OrderAscByFromAuthorization is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderAscByFromAuthorization() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("from_authorization ASC"))
}

// OrderAscById is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderAscById() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscBySpaceId is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderAscBySpaceId() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("space_id ASC"))
}

// OrderAscBySpaceTypeId is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderAscBySpaceTypeId() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("space_type_id ASC"))
}

// OrderDescByBkDataId is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderDescByBkDataId() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("bk_data_id DESC"))
}

// OrderDescByFromAuthorization is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderDescByFromAuthorization() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("from_authorization DESC"))
}

// OrderDescById is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderDescById() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescBySpaceId is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderDescBySpaceId() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("space_id DESC"))
}

// OrderDescBySpaceTypeId is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) OrderDescBySpaceTypeId() SpaceDataSourceQuerySet {
	return qs.w(qs.db.Order("space_type_id DESC"))
}

// SpaceIdEq is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdEq(spaceId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_id = ?", spaceId))
}

// SpaceIdGt is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdGt(spaceId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_id > ?", spaceId))
}

// SpaceIdGte is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdGte(spaceId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_id >= ?", spaceId))
}

// SpaceIdIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdIn(spaceId ...string) SpaceDataSourceQuerySet {
	if len(spaceId) == 0 {
		qs.db.AddError(errors.New("must at least pass one spaceId in SpaceIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("space_id IN (?)", spaceId))
}

// SpaceIdLike is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdLike(spaceId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_id LIKE ?", spaceId))
}

// SpaceIdLt is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdLt(spaceId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_id < ?", spaceId))
}

// SpaceIdLte is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdLte(spaceId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_id <= ?", spaceId))
}

// SpaceIdNe is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdNe(spaceId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_id != ?", spaceId))
}

// SpaceIdNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdNotIn(spaceId ...string) SpaceDataSourceQuerySet {
	if len(spaceId) == 0 {
		qs.db.AddError(errors.New("must at least pass one spaceId in SpaceIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("space_id NOT IN (?)", spaceId))
}

// SpaceIdNotlike is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceIdNotlike(spaceId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_id NOT LIKE ?", spaceId))
}

// SpaceTypeIdEq is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdEq(spaceTypeId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_type_id = ?", spaceTypeId))
}

// SpaceTypeIdGt is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdGt(spaceTypeId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_type_id > ?", spaceTypeId))
}

// SpaceTypeIdGte is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdGte(spaceTypeId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_type_id >= ?", spaceTypeId))
}

// SpaceTypeIdIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdIn(spaceTypeId ...string) SpaceDataSourceQuerySet {
	if len(spaceTypeId) == 0 {
		qs.db.AddError(errors.New("must at least pass one spaceTypeId in SpaceTypeIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("space_type_id IN (?)", spaceTypeId))
}

// SpaceTypeIdLike is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdLike(spaceTypeId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_type_id LIKE ?", spaceTypeId))
}

// SpaceTypeIdLt is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdLt(spaceTypeId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_type_id < ?", spaceTypeId))
}

// SpaceTypeIdLte is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdLte(spaceTypeId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_type_id <= ?", spaceTypeId))
}

// SpaceTypeIdNe is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdNe(spaceTypeId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_type_id != ?", spaceTypeId))
}

// SpaceTypeIdNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdNotIn(spaceTypeId ...string) SpaceDataSourceQuerySet {
	if len(spaceTypeId) == 0 {
		qs.db.AddError(errors.New("must at least pass one spaceTypeId in SpaceTypeIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("space_type_id NOT IN (?)", spaceTypeId))
}

// SpaceTypeIdNotlike is an autogenerated method
// nolint: dupl
func (qs SpaceDataSourceQuerySet) SpaceTypeIdNotlike(spaceTypeId string) SpaceDataSourceQuerySet {
	return qs.w(qs.db.Where("space_type_id NOT LIKE ?", spaceTypeId))
}

// SetBkDataId is an autogenerated method
// nolint: dupl
func (u SpaceDataSourceUpdater) SetBkDataId(bkDataId uint) SpaceDataSourceUpdater {
	u.fields[string(SpaceDataSourceDBSchema.BkDataId)] = bkDataId
	return u
}

// SetFromAuthorization is an autogenerated method
// nolint: dupl
func (u SpaceDataSourceUpdater) SetFromAuthorization(fromAuthorization bool) SpaceDataSourceUpdater {
	u.fields[string(SpaceDataSourceDBSchema.FromAuthorization)] = fromAuthorization
	return u
}

// SetId is an autogenerated method
// nolint: dupl
func (u SpaceDataSourceUpdater) SetId(id int) SpaceDataSourceUpdater {
	u.fields[string(SpaceDataSourceDBSchema.Id)] = id
	return u
}

// SetSpaceId is an autogenerated method
// nolint: dupl
func (u SpaceDataSourceUpdater) SetSpaceId(spaceId string) SpaceDataSourceUpdater {
	u.fields[string(SpaceDataSourceDBSchema.SpaceId)] = spaceId
	return u
}

// SetSpaceTypeId is an autogenerated method
// nolint: dupl
func (u SpaceDataSourceUpdater) SetSpaceTypeId(spaceTypeId string) SpaceDataSourceUpdater {
	u.fields[string(SpaceDataSourceDBSchema.SpaceTypeId)] = spaceTypeId
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u SpaceDataSourceUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u SpaceDataSourceUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set SpaceDataSourceQuerySet

// ===== BEGIN of SpaceDataSource modifiers

// SpaceDataSourceDBSchemaField describes database schema field. It requires for method 'Update'
type SpaceDataSourceDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f SpaceDataSourceDBSchemaField) String() string {
	return string(f)
}

// SpaceDataSourceDBSchema stores db field names of SpaceDataSource
var SpaceDataSourceDBSchema = struct {
	Id                SpaceDataSourceDBSchemaField
	SpaceTypeId       SpaceDataSourceDBSchemaField
	SpaceId           SpaceDataSourceDBSchemaField
	BkDataId          SpaceDataSourceDBSchemaField
	FromAuthorization SpaceDataSourceDBSchemaField
}{

	Id:                SpaceDataSourceDBSchemaField("id"),
	SpaceTypeId:       SpaceDataSourceDBSchemaField("space_type_id"),
	SpaceId:           SpaceDataSourceDBSchemaField("space_id"),
	BkDataId:          SpaceDataSourceDBSchemaField("bk_data_id"),
	FromAuthorization: SpaceDataSourceDBSchemaField("from_authorization"),
}

// Update updates SpaceDataSource fields by primary key
// nolint: dupl
func (o *SpaceDataSource) Update(db *gorm.DB, fields ...SpaceDataSourceDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":                 o.Id,
		"space_type_id":      o.SpaceTypeId,
		"space_id":           o.SpaceId,
		"bk_data_id":         o.BkDataId,
		"from_authorization": o.FromAuthorization,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update SpaceDataSource %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// SpaceDataSourceUpdater is an SpaceDataSource updates manager
type SpaceDataSourceUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewSpaceDataSourceUpdater creates new SpaceDataSource updater
// nolint: dupl
func NewSpaceDataSourceUpdater(db *gorm.DB) SpaceDataSourceUpdater {
	return SpaceDataSourceUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&SpaceDataSource{}),
	}
}

// ===== END of SpaceDataSource modifiers

// ===== END of all query sets
