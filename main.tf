provider "apex" {
  address = "http://localhost"
  port    = "3001"
  token   = "authToken"
}

resource "apex_example_volume" "test_volume" {
    volume_id = "vol-1633382"
    name = "New-Volume-Richmond"
    appliance_id = "adss-2943821"
    description = "Adding a test volume"
    size = 15602
}

output "rich-volume-id" {
    value = apex_example_volume.test_volume
}

resource "apex_example_server" "richang-server" {
	server_count = "12"
}

output "richang-server-id" {
    value = apex_example_server.richang-server.id

}
