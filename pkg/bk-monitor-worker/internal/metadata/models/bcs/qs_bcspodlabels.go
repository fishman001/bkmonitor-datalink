// Code generated by go-queryset. DO NOT EDIT.
package bcs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set BCSPodLabelsQuerySet

// BCSPodLabelsQuerySet is an queryset type for BCSPodLabels
type BCSPodLabelsQuerySet struct {
	db *gorm.DB
}

// NewBCSPodLabelsQuerySet constructs new BCSPodLabelsQuerySet
func NewBCSPodLabelsQuerySet(db *gorm.DB) BCSPodLabelsQuerySet {
	return BCSPodLabelsQuerySet{
		db: db.Model(&BCSPodLabels{}),
	}
}

func (qs BCSPodLabelsQuerySet) w(db *gorm.DB) BCSPodLabelsQuerySet {
	return NewBCSPodLabelsQuerySet(db)
}

func (qs BCSPodLabelsQuerySet) Select(fields ...BCSPodLabelsDBSchemaField) BCSPodLabelsQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *BCSPodLabels) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *BCSPodLabels) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) All(ret *[]BCSPodLabels) error {
	return qs.db.Find(ret).Error
}

// ClusterIDEq is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDEq(clusterID string) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("cluster_id = ?", clusterID))
}

// ClusterIDGt is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDGt(clusterID string) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("cluster_id > ?", clusterID))
}

// ClusterIDGte is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDGte(clusterID string) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("cluster_id >= ?", clusterID))
}

// ClusterIDIn is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDIn(clusterID ...string) BCSPodLabelsQuerySet {
	if len(clusterID) == 0 {
		qs.db.AddError(errors.New("must at least pass one clusterID in ClusterIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("cluster_id IN (?)", clusterID))
}

// ClusterIDLike is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDLike(clusterID string) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("cluster_id LIKE ?", clusterID))
}

// ClusterIDLt is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDLt(clusterID string) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("cluster_id < ?", clusterID))
}

// ClusterIDLte is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDLte(clusterID string) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("cluster_id <= ?", clusterID))
}

// ClusterIDNe is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDNe(clusterID string) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("cluster_id != ?", clusterID))
}

// ClusterIDNotIn is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDNotIn(clusterID ...string) BCSPodLabelsQuerySet {
	if len(clusterID) == 0 {
		qs.db.AddError(errors.New("must at least pass one clusterID in ClusterIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("cluster_id NOT IN (?)", clusterID))
}

// ClusterIDNotlike is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ClusterIDNotlike(clusterID string) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("cluster_id NOT LIKE ?", clusterID))
}

// Count is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Delete is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) Delete() error {
	return qs.db.Delete(BCSPodLabels{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(BCSPodLabels{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(BCSPodLabels{})
	return db.RowsAffected, db.Error
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) GetUpdater() BCSPodLabelsUpdater {
	return NewBCSPodLabelsUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) IDEq(ID uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) IDGt(ID uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) IDGte(ID uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) IDIn(ID ...uint) BCSPodLabelsQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) IDLt(ID uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) IDLte(ID uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) IDNe(ID uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) IDNotIn(ID ...uint) BCSPodLabelsQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// LabelEq is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) LabelEq(label uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("label = ?", label))
}

// LabelGt is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) LabelGt(label uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("label > ?", label))
}

// LabelGte is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) LabelGte(label uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("label >= ?", label))
}

// LabelIn is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) LabelIn(label ...uint) BCSPodLabelsQuerySet {
	if len(label) == 0 {
		qs.db.AddError(errors.New("must at least pass one label in LabelIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("label IN (?)", label))
}

// LabelLt is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) LabelLt(label uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("label < ?", label))
}

// LabelLte is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) LabelLte(label uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("label <= ?", label))
}

// LabelNe is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) LabelNe(label uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("label != ?", label))
}

