package main

import "unsafe"

// Vulkan základné typy
type VkInstance uintptr
type VkPhysicalDevice uintptr
type VkDevice uintptr
type VkResult int32

// Vulkan konštanty
const (
	VK_SUCCESS VkResult = 0
)

// Štruktúry pre Vulkan vlastnosti zariadenia
type VkPhysicalDeviceProperties struct {
	ApiVersion    uint32
	DriverVersion uint32
	VendorID      uint32
	DeviceID      uint32
	DeviceName    [256]byte
	DeviceType    uint32
}

// VkInstanceCreateInfo – Nevyhnutná pre vkCreateInstance
type VkInstanceCreateInfo struct {
	SType                 uint32
	PNext                 unsafe.Pointer
	Flags                 uint32
	ApplicationInfo       unsafe.Pointer
	EnabledLayerCount     uint32
	PpEnabledLayerNames   unsafe.Pointer
	EnabledExtensionCount uint32
	PpEnabledExtensionNames unsafe.Pointer
}

// Vulkan štruktúry pre RAM a VRAM
type VkMemoryHeap struct {
	Size  uint64
	Flags uint32
	_     uint32 // Padding
}

type VkPhysicalDeviceMemoryProperties struct {
	MemoryTypeCount uint32
	MemoryTypes     [32]VkMemoryType
	MemoryHeapCount uint32
	MemoryHeaps     [16]VkMemoryHeap
}

type VkMemoryType struct {
	HeapIndex uint32
	PropertyFlags uint32
}

// func type
type VkGetPhysicalDeviceMemoryPropertiesFunc func(VkPhysicalDevice, *VkPhysicalDeviceMemoryProperties)
type VkCreateInstanceFunc func(unsafe.Pointer, unsafe.Pointer, *VkInstance) VkResult
type VkEnumeratePhysicalDevicesFunc func(VkInstance, *uint32, *VkPhysicalDevice) VkResult
type VkGetPhysicalDevicePropertiesFunc func(VkPhysicalDevice, *VkPhysicalDeviceProperties)

