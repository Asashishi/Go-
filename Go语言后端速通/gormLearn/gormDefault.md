**默认情况下, gorm使用ID作为主键, 结构体的命名将被以蛇形(下划线命名法)复数形式作为表名, 字段名将以蛇形形式作为列名, 并使用CreateAt, UpdateAt字段追踪创建与删除时间
gorm有定义一个gorm.Model结构体, 包括:**
ID        uint `gorm:"primarykey"`
CreatedAt time.Time
UpdatedAt time.Time
DeletedAt DeletedAt `gorm:"index"`
**如需单独设置表名则使用:**
func (u User) TableName() {
   return "user"
}