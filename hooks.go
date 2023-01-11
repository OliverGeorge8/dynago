package dynago

type BeforeUpdate interface {
	BeforeUpdate()
}

type AfterUpdate interface {
	AfterUpdate()
}

type BeforeInsert interface {
	BeforeInsert()
}

type AfterInsert interface {
	AfterInsert()
}
