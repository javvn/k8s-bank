resource "kubernetes_namespace" "this" {
  metadata {
    name = "jawn-bank"
  }
}

resource "helm_release" "this" {
  namespace = kubernetes_namespace.this.metadata[0].name
  name      = var.name
  chart     = var.chart
  values    = local.values


  depends_on = [
    kubernetes_namespace.this
  ]
  #  version    = "6.0.1"
  #
  #  set {
  #    name  = "cluster.enabled"
  #    value = "true"
  #  }
  #
  #  set {
  #    name  = "metrics.enabled"
  #    value = "true"
  #  }
  #
  #  set {
  #    name  = "service.annotations.prometheus\\.io/port"
  #    value = "9127"
  #    type  = "string"
  #  }
}

resource "null_resource" "this" {
  triggers = {
    jenkins_release_id = helm_release.this.id
  }

  provisioner "local-exec" {
    command = "kubectl exec -n ${var.namespace} -i svc/${var.name} -c jenkins -- /bin/cat /run/secrets/additional/chart-admin-password && echo"
  }


  provisioner "local-exec" {
    command = "kubectl -n ${var.namespace} port-forward svc/${var.name} ${local.jenkins_config.controller.servicePort}:${local.jenkins_config.controller.servicePort}"
  }


}