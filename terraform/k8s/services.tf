resource "kubernetes_service" "app-master" {
    metadata {
        name = "app-master"
    }

    spec {
        selector = {
          app  = "bankingOnGolang"
        }
        port {
            name        = "http"
            port        = 80
            target_port = ${}appPort}
        }

#Expose_K8_Ports()

        type = "LoadBalancer"
    }
  
}
