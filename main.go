package main

func main() {
    LoadVulkan()
    defer UnloadVulkan()

    instance := CreateVulkanInstance()
    GetVulkanDevices(instance)
    GetVulkanDevices0(instance)
}
