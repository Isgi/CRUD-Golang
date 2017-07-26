package main

import (
	"net/http"
	"encoding/json"
	"log"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
)

type HttpResp struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

type Vm struct {
  ID           int `json:"id"`
	VmName       string `json:"vm_name"`
	Os           string `json:"os"`
	IpAddress    string `json:"ip_address"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
  AppName      string `json:"app_name"`
  Responsible  string `json:"responsible"`
  Vlan         string `json:"vlan"`
}

func GetVm(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	//initialization object connect() from mysql.go
	db := connect()
	//db closed after the completion of this function in execution
	defer db.Close()

	vms := make([]Vm, 0)

	key := r.URL.Query()["search"]
	search := "%"
	if len(key) > 0 {
		search = key[0]+"%"
	}

	stmt, err := db.Prepare("SELECT * FROM vm WHERE vm_name LIKE ? OR ip_address LIKE ?")
	if err != nil {
		log.Print(err.Error())
	}

	results, err := stmt.Query(search,search)
	defer results.Close()

	for results.Next()  {
		var vm Vm
		err = results.Scan(&vm.ID, &vm.VmName, &vm.Os, &vm.IpAddress, &vm.Port, &vm.User, &vm.Password, &vm.AppName, &vm.Responsible, &vm.Vlan)
		if err != nil {
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to select all from vms"})
		}
		vms = append(vms, vm)
	}

	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
	log.Print(vms)
	json.NewEncoder(w).Encode(vms)
}

func GetVmID(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	vmID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}

	var vm Vm
	err = db.QueryRow("SELECT * FROM vm where id = ?", vmID).Scan(&vm.ID, &vm.VmName, &vm.Os, &vm.IpAddress, &vm.Port, &vm.User, &vm.Password, &vm.AppName, &vm.Responsible, &vm.Vlan)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to select vm from database"})
	}

	json.NewEncoder(w).Encode(vm)
}

func AddVm(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var vm Vm
	err := decoder.Decode(&vm)
  log.Print(vm.VmName)
	if err != nil {
		log.Print(err.Error())
	}

	stmt, _ :=db.Prepare("INSERT INTO vm (vm_name, os, ip_address, port, user, password, app_name, responsible, vlan ) values (?,?,?,?,?,?,?,?,?)")

	res, err := stmt.Exec(vm.VmName, vm.Os, vm.IpAddress, vm.Port, vm.User, vm.Password, vm.AppName, vm.Responsible, vm.Vlan)

	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert vm into database"})
	}

	id, err := res.LastInsertId()
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to get last insert id"})
	}

	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted Vm Into the Database", Body: strconv.Itoa(int(id))})
}

func EditVm(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var vm Vm
	err := decoder.Decode(&vm)

	if err != nil {
		log.Print(err.Error())
	}

	vars := mux.Vars(r)
	vmID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}

	stmt, _ := db.Prepare("UPDATE vm SET vm_name = ?, os = ?, ip_address = ?, port = ?, user = ?, password = ?, app_name = ?, responsible = ?, vlan = ? WHERE id = ?")

	_, err = stmt.Exec(vm.VmName, vm.Os, vm.IpAddress, vm.Port, vm.User, vm.Password, vm.AppName, vm.Responsible, vm.Vlan, vmID)

	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to Update Post in the Database"})
	}
	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Update Vm in the Database"})
}

func DeleteVm(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	vmID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}

	stmt, err := db.Prepare("DELETE FROM vm where id = ?")
	if err != nil {
		log.Print(err.Error())
	}

	_, err = stmt.Exec(vmID)
	log.Print("Get Id ",err)
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to delete vm from database"})
	}

	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Deleted vm from the Database"})
}
