package main

import (
    "fmt"
    "log"

    "github.com/ebitengine/purego"
)

var (
    vulkanLib                     uintptr
    vkCreateInstance               VkCreateInstanceFunc
    vkEnumeratePhysicalDevices     VkEnumeratePhysicalDevicesFunc
    vkGetPhysicalDeviceProperties  VkGetPhysicalDevicePropertiesFunc
    vkGetPhysicalDeviceMemoryProperties VkGetPhysicalDeviceMemoryPropertiesFunc
)

// LoadVulkan – Načítanie Vulkan knižnice a funkcií
func LoadVulkan() {
    var err error
    vulkanLib, err = purego.Dlopen("libvulkan.so", purego.RTLD_NOW)
    if err != nil {
        log.Fatalf("❌ Vulkan knižnica nenájdená: %v", err)
    }
    fmt.Println("✅ Vulkan knižnica načítaná!")

    // Načítanie funkcií
    purego.RegisterLibFunc(&vkCreateInstance, vulkanLib, "vkCreateInstance")
    purego.RegisterLibFunc(&vkEnumeratePhysicalDevices, vulkanLib, "vkEnumeratePhysicalDevices")
    purego.RegisterLibFunc(&vkGetPhysicalDeviceProperties, vulkanLib, "vkGetPhysicalDeviceProperties")
    purego.RegisterLibFunc(&vkGetPhysicalDeviceMemoryProperties, vulkanLib, "vkGetPhysicalDeviceMemoryProperties")
}

// UnloadVulkan – Uvoľnenie knižnice
func UnloadVulkan() {
    if vulkanLib != 0 {
        purego.Dlclose(vulkanLib)
        fmt.Println("🔄 Vulkan knižnica uvoľnená!")
        vulkanLib = 0
    }
}

