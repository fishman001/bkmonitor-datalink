// Code generated by go-queryset. DO NOT EDIT.
package resulttable

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set LabelQuerySet

// LabelQuerySet is an queryset type for Label
type LabelQuerySet struct {
	db *gorm.DB
}

// NewLabelQuerySet constructs new LabelQuerySet
func NewLabelQuerySet(db *gorm.DB) LabelQuerySet {
	return LabelQuerySet{
		db: db.Model(&Label{}),
	}
}

func (qs LabelQuerySet) w(db *gorm.DB) LabelQuerySet {
	return NewLabelQuerySet(db)
}

func (qs LabelQuerySet) Select(fields ...LabelDBSchemaField) LabelQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *Label) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Label) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) All(ret *[]Label) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Delete is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) Delete() error {
	return qs.db.Delete(Label{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Label{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Label{})
	return db.RowsAffected, db.Error
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) GetUpdater() LabelUpdater {
	return NewLabelUpdater(qs.db)
}

// IndexEq is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexEq(index uint) LabelQuerySet {
	return qs.w(qs.db.Where("index = ?", index))
}

// IndexGt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexGt(index uint) LabelQuerySet {
	return qs.w(qs.db.Where("index > ?", index))
}

// IndexGte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexGte(index uint) LabelQuerySet {
	return qs.w(qs.db.Where("index >= ?", index))
}

// IndexIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexIn(index ...uint) LabelQuerySet {
	if len(index) == 0 {
		qs.db.AddError(errors.New("must at least pass one index in IndexIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("index IN (?)", index))
}

// IndexIsNotNull is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexIsNotNull() LabelQuerySet {
	return qs.w(qs.db.Where("index IS NOT NULL"))
}

// IndexIsNull is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexIsNull() LabelQuerySet {
	return qs.w(qs.db.Where("index IS NULL"))
}

// IndexLt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexLt(index uint) LabelQuerySet {
	return qs.w(qs.db.Where("index < ?", index))
}

// IndexLte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexLte(index uint) LabelQuerySet {
	return qs.w(qs.db.Where("index <= ?", index))
}

// IndexNe is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexNe(index uint) LabelQuerySet {
	return qs.w(qs.db.Where("index != ?", index))
}

// IndexNotIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IndexNotIn(index ...uint) LabelQuerySet {
	if len(index) == 0 {
		qs.db.AddError(errors.New("must at least pass one index in IndexNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("index NOT IN (?)", index))
}

// IsAdminOnlyEq is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IsAdminOnlyEq(isAdminOnly bool) LabelQuerySet {
	return qs.w(qs.db.Where("is_admin_only = ?", isAdminOnly))
}

// IsAdminOnlyIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IsAdminOnlyIn(isAdminOnly ...bool) LabelQuerySet {
	if len(isAdminOnly) == 0 {
		qs.db.AddError(errors.New("must at least pass one isAdminOnly in IsAdminOnlyIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("is_admin_only IN (?)", isAdminOnly))
}

// IsAdminOnlyNe is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IsAdminOnlyNe(isAdminOnly bool) LabelQuerySet {
	return qs.w(qs.db.Where("is_admin_only != ?", isAdminOnly))
}

// IsAdminOnlyNotIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) IsAdminOnlyNotIn(isAdminOnly ...bool) LabelQuerySet {
	if len(isAdminOnly) == 0 {
		qs.db.AddError(errors.New("must at least pass one isAdminOnly in IsAdminOnlyNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("is_admin_only NOT IN (?)", isAdminOnly))
}

// LabelIdEq is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdEq(labelId string) LabelQuerySet {
	return qs.w(qs.db.Where("label_id = ?", labelId))
}

// LabelIdGt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdGt(labelId string) LabelQuerySet {
	return qs.w(qs.db.Where("label_id > ?", labelId))
}

// LabelIdGte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdGte(labelId string) LabelQuerySet {
	return qs.w(qs.db.Where("label_id >= ?", labelId))
}

// LabelIdIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdIn(labelId ...string) LabelQuerySet {
	if len(labelId) == 0 {
		qs.db.AddError(errors.New("must at least pass one labelId in LabelIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("label_id IN (?)", labelId))
}

// LabelIdLike is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdLike(labelId string) LabelQuerySet {
	return qs.w(qs.db.Where("label_id LIKE ?", labelId))
}

// LabelIdLt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdLt(labelId string) LabelQuerySet {
	return qs.w(qs.db.Where("label_id < ?", labelId))
}

// LabelIdLte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdLte(labelId string) LabelQuerySet {
	return qs.w(qs.db.Where("label_id <= ?", labelId))
}

