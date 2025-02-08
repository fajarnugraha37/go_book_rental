package filter

type QueryFilter struct {
	Comparator  ComparatorType // eq, neq, like, notlike, in, notin, lt, lte, gt, gte
	Field       string
	Param       any
	SecondParam any
	Predicate   *Predicate
}
