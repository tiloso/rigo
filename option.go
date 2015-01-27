package rigo

type option func(interface{})

func Timeout(v uint32) option {
	return func(i interface{}) {
		if t, ok := i.(timeouter); ok {
			t.setTimeout(v)
		}
		// TODO panic if !ok (e.g. user tries to set option on not supported
		// entity?)
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
