package rigo

// how to make it more general? e.g. specify each option on package level once
// but allow to use for any element (e.g. Index, Bucket, ...)

// use interfaces instead and test if type assertion works?

// >> Rob's options / self referential functions
// option sets the options specified.

// idea embed rpb.Stuff in query step (e.g. Index) directly to avoid duplication
// of Fields

// example usage
// I([]byte("key_bin")).Option(rigo.PaginationSort(true))

// <<

type option func(interface{})

// How to group options in godoc without exporting option => option?
func Timeout(v uint32) option {
	return func(i interface{}) {
		if t, ok := i.(timeouter); ok {
			t.setTimeout(v)
		}
		// panic instead if !ok (e.g. user tries to set option on not supported entity?)
	}
}

type timeouter interface {
	setTimeout(uint32)
}

func ReturnTerms(v bool) option {
	return func(i interface{}) {
		if t, ok := i.(returnTermer); ok {
			t.setReturnTerms(v)
		}
	}
}

type returnTermer interface {
	setReturnTerms(bool)
}

func MaxResults(v uint32) option {
	return func(i interface{}) {
		if t, ok := i.(maxResulter); ok {
			t.setMaxResults(v)
		}
	}
}

type maxResulter interface {
	setMaxResults(uint32)
}

func PaginationSort(v bool) option {
	return func(i interface{}) {
		if t, ok := i.(paginationSorter); ok {
			t.setPaginationSort(v)
		}
	}
}

type paginationSorter interface {
	setPaginationSort(bool)
}

func Head(v bool) option {
	return func(i interface{}) {
		if t, ok := i.(header); ok {
			t.setHead(v)
		}
	}
}

type header interface {
	setHead(bool)
}

func IfModified(v []byte) option {
	return func(i interface{}) {
		if t, ok := i.(ifModifiedSetter); ok {
			t.setIfModified(v)
		}
	}
}

type ifModifiedSetter interface {
	setIfModified([]byte)
}
