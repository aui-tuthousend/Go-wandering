package main

import "fmt"

type node struct {
	noID         int
	namaKaryawan string
	prev         *node
	next         *node
}

type doublyLinkedList struct {
	head *node
	tail *node
	size int
}

func (list *doublyLinkedList) pushBack(ID int, namaK string) {
	newNode := &node{noID: ID, namaKaryawan: namaK}
	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.size++
}

func (list *doublyLinkedList) pushFront(ID int, namaK string) {
	newNode := &node{noID: ID, namaKaryawan: namaK}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
	list.size++
}

func (list *doublyLinkedList) deleteNode(value int) bool {
	node := list.head

	for node != nil {
		if node.noID == value {
			// If the node is the head, update the head pointer
			if node == list.head {
				list.head = node.next
			} else {
				node.prev.next = node.next
			}

			// If the node is the tail, update the tail pointer
			if node == list.tail {
				list.tail = node.prev
			} else {
				node.next.prev = node.prev
			}

			// Remove the node from the list
			node = nil

			list.size--

			return true
		}

		node = node.next
	}
	
	fmt.Println("ID tidak ditemukan")
	return false
}

func (list *doublyLinkedList) updateNode(aidi int) bool {
	var tempaidi int
	var tempnama string
	node := list.head

	for node != nil {
		if node.noID == aidi {
			fmt.Print("Masukkan ID baru: ")
			fmt.Scanln(&tempaidi)
			fmt.Print("Masukkan Nama baru: ")
			fmt.Scanln(&tempnama)

			node.noID = tempaidi
			node.namaKaryawan = tempnama

			return true
		}
		node = node.next
	}
	fmt.Println("ID tidak ditemukan")
	return false
}

func (list *doublyLinkedList) PrintForward() {
	current := list.head
	for current != nil {
		fmt.Print(current.noID, "     ")
		fmt.Println(current.namaKaryawan)
		current = current.next
	}
}

func (list *doublyLinkedList) viewByID (id int) bool {
	node := list.head
	
	for node != nil {
		if node.noID == id {
			fmt.Print("\nID Karyawan: ", node.noID)
			fmt.Print("\nNama Karyawan: ", node.namaKaryawan)

			return true
		}

		node = node.next
	}

	fmt.Println("ID tidak ditemukan")
	return false
}

func (list *doublyLinkedList) getSize() int {
	return list.size
}

// func (list *doublyLinkedList) isEmpty() bool {
//     return list.size == 0
// }

func menu() {
	fmt.Println("Pilihan Menu:")
	fmt.Println("[1]. Insert Depan")
	fmt.Println("[2]. Insert Belakang")
	fmt.Println("[3]. Update")
	fmt.Println("[4]. Hapus")
	fmt.Println("[5]. View")
	fmt.Println("[6]. View By ID")
	fmt.Println("[7]. Menu Utama")
}

