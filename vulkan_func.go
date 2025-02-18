package main

import (
	"fmt"
	"log"
	"unsafe"
)

// CreateVulkanInstance – Vytvorenie Vulkan instance
func CreateVulkanInstance() VkInstance {
	var instance VkInstance
	instanceCreateInfo := VkInstanceCreateInfo{} // Použijeme prázdnu štruktúru

	res := vkCreateInstance(unsafe.Pointer(&instanceCreateInfo), nil, &instance)
	if res != VK_SUCCESS {
		log.Fatalf("❌ vkCreateInstance zlyhalo s kódom: %d", res)
	}
	fmt.Println("✅ vkCreateInstance úspešne zavolané!")
	return instance
}

// GetVulkanDevices – Získanie dostupných Vulkan zariadení
func GetVulkanDevices(instance VkInstance) {
	var deviceCount uint32
	vkEnumeratePhysicalDevices(instance, &deviceCount, nil)
	fmt.Printf("🔍 Počet Vulkan zariadení: %d\n", deviceCount)

	if deviceCount == 0 {
		log.Fatal("❌ Žiadne Vulkan zariadenia!")
	}

	devices := make([]VkPhysicalDevice, deviceCount)
	vkEnumeratePhysicalDevices(instance, &deviceCount, &devices[0])

	for i, device := range devices {
		var properties VkPhysicalDeviceProperties
		vkGetPhysicalDeviceProperties(device, &properties)

		name := string(properties.DeviceName[:])
		name = name[:len(name)-1] // Odstránenie \x00

		fmt.Printf("📌 Zariadenie %d: %s\n", i+1, name)
		fmt.Printf("   - Vulkan verzia: %d\n", properties.ApiVersion)
		fmt.Printf("   - Driver verzia: %d\n", properties.DriverVersion)
		fmt.Printf("   - Vendor ID: 0x%X\n", properties.VendorID)
		fmt.Printf("   - Device ID: 0x%X\n", properties.DeviceID)
		// 🔹 Získanie RAM a VRAM
		var memProperties VkPhysicalDeviceMemoryProperties
		vkGetPhysicalDeviceMemoryProperties(device, &memProperties)

		var totalVRAM uint64
		var totalRAM uint64

		for _, heap := range memProperties.MemoryHeaps[:memProperties.MemoryHeapCount] {
			if heap.Flags&1 != 0 { // VRAM (DEVICE_LOCAL_BIT)
				totalVRAM += heap.Size
			} else { // RAM
				totalRAM += heap.Size
			}
		}

		fmt.Printf("   - VRAM: %.2f GB\n", float64(totalVRAM)/(1024*1024*1024))
		fmt.Printf("   - RAM:  %.2f GB\n", float64(totalRAM)/(1024*1024*1024))
	}
}
