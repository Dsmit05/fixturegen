package example

type ElfBuilder struct{ s *Elf }

func ElfBuilderNew() *ElfBuilder {
	return &ElfBuilder{s: &Elf{}}
}
func (b *ElfBuilder) P() *Elf {
	return b.s
}
func (b *ElfBuilder) V() Elf {
	return *b.s
}
func (b *ElfBuilder) Lifetime(v time.Time) *ElfBuilder {
	b.s.Lifetime = v
	return b
}
func (b *ElfBuilder) Age(v uint64) *ElfBuilder {
	b.s.Age = v
	return b
}
func (b *ElfBuilder) Name(v string) *ElfBuilder {
	b.s.Name = v
	return b
}
func (b *ElfBuilder) Friend(v []unsafe.Pointer) *ElfBuilder {
	b.s.Friend = v
	return b
}
func (b *ElfBuilder) Familiar(v Familiar) *ElfBuilder {
	b.s.Familiar = v
	return b
}

type FamiliarBuilder struct{ s *Familiar }

func FamiliarBuilderNew() *FamiliarBuilder {
	return &FamiliarBuilder{s: &Familiar{}}
}
func (b *FamiliarBuilder) P() *Familiar {
	return b.s
}
func (b *FamiliarBuilder) V() Familiar {
	return *b.s
}
func (b *FamiliarBuilder) Lifetime(v time.Time) *FamiliarBuilder {
	b.s.Lifetime = v
	return b
}
func (b *FamiliarBuilder) Name(v string) *FamiliarBuilder {
	b.s.Name = v
	return b
}
