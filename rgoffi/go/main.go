package main

import "C"

import (
	"fmt"
	"unsafe"
)

// TODO(threadedstream): play with that a bit
// func (
// 	deallocate_go_memory(n uintptr) *uint8 {
// 	}
// )

//export allocate_go_memory
func allocate_go_memory(n uintptr) *uint8{
	return systemAlloc(n, nil)
}

//export deallocate_go_memory
func deallocate_go_memory(v *uint8, n uintptr) {
	systemFree(v, n, nil)
}

//export persistent_alloc
func persistent_alloc(size, align uintptr, sysStat *uint64) *uint8 {
	return persistentAlloc(size, align, sysStat)
}

//export memclr_noheap_pointers
func memclr_noheap_pointers(ptr *uint8, size uintptr) {
	memclrNoHeapPointers(ptr, size)
}

//export runtime_lock2
func runtime_lock2(m *uint8) {
	lock2(m)
}

//export runtime_unlock2
func runtime_unlock2(m *uint8) {
	unlock2(m)
}

//export runtime_extend_random
func runtime_extend_random(r []byte, n int) {
	extendRandom(r, n)
}



//go:notinheap
type node struct {
	data int64 // 8
	next *node // 8
}

// //export fix_allocate
// func fix_allocate(f unsafe.Pointer) *uint8 {	
// 	falloc := (*fixalloc) (f)
// 	return falloc.fixalloc_alloc()
// }


//go:linkname runtime.mallocgc main.mallocgc
func mallocgc(size uintptr, typ *_type) *uint8 {
	return nil
}


func fixallocSample() {
	f := &fixalloc{}
	fmt.Println("initializing fixalloc\n")
	f.init(unsafe.Sizeof(node{}))
	n := (*node) (f.alloc())
	n.data = 0xcafebabe
	n.next = nil
	fmt.Printf("[data = %d, next: %p]\n", n.data, n.next)
	fmt.Println("freeing no longer needed memory\n")
	f.free(unsafe.Pointer(n))
	fmt.Printf("[data = %d, next: %p]\n", n.data, n.next)
}

func systemAllocSample() {
	fmt.Printf("attempt to allocate memory through stub...\n")
	ptr := (*uint8) (systemAlloc(256, nil))
	fmt.Printf("%p\n", ptr)
	writeContents(ptr, 16)
	printContents(ptr, 16)
	systemFree(ptr, 256, nil)
}

func main() {
	n := new(int)
	*n = 32
	fmt.Printf("%x", n)	
}