func (list *doublyLinkedList) clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func main() {
	list := &doublyLinkedList{}
	kasir := &doublyLinkedList{}
	var lanjut string
	var pilih, pilih1, pilih2 int

	for {
		var tempId, hapus, apdet int
		var tempNama string

		fmt.Print("\033[H\033[2J")

		fmt.Println("\n---------------------------------------------------")
		fmt.Println("Data Karyawan PT. Dilihat-lihat malah gak kelihatan")
		fmt.Println("---------------------------------------------------")
		fmt.Println("Data yang tersedia: ")
		fmt.Println("[1]. Member ")
		fmt.Println("[2]. Kasir")
		fmt.Print("Masukkan data yang akan diolah: ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			fmt.Println("\nData Member")
			menu()
			fmt.Print("Pilih Menu: ")
			fmt.Scanln(&pilih1)
			switch pilih1 {
			case 1:
				fmt.Print("\nMasukkan ID: ")
				fmt.Scanln(&tempId)
				fmt.Print("Masukkan Nama: ")
				fmt.Scanln(&tempNama)

				list.pushFront(tempId, tempNama)

			case 2:
				fmt.Print("\nMasukkan ID: ")
				fmt.Scanln(&tempId)
				fmt.Print("Masukkan Nama: ")
				fmt.Scanln(&tempNama)

				list.pushBack(tempId, tempNama)

			case 3:
				//update
				if list.getSize() == 0 {
					fmt.Println("\nBelum ada data")
					break
				}

				fmt.Print("\nMasukkan no ID yang ingin di Update: ")
				fmt.Scanln(&apdet)
				list.updateNode(apdet)

			case 4:
				//delete
				if list.getSize() == 0 {
					fmt.Println("\nBelum ada data")
					break
				}
				fmt.Print("\nMasukkan No ID yang akan dihapus: ")
				fmt.Scanln(&hapus)
				list.deleteNode(hapus)

			case 5:
				fmt.Println("\n----------------------------------------------------")
				fmt.Println("  Data Member PT. Dilihat-lihat Malah ga kelihatan")
				fmt.Println("----------------------------------------------------")
				fmt.Println("ID     Nama\n")
				list.PrintForward()

			case 6:
				var id int
				fmt.Print("Masukkan ID karyawan: ")
				fmt.Scanln(&id)
				list.viewByID(id)
				fmt.Println("Posisi: Member")

			case 7:
				lanjut = "y"
				continue

			default:
				fmt.Println("Maaf Menyu tida ada :))")
			}
		} else if pilih == 2 {
			fmt.Println("\nData Kasir  ")
			menu()
			fmt.Print("Pilih Menu: ")
			fmt.Scanln(&pilih2)
			switch pilih2 {
			case 1:
				fmt.Print("\nMasukkan ID: ")
				fmt.Scanln(&tempId)
				fmt.Print("Masukkan Nama: ")
				fmt.Scanln(&tempNama)

				kasir.pushFront(tempId, tempNama)

			case 2:
				fmt.Print("\nMasukkan ID: ")
				fmt.Scanln(&tempId)
				fmt.Print("Masukkan Nama: ")
				fmt.Scanln(&tempNama)

				kasir.pushBack(tempId, tempNama)

			case 3:
				//update
				if kasir.getSize() == 0 {
					fmt.Println("\nBelum ada data")
					break
				}
				fmt.Print("\nMasukkan no ID yang ingin di Update: ")
				fmt.Scanln(&apdet)
				kasir.updateNode(apdet)

			case 4:
				//delete
				if kasir.getSize() == 0 {
					fmt.Println("\nBelum ada data")
					break
				}
				fmt.Print("\nMasukkan No ID yang akan dihapus: ")
				fmt.Scanln(&hapus)
				kasir.deleteNode(hapus)

			case 5:
				fmt.Println("\n---------------------------------------------------")
				fmt.Println("  Data Kasir PT. Dilihat-lihat Malah ga kelihatan")
				fmt.Println("---------------------------------------------------")
				fmt.Println("ID     Nama\n")
				kasir.PrintForward()

			case 6:
				//view by id
				var id int
				fmt.Print("\nMasukkan ID karyawan: ")
				fmt.Scanln(&id)
				kasir.viewByID(id)
				fmt.Println("\nPosisi: Kasir")

			case 7:
				lanjut = "y"
				continue

			default:
				fmt.Println("Maaf Menyu tida ada :))")
			}
		}

		fmt.Print("\napakah ingin lanjut? ")
		fmt.Scanln(&lanjut)
		if lanjut == "t" {
			break
		}
	}

	// list.pushBack(10, "kuku")
	// list.pushBack(11, "gebyi")
	// list.pushBack(12, "kowals Q")

	// fmt.Println("List size:", list.getSize())
	// fmt.Println("List elements:")

	// list.deleteNode(3.5)

	list.clear()
	kasir.clear()

	// fmt.Println("List size:", list.getSize())
	// fmt.Println("List elements:")

}
