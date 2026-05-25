data "zedcloud_deployment" "by_name" {
  depends_on = [
    zedcloud_deployment.tf_deployment,
  ]

  name       = zedcloud_deployment.tf_deployment.name
  project_id = zedcloud_project.test_tf_project.id
}

data "zedcloud_deployment" "by_id" {
  depends_on = [
    zedcloud_deployment.tf_deployment,
  ]

  id         = zedcloud_deployment.tf_deployment.id
  project_id = zedcloud_project.test_tf_project.id
}
