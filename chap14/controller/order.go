package controller

import (
	"encoding/json"
	"fmt"
	"golang-beginner-14/order"
	"io"
	"os"
	"strings"
	"time"
)

func restoMenu() []order.Item {
	menuList := []order.Item{
		{Id: 1, Name: "Fried Rice", Price: 15000.0},
		{Id: 2, Name: "Cheese Burger", Price: 25000.0},
		{Id: 3, Name: "Mocha Float", Price: 15000.0},
		{Id: 4, Name: "Ice Tea", Price: 6000.0},
		{Id: 5, Name: "Hot Coffe", Price: 12000.0},
		{Id: 6, Name: "Tiramisu Slice Cake", Price: 25000.0},
		{Id: 7, Name: "Fried Potato", Price: 10000.0},
		{Id: 8, Name: "Mix platter", Price: 14000.0},
		{Id: 9, Name: "Avocado Juice", Price: 15000.0},
		{Id: 10, Name: "Soda", Price: 10000.0},
	}
	return menuList
}

func CreateOrder() {
	for {
		ClearScreen()
		var option int
		menus := restoMenu()
		fmt.Println("--------- Order Food ---------")
		fmt.Println("--------- Menu ---------")
		for _, item := range menus {
			fmt.Printf("%d | %s | Rp. %.2f\n", item.Id, item.Name, item.Price)
		}
		fmt.Println("99. Back")
		fmt.Println("Choose menu: ")
		fmt.Scan(&option)
		if option == 0 || option > len(menus) {
			if option == 99 {
				return
			}
			fmt.Println("Invalid option")
			continue
		} else {
			go createJsonData(option)
			return
		}
	}
}

func createJsonData(menu_id int) {
	itemToOrder := order.Item{}
	for _, item := range restoMenu() {
		if item.Id == menu_id {
			itemToOrder = order.Item{
				Id:    item.Id,
				Name:  item.Name,
				Price: item.Price,
			}
		}
	}

	file, err := os.OpenFile("orders.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	var orders []order.Order
	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&orders); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	}

	// Create the new order
	if len(orders) == 0 {
		newOrder := order.Order{
			Id:            1,
			Items:         []order.Item{itemToOrder},
			OrderStatus:   "On Process",
			TotalPrice:    itemToOrder.Price,
			PaymentStatus: false,
		}
		orders = append(orders, newOrder)
	} else {
		lastOrder := orders[len(orders)-1]
		newOrder := order.Order{
			Id:            lastOrder.Id + 1,
			Items:         []order.Item{itemToOrder},
			OrderStatus:   "On Process",
			TotalPrice:    itemToOrder.Price,
			PaymentStatus: false,
		}
		orders = append(orders, newOrder)
	}

	// Move the file pointer to the beginning for writing
	file.Seek(0, 0)  // Go back to the start of the file
	file.Truncate(0) // Clear the file content

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&orders); err != nil {
		fmt.Println("Encode error message: ", err)
		return
	}
}

func EditOrder() {
	for {
		ClearScreen()
		var option, id int

		fmt.Println("--- Add Food To Unpaid Order List ---")
		unpaidOrderList()
		fmt.Println("99. Back")
		fmt.Println("Choose order to edit: ")
		fmt.Scan(&id)

		menus := restoMenu()
		fmt.Println("--------- Menu ---------")
		for _, item := range menus {
			fmt.Printf("%d | %s | Rp. %.2f\n", item.Id, item.Name, item.Price)
		}
		fmt.Println("Choose menu: ")
		fmt.Scan(&option)
		if option == 0 || option > len(menus) {
			if option != 99 {
				fmt.Println("Invalid option")
				continue
			}
			return
		} else {
			go editJsonData(option, id, "add_item", "")
		}

		if !promptContinue("edit_order") {
			ClearScreen()
			return
		}
	}
}