// LabelIdNe is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdNe(labelId string) LabelQuerySet {
	return qs.w(qs.db.Where("label_id != ?", labelId))
}

// LabelIdNotIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdNotIn(labelId ...string) LabelQuerySet {
	if len(labelId) == 0 {
		qs.db.AddError(errors.New("must at least pass one labelId in LabelIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("label_id NOT IN (?)", labelId))
}

// LabelIdNotlike is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelIdNotlike(labelId string) LabelQuerySet {
	return qs.w(qs.db.Where("label_id NOT LIKE ?", labelId))
}

// LabelNameEq is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameEq(labelName string) LabelQuerySet {
	return qs.w(qs.db.Where("label_name = ?", labelName))
}

// LabelNameGt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameGt(labelName string) LabelQuerySet {
	return qs.w(qs.db.Where("label_name > ?", labelName))
}

// LabelNameGte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameGte(labelName string) LabelQuerySet {
	return qs.w(qs.db.Where("label_name >= ?", labelName))
}

// LabelNameIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameIn(labelName ...string) LabelQuerySet {
	if len(labelName) == 0 {
		qs.db.AddError(errors.New("must at least pass one labelName in LabelNameIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("label_name IN (?)", labelName))
}

// LabelNameLike is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameLike(labelName string) LabelQuerySet {
	return qs.w(qs.db.Where("label_name LIKE ?", labelName))
}

// LabelNameLt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameLt(labelName string) LabelQuerySet {
	return qs.w(qs.db.Where("label_name < ?", labelName))
}

// LabelNameLte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameLte(labelName string) LabelQuerySet {
	return qs.w(qs.db.Where("label_name <= ?", labelName))
}

// LabelNameNe is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameNe(labelName string) LabelQuerySet {
	return qs.w(qs.db.Where("label_name != ?", labelName))
}

// LabelNameNotIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameNotIn(labelName ...string) LabelQuerySet {
	if len(labelName) == 0 {
		qs.db.AddError(errors.New("must at least pass one labelName in LabelNameNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("label_name NOT IN (?)", labelName))
}

// LabelNameNotlike is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelNameNotlike(labelName string) LabelQuerySet {
	return qs.w(qs.db.Where("label_name NOT LIKE ?", labelName))
}

// LabelTypeEq is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeEq(labelType string) LabelQuerySet {
	return qs.w(qs.db.Where("label_type = ?", labelType))
}

// LabelTypeGt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeGt(labelType string) LabelQuerySet {
	return qs.w(qs.db.Where("label_type > ?", labelType))
}

// LabelTypeGte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeGte(labelType string) LabelQuerySet {
	return qs.w(qs.db.Where("label_type >= ?", labelType))
}

// LabelTypeIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeIn(labelType ...string) LabelQuerySet {
	if len(labelType) == 0 {
		qs.db.AddError(errors.New("must at least pass one labelType in LabelTypeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("label_type IN (?)", labelType))
}

// LabelTypeLike is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeLike(labelType string) LabelQuerySet {
	return qs.w(qs.db.Where("label_type LIKE ?", labelType))
}

// LabelTypeLt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeLt(labelType string) LabelQuerySet {
	return qs.w(qs.db.Where("label_type < ?", labelType))
}

// LabelTypeLte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeLte(labelType string) LabelQuerySet {
	return qs.w(qs.db.Where("label_type <= ?", labelType))
}

// LabelTypeNe is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeNe(labelType string) LabelQuerySet {
	return qs.w(qs.db.Where("label_type != ?", labelType))
}

// LabelTypeNotIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeNotIn(labelType ...string) LabelQuerySet {
	if len(labelType) == 0 {
		qs.db.AddError(errors.New("must at least pass one labelType in LabelTypeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("label_type NOT IN (?)", labelType))
}

// LabelTypeNotlike is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LabelTypeNotlike(labelType string) LabelQuerySet {
	return qs.w(qs.db.Where("label_type NOT LIKE ?", labelType))
}

// LevelEq is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelEq(level uint) LabelQuerySet {
	return qs.w(qs.db.Where("level = ?", level))
}

// LevelGt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelGt(level uint) LabelQuerySet {
	return qs.w(qs.db.Where("level > ?", level))
}

// LevelGte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelGte(level uint) LabelQuerySet {
	return qs.w(qs.db.Where("level >= ?", level))
}

// LevelIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelIn(level ...uint) LabelQuerySet {
	if len(level) == 0 {
		qs.db.AddError(errors.New("must at least pass one level in LevelIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("level IN (?)", level))
}

// LevelIsNotNull is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelIsNotNull() LabelQuerySet {
	return qs.w(qs.db.Where("level IS NOT NULL"))
}

// LevelIsNull is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelIsNull() LabelQuerySet {
	return qs.w(qs.db.Where("level IS NULL"))
}

// LevelLt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelLt(level uint) LabelQuerySet {
	return qs.w(qs.db.Where("level < ?", level))
}

// LevelLte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelLte(level uint) LabelQuerySet {
	return qs.w(qs.db.Where("level <= ?", level))
}

// LevelNe is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelNe(level uint) LabelQuerySet {
	return qs.w(qs.db.Where("level != ?", level))
}

// LevelNotIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) LevelNotIn(level ...uint) LabelQuerySet {
	if len(level) == 0 {
		qs.db.AddError(errors.New("must at least pass one level in LevelNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("level NOT IN (?)", level))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) Limit(limit int) LabelQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) Offset(offset int) LabelQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs LabelQuerySet) One(ret *Label) error {
	return qs.db.First(ret).Error
}

// OrderAscByIndex is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderAscByIndex() LabelQuerySet {
	return qs.w(qs.db.Order("index ASC"))
}

// OrderAscByIsAdminOnly is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderAscByIsAdminOnly() LabelQuerySet {
	return qs.w(qs.db.Order("is_admin_only ASC"))
}

// OrderAscByLabelId is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderAscByLabelId() LabelQuerySet {
	return qs.w(qs.db.Order("label_id ASC"))
}

// OrderAscByLabelName is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderAscByLabelName() LabelQuerySet {
	return qs.w(qs.db.Order("label_name ASC"))
}

// OrderAscByLabelType is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderAscByLabelType() LabelQuerySet {
	return qs.w(qs.db.Order("label_type ASC"))
}

// OrderAscByLevel is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderAscByLevel() LabelQuerySet {
	return qs.w(qs.db.Order("level ASC"))
}

// OrderAscByParentLabel is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderAscByParentLabel() LabelQuerySet {
	return qs.w(qs.db.Order("parent_label ASC"))
}

// OrderDescByIndex is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderDescByIndex() LabelQuerySet {
	return qs.w(qs.db.Order("index DESC"))
}

// OrderDescByIsAdminOnly is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderDescByIsAdminOnly() LabelQuerySet {
	return qs.w(qs.db.Order("is_admin_only DESC"))
}

// OrderDescByLabelId is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderDescByLabelId() LabelQuerySet {
	return qs.w(qs.db.Order("label_id DESC"))
}

// OrderDescByLabelName is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderDescByLabelName() LabelQuerySet {
	return qs.w(qs.db.Order("label_name DESC"))
}

// OrderDescByLabelType is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderDescByLabelType() LabelQuerySet {
	return qs.w(qs.db.Order("label_type DESC"))
}

// OrderDescByLevel is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderDescByLevel() LabelQuerySet {
	return qs.w(qs.db.Order("level DESC"))
}

// OrderDescByParentLabel is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) OrderDescByParentLabel() LabelQuerySet {
	return qs.w(qs.db.Order("parent_label DESC"))
}

// ParentLabelEq is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelEq(parentLabel string) LabelQuerySet {
	return qs.w(qs.db.Where("parent_label = ?", parentLabel))
}

// ParentLabelGt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelGt(parentLabel string) LabelQuerySet {
	return qs.w(qs.db.Where("parent_label > ?", parentLabel))
}

// ParentLabelGte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelGte(parentLabel string) LabelQuerySet {
	return qs.w(qs.db.Where("parent_label >= ?", parentLabel))
}

// ParentLabelIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelIn(parentLabel ...string) LabelQuerySet {
	if len(parentLabel) == 0 {
		qs.db.AddError(errors.New("must at least pass one parentLabel in ParentLabelIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("parent_label IN (?)", parentLabel))
}

// ParentLabelIsNotNull is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelIsNotNull() LabelQuerySet {
	return qs.w(qs.db.Where("parent_label IS NOT NULL"))
}

// ParentLabelIsNull is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelIsNull() LabelQuerySet {
	return qs.w(qs.db.Where("parent_label IS NULL"))
}

// ParentLabelLike is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelLike(parentLabel string) LabelQuerySet {
	return qs.w(qs.db.Where("parent_label LIKE ?", parentLabel))
}

