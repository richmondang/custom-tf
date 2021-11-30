resource "apex_example_server" "richang-server" {
	server_count = "1"
}

output "richang-server-id" {
    value = apex_example_server.richang-server.id

}

resource "apex_example_volume" "richang-volume"{
    resource_id = "123456"
    name = "New Volume-Richmond"
    description = "Test create new volume"
    size = 15602
}

output "rich-volume-id" {
    value = apex_example_volume.richang-volume.id
}