func unpaidOrderList() {
	var orders []order.Order
	var unpaidOrders []order.Order

	file, err := os.OpenFile("orders.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&orders); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	}

	if len(orders) == 0 {
		fmt.Println("Please create new order first")
	} else {
		for _, order := range orders {
			if !order.PaymentStatus {
				unpaidOrders = append(unpaidOrders, order)
			}
		}
		unpaidOrdersJson, err := json.MarshalIndent(unpaidOrders, "", "  ")
		if err != nil {
			fmt.Println("Marshal error message: ", err)
			return
		}

		// Print the filtered orders as JSON string
		// ClearScreen()
		fmt.Println("--------- Unpaid Order ---------")
		if len(unpaidOrdersJson) == 0 {
			fmt.Println("There is no paid order yet")
		} else {
			fmt.Println(string(unpaidOrdersJson))
		}
	}
}

func paidOrderList() {
	var orders []order.Order
	var paidOrders []order.Order

	file, err := os.OpenFile("orders.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&orders); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	}

	if len(orders) == 0 {
		fmt.Println("Please create new order first")
	} else {
		for _, order := range orders {
			if order.PaymentStatus {
				paidOrders = append(paidOrders, order)
			}
		}
		paidOrdersJson, err := json.MarshalIndent(paidOrders, "", "  ")
		if err != nil {
			fmt.Println("Marshal error message: ", err)
			return
		}

		// Print the filtered orders as JSON string
		ClearScreen()
		fmt.Println("--------- Paid Order ---------")
		if len(paidOrdersJson) == 0 {
			fmt.Println("There is no paid order yet")
		} else {
			fmt.Println(string(paidOrdersJson))
		}

	}
}

func editJsonData(menu_id, order_id int, flag, status string) {
	itemToOrder := order.Item{}
	orderFound := false
	var orders []order.Order

	if menu_id != 0 {
		for _, item := range restoMenu() {
			if item.Id == menu_id {
				itemToOrder = order.Item{
					Id:    item.Id,
					Name:  item.Name,
					Price: item.Price,
				}
			}
		}
	}

	// Open file for read and write
	file, err := os.OpenFile("orders.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&orders); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	}

	for i := range orders {
		if orders[i].Id == order_id {
			if flag == "add_item" {
				orders[i].Items = append(orders[i].Items, itemToOrder)
				orders[i].TotalPrice += itemToOrder.Price
			} else if flag == "pay_order" {
				orders[i].PaymentStatus = true
			} else if flag == "edit_status" {
				if orders[i].OrderStatus != status {
					orders[i].OrderStatus = status
				}
			}
			orderFound = true
			break
		}
	}

	if !orderFound {
		fmt.Println("--- Notification ---")
		fmt.Println("Invalid Order ID")
		return
	}

	// Move the file pointer to the beginning for writing
	file.Seek(0, 0)
	file.Truncate(0)

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&orders); err != nil {
		fmt.Println("Encode error message: ", err)
		return
	}
}

func PayOrder() {
	for {
		var order_id int

		unpaidOrderList()

		fmt.Println("99. Back")
		fmt.Println("Choose order id to pay: ")
		fmt.Scan(&order_id)

		if order_id != 0 {
			if order_id == 99 {
				return
			}
			go editJsonData(0, order_id, "pay_order", "")
		} else {
			fmt.Println("Invalid option")
			continue
		}

		if !promptContinue("pay_order") {
			ClearScreen()
			return
		}
	}
}

func PrintOrderHistory() {
	for {
		ClearScreen()
		var option int

		fmt.Println("--------- Order History ---------")
		fmt.Println("1. Unpaid Order")
		fmt.Println("2. Paid Order")
		fmt.Println("99. Back")
		fmt.Print("Choose Option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			ClearScreen()
			unpaidOrderList()
			if !promptContinue("show history again") {
				ClearScreen()
				return
			}
		case 2:
			ClearScreen()
			paidOrderList()
			if !promptContinue("show history again") {
				ClearScreen()
				return
			}
		case 99:
			return
		default:
			fmt.Println("Invalid Option")
			time.Sleep(3 * time.Second)
		}
		continue
	}
}

func EditOrderStatus() {
	for {
		var id int
		var status string
		statusEnum := []string{"on progress", "on delivery", "done"}

		fmt.Println("Paid Order")
		paidOrderList()
		fmt.Println("99. Back")
		fmt.Println("Choose order to edit the status: ")
		fmt.Scan(&id)
		if id == 99 {
			return
		}

		fmt.Println("Update status to (On progress/On delivery/Done): ")
		fmt.Scan(&status)

		if status != "" || includes(statusEnum, status) {
			go editJsonData(0, id, "edit_status", status)
			return
		} else {
			fmt.Println("Invalid status input")
			time.Sleep(3 * time.Second)
			continue
		}
	}
}

func includes(slice []string, value string) bool {
	for _, v := range slice {
		if strings.ToLower(v) == strings.ToLower(value) {
			return true
		}
	}
	return false
}
