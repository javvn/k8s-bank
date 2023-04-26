locals {

  jenkins_config = yamldecode(file("../values-jenkins.yaml"))

  values = [for k, v in var.values : file(v)]
}