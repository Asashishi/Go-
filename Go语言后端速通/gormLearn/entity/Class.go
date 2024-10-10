package entity

type Class struct {
	BaseModel
	Name string `gorm:"type:varchar(32);unique;not null"`
	Num  int

	// 外键约束
	TeacherID int
	// 外键关联 并创建级联删除主表删除时此表则同步删除相关记录 如不使用级联删除则主表删除id时必须先手动删除子表记录
	Teacher Teacher `gorm:"foreignkey:TeacherID;constraint:OnDelete:CASCADE;"`
	// 当外键结构体被外部进行 Class.Teacher时可以直接访问到对应结构体的相应实例化对象

	// 在多个结构体中声明关联中间表不会进行重复创建
	Students []Student `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}
