output "helm_jenkins" {
  value     = helm_release.this
  sensitive = true
}

#output "jenkins_config" {
#  value = local.jenkins_config
#}