// ParentLabelLt is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelLt(parentLabel string) LabelQuerySet {
	return qs.w(qs.db.Where("parent_label < ?", parentLabel))
}

// ParentLabelLte is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelLte(parentLabel string) LabelQuerySet {
	return qs.w(qs.db.Where("parent_label <= ?", parentLabel))
}

// ParentLabelNe is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelNe(parentLabel string) LabelQuerySet {
	return qs.w(qs.db.Where("parent_label != ?", parentLabel))
}

// ParentLabelNotIn is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelNotIn(parentLabel ...string) LabelQuerySet {
	if len(parentLabel) == 0 {
		qs.db.AddError(errors.New("must at least pass one parentLabel in ParentLabelNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("parent_label NOT IN (?)", parentLabel))
}

// ParentLabelNotlike is an autogenerated method
// nolint: dupl
func (qs LabelQuerySet) ParentLabelNotlike(parentLabel string) LabelQuerySet {
	return qs.w(qs.db.Where("parent_label NOT LIKE ?", parentLabel))
}

// SetIndex is an autogenerated method
// nolint: dupl
func (u LabelUpdater) SetIndex(index *uint) LabelUpdater {
	u.fields[string(LabelDBSchema.Index)] = index
	return u
}

// SetIsAdminOnly is an autogenerated method
// nolint: dupl
func (u LabelUpdater) SetIsAdminOnly(isAdminOnly bool) LabelUpdater {
	u.fields[string(LabelDBSchema.IsAdminOnly)] = isAdminOnly
	return u
}

// SetLabelId is an autogenerated method
// nolint: dupl
func (u LabelUpdater) SetLabelId(labelId string) LabelUpdater {
	u.fields[string(LabelDBSchema.LabelId)] = labelId
	return u
}

// SetLabelName is an autogenerated method
// nolint: dupl
func (u LabelUpdater) SetLabelName(labelName string) LabelUpdater {
	u.fields[string(LabelDBSchema.LabelName)] = labelName
	return u
}

// SetLabelType is an autogenerated method
// nolint: dupl
func (u LabelUpdater) SetLabelType(labelType string) LabelUpdater {
	u.fields[string(LabelDBSchema.LabelType)] = labelType
	return u
}

// SetLevel is an autogenerated method
// nolint: dupl
func (u LabelUpdater) SetLevel(level *uint) LabelUpdater {
	u.fields[string(LabelDBSchema.Level)] = level
	return u
}

// SetParentLabel is an autogenerated method
// nolint: dupl
func (u LabelUpdater) SetParentLabel(parentLabel *string) LabelUpdater {
	u.fields[string(LabelDBSchema.ParentLabel)] = parentLabel
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u LabelUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u LabelUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set LabelQuerySet

// ===== BEGIN of Label modifiers

// LabelDBSchemaField describes database schema field. It requires for method 'Update'
type LabelDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f LabelDBSchemaField) String() string {
	return string(f)
}

// LabelDBSchema stores db field names of Label
var LabelDBSchema = struct {
	LabelId     LabelDBSchemaField
	LabelName   LabelDBSchemaField
	LabelType   LabelDBSchemaField
	IsAdminOnly LabelDBSchemaField
	ParentLabel LabelDBSchemaField
	Level       LabelDBSchemaField
	Index       LabelDBSchemaField
}{

	LabelId:     LabelDBSchemaField("label_id"),
	LabelName:   LabelDBSchemaField("label_name"),
	LabelType:   LabelDBSchemaField("label_type"),
	IsAdminOnly: LabelDBSchemaField("is_admin_only"),
	ParentLabel: LabelDBSchemaField("parent_label"),
	Level:       LabelDBSchemaField("level"),
	Index:       LabelDBSchemaField("index"),
}

// Update updates Label fields by primary key
// nolint: dupl
func (o *Label) Update(db *gorm.DB, fields ...LabelDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"label_id":      o.LabelId,
		"label_name":    o.LabelName,
		"label_type":    o.LabelType,
		"is_admin_only": o.IsAdminOnly,
		"parent_label":  o.ParentLabel,
		"level":         o.Level,
		"index":         o.Index,
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

		return fmt.Errorf("can't update Label %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// LabelUpdater is an Label updates manager
type LabelUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewLabelUpdater creates new Label updater
// nolint: dupl
func NewLabelUpdater(db *gorm.DB) LabelUpdater {
	return LabelUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Label{}),
	}
}

// ===== END of Label modifiers

// ===== END of all query sets