// LabelNotIn is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) LabelNotIn(label ...uint) BCSPodLabelsQuerySet {
	if len(label) == 0 {
		qs.db.AddError(errors.New("must at least pass one label in LabelNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("label NOT IN (?)", label))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) Limit(limit int) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) Offset(offset int) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs BCSPodLabelsQuerySet) One(ret *BCSPodLabels) error {
	return qs.db.First(ret).Error
}

// OrderAscByClusterID is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) OrderAscByClusterID() BCSPodLabelsQuerySet {
	return qs.w(qs.db.Order("cluster_id ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) OrderAscByID() BCSPodLabelsQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByLabel is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) OrderAscByLabel() BCSPodLabelsQuerySet {
	return qs.w(qs.db.Order("label ASC"))
}

// OrderAscByResource is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) OrderAscByResource() BCSPodLabelsQuerySet {
	return qs.w(qs.db.Order("resource ASC"))
}

// OrderDescByClusterID is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) OrderDescByClusterID() BCSPodLabelsQuerySet {
	return qs.w(qs.db.Order("cluster_id DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) OrderDescByID() BCSPodLabelsQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByLabel is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) OrderDescByLabel() BCSPodLabelsQuerySet {
	return qs.w(qs.db.Order("label DESC"))
}

// OrderDescByResource is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) OrderDescByResource() BCSPodLabelsQuerySet {
	return qs.w(qs.db.Order("resource DESC"))
}

// ResourceEq is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ResourceEq(resource uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("resource = ?", resource))
}

// ResourceGt is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ResourceGt(resource uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("resource > ?", resource))
}

// ResourceGte is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ResourceGte(resource uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("resource >= ?", resource))
}

// ResourceIn is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ResourceIn(resource ...uint) BCSPodLabelsQuerySet {
	if len(resource) == 0 {
		qs.db.AddError(errors.New("must at least pass one resource in ResourceIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("resource IN (?)", resource))
}

// ResourceLt is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ResourceLt(resource uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("resource < ?", resource))
}

// ResourceLte is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ResourceLte(resource uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("resource <= ?", resource))
}

// ResourceNe is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ResourceNe(resource uint) BCSPodLabelsQuerySet {
	return qs.w(qs.db.Where("resource != ?", resource))
}

// ResourceNotIn is an autogenerated method
// nolint: dupl
func (qs BCSPodLabelsQuerySet) ResourceNotIn(resource ...uint) BCSPodLabelsQuerySet {
	if len(resource) == 0 {
		qs.db.AddError(errors.New("must at least pass one resource in ResourceNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("resource NOT IN (?)", resource))
}

// SetClusterID is an autogenerated method
// nolint: dupl
func (u BCSPodLabelsUpdater) SetClusterID(clusterID string) BCSPodLabelsUpdater {
	u.fields[string(BCSPodLabelsDBSchema.ClusterID)] = clusterID
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u BCSPodLabelsUpdater) SetID(ID uint) BCSPodLabelsUpdater {
	u.fields[string(BCSPodLabelsDBSchema.ID)] = ID
	return u
}

// SetLabel is an autogenerated method
// nolint: dupl
func (u BCSPodLabelsUpdater) SetLabel(label uint) BCSPodLabelsUpdater {
	u.fields[string(BCSPodLabelsDBSchema.Label)] = label
	return u
}

// SetResource is an autogenerated method
// nolint: dupl
func (u BCSPodLabelsUpdater) SetResource(resource uint) BCSPodLabelsUpdater {
	u.fields[string(BCSPodLabelsDBSchema.Resource)] = resource
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u BCSPodLabelsUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u BCSPodLabelsUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set BCSPodLabelsQuerySet

// ===== BEGIN of BCSPodLabels modifiers

// BCSPodLabelsDBSchemaField describes database schema field. It requires for method 'Update'
type BCSPodLabelsDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f BCSPodLabelsDBSchemaField) String() string {
	return string(f)
}

// BCSPodLabelsDBSchema stores db field names of BCSPodLabels
var BCSPodLabelsDBSchema = struct {
	ID        BCSPodLabelsDBSchemaField
	Resource  BCSPodLabelsDBSchemaField
	Label     BCSPodLabelsDBSchemaField
	ClusterID BCSPodLabelsDBSchemaField
}{

	ID:        BCSPodLabelsDBSchemaField("id"),
	Resource:  BCSPodLabelsDBSchemaField("resource"),
	Label:     BCSPodLabelsDBSchemaField("label"),
	ClusterID: BCSPodLabelsDBSchemaField("cluster_id"),
}

// Update updates BCSPodLabels fields by primary key
// nolint: dupl
func (o *BCSPodLabels) Update(db *gorm.DB, fields ...BCSPodLabelsDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"resource":   o.Resource,
		"label":      o.Label,
		"cluster_id": o.ClusterID,
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

		return fmt.Errorf("can't update BCSPodLabels %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// BCSPodLabelsUpdater is an BCSPodLabels updates manager
type BCSPodLabelsUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewBCSPodLabelsUpdater creates new BCSPodLabels updater
// nolint: dupl
func NewBCSPodLabelsUpdater(db *gorm.DB) BCSPodLabelsUpdater {
	return BCSPodLabelsUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&BCSPodLabels{}),
	}
}

// ===== END of BCSPodLabels modifiers

// ===== END of all query sets
