//20:20 - time completed
type OrderedStream struct {
	ptr int
	list []string
}


func Constructor(n int) OrderedStream {
	return OrderedStream{
		list: make([]string, n),
		ptr: 0,
	}
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.list[idKey-1] = value

	if this.list[this.ptr] == ""{
		return []string{}
	}

	var end = this.ptr
	for this.ptr < len(this.list) && this.list[this.ptr] != ""{
		this.ptr++
		end++
	}

	return this.list[idKey-1:end]
}
